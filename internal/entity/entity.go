package entity

import "errors"

type OrderRequest struct {
	OrderID  string  `json:"orderId"`
	CardHash string  `json:"cardHash"`
	Total    float64 `json:"total"`
}

// Validate OrderRequest

func (o *OrderRequest) Validate() error {
	if o.OrderID == "" {
		return errors.New("OrderID cannot be empty")
	}
	if o.CardHash == "" {
		return errors.New("CardHash cannot be empty")
	}
	if o.Total <= 0 {
		return errors.New("Total must be greater than 0")
	}
	return nil
}

// Processing business rule

