// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0888

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 1},
		arg2:   []int{2, 2},
		target: []int{1, 2},
	},
	{
		arg1:   []int{1, 2},
		arg2:   []int{2, 3},
		target: []int{1, 2},
	},
	{
		arg1:   []int{2},
		arg2:   []int{1, 3},
		target: []int{2, 3},
	},
	{
		arg1:   []int{1, 2, 5},
		arg2:   []int{2, 4},
		target: []int{5, 4},
	},
}

type p0888TestSuite struct {
	suite.Suite
}

func (s *p0888TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, fairCandySwap(v.arg1, v.arg2))
	}
}

func TestP0888TestSuite(t *testing.T) {
	s := &p0888TestSuite{}
	suite.Run(t, s)
}
