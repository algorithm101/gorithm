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

package p0040

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0040TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target [][]int
}

var values = []result{
	{
		arg1: []int{10, 1, 2, 7, 6, 1, 5},
		arg2: 8,
		target: [][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		},
	},
	{
		arg1: []int{2, 5, 2, 1, 2},
		arg2: 5,
		target: [][]int{
			{1, 2, 2},
			{5},
		},
	},
}

func (s *p0040TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, combinationSum2(v.arg1, v.arg2))
	}
}

func TestP0040TestSuite(t *testing.T) {
	s := &p0040TestSuite{}
	suite.Run(t, s)
}
