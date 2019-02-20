package leetCode

/*
递归
如果在root的左右子树中分别找到p, q,则root就是LCA节点
如果左子树中没有找到p和q，则p，q在root的右子树中，right就是LCA节点
同理，left是LCA节点
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

/*
Method 2:PATH
从根->p, 根->q 的path中，找到最后一个相等的节点,就是结果
*/
func findPath(root, tar *TreeNode, s []*TreeNode) ([]*TreeNode, bool) {
	if root == nil {
		return s, false
	}

	if root == tar {
		s = append(s, root)
		return s, true
	}

	s = append(s, root)
	s, l := findPath(root.Left, tar, s)
	if l {
		return s, true
	}
	s, r := findPath(root.Right, tar, s)
	if r {
		return s, true
	}

	s = s[0 : len(s)-1]
	return s, false
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {

		return root
	}

	pPath, _ := findPath(root, p, make([]*TreeNode, 0))
	qPath, _ := findPath(root, q, make([]*TreeNode, 0))

	var resp *TreeNode
	for i := 0; i < len(pPath) && i < len(qPath); i++ {
		if pPath[i] != qPath[i] {
			break
		}
		resp = pPath[i]
	}

	return resp
}
