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

package p0893

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0893TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"a", "b", "c", "a", "c", "c"},
		target: 3,
	},
	{
		arg1:   []string{"aa", "bb", "ab", "ba"},
		target: 4,
	},
	{
		arg1:   []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		target: 3,
	},
	{
		arg1:   []string{"abcd", "cdab", "adcb", "cbad"},
		target: 1,
	},
}

func (s *p0893TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numSpecialEquivGroups(v.arg1))
	}
}

func TestP0893TestSuite(t *testing.T) {
	s := &p0893TestSuite{}
	suite.Run(t, s)
}
