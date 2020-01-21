// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0463

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{0, 1, 0, 0},
			[]int{1, 1, 1, 0},
			[]int{0, 1, 0, 0},
			[]int{1, 1, 0, 0},
		},
		target: 16,
	},
}

type p0463TestSuite struct {
	suite.Suite
}

func (s *p0463TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, islandPerimeter(v.arg1))
	}
}

func TestP0463TestSuite(t *testing.T) {
	s := &p0463TestSuite{}
	suite.Run(t, s)
}
