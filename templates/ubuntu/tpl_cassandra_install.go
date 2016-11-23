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

const (
  ppa_cql = `sudo echo "deb http://debian.datastax.com/datastax-ddc 3.7 main" | sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list
  `
)


var ubuntucassandrainstall *UbuntuCassandraInstall

func init() {
	ubuntucassandrainstall = &UbuntuCassandraInstall{}
	templates.Register("UbuntuCassandraInstall", ubuntucassandrainstall)
}

type UbuntuCassandraInstall struct {
}

func (tpl *UbuntuCassandraInstall) Options(t *templates.Template) {

}

func (tpl *UbuntuCassandraInstall) Render(p urknall.Package) {
	p.AddTemplate("cassandra", &UbuntuCassandraInstallTemplate{})
}

func (tpl *UbuntuCassandraInstall) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target, tpl, inputs)
}

type UbuntuCassandraInstallTemplate struct {
}

func (m *UbuntuCassandraInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("java-install",
    Shell(ppa_cql),
    Shell("sudo curl -L https://debian.datastax.com/debian/repo_key | sudo apt-key add -"),
    Shell("sudo apt-get -y update"),
		InstallPackages("datastax-ddc"),
	)
}
