package main

import (
	"fmt"
	"codenation/models"
	"codenation/requests"
	"codenation/algoritmos"
)

var json models.Answer


func main() {
	get := requests.NewGet()
	algoritmo := algoritmos.NewAlgoritmo()

	json = get.Json()
	teste := algoritmo.DescriptografaJulioCesar(json)
	fmt.Println(teste)
}