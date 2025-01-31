package catalog

import (
	"fmt"
	"github.com/eu-micaeu/Lista3-ArqSoft/models"
)

// CatalogService gerencia o catálogo de produtos
type CatalogService struct {
	Products []models.Product
}

// NewCatalogService cria uma nova instância de CatalogService
func NewCatalogService() *CatalogService {
	return &CatalogService{
		Products: []models.Product{
			{ID: 1, Name: "Produto 1", Price: 10.00},
			{ID: 2, Name: "Produto 2", Price: 20.00},
			{ID: 3, Name: "Produto 3", Price: 30.00},
		},
	}
}

// GetAvailableProducts retorna a lista de produtos disponíveis
func (c *CatalogService) GetAvailableProducts() []models.Product {
	fmt.Println("Produtos disponíveis:")
	for _, product := range c.Products {
		fmt.Printf("ID: %d - %s - R$ %.2f\n", product.ID, product.Name, product.Price)
	}
	return c.Products
}

