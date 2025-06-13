package models

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Age   int    `json:"age"`
}
