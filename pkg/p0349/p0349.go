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

package p0349

func intersection(nums1 []int, nums2 []int) []int {
	maps := make(map[int]int, len(nums1))

	for _, v := range nums1 {
		maps[v] = 1
	}

	ret := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := maps[v]; ok {
			ret = append(ret, v)
		}
		delete(maps, v)
	}
	return ret
}
