// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0203

import (
	"testing"

	"algorithm101.io/gorithm/pkg/utils"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 6, 3, 4, 5, 6},
		arg2:   6,
		target: []int{1, 2, 3, 4, 5},
	},
}

type p0203TestSuite struct {
	suite.Suite
}

func (s *p0203TestSuite) Test() {
	for _, v := range values {
		utils.AssertLinkEqual(s.Suite, utils.Links(v.target), removeElements(utils.Links(v.arg1), v.arg2))
	}
}

func TestP0203TestSuite(t *testing.T) {
	s := &p0203TestSuite{}
	suite.Run(t, s)
}
