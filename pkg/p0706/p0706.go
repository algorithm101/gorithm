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

package p0706

type entry struct {
	key   int
	value int
}

type MyHashMap struct {
	values [][]entry
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	r := MyHashMap{
		values: make([][]entry, 100),
	}

	for i := 0; i < len(r.values); i++ {
		r.values[i] = make([]entry, 0)
	}
	return r
}

/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int) {
	i := key % 100
	v := this.values[i]
	for index, e := range v {
		if key == e.key {
			v[index] = entry{
				key:   key,
				value: value,
			}
			return
		}
	}

	this.values[i] = append(v, entry{
		key:   key,
		value: value,
	})
	return
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	i := key % 100
	v := this.values[i]
	for _, e := range v {
		if key == e.key {
			return e.value
		}
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	i := key % 100
	v := this.values[i]
	for index, e := range v {
		if key == e.key {
			v[index], v[len(v)-1] = v[len(v)-1], v[index]
			this.values[i] = v[0 : len(v)-1]
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
