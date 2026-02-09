func balanceBST(root *TreeNode) *TreeNode {
    stree := sort(root)
    return bbst(stree)
}

func sort(root *TreeNode) []*TreeNode {
    if root == nil {
        return []*TreeNode{}
    }

    left := sort(root.Left)
    right := sort(root.Right)
    left = append(left, root)
    return append(left, right...)
}

func bbst(stree []*TreeNode) *TreeNode {
    if len(stree) == 0 {
        return nil
    }

    base := stree[len(stree) / 2]
    base.Left = bbst(stree[:len(stree) / 2])
    base.Right = bbst(stree[len(stree) / 2 + 1:])

    return base
}