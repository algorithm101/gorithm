// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0811

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target []string
}

var values = []result{
	{
		arg1:   []string{"9001 discuss.leetcode.com"},
		target: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
	},
	{
		arg1:   []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		target: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
	},
}

type p0811TestSuite struct {
	suite.Suite
}

func (s *p0811TestSuite) Test() {
	for _, v := range values {
		sort.Strings(v.target)
		actual := subdomainVisits(v.arg1)
		sort.Strings(actual)
		s.Equal(v.target, actual)
	}
}

func TestP0811TestSuite(t *testing.T) {
	s := &p0811TestSuite{}
	suite.Run(t, s)
}
