package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	//cria um array baseado no tamanho do arrays passado (nao eh dinamico)
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice nao eh array
	//slice eh tipo um array dinamico
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao Alterada" //ele funciona igual um ponteiro
	fmt.Println(slice2)

	// arrays internos
	slice3 := make([]float32, 10, 15)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 20)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

}
