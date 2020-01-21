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

package p0645

func findErrorNums(nums []int) []int {
	repeated := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if nums[i] == i+1 {
			continue
		}

		if nums[v-1] == v {
			repeated = nums[v-1]
			break
		}

		nums[i], nums[v-1] = nums[v-1], nums[i]
		i--
	}

	diff := 0
	for i := 1; i <= len(nums); i++ {
		diff += (nums[i-1] - i)
	}

	return []int{repeated, repeated - diff}
}
