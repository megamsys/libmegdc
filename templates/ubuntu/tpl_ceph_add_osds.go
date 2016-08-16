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
)


var ubuntuaddosds *UbuntuAddOsds

func init() {
	ubuntuaddosds = &UbuntuAddOsds{}
	templates.Register("UbuntuAddOsds", ubuntuaddosds)
}

type UbuntuAddOsds struct {
	osds      []string
	cephuser string
  clienthostname string
}

func (tpl *UbuntuAddOsds) Options(t *templates.Template) {
	if osds, ok := t.Maps[OSDs]; ok {
		tpl.osds = osds
	}
	if cephuser, ok := t.Options[CEPHUSER]; ok {
		tpl.cephuser = cephuser
	}
  if clienthostname, ok := t.Options[CLIENTHOST]; ok {
    tpl.clienthostname = clienthostname
  }
}

func (tpl *UbuntuAddOsds) Render(p urknall.Package) {
	p.AddTemplate("zap-disk", &UbuntuAddOsdsTemplate{
		osds:     tpl.osds,
		cephuser: tpl.cephuser,
    clienthostname: tpl.clienthostname,
	})
}

func (tpl *UbuntuAddOsds) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuAddOsds{
		osds:     tpl.osds,
		cephuser: tpl.cephuser,
		clienthostname: tpl.clienthostname,

	},inputs)
}

type UbuntuAddOsdsTemplate struct {
  osds     []string
	cephuser string
  clienthostname string
}

func (m *UbuntuAddOsdsTemplate) Render(pkg urknall.Package) {
  CephUser := m.cephuser
	CephHome := UserHomePrefix + CephUser
  ClientHostName := m.clienthostname

  osds := ArraytoString(ClientHostName+":","",m.osds)

	pkg.AddCommands("add-osds",
    AsUser("root",Shell("sudo chown -R "+ CephUser +":"+ CephUser +" /etc/ceph/ceph.client.admin.keyring")),
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy osd prepare "+ osds )),
	)
	_ = RemoveAllCaches("/var/lib/urknall")
}
