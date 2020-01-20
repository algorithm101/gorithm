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

package p0221

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0221TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]byte
	target int
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("10100"),
			[]byte("10111"),
			[]byte("11111"),
			[]byte("10010"),
		},
		target: 4,
	},
}

func (s *p0221TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maximalSquare(v.arg1))
	}
}

func TestP0221TestSuite(t *testing.T) {
	s := &p0221TestSuite{}
	suite.Run(t, s)
}
