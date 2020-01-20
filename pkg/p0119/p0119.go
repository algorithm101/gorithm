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

package p0119

func getRow(rowIndex int) []int {
	_cnk := func(n int64, k int64) int {
		if k > n {
			return 0
		} else if k > n/2 {
			k = n - k
		}

		numerator, denomator := int64(1), int64(1)
		for i := int64(0); i < k; i++ {
			numerator *= (n - i)
			denomator *= (k - i)

			if numerator%denomator == 0 {
				numerator = numerator / denomator
				denomator = 1
			} else {
				vs := []int64{2, 5}
				for _, v := range vs {
					if numerator%v == 0 && denomator%v == 0 {
						numerator /= v
						denomator /= v
					}
				}
			}
		}
		return int(numerator / denomator)
	}

	ret := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		if i == 0 || i == rowIndex {
			ret[i] = 1
		} else {
			ret[i] = _cnk(int64(rowIndex), int64(i))
		}
	}

	return ret
}
