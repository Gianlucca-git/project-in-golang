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
)
