// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0110

import (
	"testing"

	"algorithm101.io/gorithm/pkg/utils"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{3, 9, 20, 0xFFFFFFFF, 0xFFFFFFFF, 15, 7},
		target: true,
	},
	{
		arg1:   []int{1, 2, 2, 3, 3, 0xFFFFFFFF, 0xFFFFFFFF, 4, 4},
		target: false,
	},
}

type p0110TestSuite struct {
	suite.Suite
}

func (s *p0110TestSuite) Test() {
	for _, v := range values {
		n := utils.Trees(v.arg1)
		s.Equal(v.target, isBalanced(n))
	}
}

func TestP0110TestSuite(t *testing.T) {
	s := &p0110TestSuite{}
	suite.Run(t, s)
}
