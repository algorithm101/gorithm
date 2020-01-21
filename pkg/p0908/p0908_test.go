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

package p0908

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0908TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{1},
		arg2:   0,
		target: 0,
	},
	{
		arg1:   []int{0, 10},
		arg2:   2,
		target: 6,
	},
	{
		arg1:   []int{1, 3, 6},
		arg2:   3,
		target: 0,
	},
}

func (s *p0908TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, smallestRangeI(v.arg1, v.arg2))
	}
}

func TestP0908TestSuite(t *testing.T) {
	s := &p0908TestSuite{}
	suite.Run(t, s)
}
