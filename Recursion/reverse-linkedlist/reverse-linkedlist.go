package main

import "fmt"

type Node struct{
	Val int
	NextNode *Node
}

func reverseLinkedlist(node, current , previous *Node) *Node{
  if current == nil{
		return nil
	}

	if current.NextNode == nil{
		fmt.Println("data is", current, current.NextNode)
		current.NextNode = previous
		fmt.Println(current.NextNode)
		return current
		 
	}

	fmt.Println("before call", current, current.NextNode)
	res := reverseLinkedlist(current, current.NextNode, current)
	fmt.Println("result ",res)
	current.NextNode = previous
	return res
}

func main(){
	node := &Node{5,nil}
	node.NextNode = &Node{6,nil}
	node.NextNode.NextNode = &Node{7,nil}

	result := reverseLinkedlist(node, node, nil)	
	
	for result != nil {
		fmt.Println("bla\n", result)
		result = result.NextNode
	}
	
}