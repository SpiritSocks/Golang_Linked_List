package main

import "fmt"

type ListNode struct {
	link *ListNode
	data int
}

// Prints the list
func Print(head *ListNode) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Println(head.data)
		head = head.link
	}
}

// Adds a node to the beggining of the list
func AddNode(headPtr **ListNode, x int) {
	temp := &ListNode{nil, x}

	if *headPtr != nil {
		temp.link = *headPtr
	}
	*headPtr = temp
}

// Adds a node at the n-th position
func AddNodeAtNPos(head **ListNode, data int, npos int) {

	node := &ListNode{nil, data}

	if npos <= 1 {
		node.link = *head
		*head = node
		return
	}

	var cur *ListNode = *head

	for i := 0; i < npos-2; i++ {
		cur = cur.link
	}

	node.link = cur.link
	cur.link = node
}

// Deletes the first node in the list
func DeleteNode(head **ListNode) {
	if head == nil || *head == nil {
		return
	}

	*head = (*head).link
}

// Delete the node located at the n-th position
func DeleteAtNPos() {

}

// Finds the middle node of linked list
func FindMiddleOfList(head *ListNode) int {
	if head == nil {
		return -1
	}

	var fast *ListNode = head
	var slow *ListNode = head

	for fast != nil && fast.link != nil {
		slow = slow.link
		fast = fast.link.link
	}
	return slow.data
}

func main() {

	var head *ListNode

	AddNode(&head, 10)
	AddNode(&head, 11)
	AddNode(&head, 12)
	AddNodeAtNPos(&head, 20, 2)

	Print(head)

	fmt.Println()

	DeleteNode(&head)

	Print(head)

	fmt.Println("The middle node is: ", FindMiddleOfList(head))

}
