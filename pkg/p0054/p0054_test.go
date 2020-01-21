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

package p0054

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0054TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target []int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		target: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		arg1: [][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
		},
		target: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		arg1: [][]int{
			[]int{7},
			[]int{9},
			[]int{6},
		},
		target: []int{7, 9, 6},
	},
}

func (s *p0054TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, spiralOrder(v.arg1))
	}
}

func TestP0054TestSuite(t *testing.T) {
	s := &p0054TestSuite{}
	suite.Run(t, s)
}
