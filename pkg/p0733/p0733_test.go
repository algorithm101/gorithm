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

package p0733

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0733TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	arg2   int
	arg3   int
	arg4   int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		},
		arg2: 1,
		arg3: 1,
		arg4: 2,
		target: [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		},
	},
}

func (s *p0733TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, floodFill(v.arg1, v.arg2, v.arg3, v.arg4))
	}
}

func TestP0733TestSuite(t *testing.T) {
	s := &p0733TestSuite{}
	suite.Run(t, s)
}
