// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0844

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
		arg1:   "ab#c",
		arg2:   "ad#c",
		target: true,
	},
	{
		arg1:   "ab##",
		arg2:   "c#d#",
		target: true,
	},
	{
		arg1:   "a##c",
		arg2:   "#a#c",
		target: true,
	},
	{
		arg1:   "a#c",
		arg2:   "b",
		target: false,
	},
}

type p0844TestSuite struct {
	suite.Suite
}

func (s *p0844TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, backspaceCompare(v.arg1, v.arg2))
	}
}

func TestP0844TestSuite(t *testing.T) {
	s := &p0844TestSuite{}
	suite.Run(t, s)
}
