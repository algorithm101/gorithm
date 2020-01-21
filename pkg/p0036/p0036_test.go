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

package p0036

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0036TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   [][]byte
	target bool
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("53..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		},
		target: true,
	},
	{
		arg1: [][]byte{
			[]byte("83..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		},
		target: false,
	},
}

func (s *p0036TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isValidSudoku(v.arg1))
	}
}

func TestP0036TestSuite(t *testing.T) {
	s := &p0036TestSuite{}
	suite.Run(t, s)
}
