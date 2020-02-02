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
package p0225

import (
	"container/list"
)

type MyStack struct {
	l *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		l: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.l.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	e := this.l.Back()
	if e != nil {
		v := e.Value.(int)
		this.l.Remove(e)
		return v
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	e := this.l.Back()
	if e != nil {
		v := e.Value.(int)
		return v
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l.Len() == 0
}
