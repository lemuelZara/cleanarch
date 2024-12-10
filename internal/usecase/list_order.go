package usecase

import (
	"github.com/lemuelZara/cleanarch/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{OrderRepository: OrderRepository}
}

func (c *ListOrderUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	items, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var orders []ListOrdersOutputDTO
	for _, order := range items {
		orders = append(orders, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orders, nil
}
