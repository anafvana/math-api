package api

type MathRequest struct {
	FirstNumber  *int    `json:"first_number"`
	SecondNumber *int    `json:"second_number"`
	Operation    *string `json:"operation"`
}

type MathResponse struct {
	Result int `json:"result"`
}
