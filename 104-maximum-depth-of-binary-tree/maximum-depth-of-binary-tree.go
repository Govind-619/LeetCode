/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    LeftDepth:= maxDepth(root.Left)
    RightDepth:= maxDepth(root.Right)
    if LeftDepth > RightDepth{
        return LeftDepth+1
    }else{
        return RightDepth+1
    }
    
}