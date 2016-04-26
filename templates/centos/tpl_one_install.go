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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
	ONE_INSTALL_LOG = "/var/log/megam/megamcib/opennebula.log"
	Slash = `\`
  Repo = `[opennebula]
name=opennebula
baseurl=http://downloads.opennebula.org/repo/4.14/CentOS/6/x86_64
enabled=1
gpgcheck=0
`
)

var centosoneinstall *CentosOneInstall

func init() {
	centosoneinstall = &CentosOneInstall{}
	templates.Register("CentosOneInstall", centosoneinstall)
}

type CentosOneInstall struct{}

func (tpl *CentosOneInstall) Render(p urknall.Package) {
	p.AddTemplate("one", &CentosOneInstallTemplate{})
}

func (tpl *CentosOneInstall) Options(t *templates.Template) {}

func (tpl *CentosOneInstall) Run(target urknall.Target) error {
	return urknall.Run(target, &CentosOneInstall{})
}

type CentosOneInstallTemplate struct{}

func (m *CentosOneInstallTemplate) Render(pkg urknall.Package) {

	ip := IP("")

	pkg.AddCommands("repository",
			WriteFile("/etc/yum.repos.d/opennebula.repo", Repo, "root", 0644),
	)

	pkg.AddCommands("one-install",
		InstallPackages("opennebula-server opennebula-sunstone "),
	)

	pkg.AddCommands("requires",
		Shell("echo 'oneadmin ALL = (root) NOPASSWD:ALL' | sudo tee /etc/sudoers.d/oneadmin"),
		Shell("sudo chmod 0440 /etc/sudoers.d/oneadmin"),
		//Shell("sudo rm /usr/share/one/install_gems"),
		//Shell("sudo cp /usr/share/megam/megdc/conf/install_gems /usr/share/one/install_gems"),
		//Shell("sudo chmod 755 /usr/share/one/install_gems"),
		Shell("wget http://ftp.ruby-lang.org/pub/ruby/2.1/ruby-2.1.1.tar.gz "),
		Shell("tar xvzf ruby-2.1.1.tar.gz"),
		Shell("cd ruby-2.1.1;./configure --prefix=/usr"),
		Shell("cd ruby-2.1.1;make"),
		Shell("cd ruby-2.1.1;make install"),
		Shell("sudo /usr/share/one/install_gems"),
		Shell("sed -i 's/^[ \t]*:host:.*/:host: "+ip+"/' /etc/one/sunstone-server.conf"),
		AsUser("oneadmin",Shell("echo 'TM_MAD=ssh' >/tmp/ds_tm_mad")),
		AsUser("oneadmin",Shell("onedatastore update 0 /tmp/ds_tm_mad")),
		AsUser("oneadmin",Shell("onedatastore update 1 /tmp/ds_tm_mad")),
		AsUser("oneadmin",Shell("onedatastore update 2 /tmp/ds_tm_mad")),
		Shell("sunstone-server start"),
		Shell("econe-server start"),
		Shell("sudo -H -u oneadmin bash -c 'one restart'"),
		Shell("service opennebula restart"),
	)
}
