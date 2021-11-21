package model

type (
	ProductRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Price       int    `json:"price" validate:"required"`
		Quantity    int    `json:"quantity" validate:"required"`
	}

	ProductResponse struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		Quantity    int    `json:"quantity"`
		CreatedAt   string `json:"created_at,omitempty"`
	}
)
