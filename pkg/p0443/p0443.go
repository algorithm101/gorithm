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

package p0443

import (
	"strconv"
)

func compress(chars []byte) int {
	_writeCount := func(chars []byte, j int, count int) int {
		if count == 1 {
			return 0
		}

		countStr := strconv.Itoa(count)
		for _, c := range countStr {
			chars[j] = byte(c)
			j++
		}

		return len(countStr)
	}

	j, count := 1, 1
	for i, last := 1, chars[0]; i < len(chars); i++ {
		if chars[i] != last {
			last = chars[i]
			j += _writeCount(chars, j, count)
			chars[j] = chars[i]
			j++
			count = 1
		} else {
			count++
		}
	}

	j += _writeCount(chars, j, count)

	return j
}
