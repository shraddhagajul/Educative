package main

import "fmt"

type Node struct{
	Val int
	NextNode *Node
}

var node Node

func lengthOfLinkedlist(node *Node,  head int) int{
	
	if node.NextNode == nil{
		head++
		return head
	}
	
	head++
	return lengthOfLinkedlist(node.NextNode,head)

}


	func main(){
		node := &Node{5,nil}
		node.NextNode = &Node{6,nil}
		node.NextNode.NextNode = &Node{7,nil}

		
	
		result := lengthOfLinkedlist(node, 0)	
		
		fmt.Println("Result :",result)
	}


