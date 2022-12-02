package mergetwosortedlists

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

type ListNode = structures.ListNode
type ListNodeCase = structures.ListNodeCase

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	l := &ListNode{}
	p := l
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}

		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return l.Next
}

func TestMergeTwoLists(t *testing.T) {
	qs := []ListNodeCase{
		{
			List1: []int{1, 2, 4},
			List2: []int{1, 3, 4},
			Res:   []int{1, 1, 2, 3, 4, 4},
		},
		{
			List1: []int{},
			List2: []int{},
		},
		{
			List1: []int{},
			List2: []int{0},
			Res:   []int{0},
		},
	}

	for i, q := range qs {
		l1 := structures.NewListNode(q.List1)
		l2 := structures.NewListNode(q.List2)
		expected := structures.NewListNode(q.Res)
		if got := mergeTwoLists2(l1, l2); !reflect.DeepEqual(got, expected) {
			log.Printf("[ERROR]  case %v failed, expected: %v, got: %v\n", i, expected, got)
		}

	}
}
