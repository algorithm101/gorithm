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

package p0867

func transpose(A [][]int) [][]int {
	if 0 == len(A) {
		return A
	}

	r, c := len(A), len(A[0])
	ret := make([][]int, c)
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, r)
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			ret[j][i] = A[i][j]
		}
	}

	return ret
}
