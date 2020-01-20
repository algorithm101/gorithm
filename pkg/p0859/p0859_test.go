// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0859

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target bool
}

var values = []result{
	{
		arg1:   "ab",
		arg2:   "ba",
		target: true,
	},
	{
		arg1:   "ab",
		arg2:   "ab",
		target: false,
	},
	{
		arg1:   "aa",
		arg2:   "aa",
		target: true,
	},
	{
		arg1:   "aaaaaaabc",
		arg2:   "aaaaaaacb",
		target: true,
	},
	{
		arg1:   "",
		arg2:   "aa",
		target: false,
	},
}

type p0859TestSuite struct {
	suite.Suite
}

func (s *p0859TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, buddyStrings(v.arg1, v.arg2))
	}
}

func TestP0859TestSuite(t *testing.T) {
	s := &p0859TestSuite{}
	suite.Run(t, s)
}
