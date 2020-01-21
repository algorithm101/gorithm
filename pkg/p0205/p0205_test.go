// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0205

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
		arg1:   "egg",
		arg2:   "add",
		target: true,
	},
	{
		arg1:   "foo",
		arg2:   "bar",
		target: false,
	},
	{
		arg1:   "paper",
		arg2:   "title",
		target: true,
	},
	{
		arg1:   "ab",
		arg2:   "aa",
		target: false,
	},
}

type p0205TestSuite struct {
	suite.Suite
}

func (s *p0205TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isIsomorphic(v.arg1, v.arg2))
	}
}

func TestP0205TestSuite(t *testing.T) {
	s := &p0205TestSuite{}
	suite.Run(t, s)
}
