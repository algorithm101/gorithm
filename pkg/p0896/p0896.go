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

package p0896

func isMonotonic(A []int) bool {
	mode := 0
	for i := 1; i < len(A); i++ {
		switch mode {
		case 0:
			if A[i] > A[i-1] {
				mode = 1
			} else if A[i] < A[i-1] {
				mode = 2
			}
		case 1:
			if A[i] < A[i-1] {
				return false
			}
		case 2:
			if A[i] > A[i-1] {
				return false
			}
		}
	}

	return true
}
