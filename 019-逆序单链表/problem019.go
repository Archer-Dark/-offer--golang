package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

//loop
func RevLoop(head *Node) *Node {
	var l,r *Node
	for head != nil {
		r = head.Next
		head.Next = l
		l = head
		head = r
	}
	return l
}

//递归
func RevRecur(l,head,r *Node) *Node {
	if head == nil {
		return l
	}
	r = head.Next
	head.Next = l
	l = head
	head = r
	return RevRecur(l,head,r)

}

//defer
func RevDefer(head *Node) *Node {
	var l *Node
	for head != nil {
		defer func(l,head *Node) {
			head.Next = l
		}(l,head)
		l = head
		head = head.Next
	}
	return l
}

func main() {
	head := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}
	rhead := RevLoop(head)
	for rhead != nil {
		fmt.Println(rhead.Data)
		rhead = rhead.Next
	}

	fmt.Println("**********************************")

	head = &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}
	rrhead := RevRecur(nil,head,nil)
	for rrhead != nil {
		fmt.Println(rrhead.Data)
		rrhead = rrhead.Next
	}

	fmt.Println("**********************************")

	head = &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}
	dfhead := RevDefer(head)
	for dfhead != nil {
		fmt.Println(dfhead.Data)
		dfhead = dfhead.Next
	}
}