package addtwonumber

import (
	"testing"
)

func addLargeTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a ListNode
	addEachNode(l1, l2, &a, 0)

	if a.Next != nil {
		a = *a.Next
	}

	return &a
}

func addEachNode(f *ListNode, s *ListNode, result *ListNode, carry int) {
	nf := &ListNode{}
	ns := &ListNode{}
	var con bool
	if f.Next != nil {
		nf = f.Next
		con = true
	}

	if s.Next != nil {
		ns = s.Next
		con = true
	}
	sum := f.Val + s.Val + carry
	nr := ListNode{
		Val: sum % 10,
	}

	carry = sum / 10

	// fmt.Println("====== ")
	// fmt.Println("First: ", f.Val)
	// fmt.Println("Second: ", s.Val)
	// fmt.Println("Carry: ", carry)
	// fmt.Println("Sum&Mod: ", nr.Val)
	// fmt.Println("====== ")
	result.Next = &nr

	if con {
		addEachNode(nf, ns, &nr, carry)
	} else if carry == 1 {
		t := &ListNode{
			Val: 1,
		}
		result.Next.Next = t
	}
}

func TestAddTwoNumbersLargNumber(t *testing.T) {
	//=============case 1=====================
	// f1 := ListNode{Val: 3}
	// f2 := ListNode{Val: 4, Next: &f1}
	// f3 := ListNode{Val: 2, Next: &f2}

	// s1 := ListNode{Val: 4}
	// s2 := ListNode{Val: 6, Next: &s1}
	// s3 := ListNode{Val: 5, Next: &s2}
	//=========================================

	//=============case 2=====================
	f1 := ListNode{Val: 9}
	f2 := ListNode{Val: 9, Next: &f1}
	f3 := ListNode{Val: 9, Next: &f2}
	f4 := ListNode{Val: 9, Next: &f3}
	f5 := ListNode{Val: 9, Next: &f4}
	f6 := ListNode{Val: 9, Next: &f5}
	f7 := ListNode{Val: 9, Next: &f6}

	s1 := ListNode{Val: 9}
	s2 := ListNode{Val: 9, Next: &s1}
	s3 := ListNode{Val: 9, Next: &s2}
	s4 := ListNode{Val: 9, Next: &s3}
	//=========================================

	a := addLargeTwoNumbers(&f7, &s4)
	printList(a)
}
