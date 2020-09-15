func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := ListNode{}
	curPtr := &r
	rem := 0
	for l1 != nil || l2 != nil {
		c := rem
		if l1 != nil {
			c += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			c += l2.Val
			l2 = l2.Next
		}
		rem = 0
		if c > 9 {
			rem = 1
			c -= 10
		}
		curPtr.Next = &ListNode{Next: nil, Val: c}
		curPtr = curPtr.Next
	}

	if rem > 0 {
		curPtr.Next = &ListNode{Next: nil, Val: 1}
	}

	return r.Next
}

