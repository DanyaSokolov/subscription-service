package model

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      string     `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}


//swagger
type SubscriptionCreateRequest struct {
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int     `json:"price" example:"400"`
	UserID      string  `json:"user_id" example:"1e4adfee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string  `json:"start_date" example:"2025-04"`
	EndDate     *string `json:"end_date,omitempty" example:"2025-08"`
}

//swagger
type SubscriptionUpdateRequest struct {
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int     `json:"price" example:"500"`
	UserID      string  `json:"user_id" example:"e9c3a059-089d-4c5b-b0cc-17b5a78bb6ef"`
	StartDate   string  `json:"start_date" example:"2025-05"`
	EndDate     *string `json:"end_date,omitempty" example:"2025-12"`
}

type TotalCostRequest struct {
	ServiceName string  `json:"service_name" example:"Spotify"`
	UserID      string  `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string  `json:"start_date" example:"2025-07"`
	EndDate     *string `json:"end_date,omitempty" example:"2025-11"`
}

type TotalCostResponse struct {
	TotalCost int `json:"total_cost"`
}