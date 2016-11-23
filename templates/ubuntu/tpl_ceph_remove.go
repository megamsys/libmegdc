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
	"os"

	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

var ubuntucephremove *UbuntuCephRemove

func init() {
	ubuntucephremove = &UbuntuCephRemove{}
	templates.Register("UbuntuCephRemove", ubuntucephremove)
}

type UbuntuCephRemove struct {
	cephuser string
}

func (tpl *UbuntuCephRemove) Options(t *templates.Template) {
	if cephuser, ok := t.Options[CephUser]; ok {
		tpl.cephuser = cephuser
	}
}
func (tpl *UbuntuCephRemove) Render(p urknall.Package) {
	p.AddTemplate("ceph", &UbuntuCephRemoveTemplate{
        cephuser: tpl.cephuser,
})
}

func (tpl *UbuntuCephRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuCephRemove{
		cephuser: tpl.cephuser,
	},inputs)
}

type UbuntuCephRemoveTemplate struct {
	cephuser string
}

func (m *UbuntuCephRemoveTemplate) Render(pkg urknall.Package) {
	host, _ := os.Hostname()

	CephUser := m.cephuser
	pkg.AddCommands("cache-clean",
    Shell("rm -r /var/lib/urknall/ceph*"),
	)
	pkg.AddCommands("purgedata",
		AsUser(CephUser, Shell("ceph-deploy purgedata "+host+"")),
	)
	pkg.AddCommands("forgetKeys",
		AsUser(CephUser, Shell("ceph-deploy forgetkeys")),
	)
	pkg.AddCommands("purge",
		AsUser(CephUser, Shell("ceph-deploy purge "+host+"")),
	)
	pkg.AddCommands("rm-sshkey",
		AsUser(CephUser, Shell("rm -r ~/.ssh")),
	)
	pkg.AddCommands("remove",
		Shell("rm -r /var/lib/ceph/"),
		Shell("rm -r "+CephUser+"/ceph-cluster"),
		Shell("apt-get -y remove ceph-deploy ceph-common ceph-mds"),
		Shell("apt-get -y purge ceph-deploy ceph-common ceph-mds"),
		Shell("apt-get -y autoremove"),
		Shell("rm -r /run/ceph"),
		Shell("rm /var/log/upstart/ceph*"),
	)

}
