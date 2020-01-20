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

package p0299

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0299TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	arg2   string
	target string
}

var values = []result{
	{
		arg1:   "1807",
		arg2:   "7810",
		target: "1A3B",
	},
	{
		arg1:   "1123",
		arg2:   "0111",
		target: "1A1B",
	},
	{
		arg1:   "1122",
		arg2:   "2211",
		target: "0A4B",
	},
}

func (s *p0299TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, getHint(v.arg1, v.arg2))
	}
}

func TestP0299TestSuite(t *testing.T) {
	s := &p0299TestSuite{}
	suite.Run(t, s)
}
