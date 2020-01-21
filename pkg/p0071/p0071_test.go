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

package p0071

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0071TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target string
}

var values = []result{
	{
		arg1:   "/home/",
		target: "/home",
	},
	{
		arg1:   "/a/./b/../../c/",
		target: "/c",
	},
	{
		arg1:   "/home/foo/.ssh/../.ssh2/authorized_keys/",
		target: "/home/foo/.ssh2/authorized_keys",
	},
}

func (s *p0071TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, simplifyPath(v.arg1))
	}
}

func TestP0071TestSuite(t *testing.T) {
	s := &p0071TestSuite{}
	suite.Run(t, s)
}
