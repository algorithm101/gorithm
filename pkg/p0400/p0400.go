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

package p0400

import (
	"strconv"
)

func findNthDigit(n int) int {
	digitType, digitSum := 1, 9
	for n > digitType*digitSum {
		n -= digitType * digitSum
		digitType++
		digitSum *= 10
	}

	indexSubRange, indexInNum := (n-1)/digitType, (n-1)%digitType

	num := 1
	for i := 1; i < digitType; i++ {
		num *= 10
	}
	num += indexSubRange

	r, _ := strconv.Atoi(string(strconv.Itoa(num)[indexInNum]))
	return r
}
