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

package p0705

import "fmt"

type MyHashSet struct {
	arr [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	r := MyHashSet{
		arr: make([][]int, 100),
	}
	for i := 0; i < len(r.arr); i++ {
		r.arr[i] = make([]int, 0)
	}

	return r
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

	i := key % 100
	v := this.arr[i]
	this.arr[i] = append(v, key)
}

func (this *MyHashSet) Remove(key int) {
	i := key % 100
	v := this.arr[i]
	for index, num := range v {
		if num == key {
			v[index], v[len(v)-1] = v[len(v)-1], v[index]
			this.arr[i] = v[0 : len(v)-1]
			fmt.Println(this.arr[i])
			break
		}
	}
}

/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
	i := key % 100
	v := this.arr[i]
	for _, num := range v {
		if num == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
