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

package p0883

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0883TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{2},
		},
		target: 5,
	},
	{
		arg1: [][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		target: 17,
	},
	{
		arg1: [][]int{
			[]int{1, 0},
			[]int{0, 2},
		},
		target: 8,
	},
	{
		arg1: [][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		target: 14,
	},
	{
		arg1: [][]int{
			[]int{2, 2, 2},
			[]int{2, 1, 2},
			[]int{2, 2, 2},
		},
		target: 21,
	},
}

func (s *p0883TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, projectionArea(v.arg1))
	}
}

func TestP0883TestSuite(t *testing.T) {
	s := &p0883TestSuite{}
	suite.Run(t, s)
}
