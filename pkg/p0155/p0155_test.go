// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0155

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p0155TestSuite struct {
	suite.Suite
}

func (s *p0155TestSuite) Test() {
	minStack := Constructor()
	minStack.Push(1)
	s.Equal(1, minStack.Top())
	s.Equal(1, minStack.GetMin())

	minStack.Push(2)
	s.Equal(2, minStack.Top())
	s.Equal(1, minStack.GetMin())

	minStack.Push(-2)
	s.Equal(-2, minStack.Top())
	s.Equal(-2, minStack.GetMin())

	minStack.Pop()
	s.Equal(2, minStack.Top())
	s.Equal(1, minStack.GetMin())
}

func TestP0155TestSuite(t *testing.T) {
	s := &p0155TestSuite{}
	suite.Run(t, s)
}
