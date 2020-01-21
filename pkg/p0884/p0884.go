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

package p0884

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	mapsa := make(map[string]int)
	for _, w := range strings.Split(A, " ") {
		count := mapsa[w]
		mapsa[w] = count + 1
	}

	mapsb := make(map[string]int)
	for _, w := range strings.Split(B, " ") {
		count := mapsb[w]
		mapsb[w] = count + 1
	}

	ret := []string{}

	for w, c := range mapsa {
		if c == 1 && 0 == mapsb[w] {
			ret = append(ret, w)
		}
	}

	for w, c := range mapsb {
		if c == 1 && 0 == mapsa[w] {
			ret = append(ret, w)
		}
	}

	return ret
}
