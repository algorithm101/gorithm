// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0219

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 1},
		arg2:   3,
		target: true,
	},
	{
		arg1:   []int{1, 0, 1, 1},
		arg2:   1,
		target: true,
	},
	{
		arg1:   []int{1, 2, 3, 1, 2, 3},
		arg2:   2,
		target: false,
	},
}

type p0219TestSuite struct {
	suite.Suite
}

func (s *p0219TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, containsNearbyDuplicate(v.arg1, v.arg2))
	}
}

func TestP0219TestSuite(t *testing.T) {
	s := &p0219TestSuite{}
	suite.Run(t, s)
}
