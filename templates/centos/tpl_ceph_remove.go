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

package centos

import (
	"os"

	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

var centoscephremove *CentosCephRemove

func init() {
	centoscephremove = &CentosCephRemove{}
	templates.Register("CentosCephRemove", centoscephremove)
}

type CentosCephRemove struct {
	cephuser string
}

func (tpl *CentosCephRemove) Options(t *templates.Template) {
	if cephuser, ok := t.Options[CephUser]; ok {
		tpl.cephuser = cephuser
	}
}
func (tpl *CentosCephRemove) Render(p urknall.Package) {
	p.AddTemplate("ceph", &CentosCephRemoveTemplate{
        cephuser: tpl.cephuser,
})
}

func (tpl *CentosCephRemove) Run(target urknall.Target) error {
	return urknall.Run(target, &CentosCephRemove{
		cephuser: tpl.cephuser,
	})
}

type CentosCephRemoveTemplate struct {
	cephuser string
}

func (m *CentosCephRemoveTemplate) Render(pkg urknall.Package) {
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
		AsUser(CephUser, Shell("rm -rf ~/.ssh")),
	)
	pkg.AddCommands("remove",
		Shell("rm -rf /var/lib/ceph/"),
		Shell("rm -rf "+CephUser+"/ceph-cluster"),
		Shell("yum -y remove ceph-deploy ceph-common ceph-mds"),
		Shell("yum -y purge ceph-deploy ceph-common ceph-mds"),
		Shell("yum -y autoremove"),
		Shell("rm -rf /run/ceph"),
		Shell("rm /var/log/upstart/ceph*"),
	)

}
