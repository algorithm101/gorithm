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

package p0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
		} else if n == 0 {
			return
		} else {
			if nums1[m-1] > nums2[n-1] {
				nums1[m+n-1] = nums1[m-1]
				m--
			} else {
				nums1[m+n-1] = nums2[n-1]
				n--
			}
		}
	}
}
