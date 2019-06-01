package easy

import "testing"

func TestMinDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	got, want := minDepth(root), 2
	if got != want {
		t.Fatalf("got: %v; want: %v", got, want)
	}
}
