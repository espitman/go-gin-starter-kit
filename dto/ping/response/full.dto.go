package dto_ping

type Full_Payload struct {
	Message string
}

type Full struct {
	Message string
	Status  int
	Payload Full_Payload
}
