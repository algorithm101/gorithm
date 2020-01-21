// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0459

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target bool
}

var values = []result{
	{
		arg1:   "abab",
		target: true,
	},
	{
		arg1:   "aba",
		target: false,
	},
	{
		arg1:   "abcabcabcabc",
		target: true,
	},
}

type p0459TestSuite struct {
	suite.Suite
}

func (s *p0459TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, repeatedSubstringPattern(v.arg1))
	}
}

func TestP0459TestSuite(t *testing.T) {
	s := &p0459TestSuite{}
	suite.Run(t, s)
}
