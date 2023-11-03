package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	current := &dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// После того как мы достигли конца одного из списков, просто присоединяем оставшийся список к current.
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Возвращаем список, начиная с узла следующего за dummy, так как dummy был искусственно добавлен.
	return dummy.Next
}

func mergeTwoListsRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Выбираем меньший узел и рекурсивно сливаем оставшиеся элементы
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {
	list11 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list21 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	fmt.Println(mergeTwoLists(&list11, &list21)) // [1,1,2,3,4,4]

	list12 := ListNode{}
	list22 := ListNode{}

	fmt.Println(mergeTwoLists(&list12, &list22)) // []

	list13 := ListNode{}
	list23 := ListNode{Val: 0}

	fmt.Println(mergeTwoLists(&list13, &list23)) // [0]
}
