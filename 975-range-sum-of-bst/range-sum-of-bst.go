/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }

    if root.Val < low {
        // All values in the left subtree will be smaller, so skip it
        return rangeSumBST(root.Right, low, high)
    }

    if root.Val > high {
        // All values in the right subtree will be larger, so skip it
        return rangeSumBST(root.Left, low, high)
    }

    // Current node is in range
    return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
