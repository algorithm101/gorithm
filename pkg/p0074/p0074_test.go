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

package p0074

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0074TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		arg2:   3,
		target: true,
	},
	{
		arg1: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		arg2:   13,
		target: false,
	},
}

func (s *p0074TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, searchMatrix(v.arg1, v.arg2))
	}
}

func TestP0074TestSuite(t *testing.T) {
	s := &p0074TestSuite{}
	suite.Run(t, s)
}
