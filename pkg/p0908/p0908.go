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

package p0908

func smallestRangeI(A []int, K int) int {
	if 0 == len(A) {
		return 0
	}

	max, min := A[0], A[0]
	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > max:
			max = A[i]
		case A[i] < min:
			min = A[i]
		}
	}

	ret := max - min - 2*K
	if ret < 0 {
		return 0
	}
	return ret
}
