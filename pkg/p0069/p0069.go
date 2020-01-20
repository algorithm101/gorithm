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

package p0069

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	begin, end, res := 1, x, 0
	for begin <= end {
		middle := (begin + end) / 2
		if middle*middle == x {
			return middle
		}

		if middle*middle < x {
			begin = middle + 1
			res = middle
		} else {
			end = middle - 1
		}
	}

	return res
}
