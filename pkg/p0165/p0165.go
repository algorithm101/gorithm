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

package p0165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	_toInt := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}

	_compare := func(v1 []string, v2 []string, fn interface{}) int {
		if 0 == len(v1) && 0 == len(v2) {
			return 0
		}

		if 0 == len(v2) {
			v2 = []string{"0"}
		}
		if 0 == len(v1) {
			v1 = []string{"0"}
		}

		v1i, v2i := _toInt(v1[0]), _toInt(v2[0])
		if v1i > v2i {
			return 1
		} else if v1i < v2i {
			return -1
		}

		_fn := fn.(func([]string, []string, interface{}) int)
		return _fn(v1[1:len(v1)], v2[1:len(v2)], fn)
	}

	return _compare(strings.Split(version1, "."), strings.Split(version2, "."), _compare)
}
