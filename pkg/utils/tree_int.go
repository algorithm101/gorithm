// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"

	"github.com/stretchr/testify/suite"
)

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trees(arg []int, parent *TreeNode, index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2 * (index + 1)

	if leftChildIndex < len(arg) && arg[leftChildIndex] != 0xFFFFFFFF {
		parent.Left = &TreeNode{
			Val:   arg[leftChildIndex],
			Left:  nil,
			Right: nil,
		}

		trees(arg, parent.Left, leftChildIndex)
	}

	if rightChildIndex < len(arg) && arg[rightChildIndex] != 0xFFFFFFFF {
		parent.Right = &TreeNode{
			Val:   arg[rightChildIndex],
			Left:  nil,
			Right: nil,
		}

		trees(arg, parent.Right, rightChildIndex)
	}
}

// Trees construct tree node
func Trees(arg []int) *TreeNode {
	if 0 == len(arg) {
		return nil
	}

	root := &TreeNode{
		Val:   arg[0],
		Left:  nil,
		Right: nil,
	}

	trees(arg, root, 0)

	return root
}

// PrintTree ...
func PrintTree(root *TreeNode, prefix string) {
	v := "nil"
	if root != nil {
		v = fmt.Sprintf("%d", root.Val)
	}
	fmt.Println(fmt.Sprintf("%s: %s", prefix, v))

	if root != nil {
		PrintTree(root.Left, prefix+"  ")
		PrintTree(root.Right, prefix+"  ")
	}
}

// AssertTreeEqual check the link node is equal
func AssertTreeEqual(t suite.Suite, expected *TreeNode, actual *TreeNode) {
	if expected == nil && actual == nil {
		return
	}
	t.NotNil(expected)
	t.NotNil(actual)

	if expected == nil {
		t.FailNow("expected value should not be nil")
	}
	if actual == nil {
		t.FailNow("actual value should not be nil")
	}

	t.Equal(expected.Val, actual.Val)

	AssertTreeEqual(t, expected.Left, actual.Left)
	AssertTreeEqual(t, expected.Right, actual.Right)
}
