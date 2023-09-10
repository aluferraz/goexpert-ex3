package usecase

import (
	"github.com/aluferraz/goexpert-ex3/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {

	orders, err := c.OrderRepository.ListAll()

	if err != nil {
		return []ListOrdersOutputDTO{}, err
	}
	dto := make([]ListOrdersOutputDTO, 0)
	for _, order := range orders {
		if err != nil {
			return []ListOrdersOutputDTO{}, err
		}
		dto = append(dto, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return dto, nil

}
