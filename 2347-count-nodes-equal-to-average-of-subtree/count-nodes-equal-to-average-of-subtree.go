/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfSubtree(root *TreeNode) int {
    answer := 0

    var dfs func(*TreeNode) (int, int)

    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0
        }

        leftSum, leftCount := dfs(node.Left)
        rightSum, rightCount := dfs(node.Right)

        sum := leftSum + rightSum + node.Val
        count := leftCount + rightCount + 1

        if node.Val == sum/count {
            answer++
        }

        return sum, count
    }

    dfs(root)
    return answer
}