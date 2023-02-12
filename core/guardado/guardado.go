package guardado

import (
	"fmt"
	"os"
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
