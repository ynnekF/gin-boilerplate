package models

type Response struct {
	Status         string      `json:"status"`
	Message        string      `json:"message"`
	Error          string      `json:"error"`
	Date           string      `json:"date"`
	Transaction_Id string      `json:"transaction_id"`
}