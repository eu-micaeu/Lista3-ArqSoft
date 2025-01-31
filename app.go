package main

import (
	"fmt"

	"github.com/eu-micaeu/Lista3-ArqSoft/pck/auth"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/catalog"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/orders"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/payments"
)

// Função principal
func main() {

	// Autenticação
	valid := auth.Auth()

	for !valid {

		valid = auth.Auth()

	} 
		
	// Catálogo de produtos
	produtos := catalog.Catalog()
	
	// Pedido
	pedidos := orders.Order(produtos)

	// Pagamento
	pagamento := payments.Payment(pedidos)

	// Verificar o estado do pagamento
	if pagamento.State == "Paid" {

		fmt.Println("\n-> Pedido CONCLUIDO!")

	} else {

		fmt.Println("\n-> Pedido CANCELADO!")

	}
	
	fmt.Println("\n-= Fim da execução =-")

}