package lectura

import (
	"io/ioutil"
	"strings"
)

func LeerTopicosEnArchivo() []string {
	filename := "topicos.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var topicos []string
	topicosLeidos := strings.Split(string(data), "\n")
	for _, tema := range topicosLeidos {
		topicos = append(topicos, tema)
	}

	return topicos
}
