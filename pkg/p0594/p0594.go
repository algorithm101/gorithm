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

package p0594

func findLHS(nums []int) int {
	maps := make(map[int]int, len(nums))
	for _, v := range nums {
		count := maps[v]
		maps[v] = count + 1
	}

	max := 0
	for k, count := range maps {
		if count2, exists := maps[k+1]; exists {
			if count+count2 > max {
				max = count + count2
			}
		}
	}

	return max
}
