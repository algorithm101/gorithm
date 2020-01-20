// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0414

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
		arg1:   []int{3, 2, 1},
		target: 1,
	},
	{
		arg1:   []int{1, 2},
		target: 2,
	},
	{
		arg1:   []int{2, 2, 3, 1},
		target: 1,
	},
	{
		arg1:   []int{1, 2, -2147483648},
		target: -2147483648,
	},
}

type p0414TestSuite struct {
	suite.Suite
}

func (s *p0414TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, thirdMax(v.arg1))
	}
}

func TestP0414TestSuite(t *testing.T) {
	s := &p0414TestSuite{}
	suite.Run(t, s)
}
