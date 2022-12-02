package partitionlist

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

type ListNode = structures.ListNode
type ListNodeCase = structures.ListNodeCase2

func partition(head *ListNode, x int) *ListNode {
	l1 := &ListNode{}
	l2 := &ListNode{}
	p1 := l1
	p2 := l2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next

		} else {
			p2.Next = head
			p2 = p2.Next

		}

		head = head.Next

	}

	p2.Next = nil
	p1.Next = l2.Next
	return l1.Next

}

func TestPartition(t *testing.T) {
	qs := []ListNodeCase{
		{
			Head: []int{1, 4, 3, 2, 5, 2},
			X:    3,
			Res:  []int{1, 2, 2, 4, 3, 5},
		},
		{
			Head: []int{2, 1},
			X:    2,
			Res:  []int{1, 2},
		},
	}

	for i, q := range qs {
		headInput := structures.NewListNode(q.Head)
		expected := structures.NewListNode(q.Res)
		if got := partition(headInput, q.X); !reflect.DeepEqual(got, expected) {
			log.Printf("[ERROR]  case %v failed, expected: %v, got: %v\n", i, expected, got)
		}

	}
}
