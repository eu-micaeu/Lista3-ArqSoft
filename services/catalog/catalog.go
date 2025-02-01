package catalog

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/fatih/color" 
)

// CatalogService gerencia o catálogo de produtos
type CatalogService struct {
	Products []models.Product
}

// NewCatalogService cria uma nova instância de CatalogService
func NewCatalogService() *CatalogService {
	return &CatalogService{
		Products: []models.Product{
			{ID: 1, Name: "Produto 1", Price: 10.00, Stock: 10},
			{ID: 2, Name: "Produto 2", Price: 20.00, Stock: 20},
			{ID: 3, Name: "Produto 3", Price: 30.00, Stock: 30},
		},
	}
}

// DisplayProducts exibe e retorna a lista de produtos disponíveis
func (c *CatalogService) DisplayProducts() []models.Product {
	fmt.Println("Produtos disponíveis:")
	for _, product := range c.Products {
		fmt.Printf("ID: %d - %s - R$ %.2f - Estoque: %d\n", product.ID, product.Name, product.Price, product.Stock)
	}
	return c.Products
}

// FindProductByID busca um produto pelo ID na lista de produtos
func (c *CatalogService) FindProductByID(products []models.Product, productID int) (models.Product, bool) {
	for _, product := range products {
		if product.ID == productID {
			return product, true
		}
	}
	return models.Product{}, false
}

// CheckStock verifica se o estoque é suficiente para a quantidade solicitada
func (c *CatalogService) CheckStock(productID int, quantity int) error {
	for _, product := range c.Products {
		if product.ID == productID {
			if product.Stock >= quantity {
				return nil
			}
			return errors.New("estoque insuficiente")
		}
	}
	return errors.New("produto não encontrado")
}

// CalculateTotal calcula o valor total dos produtos selecionados
func (c *CatalogService) CalculateTotal(selectedProducts map[int]models.Product) float64 {
	total := 0.0
	for _, product := range selectedProducts {
		total += product.Price * float64(product.Stock)
	}
	return total
}

// SelectProducts permite a seleção de produtos e retorna os produtos selecionados e o total
func (c *CatalogService) SelectProducts(reader *bufio.Reader, products []models.Product) (map[int]models.Product, float64) {
	selectedProducts := make(map[int]models.Product)
	total := 0.0

	for {
		fmt.Printf("\nDigite o ID do produto que deseja adicionar (ou 0 para finalizar): ")
		productIDInput, _ := reader.ReadString('\n')
		productID, err := strconv.Atoi(strings.TrimSpace(productIDInput))
		if err != nil || productID < 0 {
			fmt.Println("ID inválido!")
			continue
		}

		if productID == 0 {
			break
		}

		selectedProduct, found := c.FindProductByID(products, productID)
		if !found {
			fmt.Println("Produto não encontrado!")
			continue
		}

		fmt.Printf("Digite a quantidade para o produto %s: ", selectedProduct.Name)
		quantityInput, _ := reader.ReadString('\n')
		quantity, err := strconv.Atoi(strings.TrimSpace(quantityInput))
		if err != nil || quantity <= 0 {
			fmt.Println("Quantidade inválida!")
			continue
		}

		if existingProduct, exists := selectedProducts[productID]; exists {
			quantity += existingProduct.Stock
		}

		err = c.CheckStock(selectedProduct.ID, quantity)
		if err != nil {
			fmt.Println("Erro ao adicionar produto:", err)
			continue
		}

		selectedProduct.Stock = quantity
		selectedProducts[productID] = selectedProduct
		total = c.CalculateTotal(selectedProducts)

		// Log de sucesso para produto adicionado com cor verde
		color.New(color.FgGreen).Println("Produto adicionado com sucesso!")
		
	}

	return selectedProducts, total
}
