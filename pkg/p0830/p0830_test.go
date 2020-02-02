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

package p0830

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target [][]int
}

var values = []result{
	{
		arg1: "abbxxxxzzy",
		target: [][]int{
			{3, 6},
		},
	},
	{
		arg1:   "abc",
		target: [][]int{},
	},
	{
		arg1: "abcdddeeeeaabbbcd",
		target: [][]int{
			{3, 5},
			{6, 9},
			{12, 14},
		},
	},
}

type p0830TestSuite struct {
	suite.Suite
}

func (s *p0830TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largeGroupPositions(v.arg1))
	}
}

func TestP0830TestSuite(t *testing.T) {
	s := &p0830TestSuite{}
	suite.Run(t, s)
}
