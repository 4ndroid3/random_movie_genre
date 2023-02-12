package main

import (
	"fmt"

	"github.com/4ndroid3/random_movie_genre/guardado"
	"github.com/4ndroid3/random_movie_genre/objeto"
	"github.com/4ndroid3/random_movie_genre/random"
)

func main() {

	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("Inicia Randomizador de Peliculas")
	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("==--==--==--==--==--==--==--==--==")

	var ingreso string
	// agregar una variable tipo slice con 5 nombres de topicos
	topicos := []string{"accion", "comedia", "drama", "terror", "romance"}

	i := "si"

	/*Realiza un bucle while do
	preguntando el valor a ingresa y luego
	pregunta si desea seguir agregando datos. */
	for {
		{
			if i == "no" || i == "No" || i == "NO" || i == "nO" {
				break
			}
			fmt.Println("ingresar dato: ")
			fmt.Scanf("%s", &ingreso)

			topicos = append(topicos, objeto.AgregaTopico(ingreso))

			fmt.Println("Desea Ingresar otro dato mas?: ")
			fmt.Scanf("%s", &i)
		}
	}

	// Funcion para guardar la lista de topicos en un archivo .txt

	objeto.GenerarTopicos(topicos)
	guardado.GuardarTopicos(topicos)
	random.GenerarObjetoRandom(topicos)
}
