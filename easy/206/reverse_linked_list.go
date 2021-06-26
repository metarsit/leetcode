package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var end *ListNode

	for head != nil {
		next := head.Next
		head.Next = end

		end = head
		head = next
	}

	return end
}
