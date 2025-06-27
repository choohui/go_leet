package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumLL, tempLL *ListNode // Define as pointers
	rem := 0

	for l1 != nil || l2 != nil || rem != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if rem != 0 {
			sum += rem
		}

		rem = sum / 10
		nn := &ListNode{Val: sum % 10, Next: nil}
		if sumLL == nil {
			sumLL = nn
			tempLL = nn
		} else {
			tempLL.Next = nn
			tempLL = tempLL.Next
		}

	}

	return sumLL
}
