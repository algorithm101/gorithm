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

package p0011

func maxArea(height []int) int {
	minInt := func(v1 int, v2 int) int {
		if v1 < v2 {
			return v1
		}
		return v2
	}

	left, right, max := 0, len(height)-1, 0

	for left < right {
		curr := (right - left) * minInt(height[right], height[left])
		if curr > max {
			max = curr
		}

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return max
}
