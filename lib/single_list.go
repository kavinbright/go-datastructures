package lib

import "fmt"

// SingleList 单项链表
// 结构：1. 头部节点； 2. 链表长度
/* 操作：
* 1. 初始化一个链表
* 2. 追加节点
* 3. 在指定位置插入节点
* 4. 删除节点
* 5. 获取制定节点
* 6. 遍历节点信息
* 7. 获取链表长度
 */

// Object 节点数据
type Object interface{}

// Node 节点信息
type Node struct {
	data Object
	next *Node
}

// SingleList 链表结构
type SingleList struct {
	size int
	head *Node
}

// CreateSingleList 创建
func CreateSingleList() *SingleList {
	var sl *SingleList = &SingleList{
		size: 0,
		head: nil,
	}
	return sl
}

// Length 长度
func (s *SingleList) Length() int {
	return s.size
}

// Append 追加
func (s *SingleList) Append(data Object) {
	node := &Node{
		data: data,
		next: nil,
	}
	if s.size == 0 {
		s.head = node
		s.size = 1
		return
	}

	var tailNode *Node = s.head
	for i := 0; i < s.size-1; i++ {
		tailNode = tailNode.next
	}

	tailNode.next = node
	s.size++
}

// Head 获取头部节点
func (s *SingleList) Head() *Node {
	return s.head
}

// Node 获取指定节点， 从0开始
func (s *SingleList) Node(i int) *Node {
	if s.size == 0 {
		return nil
	}

	if i > s.size-1 {
		return nil
	}

	var targetNode *Node = s.head
	for j := 1; j <= i; j++ {
		targetNode = targetNode.next
	}

	return targetNode
}

// Insert 在指定位置插入节点,i 从0开始
func (s *SingleList) Insert(data Object, i int) {
	node := &Node{
		data: data,
		next: nil,
	}
	if s.size == 0 {
		s.head = node
		s.size = 1
		return
	}

	if i > s.size-1 {
		s.Append(data)
		return
	}

	if i == 0 {
		nextNode := s.head
		s.head = node
		s.head.next = nextNode
		s.size++
		return
	}

	var targetNode *Node
	for j := 0; j < i-1; j++ {
		if j == 0 {
			targetNode = s.head
		}

		targetNode = targetNode.next
	}

	// 记录好之前nextnode
	nextNode := targetNode.next
	targetNode.next = node
	targetNode.next.next = nextNode
	s.size++
}

// Range 遍历
func (s *SingleList) Range() {
	var node *Node
	for i := 0; i < s.size; i++ {
		if i == 0 {
			node = s.head
		} else {
			node = node.next
		}
		fmt.Printf("------->s-%d: %+v \n", i, node.data)
	}
}

// Delete 删除
func (s *SingleList) Delete(i int) {
	if s.size == 0 {
		return
	}

	if i == 0 {
		thisNode := s.head
		s.head = thisNode.next
		s.size = s.size - 1
		return
	}

	if i > s.size-1 {
		return
	}

	var node *Node
	for j := 0; j < i; j++ {
		if j == 0 {
			node = s.head
		} else {
			node = node.next
		}
	}
	preNode := node
	thisNode := node.next
	nextNode := thisNode.next
	preNode.next = nextNode
	s.size = s.size - 1
}
