package AddTwoNumbers

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	v := addTwoNumbers(&l1, &l2)

	out, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result: " + string(out))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var r *ListNode
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = 0
		if sum >= 10 {
			carry = sum / 10
			sum = sum - (carry * 10)
		}

		temp := &ListNode{Val: sum}
		if r == nil {
			r = temp
			result = temp
		} else {
			r.Next = temp
			r = temp
		}
	}
	return result
}
