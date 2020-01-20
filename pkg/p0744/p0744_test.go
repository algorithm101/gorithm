// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0744

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []byte
	arg2   byte
	target byte
}

var values = []result{
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'a',
		target: 'c',
	},
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'c',
		target: 'f',
	},
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'd',
		target: 'f',
	},
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'g',
		target: 'j',
	},
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'j',
		target: 'c',
	},
	{
		arg1:   []byte{'c', 'f', 'j'},
		arg2:   'k',
		target: 'c',
	},
}

type p0744TestSuite struct {
	suite.Suite
}

func (s *p0744TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, nextGreatestLetter(v.arg1, v.arg2))
	}
}

func TestP0744TestSuite(t *testing.T) {
	s := &p0744TestSuite{}
	suite.Run(t, s)
}
