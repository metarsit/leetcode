package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		head   *ListNode
		expect *ListNode
	}{
		{&ListNode{}, &ListNode{}},
		{&ListNode{1, &ListNode{2, nil}}, &ListNode{2, &ListNode{1, nil}}},
	}

	for _, data := range cases {
		assert.Equal(t, reverseList(data.head), data.expect)
	}
}
