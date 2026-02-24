/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    var dfs func(node *TreeNode, num int) int
    dfs = func(node *TreeNode, num int) int {
        num = (num << 1) | node.Val
        if node.Left == nil && node.Right == nil {
            return num
        }
        sum := 0
        if node.Left != nil {
            sum += dfs(node.Left, num)
        }
        if node.Right != nil {
            sum += dfs(node.Right, num)
        }
        return sum
    }
    return dfs(root, 0)
}