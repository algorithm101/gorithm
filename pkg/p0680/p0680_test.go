// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0680

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
		arg1:   "aba",
		target: true,
	},
	{
		arg1:   "abca",
		target: true,
	},
}

type p0680TestSuite struct {
	suite.Suite
}

func (s *p0680TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, validPalindrome(v.arg1))
	}
}

func TestP0680TestSuite(t *testing.T) {
	s := &p0680TestSuite{}
	suite.Run(t, s)
}
