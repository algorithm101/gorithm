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

package p0641

type MyCircularDeque struct {
	values []int
	start  int
	end    int
	size   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		values: make([]int, k),
		start:  0,
		end:    0,
		size:   0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size >= len(this.values) {
		return false
	}

	this.start--
	if this.start < 0 {
		this.start = len(this.values) - 1
	}
	this.values[this.start] = value
	this.size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size >= len(this.values) {
		return false
	}

	this.values[this.end] = value
	this.size++
	this.end++

	if this.end >= len(this.values) {
		this.end = this.end % len(this.values)
	}

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
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

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}

	this.size--
	this.end--
	if this.end < 0 {
		this.end = len(this.values) - 1
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}

	return this.values[this.start]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}

	end := this.end - 1
	if end < 0 {
		end = len(this.values) - 1
	}
	return this.values[end]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.values)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
