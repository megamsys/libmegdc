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
	"fmt"
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
	ONE_INSTALL_LOG = "/var/log/megam/megamcib/opennebula.log"
	Slash = `\`
	HOSTNODE = "HostNode"
	sedUrl = `sed -i 's/^[ \t]*:one_xmlrpc:.*/:one_xmlrpc: http:\/\/%s:2633\/RPC2/' /etc/one/sunstone-server.conf`
	KnownHostsList = `
	ConnectTimeout 5
	Host *
	StrictHostKeyChecking no
	`
)




var ubuntuoneinstall *UbuntuOneInstall

func init() {
	ubuntuoneinstall = &UbuntuOneInstall{}
	templates.Register("UbuntuOneInstall", ubuntuoneinstall)
}

type UbuntuOneInstall struct{
	hostip string
}

func (tpl *UbuntuOneInstall) Render(p urknall.Package) {
	p.AddTemplate("onemaster", &UbuntuOneInstallTemplate{
			hostip:  tpl.hostip,
	})
}

func (tpl *UbuntuOneInstall) Options(t *templates.Template) {
	if host, ok := t.Options[HOST]; ok {
			tpl.hostip = host
	}
}

func (tpl *UbuntuOneInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, tpl,inputs)
}

type UbuntuOneInstallTemplate struct{
		hostip string
}

func (m *UbuntuOneInstallTemplate) Render(pkg urknall.Package) {

	ip := m.hostip
	pkg.AddCommands("install",
		InstallPackages("opennebula opennebula-sunstone"),
	)

	pkg.AddCommands("prepare",
		WriteFile("/var/lib/one/.ssh/config",KnownHostsList,"oneadmin", 0755),
		Shell("echo 'oneadmin ALL = (root) NOPASSWD:ALL' | sudo tee /etc/sudoers.d/oneadmin"),
		Shell("sudo chmod 0440 /etc/sudoers.d/oneadmin"),
		Shell("sed -i 's/^[ \t]*:host:.*/:host: "+ip+"/' /etc/one/sunstone-server.conf"),
	  Shell(fmt.Sprintf(sedUrl, ip)),
		Shell("service opennebula start"),
		Shell("service opennebula-sunstone start"),
	)
}
