package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var end *ListNode
	m := make(map[*ListNode]int)

	for head != nil {
		next := head.Next
		head.Next = end
		end = head

		if _, ok := m[end]; ok {
			return true
		} else {
			m[end] = end.Val
		}

		head = next
	}

	return false
}
