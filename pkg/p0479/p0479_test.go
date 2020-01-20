// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0479

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   2,
		target: 987,
	},
	{
		arg1:   6,
		target: 1218,
	},
	{
		arg1:   1,
		target: 9,
	},
}

type p0479TestSuite struct {
	suite.Suite
}

func (s *p0479TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestPalindrome(v.arg1))
	}
}

func TestP0479TestSuite(t *testing.T) {
	s := &p0479TestSuite{}
	suite.Run(t, s)
}
