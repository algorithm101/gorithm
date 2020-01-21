// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0448

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{4, 3, 2, 7, 8, 2, 3, 1},
		target: []int{5, 6},
	},
}

type p0448TestSuite struct {
	suite.Suite
}

func (s *p0448TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findDisappearedNumbers(v.arg1))
	}
}

func TestP0448TestSuite(t *testing.T) {
	s := &p0448TestSuite{}
	suite.Run(t, s)
}
