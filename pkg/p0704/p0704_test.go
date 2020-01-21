// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0704

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{-1, 0, 3, 5, 9, 12},
		arg2:   9,
		target: 4,
	},
	{
		arg1:   []int{-1, 0, 3, 5, 9, 12},
		arg2:   2,
		target: -1,
	},
	{
		arg1:   []int{-1, 0, 3, 5, 9, 12},
		arg2:   -1,
		target: 0,
	},
	{
		arg1:   []int{-1, 0, 3, 5, 9, 12},
		arg2:   12,
		target: 5,
	},
}

type p0704TestSuite struct {
	suite.Suite
}

func (s *p0704TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, search(v.arg1, v.arg2))
	}
}

func TestP0704TestSuite(t *testing.T) {
	s := &p0704TestSuite{}
	suite.Run(t, s)
}
