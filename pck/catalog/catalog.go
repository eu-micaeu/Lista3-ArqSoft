package catalog

import (

	"github.com/eu-micaeu/Lista3-ArqSoft/models"

)

// Função que simula um catálogo de produtos
func Catalog() []models.Product {

	products := []models.Product{
		{

			ID: 1,
			Name: "Produto 1",
			Price: 10.0,
			Stock: 10,

		},
		{

			ID: 2,
			Name: "Produto 2",
			Price: 20.0,
			Stock: 20,

		},
		{

			ID: 3,
			Name: "Produto 3",
			Price: 30.0,
			Stock: 30,

		},

	}

	return products

}