package payments

import (
	"fmt"
	"time"

	//"github.com/eu-micaeu/Lista3-ArqSoft/models"
)

// Função que simula um pagamento
func Payment() {

	fmt.Println("Pagamento sendo concluido...")

	// Timer de 3 segundos

	time.Sleep(3 * time.Second)

	fmt.Println("Pagamento concluido com sucesso!")

}