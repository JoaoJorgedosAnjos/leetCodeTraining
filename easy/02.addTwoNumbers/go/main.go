package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Cria um nó dummy para simplificar a construção da lista resultado
    head := &ListNode{Val: 0}
    current := head
    carry := 0
    
    // Continua enquanto houver nós em l1, l2 ou carry
    for l1 != nil || l2 != nil || carry != 0 {
        // Obtém os valores dos nós atuais (0 se o nó for nil)
        v1 := 0
        if l1 != nil {
            v1 = l1.Val
        }
        
        v2 := 0
        if l2 != nil {
            v2 = l2.Val
        }
        
        // Calcula a soma total incluindo o carry
        sum := v1 + v2 + carry
        digit := sum % 10    // Dígito atual (resto da divisão por 10)
        carry = sum / 10     // Novo carry (divisão inteira por 10)
        
        // Cria novo nó com o dígito calculado
        current.Next = &ListNode{Val: digit}
        current = current.Next
        
        // Avança para os próximos nós (se existirem)
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    
    // Retorna o próximo do nó dummy (início da lista real)
    return head.Next
}

// Funções auxiliares para testar
func createLinkedList(arr []int) *ListNode {
    // Cria uma linked list a partir de um slice
    if len(arr) == 0 {
        return nil
    }
    
    head := &ListNode{Val: arr[0]}
    current := head
    
    for i := 1; i < len(arr); i++ {
        current.Next = &ListNode{Val: arr[i]}
        current = current.Next
    }
    
    return head
}

func linkedListToSlice(head *ListNode) []int {
    // Converte uma linked list para slice
    var result []int
    current := head
    
    for current != nil {
        result = append(result, current.Val)
        current = current.Next
    }
    
    return result
}

func printLinkedList(head *ListNode) {
    // Imprime a linked list de forma legível
    result := linkedListToSlice(head)
    fmt.Printf("%v\n", result)
}

func main() {
    // Teste 1: [2,4,3] + [5,6,4] = [7,0,8]
    fmt.Println("Teste 1:")
    l1 := createLinkedList([]int{2, 4, 3}) // representa 342
    l2 := createLinkedList([]int{5, 6, 4}) // representa 465
    result := addTwoNumbers(l1, l2)
    fmt.Print("Resultado: ")
    printLinkedList(result) // [7,0,8] = 807
    
    // Teste 2: [0] + [0] = [0]
    fmt.Println("\nTeste 2:")
    l1 = createLinkedList([]int{0})
    l2 = createLinkedList([]int{0})
    result = addTwoNumbers(l1, l2)
    fmt.Print("Resultado: ")
    printLinkedList(result) // [0]
    
    // Teste 3: [9,9,9,9,9,9,9] + [9,9,9,9] = [8,9,9,9,0,0,0,1]
    fmt.Println("\nTeste 3:")
    l1 = createLinkedList([]int{9,9,9,9,9,9,9}) // representa 9999999
    l2 = createLinkedList([]int{9,9,9,9})        // representa 9999
    result = addTwoNumbers(l1, l2)
    fmt.Print("Resultado: ")
    printLinkedList(result) // [8,9,9,9,0,0,0,1] = 10009998
    
    // Demonstração da lógica reversa
    fmt.Println("\nExplicação do Teste 1:")
    fmt.Println("l1 = [2,4,3] representa o número 342")
    fmt.Println("l2 = [5,6,4] representa o número 465") 
    fmt.Println("342 + 465 = 807")
    fmt.Println("807 em ordem reversa = [7,0,8]")
}