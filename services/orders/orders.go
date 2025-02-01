package orders

import (
	"fmt"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/payments"
	"github.com/fatih/color" 
)

// OrderService gerencia os pedidos
type OrderService struct {
	Orders []models.Order
}

// NewOrderService cria uma nova instância de OrderService
func NewOrderService() *OrderService {
	return &OrderService{}
}

// CreateAndProcessOrder cria e processa um novo pedido
func (o *OrderService) CreateAndProcessOrder(userID int, selectedProducts map[int]models.Product, total float64) *models.Order {
	order := o.CreateOrder(userID, selectedProducts)
	o.ProcessOrder(order, total)
	return order
}

// CreateOrder cria um novo pedido
func (o *OrderService) CreateOrder(userID int, selectedProducts map[int]models.Product) *models.Order {
	var productList []models.Product
	for _, p := range selectedProducts {
		productList = append(productList, p)
	}
	order := models.Order{
		ID:      len(o.Orders) + 1,
		Products: productList,
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
		// Log de pagamento com sucesso em verde
		color.New(color.FgGreen).Printf("Status do pedido %d: %s\n", order.ID, order.State)
	} else {
		order.State = "Pagamento Falhou"
		fmt.Printf("Status do pedido %d: %s\n", order.ID, order.State)
	}
}

// PrintOrderSummary exibe o resumo do pedido
func (o *OrderService) PrintOrderSummary(order *models.Order, total float64) {
	fmt.Println("\nResumo do Pedido:")
	for _, p := range order.Products {
		fmt.Printf("\n- %s: %d x R$ %.2f = R$ %.2f", p.Name, p.Stock, p.Price, p.Price*float64(p.Stock))
	}
	fmt.Printf("\nTotal do Pedido: R$ %.2f\n", total)
}
