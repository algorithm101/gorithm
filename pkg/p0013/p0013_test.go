// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0013

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
		arg1:   "III",
		target: 3,
	},
	{
		arg1:   "IV",
		target: 4,
	},
	{
		arg1:   "IX",
		target: 9,
	},
	{
		arg1:   "LVIII",
		target: 58,
	},
	{
		arg1:   "MCMXCIV",
		target: 1994,
	},
}

type p0013TestSuite struct {
	suite.Suite
}

func (s *p0013TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, romanToInt(v.arg1))
	}
}

func TestP0013TestSuite(t *testing.T) {
	s := &p0013TestSuite{}
	suite.Run(t, s)
}
