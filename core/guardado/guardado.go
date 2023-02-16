package guardado

import (
	"fmt"
	"os"

	"github.com/4ndroid3/random_movie_genre/objeto"
)

/*
Lee el archivo topicos.txt y guarda los datos en un slice de strings.
*/
func GuardarTopicos(listaTopicos []string) {
	filename := "topicos.txt"        // especificamos el nombre del archivo
	file, err := os.Create(filename) // creamos el archivo
	if err != nil {
		panic(err) // si ocurre un error al crear el archivo, lo mostramos y salimos de la función
	}
	defer file.Close() // aseguramos que se cierre el archivo al final de la función, incluso si ocurre un error

	for i, tema := range listaTopicos { // recorremos la lista de temas
		_, err := fmt.Fprint(file, tema) // escribimos el tema en el archivo
		if err != nil {
			panic(err) // si ocurre un error al escribir en el archivo, lo mostramos y salimos de la función
		}

		if i < len(listaTopicos)-1 { // si no es el último tema, escribimos una nueva línea para separarlo del siguiente tema
			_, err := fmt.Fprintln(file) // escribimos una nueva línea en el archivo
			if err != nil {
				panic(err) // si ocurre un error al escribir en el archivo, lo mostramos y salimos de la función
			}
		}
	}
}

/*
Pregunta al usuario si desea ingresar datos,
luego permite ingresar el topico a guardar.
*/
func IngresoDeDatos(topicosLeidos []string) []string {
	for {
		respuesta := preguntaSiContinua()
		if respuesta == "no" || respuesta == "n" {
			break
		}
		fmt.Print("Ingresar dato: ")
		var ingreso string
		fmt.Scanln(&ingreso)
		topicosLeidos = append(topicosLeidos, objeto.AgregaTopico(ingreso).Nombre)
	}
	return topicosLeidos
}

func preguntaSiContinua() string {
	fmt.Print("¿Desea ingresar otro dato? (si/no): ")
	var respuesta string
	fmt.Scanln(&respuesta)
	return respuesta
}
