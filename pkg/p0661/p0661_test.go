// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0661

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		target: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
	},
}

type p0661TestSuite struct {
	suite.Suite
}

func (s *p0661TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, imageSmoother(v.arg1))
	}
}

func TestP0661TestSuite(t *testing.T) {
	s := &p0661TestSuite{}
	suite.Run(t, s)
}
