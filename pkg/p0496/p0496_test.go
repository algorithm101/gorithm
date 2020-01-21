// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0496

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
		arg1:   []int{4, 1, 2},
		arg2:   []int{1, 3, 4, 2},
		target: []int{-1, 3, -1},
	},
	{
		arg1:   []int{2, 4},
		arg2:   []int{1, 2, 3, 4},
		target: []int{3, -1},
	},
}

type p0496TestSuite struct {
	suite.Suite
}

func (s *p0496TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, nextGreaterElement(v.arg1, v.arg2))
	}
}

func TestP0496TestSuite(t *testing.T) {
	s := &p0496TestSuite{}
	suite.Run(t, s)
}
