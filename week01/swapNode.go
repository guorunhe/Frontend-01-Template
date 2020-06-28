// 交互链表节点.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    firstNode := head
    secondNode := head.Next
    firstNode.Next = swapPairs(secondNode.Next)
    secondNode.Next = firstNode
    
    return secondNode

