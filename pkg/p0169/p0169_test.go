// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0169

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
		arg1:   []int{3, 2, 3},
		target: 3,
	},
	{
		arg1:   []int{2, 2, 1, 1, 1, 2, 2},
		target: 2,
	},
}

type p0169TestSuite struct {
	suite.Suite
}

func (s *p0169TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, majorityElement(v.arg1))
	}
}

func TestP0169TestSuite(t *testing.T) {
	s := &p0169TestSuite{}
	suite.Run(t, s)
}
