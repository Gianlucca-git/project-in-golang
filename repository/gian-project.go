package repository

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"strings"

	"IMPORTS/model/dto"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"log"
)

// RepositoryManager constructs a new NewRepositoryManager
type RepositoryManager interface {
	GetEmployees(ctx context.Context, request *dto.GetUsersRequest) (error, *dto.UsersResponse)
	InsertUser(ctx context.Context, user *dto.User) error
}

func NewRepositoryManager(repository Type) RepositoryManager {
	switch repository {
	case PostgresSQL:
		return &RepositoryStruct{DB: NewSQLConnection()}
	}

	return nil
}

type RepositoryStruct struct {
	*sql.DB
}

func (rh *RepositoryStruct) GetEmployees(ctx context.Context, request *dto.GetUsersRequest) (error, *dto.UsersResponse) {
	log.Print("[INFO] init: Repository GetUsers()")
	prepare, err := rh.DB.PrepareContext(ctx, selectEmployees)
	if err != nil {
		return err, nil
	}

	rows, err := prepare.QueryContext(
		ctx,
		request.Search,
		pq.Array(request.Countries),
		pq.Array(request.IdentificationsTypes),
		pq.Array(request.Departments),
		request.Status,
		request.Cursor,
		request.Limit,
	)
	if err != nil {
		return err, nil
	}
	defer func() { _ = rows.Close() }()

	var response dto.UsersResponse
	var emp dto.User
	var total int

	for rows.Next() {

		var othersNames sql.NullString
		err = rows.Scan(
			&total,
			&emp.Id,
			&emp.Name,
			&othersNames,
			&emp.LastName,
			&emp.SecondLastName,
			&emp.Country,
			&emp.IdentificationType,
			&emp.IdentificationNumber,
			&emp.Email,
			&emp.Department,
			&emp.Status,
		)
		if err != nil {
			return err, nil
		}

		emp.OthersNames = othersNames.String

		response.Users = append(response.Users, emp)
	}

	if len(response.Users) == 0 {
		return nil, nil //No content
	}

	response.LastCursor = b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("('%s','%s')", emp.Name, emp.Id)))
	response.TotalRegisters = total

	return nil, &response
}

func (rh *RepositoryStruct) InsertUser(ctx context.Context, user *dto.User) error {
	log.Print("[INFO] init: Repository InsertUser()")
	prepare, err := rh.DB.PrepareContext(ctx, insertEmployees)
	if err != nil {
		return err
	}

	rows, err := prepare.QueryContext(
		ctx,
		user.Id,
		strings.ToLower(user.Name),
		strings.ToLower(user.OthersNames),
		strings.ToLower(user.LastName),
		strings.ToLower(user.SecondLastName),
		user.CountryId,
		user.IdentificationTypeId,
		user.IdentificationNumber,
		user.Email,
		user.DepartmentId,
	)

	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	var response string
	if rows.Next() {
		err = rows.Scan(&response)
		if err != nil {
			return err
		}
	}

	log.Printf("[INFO] init: Repository InsertUser RESPONSE (%s)", response)
	if response == "finished successfully" {
		return nil
	}

	return errors.New(response)
}
