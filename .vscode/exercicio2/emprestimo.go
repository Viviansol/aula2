package main

import "fmt"

func main() {
	idade := 19
	emprego := true
	tempoTrabalho := 3
	salario := 100000

	if idade >= 22 && emprego == true && tempoTrabalho >= 1 {

		if salario >= 100000 {
			fmt.Println("Você  pode contratar empréstimo e não precisa pagar juros!")
		} else {
			fmt.Println("Você pode contratar  um empréstimo")
		}
	} else {
		fmt.Println("Você não pode contratar um empréstimo")
	}
}
