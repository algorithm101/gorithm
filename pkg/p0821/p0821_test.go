// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0821

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   byte
	target []int
}

var values = []result{
	{
		arg1:   "loveleetcode",
		arg2:   'e',
		target: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
	},
}

type p0821TestSuite struct {
	suite.Suite
}

func (s *p0821TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, shortestToChar(v.arg1, v.arg2))
	}
}

func TestP0821TestSuite(t *testing.T) {
	s := &p0821TestSuite{}
	suite.Run(t, s)
}
