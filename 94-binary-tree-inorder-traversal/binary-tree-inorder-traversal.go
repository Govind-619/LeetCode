/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    // Step 1: Traverse left subtree
    res = append(res, inorderTraversal(root.Left)...)

    // Step 2: Visit root node
    res = append(res, root.Val)

    // Step 3: Traverse right subtree
    res = append(res, inorderTraversal(root.Right)...)

    return res
}
