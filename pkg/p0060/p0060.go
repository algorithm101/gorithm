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

package p0060

func getPermutation(n int, k int) string {
	ret, w := make([]byte, n), 0

	fac := make([]int, n)
	fac[0] = 1
	for i := 1; i < n; i++ {
		fac[i] = (i) * fac[i-1]
	}
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = byte(i) + '1'
	}

	k--
	for i := n - 1; i >= 0; i-- {
		index, b := k/fac[i], byte('0') // 取第 index 大的数值
		for j := 0; j < n; j++ {
			if v[j] == '0' {
				continue
			}

			if index == 0 {
				b = v[j]
				v[j] = '0'
				break
			}
			index--
		}

		ret[w] = b
		w++

		k %= fac[i]
	}

	return string(ret)
}
