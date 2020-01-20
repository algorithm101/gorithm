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

package p0073

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0073TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		target: [][]int{
			[]int{1, 0, 1},
			[]int{0, 0, 0},
			[]int{1, 0, 1},
		},
	},
	{
		arg1: [][]int{
			[]int{0, 1, 2, 0},
			[]int{3, 4, 5, 2},
			[]int{1, 3, 1, 5},
		},
		target: [][]int{
			[]int{0, 0, 0, 0},
			[]int{0, 4, 5, 0},
			[]int{0, 3, 1, 0},
		},
	},
}

func (s *p0073TestSuite) Test() {
	for _, v := range values {
		setZeroes(v.arg1)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0073TestSuite(t *testing.T) {
	s := &p0073TestSuite{}
	suite.Run(t, s)
}
