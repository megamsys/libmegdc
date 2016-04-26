/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package templates

import (

	//"fmt"
	"gopkg.in/check.v1"
"testing"
	//log "github.com/Sirupsen/logrus"
	//"github.com/megamsys/urknall"
)
type S struct{}

var _ = check.Suite(&S{})

func Test(t *testing.T) { check.TestingT(t) }
func (s *S) TestNewTemplate(c *check.C) {
  b :=NewTemplate()
c.Assert(b, check.NotNil )
}

func (s *S) TestRunTemplate(c *check.C) {
  a := Template{Name: "example", Host: "192.168.1.243", UserName: "root", Password: "xxx"}
  c.Assert(a, check.NotNil)

  d :=a.Run()
  c.Assert(d, check.NotNil)
}
