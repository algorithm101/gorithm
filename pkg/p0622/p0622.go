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

package p0622

import "fmt"

type MyCircularQueue struct {
	values []int
	start  int
	end    int
	size   int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		values: make([]int, k),
		start:  0,
		end:    0,
		size:   0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size >= len(this.values) {
		return false
	}

	this.values[this.end] = value
	this.size++
	this.end++

	if this.end >= len(this.values) {
		this.end = this.end % len(this.values)
	}

	fmt.Println("EnQueue: ", this.end)

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.size <= 0 {
		return false
	}
	this.start++
	if this.start >= len(this.values) {
		this.start = this.start % len(this.values)
	}
	this.size--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}

	return this.values[this.start]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}

	fmt.Println("Rear: ", this.end)
	end := this.end - 1
	if end < 0 {
		end = len(this.values) - 1
	}
	return this.values[end]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.values)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
