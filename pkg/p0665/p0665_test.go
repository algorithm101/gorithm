// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0665

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{4, 2, 3},
		target: true,
	},
	{
		arg1:   []int{4, 2, 1},
		target: false,
	},
	{
		arg1:   []int{3, 4, 2, 3},
		target: false,
	},
}

type p0665TestSuite struct {
	suite.Suite
}

func (s *p0665TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, checkPossibility(v.arg1))
	}
}

func TestP0665TestSuite(t *testing.T) {
	s := &p0665TestSuite{}
	suite.Run(t, s)
}
