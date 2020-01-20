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

package p0575

func distributeCandies(candies []int) int {
	maps := make(map[int]int, len(candies))
	for _, v := range candies {
		maps[v] = 1
	}

	count := len(maps)
	if count >= len(candies)/2 {
		return len(candies) / 2
	}

	return count
}
