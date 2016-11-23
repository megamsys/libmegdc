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


var ubuntucephadminkeyring *UbuntuCephAdminKeyring

func init() {
	ubuntucephadminkeyring = &UbuntuCephAdminKeyring{}
	templates.Register("UbuntuCephAdminKeyring", ubuntucephadminkeyring)
}

type UbuntuCephAdminKeyring struct {
	cephuser string
  clienthostname string
}

func (tpl *UbuntuCephAdminKeyring) Options(t *templates.Template) {
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
  if clienthostname, ok := t.Options[CLIENTHOST]; ok {
    tpl.clienthostname = clienthostname
  }
}

func (tpl *UbuntuCephAdminKeyring) Render(p urknall.Package) {
	p.AddTemplate("add-osds", &UbuntuCephAdminKeyringTemplate{
		cephuser: tpl.cephuser,
    clienthostname: tpl.clienthostname,
	})
}

func (tpl *UbuntuCephAdminKeyring) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target,tpl,inputs)
}

type UbuntuCephAdminKeyringTemplate struct {
	cephuser string
  clienthostname string
}

func (m *UbuntuCephAdminKeyringTemplate) Render(pkg urknall.Package) {
  CephUser := m.cephuser
	CephHome := UserHomePrefix + CephUser
  ClientHostName := m.clienthostname
	if m.cephuser == "root" {
		CephHome = "/root"
	} else {
		CephHome = UserHomePrefix + m.cephuser
	}

	pkg.AddCommands("ceph-admin",
      AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy admin "+ ClientHostName )),
	)

}
