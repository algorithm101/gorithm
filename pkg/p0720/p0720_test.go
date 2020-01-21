// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0720

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target string
}

var values = []result{
	// {
	// 	arg1:   []string{"w", "wo", "wor", "worl", "world"},
	// 	target: "world",
	// },
	// {
	// 	arg1:   []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
	// 	target: "apple",
	// },
	{
		arg1:   []string{"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast", "l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"},
		target: "breakfast",
	},
}

type p0720TestSuite struct {
	suite.Suite
}

func (s *p0720TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, longestWord(v.arg1))
	}
}

func TestP0720TestSuite(t *testing.T) {
	s := &p0720TestSuite{}
	suite.Run(t, s)
}
