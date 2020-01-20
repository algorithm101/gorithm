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

package p0200

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0200TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]byte
	target int
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		},
		target: 1,
	},
	{
		arg1: [][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		},
		target: 3,
	},
}

func (s *p0200TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numIslands(v.arg1))
	}
}

func TestP0200TestSuite(t *testing.T) {
	s := &p0200TestSuite{}
	suite.Run(t, s)
}
