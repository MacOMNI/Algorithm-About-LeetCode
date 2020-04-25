package LeetCode

var nodes []int
var counts []int

func bstToGst(root *TreeNode) *TreeNode {

	//nodes := make([]int,0)
	counts = make([]int, 0)
	nodes = treeOrder(root)
	// fmt.Println(nodes)
	counts = append(counts, nodes[0])
	for i := 1; i < len(nodes); i++ {
		counts = append(counts, counts[i-1]+nodes[i])
	}
	//  fmt.Println(counts)
	updateTree(root)
	return root
}
func updateTree(root *TreeNode) {
	if root == nil {
		return
	}
	beforeIndex := nodesBinarySearch(root.Val)
	// fmt.Println(root.Val," ",beforeIndex)

	root.Val = counts[beforeIndex-1]

	updateTree(root.Left)
	updateTree(root.Right)

}
func treeOrder(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}
	res = append(res, treeOrder(root.Right)...)
	res = append(res, root.Val)
	res = append(res, treeOrder(root.Left)...)
	return res
}
func nodesBinarySearch(val int) int {
	left := 0
	right := len(nodes)
	mid := (left + right) >> 1
	for left < right {
		if nodes[mid] >= val {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) >> 1
	}
	return mid
}
