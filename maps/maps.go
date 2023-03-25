package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Fabio",
		"sobrenome": "Sarmento",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Fabio",
			"ultimo":   "Sarmento",
		},
		"curso": {"nome": "engenharia", "campus": "Campus 1"},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Peixes",
	}
	fmt.Println(usuario2)
}
