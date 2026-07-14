package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {

	numbers := [10]int64{}
	sortedNumber := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {

		fmt.Println("Digite um número:")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("Digite um número válido!")
			i--
			continue
		}
		numbers[i] = chuteInt

		switch {
		case chuteInt == sortedNumber:
			fmt.Println("Numero Correto!")
			break
		case chuteInt > sortedNumber:
			fmt.Println("Numero Maior que o sorteado!")
		case chuteInt < sortedNumber:
			fmt.Println("Numero Menor que o sorteado!")
		}

	}

	fmt.Println("Você esgotou as tentativas!")
	fmt.Println("Números digitados: ", numbers)
	fmt.Println("Número sorteado: ", sortedNumber)
}
