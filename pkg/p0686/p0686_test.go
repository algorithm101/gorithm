// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0686

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target int
}

var values = []result{
	{
		arg1:   "abcd",
		arg2:   "cdabcdab",
		target: 3,
	},
	{
		arg1:   "abcd",
		arg2:   "da",
		target: 2,
	},
	{
		arg1:   "abc",
		arg2:   "cabcabca",
		target: 4,
	},
	{
		arg1:   "a",
		arg2:   "aa",
		target: 2,
	},
}

type p0686TestSuite struct {
	suite.Suite
}

func (s *p0686TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, repeatedStringMatch(v.arg1, v.arg2))
	}
}

func TestP0686TestSuite(t *testing.T) {
	s := &p0686TestSuite{}
	suite.Run(t, s)
}
