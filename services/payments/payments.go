package payments

import (
	"fmt"
	"math/rand"
	"github.com/eu-micaeu/Lista3-ArqSoft/models"
)

// NewPayment cria uma nova instância de Payment
func NewPayment(orderID int, amount float64) *models.Payment {
	return &models.Payment{
		ID:     orderID,
		Amount: amount,
		Status: "Pendente",
	}
}

// ProcessPayment simula o processamento de um pagamento
func ProcessPayment(payment *models.Payment) bool {
	// Simula uma falha de pagamento com 10% de chance
	if rand.Intn(10) == 0 {
		payment.Status = "Falho"
		fmt.Printf("Pagamento falhou para o pedido %d\n", payment.ID)
		return false
	}

	payment.Status = "Pago"
	fmt.Printf("Pagamento concluído com sucesso para o pedido %d\n", payment.ID)
	return true
}