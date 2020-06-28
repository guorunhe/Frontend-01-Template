// 环形链表2.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            break
        }
    }
    if fast == nil || fast.Next == nil { 
        return nil
    }
    for head != fast {
        head = head.Next
        fast = fast.Next
    }
    return head

