/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    m := make(map[int]bool)
    for _, v := range nums {
        m[v] = true
    }

    dummy := &ListNode{}   // use pointer, not value
    prev := dummy
    current := head

    for current != nil {
        if !m[current.Val] {        // keep nodes not in nums
            prev.Next = current
            prev = current
        }
        current = current.Next
    }

    prev.Next = nil                 // cut off tail if last node was removed
    return dummy.Next
}
