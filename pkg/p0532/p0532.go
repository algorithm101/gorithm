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

package p0532

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	maps := make(map[int]int, len(nums))
	for _, v := range nums {
		count := maps[v]
		maps[v] = count + 1
	}

	count := 0
	for v, c := range maps {
		delete(maps, v)

		if k == 0 && c > 1 {
			count++
			continue
		}

		if _, exists := maps[v+k]; exists {
			count++
		}
		if _, exists := maps[v-k]; exists {
			count++
		}
	}

	return count
}
