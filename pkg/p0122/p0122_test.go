// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0122

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
		arg1:   []int{7, 1, 5, 3, 6, 4},
		target: 7,
	},
	{
		arg1:   []int{1, 2, 3, 4, 5},
		target: 4,
	},
	{
		arg1:   []int{7, 6, 4, 3, 1},
		target: 0,
	},
	{
		arg1:   []int{},
		target: 0,
	},
}

type p0122TestSuite struct {
	suite.Suite
}

func (s *p0122TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maxProfit(v.arg1))
	}
}

func TestP0122TestSuite(t *testing.T) {
	s := &p0122TestSuite{}
	suite.Run(t, s)
}
