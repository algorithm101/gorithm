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

package p0482

func licenseKeyFormatting(S string, K int) string {
	count := 0
	for i := 0; i < len(S); i++ {
		if S[i] != '-' {
			count++
		}
	}
	if count == 0 {
		return ""
	}

	dest := make([]byte, count+count/K)
	j, w := len(dest)-1, 0

	for i := len(S) - 1; i >= 0; i-- {
		if w == K {
			dest[j] = '-'
			j--
			w = 0
		}

		if S[i] != '-' {
			dest[j] = S[i]
			if S[i] >= 'a' && S[i] <= 'z' {
				dest[j] = S[i] - 'a' + 'A'
			}

			j--
			w++
		}
	}

	if w == 0 {
		return string(dest[j+2:])
	}

	return string(dest[j+1:])
}
