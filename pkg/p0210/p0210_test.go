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

package p0210

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0210TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   [][]int
	target []int
}

var values = []result{
	{
		arg1: 2,
		arg2: [][]int{
			[]int{1, 0},
		},
		target: []int{0, 1},
	},
	{
		arg1: 4,
		arg2: [][]int{
			[]int{1, 0},
			[]int{2, 0},
			[]int{3, 1},
			[]int{3, 2},
		},
		target: []int{0, 1, 2, 3},
	},
}

func (s *p0210TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findOrder(v.arg1, v.arg2))
	}
}

func TestP0210TestSuite(t *testing.T) {
	s := &p0210TestSuite{}
	suite.Run(t, s)
}
