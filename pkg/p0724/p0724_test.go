// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0724

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
		arg1:   []int{1, 7, 3, 6, 5, 6},
		target: 3,
	},
	{
		arg1:   []int{1, 2, 3},
		target: -1,
	},
	{
		arg1:   []int{-1, -1, -1, 0, 1, 1},
		target: 0,
	},
}

type p0724TestSuite struct {
	suite.Suite
}

func (s *p0724TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, pivotIndex(v.arg1))
	}
}

func TestP0724TestSuite(t *testing.T) {
	s := &p0724TestSuite{}
	suite.Run(t, s)
}
