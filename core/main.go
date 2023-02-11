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
	fmt.Print("ingresar dato: ")
	fmt.Scanf("%s", &ingreso)
	fmt.Println(ingreso)
	objeto.GenerarTopicos(ingreso)
	random.GenerarObjetoRandom()
}
