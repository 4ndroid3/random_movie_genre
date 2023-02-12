package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerarObjetoRandom(topicos []string) {
	// Cambiando el Seed del randomizador se obtienen numeros mas aleatorios.
	rand.Seed(time.Now().UnixNano())
	topicoRandom := topicos[rand.Intn(len(topicos))]
	fmt.Println("El topico random es: ", topicoRandom)
}
