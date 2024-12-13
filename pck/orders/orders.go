package orders

import (
	"fmt"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/catalog"
)

// Função que simula um pedido
func Order() {
	// Produtos disponíveis
	produtosDisponiveis := catalog.Catalog()

	// Pedido
	order := models.Order{}

	for {
		// Exibir todos os produtos disponíveis
		fmt.Println("\nProdutos disponíveis:")
		for _, produto := range produtosDisponiveis {
			fmt.Printf("ID: %d - %s - R$ %.2f - Estoque: %d\n", produto.ID, produto.Name, produto.Price, produto.Stock)
		}

		// Pedir para o usuário escolher um produto pelo ID
		var idEscolhido int
		fmt.Printf("\nDigite o ID do produto que deseja adicionar ao pedido: ")
		fmt.Scan(&idEscolhido)

		// Encontrar o produto pelo ID
		var produtoSelecionado *models.Product
		for i, produto := range produtosDisponiveis {
			if produto.ID == idEscolhido {
				produtoSelecionado = &produtosDisponiveis[i]
				break
			}
		}

		// Verificar se o produto foi encontrado
		if produtoSelecionado == nil {
			fmt.Println("ID inválido. Tente novamente.")
			continue
		}

		// Pedir a quantidade desejada
		var quantidade int
		fmt.Printf("Digite a quantidade para o produto %s (Estoque: %d): ", produtoSelecionado.Name, produtoSelecionado.Stock)
		fmt.Scan(&quantidade)

		// Verificar se a quantidade é válida
		if quantidade <= 0 || quantidade > produtoSelecionado.Stock {
			fmt.Println("Quantidade inválida. Tente novamente.")
			continue
		}

		// Adicionar o produto ao pedido
		order.Products = append(order.Products, models.Product{
			ID:    produtoSelecionado.ID,
			Name:  produtoSelecionado.Name,
			Price: produtoSelecionado.Price,
			Stock: quantidade,
		})

		// Atualizar o estoque localmente
		produtoSelecionado.Stock -= quantidade

		fmt.Println("Produto adicionado ao pedido com sucesso!")

		// Perguntar se deseja adicionar outro produto
		var resposta string
		fmt.Printf("Deseja adicionar outro produto? (sim/nao): ")
		fmt.Scan(&resposta)
		if resposta != "sim" {
			break
		}
	}

	// Mostrar o pedido final
	fmt.Println("\nResumo do Pedido:")

	var total float64
	
	for _, produto := range order.Products {
		fmt.Printf("Produto: %s - R$ %.2f - Quantidade: %d\n", produto.Name, produto.Price, produto.Stock)
		total += produto.Price * float64(produto.Stock)
	}
	fmt.Printf("\nTotal do Pedido: R$ %.2f\n", total)
}