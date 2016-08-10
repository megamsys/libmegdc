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
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

var ubuntuenablecephaccess *UbuntuEnableCephAccess

func init() {
	ubuntuenablecephaccess = &UbuntuEnableCephAccess{}
	templates.Register("UbuntuEnableCephAccess", ubuntuenablecephaccess)
}

type UbuntuEnableCephAccess struct {
	cephuser       string
	clientuser     string
	clienthostname string
	clientip       string
  clientpwd       string
}

func (tpl *UbuntuEnableCephAccess) Render(p urknall.Package) {
	p.AddTemplate("enablecephaccess", &UbuntuEnableCephAccessTemplate{
		cephuser:       tpl.cephuser,
		clientuser:     tpl.clientuser,
		clienthostname: tpl.clienthostname,
		clientip:       tpl.clientip,
    clientpwd:       tpl.clientpwd,
	})
}

func (tpl *UbuntuEnableCephAccess) Options(t *templates.Template) {
	if cephuser, ok := t.Options["CEPHUSER"]; ok {
		tpl.cephuser = cephuser
	}
	if clientip, ok := t.Options["CLIENTIP"]; ok {
		tpl.clientip = clientip
	}

	if clientuser, ok := t.Options["CLIENTUSER"]; ok {
		tpl.clientuser = clientuser
	}
	if clienthostname, ok := t.Options["CLIENTHOST"]; ok {
		tpl.clienthostname = clienthostname
	}
  if clientpwd, ok := t.Options["CLIENTPASSWORD"]; ok {
    tpl.clientpwd = clientpwd
  }
}

func (tpl *UbuntuEnableCephAccess) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, &UbuntuEnableCephAccess{
		cephuser:       tpl.cephuser,
		clientuser:     tpl.clientuser,
		clienthostname: tpl.clienthostname,
		clientip:       tpl.clientip,
    clientpwd:       tpl.clientpwd,
	}, inputs)
}

type UbuntuEnableCephAccessTemplate struct {
	cephuser       string
	clientuser     string
	clienthostname string
  clientpwd      string
	clientip       string
}

func (m *UbuntuEnableCephAccessTemplate) Render(pkg urknall.Package) {

	ClientIP := m.clientip
	ClientHostName := m.clienthostname
	ClientUser := m.clientuser
	CephUser := m.cephuser
	CephHome :=  UserHomePrefix + m.cephuser
  ClientPassword := m.clientpwd
  ClientHome := m.clientuser

	pkg.AddCommands("install-sshpass",
		InstallPackages("sshpass"),
	)
	pkg.AddCommands("SSHPass",
		Shell("echo '"+ClientIP+"  "+ClientHostName+"' >> /etc/hosts"),
		WriteFile(CephHome +"/.ssh/config", KnownHostsList, CephUser, 0755),
		AsUser(CephHome, Shell("sshpass -p " + ClientPassword +" scp -o StrictHostKeyChecking=no " + CephHome + "/.ssh/id_rsa.pub " + ClientUser + "@" + ClientIP + ":" + ClientHome + "/.ssh/authorized_keys")),
	)
}
