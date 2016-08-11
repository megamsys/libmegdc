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

package ubuntu
/*
import (
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"gopkg.in/check.v1"
	//"testing"
	//constants"github.com/megamsys/libgo/utils"
	//"github.com/megamsys/libgo/events"
	"io"
	"os"
	"bytes"
)

var _ = check.Suite(&S{})
var ubuntudeletepartition *UbuntuDeletePartitions

//func Test(t *testing.T) { check.TestingT(t) }

func (s *S) TestNewTemplate(c *check.C) {
	b := templates.NewTemplate()
	c.Assert(b, check.NotNil)
}

func (s *S) TestZapTemplate(c *check.C) {
	ubuntudeletepartition = &UbuntuDeletePartitions{}
	c.Assert(ubuntudeletepartition, check.NotNil)
	templates.Register("UbuntuDeletePartitions", ubuntudeletepartition)

	a := templates.Template{Name: "UbuntuDeletePartitions", Host: "192.168.0.103", UserName: "rajthilak", Password: "team4megam"}
	m := make(map[string]string)
	m["Disk"] = "sda"
	a.Options = m
	c.Assert(a, check.NotNil)
	abc := []string{"email=rajeshr@megam.io"}
	var t io.Writer
	var outBuffer bytes.Buffer
	writer := io.MultiWriter(t, &outBuffer, os.Stdout)
	d := a.Run(writer, abc)
	fmt.Println("-------run-----")

	fmt.Println(outBuffer.String())

	fmt.Println(d)
	c.Assert(d, check.NotNil)
}

/*
func SetEventsWrap() error {
    mi := make(map[string]map[string]string)
    mp := make(map[string]string)
    mp["scylla_host"] = "192.168.0.116"
    mp["scylla_keyspace"] = "megdc"
    mi[constants.META] = mp
    return events.NewWrap(mi)
}
*/
