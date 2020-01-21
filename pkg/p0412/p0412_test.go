// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0412

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target []string
}

var values = []result{
	{
		arg1: 15,
		target: []string{
			"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		},
	},
}

type p0412TestSuite struct {
	suite.Suite
}

func (s *p0412TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, fizzBuzz(v.arg1))
	}
}

func TestP0412TestSuite(t *testing.T) {
	s := &p0412TestSuite{}
	suite.Run(t, s)
}
