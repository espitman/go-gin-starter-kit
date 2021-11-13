package dto_test_response

type Full_Payload struct {
	ID string `json:"ID"`
}
type Full struct {
	Message string       `json:"message"`
	Status  int          `json:"status"`
	Payload Full_Payload `json:"payload"`
}
