package models

type Customer struct {
    CustomerId      string `json:"customer_id"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Email           string `json:"email"`
    Phone           string `json:"phone"`
    Address         string `json:"address"`
}