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

package p0784

import (
	"math"
)

func letterCasePermutation(S string) []string {
	count, bak := 0, make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		if (S[i] >= 'a' && S[i] <= 'z') || (S[i] >= 'A' && S[i] <= 'Z') {
			count++
		}
		if S[i] >= 'A' && S[i] <= 'Z' {
			bak[i] = S[i] - 'A' + 'a'
		} else {
			bak[i] = S[i]
		}
	}

	_lower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b - 'A' + 'a'
		}

		return b
	}

	ret, w := make([]string, int(math.Pow(float64(2), float64(count)))), 0
	_helper := func(i int, fn interface{}) {
		for ; i < len(bak); i++ {
			if bak[i] >= 'a' && bak[i] <= 'z' {
				break
			}
		}

		if i >= len(bak) {
			ret[w] = string(bak)
			w++
			return
		}

		for j := i; j < len(bak); j++ {
			bak[j] = _lower(bak[j])
		}

		_fn := fn.(func(int, interface{}))
		bak[i] = _lower(bak[i])
		_fn(i+1, fn)

		for j := i; j < len(bak); j++ {
			bak[j] = _lower(bak[j])
		}

		bak[i] = _lower(bak[i]) - 'a' + 'A'
		_fn(i+1, fn)
	}

	_helper(0, _helper)

	return ret
}
