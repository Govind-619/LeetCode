/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
  leftDepth, rightDepth := maxDepth(root.Left, 0), maxDepth(root.Right, 0)

  switch {
  case leftDepth == rightDepth:
    return root
  case leftDepth > rightDepth:
    return subtreeWithAllDeepest(root.Left)
  default:
    return subtreeWithAllDeepest(root.Right)
  }
}

func maxDepth(root *TreeNode, level int) int {
  if root == nil {
    return level
  }

  return max(maxDepth(root.Left, level+1), maxDepth(root.Right, level+1))
}