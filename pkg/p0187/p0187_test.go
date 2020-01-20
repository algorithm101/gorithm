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

package p0187

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0187TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   string
	target []string
}

var values = []result{
	{
		arg1:   "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		target: []string{"AAAAACCCCC", "CCCCCAAAAA"},
	},
}

func (s *p0187TestSuite) Test() {
	for _, v := range values {
		r := findRepeatedDnaSequences(v.arg1)
		sort.Strings(r)
		sort.Strings(v.target)
		s.Equal(v.target, r)
	}
}

func TestP0187TestSuite(t *testing.T) {
	s := &p0187TestSuite{}
	suite.Run(t, s)
}
