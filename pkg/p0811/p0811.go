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

package p0811

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	_subdomains := func(domain string) []string {
		segs := strings.Split(domain, ".")
		switch len(segs) {
		case 1:
			return segs
		case 2:
			return []string{domain, segs[1]}
		case 3:
			return []string{domain, segs[1] + "." + segs[2], segs[2]}
		}

		panic("Unexpected domain")
	}

	_parse := func(domain string) (int, []string) {
		segs := strings.Split(domain, " ")
		count, _ := strconv.Atoi(segs[0])
		return count, _subdomains(segs[1])
	}

	maps := make(map[string]int)
	for _, domain := range cpdomains {
		count, subdomains := _parse(domain)
		for _, subdomain := range subdomains {
			maps[subdomain] = maps[subdomain] + count
		}
	}

	ret, i := make([]string, len(maps)), 0
	for k, v := range maps {
		ret[i] = strconv.Itoa(v) + " " + k
		i++
	}

	return ret
}
