package models

type Service struct {
    ServiceId   string `json:"service_id"`
    Name        string `json:"name"`
    Price       float32 `json:"price" sql:"type:decimal(12,2)"`
}