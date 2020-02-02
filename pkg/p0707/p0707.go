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

//nolint
package p0707

type node struct {
	next *node
	prev *node
	val  int
}

type MyLinkedList struct {
	head *node
	tail *node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	r := MyLinkedList{
		head: &node{},
		size: 0,
	}
	r.tail = r.head
	r.tail.next, r.head.next = r.head, r.head
	return r
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	n := this.head
	for index >= 0 {
		n = n.next
		index--
	}

	return n.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	n := &node{
		next: this.head.next,
		prev: this.head,
		val:  val,
	}

	this.head.next = n
	n.next.prev = n
	this.size++

	this.Print()
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	n := &node{
		next: this.tail,
		prev: this.tail.prev,
		val:  val,
	}
	this.tail.prev = n
	n.prev.next = n

	this.size++

	this.Print()
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	this.size++
	n := this.head
	for index > 0 {
		n = n.next
		index--
	}

	ne := &node{
		next: n.next,
		prev: n,
		val:  val,
	}
	n.next = ne
	ne.next.prev = ne

	this.Print()
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	this.size--

	n := this.head
	for index > 0 {
		n = n.next
		index--
	}

	n.next = n.next.next
	n.next.prev = n
}

func (this *MyLinkedList) Print() {
	// 	n := this.head.next
	// 	for n != this.tail {
	// 		fmt.Printf("%d, ", n.val)
	// 		n = n.next
	// 	}

	// 	fmt.Println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
