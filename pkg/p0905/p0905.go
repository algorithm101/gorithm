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

package p0905

func sortArrayByParity(A []int) []int {
	_isOdd := func(v int) bool {
		return v&0x01 == 1
	}

	i, j := 0, len(A)-1
	for i < j {
		for i < j && !_isOdd(A[i]) {
			i++
		}
		for i < j && _isOdd(A[j]) {
			j--
		}
		if i >= j {
			break
		}

		A[i], A[j] = A[j], A[i]
	}
	return A
}
