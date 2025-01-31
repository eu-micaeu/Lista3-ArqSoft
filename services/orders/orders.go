package orders

import (
	"fmt"
	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/payments"
)

// OrderService gerencia os pedidos
type OrderService struct {
	Orders []models.Order
}

// NewOrderService cria uma nova instância de OrderService
func NewOrderService() *OrderService {
	return &OrderService{}
}

// CreateOrder cria um novo pedido
func (o *OrderService) CreateOrder(userID int, products []models.Product) *models.Order {
	order := models.Order{
		ID:      len(o.Orders) + 1,
		Products: products,
		State:   "created",
	}
	o.Orders = append(o.Orders, order)
	return &order
}

// ProcessOrder processa o pedido e chama o serviço de pagamentos
func (o *OrderService) ProcessOrder(order *models.Order, total float64) {
	payment := payments.NewPayment(order.ID, total)
	if payments.ProcessPayment(payment) {
		order.State = "Pago"
	} else {
		order.State = "Pagamento Falhou"
	}
	fmt.Printf("Status do pedido %d: %s\n", order.ID, order.State)
}