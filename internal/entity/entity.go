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

func (o *OrderRequest) Process() (*OrderResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}
	orderResponse := NewOrderResponse(o.OrderID, "failed")
	// rule, if total is less than 100.00, set status as paid
	if o.Total < 100.00 {
		orderResponse.Status = "paid"
	}
	return orderResponse, nil
}

// Response for Order
type OrderResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"` // types of the status: paid, failed
}

func NewOrderResponse(orderID, status string) *OrderResponse {
	return &OrderResponse{
		OrderID: orderID,
		Status:  status,
	}
}