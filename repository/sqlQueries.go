package repository

const (
	selectEmployees = "SELECT * FROM select_users($1,$2,$3,$4,$5,$6,$7)"
	insertEmployees = "SELECT * FROM insert_user($1::uuid,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
)
