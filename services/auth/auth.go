package auth

import (
	"fmt"
	"github.com/eu-micaeu/Lista3-ArqSoft/models"
)

// AuthService gerencia a autenticação de usuários
type AuthService struct{}

// NewAuthService cria uma nova instância de AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Login simula o login de um usuário
func (a *AuthService) Login(email, password string) *models.User {
	// Simulação de autenticação
	if email == "teste@teste.com" && password == "teste123" {
		fmt.Println("Usuário autenticado")
		return &models.User{ID: 1, Name: "Teste", Email: email, Password: password}
	}
	fmt.Println("Falha na autenticação")
	return nil
}