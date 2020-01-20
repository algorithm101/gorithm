// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0313

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   []int
	target int
}

var values = []result{
	{
		arg1:   12,
		arg2:   []int{2, 7, 13, 19},
		target: 32,
	},
	{
		arg1:   15,
		arg2:   []int{3, 5, 7, 11, 19, 23, 29, 41, 43, 47},
		target: 35,
	},
}

type p0313TestSuite struct {
	suite.Suite
}

func (s *p0313TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, nthSuperUglyNumber(v.arg1, v.arg2))
	}
}

func TestP0313TestSuite(t *testing.T) {
	s := &p0313TestSuite{}
	suite.Run(t, s)
}
