package algorithm

import (
	"strconv"
)

/**
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，
 * 并且它们的每个节点只能存储 一位 数字。
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/add-two-numbers
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// InitListNode 初始化链表
func InitListNode(n int) *ListNode {
	str := strconv.Itoa(n)
	var lastNode *ListNode
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			panic(err)
		}
		currentNode := ListNode{}
		currentNode.Val = num

		if lastNode != nil {
			currentNode.Next = lastNode
		}

		lastNode = &currentNode
	}
	return lastNode

}

// GetIntFromListNode 把链表转化为整数
func GetIntFromListNode(l *ListNode) int {
	m := 1
	r := 0
	for {
		if l == nil {
			goto Loop
		}

		r += l.Val * m
		m = m * 10

		l = l.Next
	}
Loop:
	return r
}

// AddTwoNumbers 两个数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	node := &ListNode{}
	begin := node
	for {
		if flag == 0 && l1 == nil && l2 == nil {
			goto Loop
		}

		v := 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		v += flag

		flag = v / 10
		v %= 10

		node.Val = v
		if l1 != nil || l2 != nil || flag != 0 {
			node.Next = &ListNode{}
			node = node.Next
		}

	}

Loop:
	return begin
}
