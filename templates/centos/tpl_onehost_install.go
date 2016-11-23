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

package centos

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"

)

const (
	ONEHOST_INSTALL_LOG = "/var/log/megam/megamcib/opennebulahost.log"
)

var centosonehostinstall *CentosOneHostInstall

func init() {
	centosonehostinstall = &CentosOneHostInstall{}
	templates.Register("CentosOneHostInstall", centosonehostinstall)
}

type CentosOneHostInstall struct{}

func (tpl *CentosOneHostInstall) Render(p urknall.Package) {
	p.AddTemplate("onehost", &CentosOneHostInstallTemplate{})
}

func (tpl *CentosOneHostInstall) Options(t *templates.Template) {
}

func (tpl *CentosOneHostInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosOneHostInstall{},inputs)
}

type CentosOneHostInstallTemplate struct{}

func (m *CentosOneHostInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("repository",
	Shell("if [ -f /etc/os-release ] ; then  echo '[opennebula]' >/etc/yum.repos.d/opennebula.repo; echo 'name=opennebula' >>/etc/yum.repos.d/opennebula.repo; echo 'baseurl=http://downloads.opennebula.org/repo/4.12/CentOS/7/x86_64/' >>/etc/yum.repos.d/opennebula.repo; echo 'enabled=1' >>/etc/yum.repos.d/opennebula.repo; echo 'gpgcheck=0' >>/etc/yum.repos.d/opennebula.repo; else echo '[opennebula]' >/etc/yum.repos.d/opennebula.repo; echo 'name=opennebula' >>/etc/yum.repos.d/opennebula.repo; echo 'baseurl=http://downloads.opennebula.org/repo/4.14/CentOS/6/x86_64/' >>/etc/yum.repos.d/opennebula.repo; echo 'enabled=1' >>/etc/yum.repos.d/opennebula.repo; echo 'gpgcheck=0' >>/etc/yum.repos.d/opennebula.repo ; fi"),
	)
	pkg.AddCommands("depends",
		InstallPackages("build-essential genromfs autoconf libtool qemu-utils libvirt0 bridge-utils lvm2 ssh iproute iputils-arping make"),
	)

	pkg.AddCommands("verify",
		InstallPackages("qemu-system-x86 qemu-kvm cpu-checker"),
		And("grep -E 'svm|vmx' /proc/cpuinfo"),
	)

	pkg.AddCommands("one-node",
		InstallPackages("opennebula-node-kvm"),
	)
  pkg.AddCommands("node",
		Shell("sudo usermod -p $(echo oneadmin | openssl passwd -1 -stdin) oneadmin"),
	)
	pkg.AddCommands("vswitch",
		InstallPackages("openvswitch-common openvswitch-switch bridge-utils"),
	)
  pkg.AddCommands("start",
		Shell("systemctl start messagebus.service"),
   Shell("systemctl start libvirtd.service"),
	 Shell("systemctl start libvirtd.service"),
	 Shell("systemctl start nfs.service"),
	)
}
