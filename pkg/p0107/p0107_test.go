// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0107

import (
	"testing"

	"algorithm101.io/gorithm/pkg/utils"
	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target [][]int
}

var values = []result{
	{
		arg1: []int{3, 9, 20, 0xFFFFFFFF, 0xFFFFFFFF, 15, 7},
		target: [][]int{
			[]int{15, 7},
			[]int{9, 20},
			[]int{3},
		},
	},
}

type p0107TestSuite struct {
	suite.Suite
}

func (s *p0107TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, levelOrderBottom(utils.Trees(v.arg1)))
	}
}

func TestP0107TestSuite(t *testing.T) {
	s := &p0107TestSuite{}
	suite.Run(t, s)
}
