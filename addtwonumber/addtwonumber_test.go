package addtwonumber

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	f := getValListNode(l1)
	s := getValListNode(l2)

	fi, _ := strconv.Atoi(f)
	si, _ := strconv.Atoi(s)

	as := strconv.Itoa(fi + si)
	al := convertStringToListNode(as)

	printList(al)
	return al
}

func getValListNode(l *ListNode) string {
	if l.Next != nil {
		return getValListNode(l.Next) + strconv.Itoa(l.Val)
	} else {
		return strconv.Itoa(l.Val)
	}
}

func convertStringToListNode(a string) *ListNode {
	r := ListNode{}
	n := &r
	for i := len(a) - 1; i >= 0; i-- {
		e := ListNode{}
		e.Val, _ = strconv.Atoi(string(a[i]))
		n.Next = &e
		n = n.Next
	}
	if r.Next != nil {
		r = *r.Next
	}
	return &r
}

func printList(l *ListNode) {
	fmt.Print(l.Val)
	if l.Next != nil {
		printList(l.Next)
	}
}

// func convertStringToListNode(a string) *ListNode {
// 	r := ListNode{}
// 	for i := len(a)-1; i > 0; i-- {
// 		// e := ListNode{}
// 		r.Val, _ = strconv.Atoi(string(a[i]))
// 	}
// }
// func getVal(*ListNode)

func TestAddTwoNumbers(t *testing.T) {
	f1 := ListNode{Val: 3}
	f2 := ListNode{Val: 4, Next: &f1}
	f3 := ListNode{Val: 2, Next: &f2}

	s1 := ListNode{Val: 4}
	s2 := ListNode{Val: 6, Next: &s1}
	s3 := ListNode{Val: 5, Next: &s2}
	addTwoNumbers(&f3, &s3)
}
