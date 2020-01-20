// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0009

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target bool
}

var values = []result{
	{
		arg1:   121,
		target: true,
	},
	{
		arg1:   -121,
		target: false,
	},
	{
		arg1:   10,
		target: false,
	},
}

type p0009TestSuite struct {
	suite.Suite
}

func (s *p0009TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isPalindrome(v.arg1))
	}
}

func TestP0009TestSuite(t *testing.T) {
	s := &p0009TestSuite{}
	suite.Run(t, s)
}
