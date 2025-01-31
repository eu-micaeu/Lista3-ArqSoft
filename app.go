package main

import (

	"github.com/eu-micaeu/Lista3-ArqSoft/pck/auth"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/catalog"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/orders"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/payments"
)

func main() {

	valid := auth.Auth()

	for !valid {

		valid = auth.Auth()

	} 
		
	produtos := catalog.Catalog()
	
	orders.Order(produtos)

	payments.Payment()

}