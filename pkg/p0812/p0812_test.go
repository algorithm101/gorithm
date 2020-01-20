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

package p0812

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target float64
}

var values = []result{
	{
		arg1: [][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{1, 0},
			[]int{0, 2},
			[]int{2, 0},
		},
		target: 2,
	},
}

type p0812TestSuite struct {
	suite.Suite
}

func (s *p0812TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestTriangleArea(v.arg1))
	}
}

func TestP0812TestSuite(t *testing.T) {
	s := &p0812TestSuite{}
	suite.Run(t, s)
}
