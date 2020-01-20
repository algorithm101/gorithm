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

package p0078

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0078TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target [][]int
}

var values = []result{
	{
		arg1: []int{1, 2, 3},
		target: [][]int{
			[]int{},
			[]int{1},
			[]int{2},
			[]int{1, 2},
			[]int{3},
			[]int{1, 3},
			[]int{2, 3},
			[]int{1, 2, 3},
		},
	},
	{
		arg1: []int{},
		target: [][]int{
			[]int{},
		},
	},
}

func (s *p0078TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, subsets(v.arg1))
	}
}

func TestP0078TestSuite(t *testing.T) {
	s := &p0078TestSuite{}
	suite.Run(t, s)
}
