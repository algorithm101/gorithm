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

package p0034

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0034TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []result{
	{
		arg1:   []int{5, 7, 7, 8, 8, 10},
		arg2:   8,
		target: []int{3, 4},
	},
	{
		arg1:   []int{5, 7, 7, 8, 8, 10},
		arg2:   6,
		target: []int{-1, -1},
	},
	{
		arg1:   []int{},
		arg2:   0,
		target: []int{-1, -1},
	},
	{
		arg1:   []int{1},
		arg2:   0,
		target: []int{-1, -1},
	},
}

func (s *p0034TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, searchRange(v.arg1, v.arg2))
	}
}

func TestP0034TestSuite(t *testing.T) {
	s := &p0034TestSuite{}
	suite.Run(t, s)
}
