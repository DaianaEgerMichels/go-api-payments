package entity

type OrderRequest struct {
	OrderID  string  `json:"orderId"`
	CardHash string  `json:"cardHash"`
	Total    float64 `json:"total"`
}
