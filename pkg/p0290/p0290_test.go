// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0290

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
		arg1:   "abba",
		arg2:   "dog cat cat dog",
		target: true,
	},
	{
		arg1:   "abba",
		arg2:   "dog cat cat fish",
		target: false,
	},
	{
		arg1:   "aaaa",
		arg2:   "dog cat cat dog",
		target: false,
	},
	{
		arg1:   "abba",
		arg2:   "dog dog dog dog",
		target: false,
	},
}

type p0290TestSuite struct {
	suite.Suite
}

func (s *p0290TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, wordPattern(v.arg1, v.arg2))
	}
}

func TestP0290TestSuite(t *testing.T) {
	s := &p0290TestSuite{}
	suite.Run(t, s)
}
