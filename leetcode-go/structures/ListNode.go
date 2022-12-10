package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}

	return l.Next
}

type ListNodeCase struct {
	List1, List2, Res []int
}

type ListNodeCase2 struct {
	Head, Res []int
	X         int
}

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}
