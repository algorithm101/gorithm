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

package p0047

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0047TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target [][]int
}

var values = []result{
	{
		arg1: []int{1, 1, 2},
		target: [][]int{
			[]int{1, 1, 2},
			[]int{1, 2, 1},
			[]int{2, 1, 1},
		},
	},
	{
		arg1: []int{1, 1, 2, 2},
		target: [][]int{
			[]int{1, 1, 2, 2},
			[]int{1, 2, 1, 2},
			[]int{1, 2, 2, 1},
			[]int{2, 1, 1, 2},
			[]int{2, 1, 2, 1},
			[]int{2, 2, 1, 1},
		},
	},
	{
		arg1: []int{0, 1, 0, 0, 9},
		target: [][]int{
			[]int{0, 0, 0, 1, 9},
			[]int{0, 0, 0, 9, 1},
			[]int{0, 0, 1, 0, 9},
			[]int{0, 0, 1, 9, 0},
			[]int{0, 0, 9, 1, 0},
			[]int{0, 0, 9, 0, 1},
			[]int{0, 1, 0, 0, 9},
			[]int{0, 1, 0, 9, 0},
			[]int{0, 1, 9, 0, 0},
			[]int{0, 9, 0, 1, 0},
			[]int{0, 9, 0, 0, 1},
			[]int{0, 9, 1, 0, 0},
			[]int{1, 0, 0, 0, 9},
			[]int{1, 0, 0, 9, 0},
			[]int{1, 0, 9, 0, 0},
			[]int{1, 9, 0, 0, 0},
			[]int{9, 0, 0, 1, 0},
			[]int{9, 0, 0, 0, 1},
			[]int{9, 0, 1, 0, 0},
			[]int{9, 1, 0, 0, 0},
		},
	},
}

func (s *p0047TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, permuteUnique(v.arg1))
	}
}

func TestP0047TestSuite(t *testing.T) {
	s := &p0047TestSuite{}
	suite.Run(t, s)
}
