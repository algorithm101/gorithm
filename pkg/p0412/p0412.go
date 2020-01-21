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

package p0412

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	_toString := func(i int) string {
		switch {
		case i%15 == 0:
			return "FizzBuzz"
		case i%5 == 0:
			return "Buzz"
		case i%3 == 0:
			return "Fizz"
		default:
			return fmt.Sprintf("%d", i)
		}
	}

	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		ret[i-1] = _toString(i)
	}
	return ret
}
