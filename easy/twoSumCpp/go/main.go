package main

import "fmt"

func main() {
	lista := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(lista, target))
	
}

func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {

		complement := target - num
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}
		seen[num] = i
	}
	return nil

}


/*
A variável index é atribuída dentro da expressão index, ok := seen[complement].

Aqui está como e quando isso acontece, passo a passo:

    Quando a atribuição ocorre: A atribuição de index acontece somente se a chave complement for encontrada no mapa seen.

    Qual valor é atribuído: O valor atribuído a index é o valor associado a essa chave no mapa.

Lembre-se de que o mapa seen armazena os números que já vimos (num) como chaves, e seus respectivos índices (i) como valores.

Portanto, quando o código executa index, ok := seen[complement]:

    Se o complement (o número que precisamos) for encontrado como uma chave no mapa, o valor que está guardado nessa chave (que é o índice daquele número) é atribuído à variável index.

    O ok é a flag booleana que nos confirma que essa atribuição realmente aconteceu porque a chave existia.

Exemplo Prático

Vamos simular o código com nums = [2, 7, 11, 15] e target = 9.

    1ª Iteração (i=0, num=2):

        complement = 9 - 2 = 7.

        seen[7] é buscado. O mapa está vazio.

        ok é false. O if não é executado.

        O código adiciona seen[2] = 0 ao mapa.

    2ª Iteração (i=1, num=7):

        complement = 9 - 7 = 2.

        seen[2] é buscado. A chave 2 é encontrada!

        O valor associado à chave 2 é 0.

        O Go atribui index = 0 e ok = true.

        Como ok é true, o if é executado.

        A função retorna []int{index, i}, que se traduz para []int{0, 1}.

Neste exemplo, a variável index foi atribuída com o valor 0 na segunda iteração, pois esse era o valor (índice) que estava armazenado no mapa para a chave 2.
*/