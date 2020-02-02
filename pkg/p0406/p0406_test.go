// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0406

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target [][]int
}

var values = []result{
	{
		arg1: [][]int{
			{7, 0},
			{4, 4},
			{7, 1},
			{5, 0},
			{6, 1},
			{5, 2},
		},
		target: [][]int{
			{5, 0},
			{7, 0},
			{5, 2},
			{6, 1},
			{4, 4},
			{7, 1},
		},
	},
}

type p0406TestSuite struct {
	suite.Suite
}

func (s *p0406TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reconstructQueue(v.arg1))
	}
}

func TestP0406TestSuite(t *testing.T) {
	s := &p0406TestSuite{}
	suite.Run(t, s)
}
