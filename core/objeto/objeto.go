package objeto

import "fmt"

func AgregaTopico(topico string) *Topico {
	newTopico := &Topico{Nombre: topico}
	fmt.Printf("Dato ingresado: %v \n \n", newTopico.Nombre)
	return newTopico
}

type Topico struct {
	Nombre string
	Tipo   string
}
