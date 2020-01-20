// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0581

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
		arg1:   []int{2, 6, 4, 8, 10, 9, 15},
		target: 5,
	},
	{
		arg1:   []int{3, 2},
		target: 2,
	},
	{
		arg1:   []int{2},
		target: 0,
	},
	{
		arg1:   []int{1, 2, 3, 4},
		target: 0,
	},
	{
		arg1:   []int{1, 3, 2, 2, 2},
		target: 4,
	},
}

type p0581TestSuite struct {
	suite.Suite
}

func (s *p0581TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findUnsortedSubarray(v.arg1))
	}
}

func TestP0581TestSuite(t *testing.T) {
	s := &p0581TestSuite{}
	suite.Run(t, s)
}
