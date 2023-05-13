package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular    string
	nr_agencia int
	nr_conta   int
	saldo      float64
}

func SemParametro() string {
	return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
	return texto
}

func DoisParametros(texto string, numero int) (string, int) {
	return texto, numero
}
func (c *ContaCorrente) Sacar()
func main() {
	fmt.Println(SemParametro())
	fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
	fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))
}
