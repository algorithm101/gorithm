// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0345

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target string
}

var values = []result{
	{
		arg1:   "hello",
		target: "holle",
	},
	{
		arg1:   "leetcode",
		target: "leotcede",
	},
	{
		arg1:   "aA",
		target: "Aa",
	},
}

type p0345TestSuite struct {
	suite.Suite
}

func (s *p0345TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reverseVowels(v.arg1))
	}
}

func TestP0345TestSuite(t *testing.T) {
	s := &p0345TestSuite{}
	suite.Run(t, s)
}
