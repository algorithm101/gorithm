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

package p0216

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0216TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   int
	arg2   int
	target [][]int
}

var values = []result{
	{
		arg1: 3,
		arg2: 7,
		target: [][]int{
			[]int{1, 2, 4},
		},
	},
	{
		arg1: 3,
		arg2: 9,
		target: [][]int{
			[]int{1, 2, 6},
			[]int{1, 3, 5},
			[]int{2, 3, 4},
		},
	},
}

func (s *p0216TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, combinationSum3(v.arg1, v.arg2))
	}
}

func TestP0216TestSuite(t *testing.T) {
	s := &p0216TestSuite{}
	suite.Run(t, s)
}
