// K 个一组翻转链表.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 分组.
    hair := &ListNode{Next : head}
    pre := hair
    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if (tail == nil) {
                return hair.Next
            }
        }
        next := tail.Next
        head, tail = reverse(pre, tail)
        pre.Next = head
        tail.Next = next
        pre = tail
        head = tail.Next
    }
    
    return hair.Next
}

func reverse(pre, tail *ListNode) (*ListNode, *ListNode) {
    head := pre.Next
    var tmp, next *ListNode
    nextGroup := tail.Next
    for head != nextGroup {
        next = head.Next
        head.Next = tmp
        tmp = head
        head = next
    }

    return tmp, pre.Next
}
