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
package p0155

type MinStack struct {
	values []int
	mins   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)

	if 1 == len(this.values) {
		this.mins = append(this.mins, x)
	} else {
		min := this.mins[len(this.mins)-1]
		if x < min {
			this.mins = append(this.mins, x)
		} else {
			this.mins = append(this.mins, min)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.values) > 0 {
		this.values = this.values[0 : len(this.values)-1]
		this.mins = this.mins[0 : len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	if 0 == len(this.values) {
		return 0
	}

	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	if 0 == len(this.mins) {
		return 0
	}

	return this.mins[len(this.mins)-1]
}
