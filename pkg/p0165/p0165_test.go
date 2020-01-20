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

package p0165

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0165TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	arg2   string
	target int
}

var values = []result{
	{
		arg1:   "0.1",
		arg2:   "1.1",
		target: -1,
	},
	{
		arg1:   "1.0.1",
		arg2:   "1",
		target: 1,
	},
	{
		arg1:   "7.5.2.4",
		arg2:   "7.5.3",
		target: -1,
	},
	{
		arg1:   "1.0",
		arg2:   "1",
		target: 0,
	},
}

func (s *p0165TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, compareVersion(v.arg1, v.arg2))
	}
}

func TestP0165TestSuite(t *testing.T) {
	s := &p0165TestSuite{}
	suite.Run(t, s)
}
