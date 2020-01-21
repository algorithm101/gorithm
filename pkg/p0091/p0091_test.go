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

package p0091

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0091TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "12",
		target: 2,
	},
	{
		arg1:   "226",
		target: 3,
	},
	{
		arg1:   "0",
		target: 0,
	},
	{
		arg1:   "00",
		target: 0,
	},
	{
		arg1:   "01",
		target: 0,
	},
	{
		arg1:   "10",
		target: 1,
	},
}

func (s *p0091TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numDecodings(v.arg1))
	}
}

func TestP0091TestSuite(t *testing.T) {
	s := &p0091TestSuite{}
	suite.Run(t, s)
}
