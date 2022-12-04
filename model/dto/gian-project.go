package dto

type (
	ClassifiedList struct {
		Unclassified []int `json:"sin clasificar"`
		Classified   []int `json:" clasificado"`
	}
)
