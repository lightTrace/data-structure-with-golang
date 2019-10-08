package main

import "fmt"

type Node struct {
	Data  interface{}
	Next  *Node
}

type List struct {
	HeadNode *Node
}

//判断链表是否为nil
func(list *List) isEmpty() bool{
	if list.HeadNode == nil{
		return true
	}else {
		return false
	}
}

//计算链表的长度
func(list *List) length() int{
	cur := list.HeadNode
	count := 0
	for cur != nil{
		count ++
		cur = cur.Next
	}
	return count
}
//头部插入数据
func(list *List) add(e interface{}){
	node := &Node{Data:e}
	node.Next = list.HeadNode
	list.HeadNode = node
}
//从尾部插入数据
func(list *List) append(e interface{}){
	node := &Node{Data:e}
	if list ==  nil{
		list.HeadNode = node
	}else {
		cur := list.HeadNode
		for  cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}

}
//在指定位置插入元素
func(list *List) insert(e interface{},i int){
	if i <0 {
		list.add(e)
	}else if i > list.length(){
		list.append(e)
	}else{
		pre := list.HeadNode
		count := 0
		for count < i-1{
			pre = pre.Next
			count ++
		}
		node := &Node{Data:e}
		node.Next = pre.Next
		pre.Next = node
	}
}

func main() {
	var list List
	fmt.Println(list.isEmpty())
	list.add(1)
	list.add(2)
	fmt.Println(list.isEmpty())
	fmt.Println(list.length())
	cur := list.HeadNode
	for cur != nil{
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}