package model

type Return struct {
	Status 	bool 		`json:"status"`
	Data   	interface{}	`json:"data"`
	Message string		`json:"message"`
}
