package objeto

import "fmt"

func GenerarTopicos(ingresado []string) []string {
	fmt.Printf("Dato ingresado: %v \n", ingresado)

	return ingresado
}

func AgregaTopico(topico string) string {
	fmt.Printf("Dato ingresado: %v \n \n", topico)
	return topico
}
