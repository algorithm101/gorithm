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

package p0867

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0867TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		target: [][]int{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	},
	{
		arg1: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		},
		target: [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	},
}

func (s *p0867TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, transpose(v.arg1))
	}
}

func TestP0867TestSuite(t *testing.T) {
	s := &p0867TestSuite{}
	suite.Run(t, s)
}
