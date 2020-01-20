// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0409

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "abccccdd",
		target: 7,
	},
}

type p0409TestSuite struct {
	suite.Suite
}

func (s *p0409TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, longestPalindrome(v.arg1))
	}
}

func TestP0409TestSuite(t *testing.T) {
	s := &p0409TestSuite{}
	suite.Run(t, s)
}
