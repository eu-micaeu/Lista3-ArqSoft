package auth

import (

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"fmt"

)

// Função que simula uma autenticação
func Auth() bool {
	
	user := models.User{}

	fmt.Println("\nDigite o email: ")

	fmt.Scan(&user.Email)

	fmt.Println("\nDigite a senha: ")

	fmt.Scan(&user.Password)

	if(user.Email == "teste@teste.com" && user.Password == "teste123") {

		fmt.Println("\nUsuário autenticado")

		return true

	} else {

		fmt.Println("\nUsuário ou senha inválidos")

		return false

	}

}