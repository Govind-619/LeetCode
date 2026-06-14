/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {

    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode = nil

    for slow != nil {
        nextNode := slow.Next
        slow.Next = prev
        prev = slow
        slow = nextNode
    }

    ans := 0
    first := head
    second := prev

    for second != nil {
        sum := first.Val + second.Val

        if sum > ans {
            ans = sum
        }

        first = first.Next
        second = second.Next
    }

    return ans
}