// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0628

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
		arg1:   []int{1, 2, 3},
		target: 6,
	},
	{
		arg1:   []int{1, 2, 3, 4},
		target: 24,
	},
	{
		arg1:   []int{-3, -8, 0, 7, 2, 4},
		target: 168,
	},
}

type p0628TestSuite struct {
	suite.Suite
}

func (s *p0628TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maximumProduct(v.arg1))
	}
}

func TestP0628TestSuite(t *testing.T) {
	s := &p0628TestSuite{}
	suite.Run(t, s)
}
