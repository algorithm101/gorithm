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

package p0566

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0566TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	arg2   int
	arg3   int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		arg2: 1,
		arg3: 4,
		target: [][]int{
			[]int{1, 2, 3, 4},
		},
	},
	{
		arg1: [][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		arg2: 2,
		arg3: 4,
		target: [][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
	},
}

func (s *p0566TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, matrixReshape(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0566TestSuite(t *testing.T) {
	s := &p0566TestSuite{}
	suite.Run(t, s)
}
