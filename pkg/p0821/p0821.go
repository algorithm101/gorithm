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

package p0821

func shortestToChar(S string, C byte) []int {
	ret, last := make([]int, len(S)), 0-len(S)

	for i := 0; i < len(S); i++ {
		if S[i] == C {
			last = i
			ret[i] = 0

			for j := i - 1; j >= 0 && S[j] != C; j-- {
				if last-j < ret[j] {
					ret[j] = last - j
				}
			}
		} else {
			ret[i] = i - last
		}
	}

	return ret
}
