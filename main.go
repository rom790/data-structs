package main

import (
	"fmt"
	// "reflect"
)

type Node struct {
	data int
	next *Node
}
type List struct {
	head *Node
}

func (list *List) push(data int) {

	new_node := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = new_node
		return
	}

	current_el := list.head

	for current_el.next != nil {
		current_el = current_el.next
	}
	current_el.next = new_node
}

func (list *List) get_len() int {
	if list.head == nil {
		return 0
	}

	current_el := list.head

	amount := 1

	for current_el.next != nil {
		current_el = current_el.next
		amount++
	}

	return amount
}

func (list *List) print_elem(index int) {
	if index < 0 || index >= list.get_len() {
		fmt.Println("incorrect index")
		return
	}
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	current_el := list.head

	for i := 0; i < index; i++ {
		current_el = current_el.next
	}

	fmt.Println(current_el.data)

}

func (list *List) pop() {
	if list.head == nil {
		fmt.Println("list is already empty")
		return
	}

	current_el := list.head

	if current_el.next == nil {
		list.head = nil
		return
	}

	for current_el.next.next != nil {
		current_el = current_el.next
	}

	current_el.next = nil
}

func (list *List) shift() {
	if list.head == nil {
		fmt.Println("list is already empty")
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	list.head = list.head.next
}

func (list *List) slice(position, count int) {

	len := list.get_len()

	if list.head == nil {
		fmt.Println("list is already empty")
		return
	}

	if position < 0 || position >= len {
		fmt.Println("incorrect index")
		return
	}

	if count > len {
		fmt.Println("incorrect count")
		return
	}

	current_el := list.head

	for i := 0; i < position; i++ {
		current_el = current_el.next
	}

	end_el := current_el

	for i := 0; i < count; i++ {
		end_el = end_el.next
	}

	if position == 0 {
		list.head = end_el
	} else {
		current_el.next = end_el
	}

}

func (list *List) input(index int, data ...int) {

	if index > list.get_len() || index < 0 {
		fmt.Println("index is incorrect")
		return
	}

	current_el := list.head

	for i := 0; i < index-1; i++ {
		current_el = current_el.next
	}

	close_el := current_el.next
	if index == 0 {
		close_el = current_el

	}

	for i := 0; i < len(data); i++ {
		if index == 0 && i == 0 {
			list.head = &Node{data: data[i], next: nil}
			current_el = list.head
			continue
		}
		current_el.next = &Node{data: data[i], next: nil}
		current_el = current_el.next
	}

	current_el.next = close_el
}

func (list *List) print_list() {
	if list.get_len() == 0 {
		fmt.Println("list is empty")
		return
	}

	current_el := list.head

	for current_el != nil {
		fmt.Println(current_el.data)
		current_el = current_el.next
	}

}

func main() {

	linked_list := List{}

	linked_list.push(1432)
	linked_list.push(433)
	linked_list.push(1)
	linked_list.push(632)
	linked_list.push(9889)
	linked_list.push(8448484)
	linked_list.push(100)
	// linked_list.push(100)
	// linked_list.push(100)

	// fmt.Println(linked_list.get_len())

	// linked_list.slice(9, 4)
	linked_list.input(7, 2, 4, 6, 8)
	linked_list.print_list()
	// fmt.Println(linked_list.get_len())

	// linked_list.print_elem(1)

}
