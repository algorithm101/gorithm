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

package p0304

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0304TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{10, 9, 2, 5, 3, 7, 101, 18},
		target: 4,
	},
	{
		arg1:   []int{2, 2},
		target: 1,
	},
}

func (s *p0304TestSuite) Test() {
	m := Constructor([][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	})
	s.Equal(8, m.SumRegion(2, 1, 4, 3))
	s.Equal(11, m.SumRegion(1, 1, 2, 2))
	s.Equal(12, m.SumRegion(1, 2, 2, 4))
}

func TestP0304TestSuite(t *testing.T) {
	s := &p0304TestSuite{}
	suite.Run(t, s)
}
