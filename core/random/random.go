package random

import (
	"fmt"
	"math/rand"
)

func GenerarObjetoRandom(topicos []string) {
	topicoRandom := rand.Intn(len(topicos))
	fmt.Printf("El Topico Elegido Aleatoriamente es: %v \n", topicos[topicoRandom])
}
