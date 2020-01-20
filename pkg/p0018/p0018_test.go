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

package p0018

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target [][]int
}

var values = []result{
	{
		arg1: []int{1, 0, -1, 0, -2, 2},
		arg2: 0,
		target: [][]int{
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
			[]int{-1, 0, 0, 1},
		},
	},
}

type p0018TestSuite struct {
	suite.Suite
}

func (s *p0018TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, fourSum(v.arg1, v.arg2))
	}
}

func TestP0018TestSuite(t *testing.T) {
	s := &p0018TestSuite{}
	suite.Run(t, s)
}
