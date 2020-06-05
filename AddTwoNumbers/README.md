# Add Two Numbers
`medium`

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## Starting Point
```go
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //Your code here.
}
```

### Contributor
LeetCode

### Related Topics
`Linked List` `Math`

### Similar Questions

Question|Difficulty
----|----
Multiply Strings|`Medium`
Add Binary|`Easy`
Sum of Two Integers|`Easy`
Add Strings|`Easy`
Add Two Numbers II|`Medium`
Add to Array-Form of Integer|`Easy`