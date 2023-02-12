package guardado

import (
	"fmt"
	"os"

	"github.com/4ndroid3/random_movie_genre/objeto"
)

func GuardarTopicos(listaTopicos []string) {
	fmt.Println("Guardando datos en archivo .txt")
	filename := "topicos.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i, tema := range listaTopicos {
		_, err := file.WriteString(tema)
		if err != nil {
			panic(err)
		}

		if i < len(listaTopicos)-1 {
			_, err := file.WriteString("\n")
			if err != nil {
				panic(err)
			}
		}
	}
}

/*
Pregunta al usuario si desea ingresar datos,
luego permite ingresar el topico a guardar.
*/
func IngresoDeDatos(topicosLeidos []string) []string {
	i := "si"
	var ingreso string

	i = preguntaSiContinua()

	for {
		{
			if i == "no" || i == "No" || i == "NO" || i == "nO" {
				break
			}
			fmt.Println("ingresar dato: ")
			fmt.Scanf("%s", &ingreso)

			topicosLeidos = append(topicosLeidos, objeto.AgregaTopico(ingreso))

			i = preguntaSiContinua()
		}
	}

	return topicosLeidos
}

func preguntaSiContinua() string {
	var i string
	fmt.Println("Desea Ingresar otro dato mas?: ")
	fmt.Scanf("%s", &i)
	return i
}
