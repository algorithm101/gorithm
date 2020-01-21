// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0189

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 4, 5, 6, 7},
		arg2:   3,
		target: []int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		arg1:   []int{-1, -100, 3, 99},
		arg2:   2,
		target: []int{3, 99, -1, -100},
	},
	{
		arg1:   []int{-1},
		arg2:   2,
		target: []int{-1},
	},
}

type p0189TestSuite struct {
	suite.Suite
}

func (s *p0189TestSuite) Test() {
	for _, v := range values {
		rotate(v.arg1, v.arg2)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0189TestSuite(t *testing.T) {
	s := &p0189TestSuite{}
	suite.Run(t, s)
}
