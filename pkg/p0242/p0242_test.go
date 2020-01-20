// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0242

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
		arg1:   "anagram",
		arg2:   "nagaram",
		target: true,
	},
	{
		arg1:   "rat",
		arg2:   "car",
		target: false,
	},
}

type p0242TestSuite struct {
	suite.Suite
}

func (s *p0242TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isAnagram(v.arg1, v.arg2))
	}
}

func TestP0242TestSuite(t *testing.T) {
	s := &p0242TestSuite{}
	suite.Run(t, s)
}
