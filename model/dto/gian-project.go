package dto

type (
	ExampleRequest struct {
		Id          string         // url
		Name        string         `json:"name"`         // query params
		LastName    string         `json:"last_name"`    // query params
		YearsOld    int            `json:"years_old"`    // query params
		Status      int            `json:"status"`       // query params binary
		AddressInfo AddressRequest `json:"address_info"` // body
	}
	AddressRequest struct {
		City         string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		Address      string `json:"address"`
		PostalCode   int    `json:"postal_code"`
	}

	ExampleResponse struct {
		DataPerson string      `json:"person"`
		Message    string      `json:"message"`
		DataResult interface{} `json:"data_result"`
	}
)
