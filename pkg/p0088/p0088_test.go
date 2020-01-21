// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0088

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	arg3   []int
	arg4   int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 0, 0, 0},
		arg2:   3,
		arg3:   []int{2, 5, 6},
		arg4:   3,
		target: []int{1, 2, 2, 3, 5, 6},
	},
}

type p0088TestSuite struct {
	suite.Suite
}

func (s *p0088TestSuite) Test() {
	for _, v := range values {
		merge(v.arg1, v.arg2, v.arg3, v.arg4)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0088TestSuite(t *testing.T) {
	s := &p0088TestSuite{}
	suite.Run(t, s)
}
