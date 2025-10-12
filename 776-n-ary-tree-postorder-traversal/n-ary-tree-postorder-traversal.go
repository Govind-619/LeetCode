/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    res:= []int{} 
    var order func(node *Node)
    order = func(node *Node){
        if node == nil{
            return
        }
        for _,n:= range node.Children{
            order(n)
        }
        res= append(res,node.Val)
        return
    }
    order(root)
    return res
}