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

package p0139

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0139TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	arg2   []string
	target bool
}

var values = []result{
	{
		arg1:   "leetcode",
		arg2:   []string{"leet", "code"},
		target: true,
	},
	{
		arg1:   "applepenapple",
		arg2:   []string{"apple", "pen"},
		target: true,
	},
	{
		arg1:   "catsandog",
		arg2:   []string{"cats", "dog", "sand", "and", "cat"},
		target: false,
	},
	{
		arg1: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		arg2: []string{
			"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		target: false,
	},
}

func (s *p0139TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, wordBreak(v.arg1, v.arg2))
	}
}

func TestP0139TestSuite(t *testing.T) {
	s := &p0139TestSuite{}
	suite.Run(t, s)
}
