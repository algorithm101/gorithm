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

package p0043

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0043TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	arg2   string
	target string
}

var values = []result{
	{
		arg1:   "2",
		arg2:   "3",
		target: "6",
	},
	{
		arg1:   "123",
		arg2:   "456",
		target: "56088",
	},
	{
		arg1:   "0",
		arg2:   "0",
		target: "0",
	},
}

func (s *p0043TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, multiply(v.arg1, v.arg2))
	}
}

func TestP0043TestSuite(t *testing.T) {
	s := &p0043TestSuite{}
	suite.Run(t, s)
}
