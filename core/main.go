package main

import (
	"fmt"

	"github.com/4ndroid3/random_movie_genre/guardado"
	"github.com/4ndroid3/random_movie_genre/lectura"
	"github.com/4ndroid3/random_movie_genre/random"
	"github.com/4ndroid3/random_movie_genre/users"
)

func main() {

	fmt.Println()
	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("Inicia Randomizador de Peliculas")
	fmt.Println("==--==--==--==--==--==--==--==--==")
	fmt.Println("==--==--==--==--==--==--==--==--== ")
	fmt.Println()

	topicos := guardado.IngresoDeDatos(lectura.LeerTopicosEnArchivo())
	guardado.GuardarTopicos(topicos)
	random.GenerarObjetoRandom(topicos)

	empleado := users.NewEmployee("Pompeyo")
	fmt.Println(empleado.Name, empleado.Id, empleado.Active, empleado.CalcTotal(1.25, 2.50, 3.75))
}
