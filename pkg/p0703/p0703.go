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
package p0703

import "container/heap"

type KthLargest struct {
	arr []int
	cap int
}

func (k *KthLargest) Len() int {
	return len(k.arr)
}

func (k *KthLargest) Less(i int, j int) bool {
	return k.arr[i] < k.arr[j]
}

func (k *KthLargest) Swap(i int, j int) {
	k.arr[i], k.arr[j] = k.arr[j], k.arr[i]
}

func (k *KthLargest) Push(x interface{}) {
	v := x.(int)
	k.arr = append(k.arr, v)
}

func (k *KthLargest) Pop() interface{} {
	x := k.arr[len(k.arr)-1]
	k.arr = k.arr[0 : len(k.arr)-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		cap: k,
	}

	if len(nums) >= k {
		kth.arr = nums[0:k]
	} else {
		kth.arr = nums
	}

	heap.Init(&kth)

	for i := k; i < len(nums); i++ {
		kth.Add(nums[i])
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) >= this.cap {
		if this.arr[0] < val {
			this.arr[0] = val
			heap.Fix(this, 0)
		}
	} else {
		this.arr = append(this.arr, val)
		heap.Fix(this, len(this.arr)-1)
	}

	return this.arr[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
