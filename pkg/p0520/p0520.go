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

package p0520

func detectCapitalUse(word string) bool {
	_isUppercase := func(b byte) bool {
		return 'A' <= b && b <= 'Z'
	}

	statFuncs := make([]func(byte) bool, 4)
	state := 0
	statFuncs[0] = func(b byte) bool {
		if _isUppercase(b) {
			state = 2
		} else {
			state = 1
		}
		return true
	}
	statFuncs[1] = func(b byte) bool {
		return !_isUppercase(b)
	}
	statFuncs[2] = func(b byte) bool {
		if _isUppercase(b) {
			state = 3
		} else {
			state = 1
		}
		return true
	}
	statFuncs[3] = func(b byte) bool {
		return _isUppercase(b)
	}

	for i := 0; i < len(word); i++ {
		if !statFuncs[state](word[i]) {
			return false
		}
	}

	return true
}
