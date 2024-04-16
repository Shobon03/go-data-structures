package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LinkedList struct {
	value int
	next *LinkedList
}

var (
	lastNode *LinkedList
	node *LinkedList
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of items: ")

	text, _ := reader.ReadString('\n')
	
	text = strings.Trim(text, "\n")
	totalItems, _ := strconv.Atoi(text)
	
	for i := 0; i < totalItems; i++ {
		fmt.Print("Item: ")
		readItem, _ := reader.ReadString('\n')
		readItem = strings.Trim(readItem, "\n")

		item, _ := strconv.Atoi(readItem)
		
		if i == 0 {
			Insert(item, true)
		} else {
			Insert(item, false)
		}

		lastNode = node
	}

	fmt.Println("\nPrinting linked list:")
	Remove(node, 1)
	PrintLinkedList(node)
}

func Insert(value int, isFirst bool) {
	if isFirst {
		lastNode = nil
	}

	node = &LinkedList{
		value,
		lastNode,
	}

	lastNode = node
}

func Remove(list *LinkedList, value int) {
	prevList := list

	for list.next != nil {
		if list.value == value {
			prevList.next = list.next
		}

		list = list.next
	}
}

func PrintLinkedList(list *LinkedList) {
	count := 1 
	fmt.Printf("[ ")
	for list.next != nil {
		fmt.Printf("%d -> ", list.value)
		list = list.next
		count++
	}

	fmt.Printf("%d -> x", list.value)
	fmt.Printf(" ] <---> Total items: %d\n", count)
}
