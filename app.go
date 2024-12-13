package main

import (

	"github.com/eu-micaeu/Lista3-ArqSoft/pck/auth"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/catalog"
	"github.com/eu-micaeu/Lista3-ArqSoft/pck/orders"

)

func main() {

	auth.Auth()

	catalog.Catalog()

	orders.Order()
	
}