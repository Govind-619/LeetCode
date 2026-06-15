/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next==nil{
        return nil
    }
    current:= head
    n:=1
    for current!= nil{
        current= current.Next
        n++
    }
    current= head
    if n%2==0{
        for i:=1; i<n/2-1; i++{
        current= current.Next
        }
    }else{
        for i:=1; i<n/2; i++{
        current= current.Next
        }
    }
    current.Next= current.Next.Next
    
    
    return head
    
}