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

package p0008

import (
	"math"
)

func myAtoi(str string) int {
	isNegative, ret, mode := false, int(0), 0

	for i := 0; i < len(str); i++ {
		v := ret
		if isNegative {
			v = 0 - v
		}
		if v > math.MaxInt32 {
			return math.MaxInt32
		} else if v < math.MinInt32 {
			return math.MinInt32
		}

		switch mode {
		case 0:
			switch {
			case str[i] == '+':
				mode = 1
			case str[i] == '-':
				mode = 1
				isNegative = true
			case str[i] >= '0' && str[i] <= '9':
				mode = 2
				ret = int(str[i] - '0')
			case str[i] != ' ':
				return 0
			}
		case 1:
			if str[i] >= '0' && str[i] <= '9' {
				mode = 2
				ret = int(str[i] - '0')
			} else {
				return 0
			}
		case 2:
			if str[i] >= '0' && str[i] <= '9' {
				ret = ret*10 + int(str[i]-'0')
			} else {
				mode = 3
			}
		}
	}

	if isNegative {
		ret = 0 - ret
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	} else if ret < math.MinInt32 {
		return math.MinInt32
	}
	return ret
}
