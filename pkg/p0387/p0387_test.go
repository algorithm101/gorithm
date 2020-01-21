// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0387

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
		arg1:   "leetcode",
		target: 0,
	},
	{
		arg1:   "loveleetcode",
		target: 2,
	},
}

type p0387TestSuite struct {
	suite.Suite
}

func (s *p0387TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, firstUniqChar(v.arg1))
	}
}

func TestP0387TestSuite(t *testing.T) {
	s := &p0387TestSuite{}
	suite.Run(t, s)
}
