// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0506

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []string
}

var values = []result{
	{
		arg1:   []int{5, 4, 3, 2, 1},
		target: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	},
}

type p0506TestSuite struct {
	suite.Suite
}

func (s *p0506TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findRelativeRanks(v.arg1))
	}
}

func TestP0506TestSuite(t *testing.T) {
	s := &p0506TestSuite{}
	suite.Run(t, s)
}
