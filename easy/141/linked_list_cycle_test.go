package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	// Construct cycle linked list
	h4 := ListNode{-4, nil}
	h3 := ListNode{0, &h4}
	h2 := ListNode{2, &h3}
	h1 := ListNode{3, &h2}
	h4.Next = &h2

	cases := []struct {
		head   *ListNode
		expect bool
	}{
		{&ListNode{1, nil}, false},
		{&h1, true},
	}

	for _, data := range cases {
		assert.Equal(t, hasCycle(data.head), data.expect)
	}
}
