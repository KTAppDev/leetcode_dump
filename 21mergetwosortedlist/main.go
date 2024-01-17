package main

import (
	"fmt"
)

type ListNode struct { // a Node
	Next *ListNode
	Val  int
}

func main() {
	list1 := []int{1, 2, 4} // since these are slices i will need to single them out to linked lists
	list2 := []int{1, 3, 4}

	l1 := sliceToList(list1) // function to create ListNodes of the values
	l2 := sliceToList(list2)

	mergedList := mergeTwoLists(l1, l2)
	printList(mergedList)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // dummy helps to setup and basically keeps a link to the head dummy.next
	current := dummy     // current will go on the journey

	for l1 != nil && l2 != nil { // once the lists  still have values then keep looping
		if l1.Val < l2.Val {
			current.Next = l1 // Add the node to the list
			l1 = l1.Next      // move to the next node in l1
		} else { // else do the adding and moving to the next on on the other list
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next // move to the next node4 so we can assign to the next node
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next // dummy.next is the head
}

func sliceToList(slice []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range slice {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
