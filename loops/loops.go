package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando ", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Pedro", "Maria"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	usuario := map[string]string{
		"nome":      "Fabio",
		"sobrenome": "sarmento",
	}

	for index, value := range usuario {
		fmt.Println(index, value)
	}

	// for {
	// 	fmt.Println("Executa eternamente.")
	// 	time.Sleep(time.Second)
	// }
}
