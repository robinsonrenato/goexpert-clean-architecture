package usecase

import "github.com/robinsonrenato/goexpert-clean-architecture/internal/entity"

type OrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(repo entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: repo}
}

func (uc *ListOrdersUseCase) Execute() ([]OrdersOutputDTO, error) {
	orders, err := uc.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	ordersDTO := make([]OrdersOutputDTO, len(orders))
	for i, order := range orders {
		ordersDTO[i] = OrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}
	return ordersDTO, nil
}
