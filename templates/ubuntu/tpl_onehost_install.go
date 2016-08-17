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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
	ONEHOST_INSTALL_LOG = "/var/log/megam/megamcib/opennebulahost.log"
)

var ubuntuonehostinstall *UbuntuOneHostInstall

func init() {
	ubuntuonehostinstall = &UbuntuOneHostInstall{}
	templates.Register("UbuntuOneHostInstall", ubuntuonehostinstall)
}

type UbuntuOneHostInstall struct{}

func (tpl *UbuntuOneHostInstall) Render(p urknall.Package) {
	p.AddTemplate("onehost", &UbuntuOneHostInstallTemplate{})
}

func (tpl *UbuntuOneHostInstall) Options(t *templates.Template) {
}

func (tpl *UbuntuOneHostInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuOneHostInstall{},inputs)
}

type UbuntuOneHostInstallTemplate struct{}

func (m *UbuntuOneHostInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("depends",
		InstallPackages("build-essential genromfs autoconf libtool qemu-utils libvirt0 bridge-utils lvm2 ssh iproute iputils-arping make"),
	)

	pkg.AddCommands("one-node",
		InstallPackages("opennebula-node"),
	)

  pkg.AddCommands("node",
		Shell("sudo usermod -p $(echo oneadmin | openssl passwd -1 -stdin) oneadmin"),
	)
}
