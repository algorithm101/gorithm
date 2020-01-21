// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0806

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   string
	target []int
}

var values = []result{
	{
		arg1:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		arg2:   "abcdefghijklmnopqrstuvwxyz",
		target: []int{3, 60},
	},
	{
		arg1:   []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		arg2:   "bbbcccdddaaa",
		target: []int{2, 4},
	},
}

type p0806TestSuite struct {
	suite.Suite
}

func (s *p0806TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numberOfLines(v.arg1, v.arg2))
	}
}

func TestP0806TestSuite(t *testing.T) {
	s := &p0806TestSuite{}
	suite.Run(t, s)
}
