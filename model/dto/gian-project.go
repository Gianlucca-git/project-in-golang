package dto

type (
	ClassifiedList struct {
		Unclassified []int `json:"sin-clasificar"`
		Classified   []int `json:"clasificado"`
	}

	BalanceRequest struct {
		Months []string `json:"meses"`
		Sales  []int    `json:"ventas"`
		Bills  []int    `json:"gastos"`
	}
	BalanceGeneralResponse struct {
		Balances []BalanceInMonth `json:"balances"`
	}
	BalanceInMonth struct {
		Months  string `json:"mes"`
		Sales   int    `json:"ventas"`
		Bills   int    `json:"gastos"`
		Balance int    `json:"balance"`
	}

	GetUsersRequest struct {
		Search               string   `json:"search"`
		Countries            []string `json:"countries"`
		IdentificationsTypes []string `json:"identifications_types"`
		Departments          []string `json:"departments"`
		Status               string   `json:"status"`
		Cursor               string   `json:"cursor"`
		Limit                string   `json:"limit"`
		LimitInt             int      `json:"-"`
	}
	UsersResponse struct {
		LastCursor     string `json:"last_cursor"`
		TotalRegisters int    `json:"total_registers"`
		Users          []User `json:"users"`
	}
	User struct {
		Id                   string `json:"id,omitempty"`
		Name                 string `json:"name,omitempty"`
		OthersNames          string `json:"others_names,omitempty"`
		LastName             string `json:"last_name,omitempty"`
		SecondLastName       string `json:"second_last_name,omitempty"`
		CountryId            int    `json:"country_id,omitempty"`
		Country              string `json:"country,omitempty,omitempty"`
		IdentificationTypeId int    `json:"identification_type_id,omitempty"`
		IdentificationType   string `json:"identification_type,omitempty"`
		IdentificationNumber string `json:"identification_number,omitempty"`
		Admission            string `json:"admission,omitempty"`
		RegistrationDate     string `json:"registration_date,omitempty"`
		RegistrationHours    string `json:"registration_hours,omitempty"`
		Email                string `json:"email,omitempty"`
		DepartmentId         int    `json:"department_id,omitempty"`
		Department           string `json:"department,omitempty"`
		Status               string `json:"status,omitempty"`
	}
)
