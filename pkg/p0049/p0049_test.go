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

package p0049

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0049TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []string
	target [][]string
}

var values = []result{
	{
		arg1: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		target: [][]string{
			{"bat"},
			{"ate", "eat", "tea"},
			{"nat", "tan"},
		},
	},
}

func (s *p0049TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, groupAnagrams(v.arg1))
	}
}

func TestP0049TestSuite(t *testing.T) {
	s := &p0049TestSuite{}
	suite.Run(t, s)
}
