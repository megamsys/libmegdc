/*
** Copyright [2013-2016] [Megam Systems]
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


import (
	//"os"
	//"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)


var ubuntuzapdisks *UbuntuZapDisks

func init() {
	ubuntuzapdisks = &UbuntuZapDisks{}
	templates.Register("UbuntuZapDisks", ubuntuzapdisks)
}

type UbuntuZapDisks struct {
	osds      []string
	cephuser string
  clienthostname string
}

func (tpl *UbuntuZapDisks) Options(t *templates.Template) {
	if osds, ok := t.Maps[OSDs]; ok {
		tpl.osds = osds
	}
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
  if clienthostname, ok := t.Options[CLIENTHOST]; ok {
    tpl.clienthostname = clienthostname
  }
}

func (tpl *UbuntuZapDisks) Render(p urknall.Package) {
	p.AddTemplate("zap-disk", &UbuntuZapDisksTemplate{
		osds:     tpl.osds,
		cephuser: tpl.cephuser,
    clienthostname: tpl.clienthostname,
	})
}

func (tpl *UbuntuZapDisks) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuZapDisks{
		osds:     tpl.osds,
		cephuser: tpl.cephuser,
		clienthostname: tpl.clienthostname,

	},inputs)
}

type UbuntuZapDisksTemplate struct {
  osds     []string
	cephuser string
  clienthostname string
}

func (m *UbuntuZapDisksTemplate) Render(pkg urknall.Package) {
  CephUser := m.cephuser
  ClientHostName := m.clienthostname
	if m.cephuser == "root" {
		CephHome = "/root"
	} else {
		CephHome = UserHomePrefix + m.cephuser
	}
  osds := ArraytoString(ClientHostName+":","",m.osds)

	pkg.AddCommands("zap-disks",
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy disk zap "+ osds )),
	)
}
