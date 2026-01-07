/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
    const MOD int64 = 1000000007
    var total, best int64

    var sum func(*TreeNode) int64
    sum = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }
        return int64(node.Val) + sum(node.Left) + sum(node.Right)
    }

    total = sum(root)

    var dfs func(*TreeNode) int64
    dfs = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }
        s := int64(node.Val) + dfs(node.Left) + dfs(node.Right)
        prod := s * (total - s)
        if prod > best {
            best = prod
        }
        return s
    }

    dfs(root)
    return int(best % MOD)
}