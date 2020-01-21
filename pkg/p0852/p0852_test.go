// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0852

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
		arg1:   []int{0, 1, 0},
		target: 1,
	},
	{
		arg1:   []int{0, 2, 1, 0},
		target: 1,
	},
}

type p0852TestSuite struct {
	suite.Suite
}

func (s *p0852TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, peakIndexInMountainArray(v.arg1))
	}
}

func TestP0852TestSuite(t *testing.T) {
	s := &p0852TestSuite{}
	suite.Run(t, s)
}
