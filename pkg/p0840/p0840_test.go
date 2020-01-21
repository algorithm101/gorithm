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

package p0840

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0840TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{4, 3, 8, 4},
			[]int{9, 5, 1, 9},
			[]int{2, 7, 6, 2},
		},
		target: 1,
	},
	{
		arg1: [][]int{
			[]int{10, 3, 5},
			[]int{1, 6, 11},
			[]int{7, 9, 2},
		},
		target: 0,
	},
}

func (s *p0840TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numMagicSquaresInside(v.arg1))
	}
}

func TestP0840TestSuite(t *testing.T) {
	s := &p0840TestSuite{}
	suite.Run(t, s)
}
