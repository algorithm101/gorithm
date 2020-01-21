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

package p0888

func fairCandySwap(A []int, B []int) []int {
	sum, sumA, sumB := 0, 0, 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		sumA += A[i]
	}

	mapsb := make(map[int]bool)
	for i := 0; i < len(B); i++ {
		sum += B[i]
		sumB += B[i]
		mapsb[B[i]] = true
	}

	sum /= 2

	for i := 0; i < len(A); i++ {
		b := sum - (sumA - A[i])
		if mapsb[b] {
			return []int{A[i], b}
		}
	}

	return nil
}
