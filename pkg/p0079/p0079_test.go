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

package p0079

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0079TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]byte
	arg2   string
	target bool
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("ABCD"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		},
		arg2:   "ABCCED",
		target: true,
	},
	{
		arg1: [][]byte{
			[]byte("ABCD"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		},
		arg2:   "SEE",
		target: true,
	},
	{
		arg1: [][]byte{
			[]byte("ABCD"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		},
		arg2:   "ABCB",
		target: false,
	},
}

func (s *p0079TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, exist(v.arg1, v.arg2))
	}
}

func TestP0079TestSuite(t *testing.T) {
	s := &p0079TestSuite{}
	suite.Run(t, s)
}
