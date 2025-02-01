package auth

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/eu-micaeu/Lista3-ArqSoft/models"
	"github.com/fatih/color" 
)

// AuthService gerencia a autenticação de usuários
type AuthService struct{}

// NewAuthService cria uma nova instância de AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// AuthenticateUser gerencia o processo de login
func (a *AuthService) AuthenticateUser(reader *bufio.Reader) *models.User {
	fmt.Println("Digite o email:")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Println("Digite a senha:")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	return a.Login(email, password)
}

// Login simula o login de um usuário
func (a *AuthService) Login(email, password string) *models.User {
	if email == "teste@teste.com" && password == "teste123" {
		// Log com cor verde para sucesso
		color.New(color.FgGreen).Println("Usuário autenticado com sucesso")
		return &models.User{ID: 1, Name: "Teste", Email: email, Password: password}
	}
	fmt.Println("Falha na autenticação")
	return nil
}
