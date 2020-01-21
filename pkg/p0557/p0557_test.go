// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0557

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
		arg1:   "Let's take LeetCode contest",
		target: "s'teL ekat edoCteeL tsetnoc",
	},
}

type p0557TestSuite struct {
	suite.Suite
}

func (s *p0557TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reverseWords(v.arg1))
	}
}

func TestP0557TestSuite(t *testing.T) {
	s := &p0557TestSuite{}
	suite.Run(t, s)
}
