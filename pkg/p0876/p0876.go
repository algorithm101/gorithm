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

package p0876

import (
	"algorithm101.io/gorithm/pkg/utils"
)

// ListNode for singly-linked list
type ListNode = utils.ListNode

func middleNode(head *ListNode) *ListNode {
	head2 := head

	for head2 != nil {
		head2 = head2.Next
		if head2 == nil {
			break
		}
		head2 = head2.Next

		head = head.Next
	}

	return head
}
