package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eu-micaeu/Lista3-ArqSoft/services/auth"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/catalog"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/orders"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Autenticação
	authService := auth.NewAuthService()
	user := authService.AuthenticateUser(reader)
	if user == nil {
		return
	}

	// Catálogo de produtos
	catalogService := catalog.NewCatalogService()
	products := catalogService.DisplayProducts()

	// Seleção de produtos
	selectedProducts, total := catalogService.SelectProducts(reader, products)

	if len(selectedProducts) == 0 {
		fmt.Println("Nenhum produto selecionado. Saindo...")
		return
	}

	// Criar e processar pedido
	orderService := orders.NewOrderService()
	order := orderService.CreateAndProcessOrder(user.ID, selectedProducts, total)
	orderService.PrintOrderSummary(order, total)
}
