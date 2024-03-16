package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f) // f is the in orderfunc
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
