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

package p0006

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   int
	target string
}

var values = []result{
	{
		arg1:   "PAYPALISHIRING",
		arg2:   3,
		target: "PAHNAPLSIIGYIR",
	},
	{
		arg1:   "PAYPALISHIRING",
		arg2:   4,
		target: "PINALSIGYAHRPI",
	},
	{
		arg1:   "abcdefghijklmnopqrstuvwxyz",
		arg2:   4,
		target: "agmsybfhlnrtxzceikoquwdjpv",
	},
	{
		arg1:   "A",
		arg2:   1,
		target: "A",
	},
}

type p0006TestSuite struct {
	suite.Suite
}

func (s *p0006TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, convert(v.arg1, v.arg2))
	}
}

func TestP0006TestSuite(t *testing.T) {
	s := &p0006TestSuite{}
	suite.Run(t, s)
}
