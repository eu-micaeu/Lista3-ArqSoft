package payments

import (
	"log"
	"math/rand"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/fatih/color" 
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
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		success := rand.Float32() > 0.5
		if success {
			payment.Status = "Aprovado"
			// Log de pagamento aprovado com cor verde
			color.New(color.FgGreen).Printf("Pagamento aprovado para o pedido %d\n", payment.ID)
			return true
		}
		log.Printf("Tentativa %d falhou. Tentando novamente...\n", i+1)
	}
	payment.Status = "Falhado"
	color.New(color.FgRed).Printf("Pagamento recusado para o pedido %d\n", payment.ID)
	return false
}
