package main

import (
	"fmt"

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
	var topicos []string
	i := 0
	for i < 2 {
		fmt.Println("ingresar dato: ")
		fmt.Scanf("%s", &ingreso)
		topicos = append(topicos, ingreso)
		i++
	}
	objeto.GenerarTopicos(topicos)
	random.GenerarObjetoRandom()
}
