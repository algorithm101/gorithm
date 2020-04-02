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

package p1365

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p1365TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{6, 5, 4, 8},
		target: []int{2, 1, 0, 3},
	},
	{
		arg1:   []int{7, 7, 7, 7},
		target: []int{0, 0, 0, 0},
	},
	{
		arg1:   []int{5, 0, 10, 0, 10, 6},
		target: []int{2, 0, 4, 0, 4, 3},
	},
}

func (s *p1365TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, smallerNumbersThanCurrent(v.arg1))
	}
}

func TestP1365TestSuite(t *testing.T) {
	s := &p1365TestSuite{}
	suite.Run(t, s)
}
