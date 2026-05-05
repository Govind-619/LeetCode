func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    // 1. Calculate length and find the actual tail node
    length := 1
    tail := head
    for tail.Next != nil {
        length++
        tail = tail.Next
    }

    // 2. Adjust k and check if rotation is even necessary
    k = k % length
    if k == 0 {
        return head
    }

    // 3. Connect tail to head to make it circular temporarily
    tail.Next = head

    // 4. Find the new tail: (length - k) steps from the start
    stepsToNewTail := length - k
    newTail := head
    for i := 1; i < stepsToNewTail; i++ {
        newTail = newTail.Next
    }

    // 5. Break the circle
    newHead := newTail.Next
    newTail.Next = nil

    return newHead
}