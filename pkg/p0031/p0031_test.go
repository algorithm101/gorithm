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

package p0031

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3},
		target: []int{1, 3, 2},
	},
	{
		arg1:   []int{3, 2, 1},
		target: []int{1, 2, 3},
	},
	{
		arg1:   []int{1, 1, 5},
		target: []int{1, 5, 1},
	},
	{
		arg1:   []int{1, 3, 5, 4, 2},
		target: []int{1, 4, 2, 3, 5},
	},
}

type p0031TestSuite struct {
	suite.Suite
}

func (s *p0031TestSuite) Test() {
	for _, v := range values {
		arg := make([]int, len(v.arg1))
		copy(arg, v.arg1)
		nextPermutation(arg)
		s.Equal(v.target, arg)
	}
}

func TestP0031TestSuite(t *testing.T) {
	s := &p0031TestSuite{}
	suite.Run(t, s)
}
