// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0045

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
		arg1:   []int{2, 3, 1, 1, 4},
		target: 2,
	},
	{
		arg1:   []int{1, 3, 2},
		target: 2,
	},
	{
		arg1:   []int{3, 2, 1},
		target: 1,
	},
	{
		arg1:   []int{2, 3, 1},
		target: 1,
	},
}

type p0045TestSuite struct {
	suite.Suite
}

func (s *p0045TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, jump(v.arg1))
	}
}

func TestP0045TestSuite(t *testing.T) {
	s := &p0045TestSuite{}
	suite.Run(t, s)
}
