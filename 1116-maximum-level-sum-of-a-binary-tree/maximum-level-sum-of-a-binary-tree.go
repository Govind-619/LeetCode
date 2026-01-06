/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := []*TreeNode{root}
    level := 0
    maxSum := math.MinInt32
    resultLevel := math.MaxInt32
    
    for len(queue) > 0 {
        sizeLen := len(queue)
        sum := 0

        for i := 0; i < sizeLen; i++ {
            node := queue[0]
            queue = queue[1:]
            sum += node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        level++

        if sum > maxSum {
            maxSum = sum
            resultLevel = level
        }
    } 

    return resultLevel
}