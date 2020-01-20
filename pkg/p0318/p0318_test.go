// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0318

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
		target: 16,
	},
	{
		arg1:   []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
		target: 4,
	},
	{
		arg1:   []string{"a", "aa", "aaa", "aaaa"},
		target: 0,
	},
}

type p0318TestSuite struct {
	suite.Suite
}

func (s *p0318TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maxProduct(v.arg1))
	}
}

func TestP0318TestSuite(t *testing.T) {
	s := &p0318TestSuite{}
	suite.Run(t, s)
}
