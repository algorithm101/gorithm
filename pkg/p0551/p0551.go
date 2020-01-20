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

package p0551

func checkRecord(s string) bool {
	if 0 == len(s) || 1 == len(s) {
		return true
	}

	acount, llcount := 0, 0
	if s[0] == 'A' {
		acount++
	}

	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == 'A':
			acount++
		case s[i] == 'L' && s[i-1] == 'L' && (i >= 2 && s[i-2] == 'L'):
			llcount++
		}
	}

	return acount <= 1 && llcount <= 0
}
