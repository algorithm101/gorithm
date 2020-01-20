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
	"bytes"
	"fmt"

	"github.com/stretchr/testify/suite"
)

// ListNode for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Links construct ListNode from []int
func Links(arg []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, v := range arg {
		tail.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		tail = tail.Next
	}

	return head.Next
}

// FindListNode returns the Node.Val equals target or nil
func FindListNode(n *ListNode, target int) *ListNode {
	for n != nil && n.Val != target {
		n = n.Next
	}

	return n
}

// PrintList ...
func PrintList(l *ListNode) {
	b := bytes.Buffer{}
	b.WriteString("[")

	for l != nil {
		b.WriteString(fmt.Sprintf("%d", l.Val))
		l = l.Next
		if l != nil {
			b.WriteString(", ")
		}
	}

	b.WriteString("]")

	fmt.Println(b.String())
}

// AssertLinkEqual check the link node is equal
func AssertLinkEqual(t suite.Suite, expected *ListNode, actual *ListNode) {
	for expected != nil && actual != nil {
		t.Equal(expected.Val, actual.Val)
		expected = expected.Next
		actual = actual.Next
	}

	t.Nil(expected)
	t.Nil(actual)
}
