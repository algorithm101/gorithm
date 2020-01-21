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

package p0232

import "container/list"

type MyQueue struct {
	list  *list.List
	list2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		list:  list.New(),
		list2: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.list.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for this.list.Len() != 0 {
		e := this.list.Back()
		this.list2.PushBack(e.Value)
		this.list.Remove(e)
	}

	e := this.list2.Back()
	this.list2.Remove(e)

	for this.list2.Len() != 0 {
		e := this.list2.Back()
		this.list.PushBack(e.Value)
		this.list2.Remove(e)
	}

	if e != nil {
		return e.Value.(int)
	}
	return 0
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for this.list.Len() != 0 {
		e := this.list.Back()
		this.list2.PushBack(e.Value)
		this.list.Remove(e)
	}

	e := this.list2.Back()
	// this.list2.Remove(e)

	for this.list2.Len() != 0 {
		e := this.list2.Back()
		this.list.PushBack(e.Value)
		this.list2.Remove(e)
	}

	if e != nil {
		return e.Value.(int)
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.list.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
