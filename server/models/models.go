package models

type Method uint8

type Information struct {
	Method Method `json:"mtd"`
	Data   string `json:"dat"`
}
