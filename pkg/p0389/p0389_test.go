// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0389

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target byte
}

var values = []result{
	{
		arg1:   "abcd",
		arg2:   "abcde",
		target: 'e',
	},
}

type p0389TestSuite struct {
	suite.Suite
}

func (s *p0389TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findTheDifference(v.arg1, v.arg2))
	}
}

func TestP0389TestSuite(t *testing.T) {
	s := &p0389TestSuite{}
	suite.Run(t, s)
}
