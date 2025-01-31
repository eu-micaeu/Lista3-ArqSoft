package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eu-micaeu/Lista3-ArqSoft/services/auth"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/catalog"
	"github.com/eu-micaeu/Lista3-ArqSoft/services/orders"
	"github.com/eu-micaeu/Lista3-ArqSoft/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Autenticação
	authService := auth.NewAuthService()
	
	fmt.Println("Digite o email:")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	
	fmt.Println("Digite a senha:")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := authService.Login(email, password)
	if user == nil {
		fmt.Println("Falha na autenticação")
		return
	}

	// Catálogo de produtos
	catalogService := catalog.NewCatalogService()
	products := catalogService.GetAvailableProducts()

	// Seleção de produtos
	selectedProducts := make([]models.Product, 0)
	total := 0.0

	for {
		fmt.Println("\nDigite o ID do produto que deseja adicionar (ou 0 para finalizar):")
		productIDInput, _ := reader.ReadString('\n')
		productID, err := strconv.Atoi(strings.TrimSpace(productIDInput))
		if err != nil || productID < 0 {
			fmt.Println("ID inválido!")
			continue
		}
		
		if productID == 0 {
			break
		}

		// Encontrar produto
		var selectedProduct models.Product
		found := false
		for _, p := range products {
			if p.ID == productID {
				selectedProduct = p
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Produto não encontrado!")
			continue
		}

		// Solicitar quantidade
		fmt.Printf("Digite a quantidade para o produto %s: ", selectedProduct.Name)
		quantityInput, _ := reader.ReadString('\n')
		quantity, err := strconv.Atoi(strings.TrimSpace(quantityInput))
		if err != nil || quantity <= 0 {
			fmt.Println("Quantidade inválida!")
			continue
		}

		// Adicionar produto ao pedido
		for i := 0; i < quantity; i++ {
			selectedProducts = append(selectedProducts, selectedProduct)
			total += selectedProduct.Price
		}
		
		fmt.Println("Produto adicionado com sucesso!")
	}

	if len(selectedProducts) == 0 {
		fmt.Println("Nenhum produto selecionado. Saindo...")
		return
	}

	// Criar e processar pedido
	orderService := orders.NewOrderService()
	order := orderService.CreateOrder(user.ID, selectedProducts)
	
	fmt.Printf("\nResumo do Pedido:")
	for _, p := range selectedProducts {
		fmt.Printf("\n- %s: R$ %.2f", p.Name, p.Price)
	}
	fmt.Printf("\nTotal do Pedido: R$ %.2f\n", total)

	fmt.Println("\nProcessando pagamento...")
	orderService.ProcessOrder(order, total)
}