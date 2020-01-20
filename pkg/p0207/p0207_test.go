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

package p0207

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0207TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   [][]int
	target bool
}

var values = []result{
	{
		arg1: 2,
		arg2: [][]int{
			[]int{1, 0},
		},
		target: true,
	},
	{
		arg1: 2,
		arg2: [][]int{
			[]int{1, 0},
			[]int{0, 1},
		},
		target: false,
	},
	{
		arg1: 3,
		arg2: [][]int{
			[]int{1, 0},
			[]int{1, 2},
			[]int{0, 1},
		},
		target: false,
	},
}

func (s *p0207TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canFinish(v.arg1, v.arg2))
	}
}

func TestP0207TestSuite(t *testing.T) {
	s := &p0207TestSuite{}
	suite.Run(t, s)
}
