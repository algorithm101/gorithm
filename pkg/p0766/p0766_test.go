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

package p0766

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0766TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target bool
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 1, 2, 3},
			[]int{9, 5, 1, 2},
		},
		target: true,
	},
	{
		arg1: [][]int{
			[]int{1, 2},
			[]int{2, 2},
		},
		target: false,
	},
	{
		arg1: [][]int{
			[]int{65, 98, 57},
		},
		target: true,
	},
}

func (s *p0766TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isToeplitzMatrix(v.arg1))
	}
}

func TestP0766TestSuite(t *testing.T) {
	s := &p0766TestSuite{}
	suite.Run(t, s)
}
