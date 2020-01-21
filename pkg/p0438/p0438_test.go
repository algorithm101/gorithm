// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0438

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target []int
}

var values = []result{
	{
		arg1:   "cbaebabacd",
		arg2:   "abc",
		target: []int{0, 6},
	},
	{
		arg1:   "abab",
		arg2:   "ab",
		target: []int{0, 1, 2},
	},
}

type p0438TestSuite struct {
	suite.Suite
}

func (s *p0438TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findAnagrams(v.arg1, v.arg2))
	}
}

func TestP0438TestSuite(t *testing.T) {
	s := &p0438TestSuite{}
	suite.Run(t, s)
}
