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

package p0091

func numDecodings(s string) int {
	if 0 == len(s) || s[0] == '0' {
		return 0
	}

	dp0, dp1 := 1, 1
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && (s[i-1] > '2' || s[i-1] == '0'):
			return 0
		case s[i] == '0' || s[i-1] == '0':
			dp1, dp0 = dp0, dp1
		case s[i-1:i+1] <= "26":
			dp1, dp0 = dp1+dp0, dp1
		default:
			dp1, dp0 = dp1, dp1
		}
	}

	return dp1
}
