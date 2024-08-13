package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gerar [cpf|cnpj]")
		return
	}

	switch os.Args[1] {
	case "cpf":
		cpf := gerarCPF()
		fmt.Printf("CPF gerado: %s\n", cpf)
		err := clipboard.WriteAll(strings.ReplaceAll(strings.ReplaceAll(cpf, ".", ""), "-", ""))
		if err != nil {
			fmt.Println("Erro ao copiar para o clipboard:", err)
		} else {
			fmt.Println("CPF copiado para o clipboard.")
		}
	case "cnpj":
		cnpj := gerarCNPJ()
		fmt.Printf("CNPJ gerado: %s\n", cnpj)
		err := clipboard.WriteAll(strings.ReplaceAll(strings.ReplaceAll(cnpj, ".", ""), "/", ""))
		if err != nil {
			fmt.Println("Erro ao copiar para o clipboard:", err)
		} else {
			fmt.Println("CNPJ copiado para o clipboard.")
		}
	default:
		fmt.Println("Invalid option. Use 'cpf' or 'cnpj'.")
	}
}

func gerarCPF() string {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 9)
	for i := 0; i < 9; i++ {
		numbers[i] = rand.Intn(10)
	}

	numbers = append(numbers, calcDVCPF(numbers[:9]))
	numbers = append(numbers, calcDVCPF(numbers))

	return fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d", toInterfaceSlice(numbers)...)
}

func gerarCNPJ() string {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 12)
	for i := 0; i < 8; i++ {
		numbers[i] = rand.Intn(10)
	}
	numbers[8] = 0
	numbers[9] = 0
	numbers[10] = 0
	numbers[11] = 1

	numbers = append(numbers, calcDVCNPJ(numbers[:12]))
	numbers = append(numbers, calcDVCNPJ(numbers))

	return fmt.Sprintf("%d%d.%d%d%d.%d%d%d/%d%d%d%d-%d%d", toInterfaceSlice(numbers)...)
}

func calcDVCPF(numbers []int) int {
	weights := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	if len(numbers) == 10 {
		weights = append([]int{11}, weights...)
	}
	sum := 0
	for i, v := range numbers {
		sum += v * weights[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func calcDVCNPJ(numbers []int) int {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	if len(numbers) == 13 {
		weights = append([]int{6}, weights...)
	}
	sum := 0
	for i, v := range numbers {
		sum += v * weights[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func toInterfaceSlice(intSlice []int) []interface{} {
	interfaceSlice := make([]interface{}, len(intSlice))
	for i, d := range intSlice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
