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

package p0004

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   []int
	target float64
}

var values = []result{
	{
		arg1:   []int{1, 3},
		arg2:   []int{2},
		target: 2,
	},
	{
		arg1:   []int{1, 2},
		arg2:   []int{3, 4},
		target: 2.5,
	},
	{
		arg1:   []int{},
		arg2:   []int{1, 2, 3, 4, 5},
		target: 3,
	},
	{
		arg1:   []int{3},
		arg2:   []int{1, 2},
		target: 2,
	},
}

type p0004TestSuite struct {
	suite.Suite
}

func (s *p0004TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMedianSortedArrays(v.arg1, v.arg2))
	}
}

func TestP0004TestSuite(t *testing.T) {
	s := &p0004TestSuite{}
	suite.Run(t, s)
}
