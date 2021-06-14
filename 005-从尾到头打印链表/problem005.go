package main

import "fmt"

/*
这道题目的输入是一个链表，输出是一个数组。
链表具有的特点：
事先链表的长度是不知道的，也就是说事先你是不知道链表中存在多少节点的
链表只能从头节点开始往后遍历，而不能从非头节点往前遍历
数组具有的特点：
数组初始化的时候就必须指定长度
我们不仅可以从数组的第一个元素往后遍历访问数组，而且也可以从最后一个元素往前遍历数组
所以，我们要返回一个数组，首先需要计算出这个数组的长度，实际上数组的长度就等于链表的节点的个数，那么问题就变成了求链表的节点个数了，我们可以这样求解：

先初始化一个整形变量 length
从链表的头节点往后遍历，每遇到一个节点，那么 length 就自加 1，一直到链表的空节点为止
得到了数组的长度后，我们就可以初始化一个结果数组，然后再次从链表的头节点开始往后遍历，每遇到一个节点，那么就将节点的值放到数组的后端 (这里利用了数组可以从最后一个元素往前访问的特点)
*/


type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	if head == nil{
		return nil
	}
	count := 0
	newHead := head
	for head != nil{
		count++
		head = head.Next
	}
	res := make([]int,count)
	for i:=0;newHead != nil;i++{
		res[count-i-1] = newHead.Val
		newHead = newHead.Next
	}
	return res
}

//var res2 []int
func ReversePrint2(head *ListNode) []int {
	if head == nil{
		return nil
	}
	var res2 []int
	count := 0
	for {
		defer func(i,v int) {
			res2[i] = v
		}(count,head.Val)
		count++
		if head.Next == nil {
			res2 = make([]int,count)
			break
		}
		head = head.Next
	}
	return res2
}


func ReversePrint3(head *ListNode) (res []int) {
	if head == nil{
		return nil
	}
	//res = make([]int,0,3)
	for head != nil{
		defer func(v int) {
			res = append(res,v)
			//fmt.Println(res)
		}(head.Val)
		head = head.Next
	}
	return res
}

func main()  {
	var head *ListNode
	for i := 0; i < 10; i++ {
		temp := &ListNode{i,head}
		head = temp
	}

	//for head != nil {
	//	fmt.Println(head)
	//	head = head.Next
	//}

	//fmt.Println(ReversePrint3(head))
	fmt.Println(ReversePrint2(head))
	//fmt.Println(ReversePrint(head))
}