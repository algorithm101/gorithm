// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0453

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3},
		target: 3,
	},
}

type p0453TestSuite struct {
	suite.Suite
}

func (s *p0453TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minMoves(v.arg1))
	}
}

func TestP0453TestSuite(t *testing.T) {
	s := &p0453TestSuite{}
	suite.Run(t, s)
}
