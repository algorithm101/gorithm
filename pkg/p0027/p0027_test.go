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

package p0027

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1    []int
	arg2    int
	target  int
	target2 []int
}

var values = []result{
	{
		arg1:    []int{3, 2, 2, 3},
		arg2:    3,
		target:  2,
		target2: []int{2, 2},
	},
	{
		arg1:    []int{0, 1, 2, 2, 3, 0, 4, 2},
		arg2:    2,
		target:  5,
		target2: []int{0, 1, 3, 0, 4},
	},
}

type p0027TestSuite struct {
	suite.Suite
}

func (s *p0027TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target2, v.arg1[:removeElement(v.arg1, v.arg2)])
	}
}

func TestP0027TestSuite(t *testing.T) {
	s := &p0027TestSuite{}
	suite.Run(t, s)
}
