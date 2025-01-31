package payments

import (
	"fmt"
	"time"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/orders"
)

// Função que simula um pagamento
func Payment(order models.Order) models.Order {

	var option string

	// Perguntar se o usuário deseja pagar o pedido
	fmt.Print("\nDeseja pagar o pedido? (sim/nao): ")
	fmt.Scan(&option)

	// Verificar a opção escolhida
	if option == "sim" {

		fmt.Println("\nPagamento sendo concluido...")

		// Timer de 3 segundos

		time.Sleep(3 * time.Second)

		fmt.Println("Pagamento concluido com sucesso!")

		// Alterar o estado do pedido para "Pago"
		order = orders.PayOrder(order)
		
		} else {
			
			fmt.Println("\nPagamento cancelado.")

		}
		
		return order

}