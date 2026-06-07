/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    
    nodes := make(map[int]*TreeNode)
    
    children := make(map[int]bool)
    
    for _, d := range descriptions {
        parent := d[0]
        child := d[1]
        isLeft := d[2]
        
        if _, exists := nodes[parent]; !exists {
            nodes[parent] = &TreeNode{Val: parent}
        }
        
        if _, exists := nodes[child]; !exists {
            nodes[child] = &TreeNode{Val: child}
        }
        
        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }
        
        children[child] = true
    }
    
    for value, node := range nodes {
        if !children[value] {
            return node
        }
    }
    
    return nil
}