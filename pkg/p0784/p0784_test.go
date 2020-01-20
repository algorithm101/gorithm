// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0784

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target []string
}

var values = []result{
	{
		arg1:   "a1b2",
		target: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
	},
	{
		arg1:   "3z4",
		target: []string{"3z4", "3Z4"},
	},
	{
		arg1:   "12345",
		target: []string{"12345"},
	},
}

type p0784TestSuite struct {
	suite.Suite
}

func (s *p0784TestSuite) Test() {
	for _, v := range values {
		sort.Strings(v.target)
		actual := letterCasePermutation(v.arg1)
		sort.Strings(actual)
		s.Equal(v.target, actual)
	}
}

func TestP0784TestSuite(t *testing.T) {
	s := &p0784TestSuite{}
	suite.Run(t, s)
}
