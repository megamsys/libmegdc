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
  ppa_jdk = `sudo apt-add-repository -y ppa:openjdk-r/ppa
  `
)
var ubuntujavainstall *UbuntuJavaInstall

func init() {
	ubuntujavainstall = &UbuntuJavaInstall{}
	templates.Register("UbuntuJavaInstall", ubuntujavainstall)
}

type UbuntuJavaInstall struct {
}

func (tpl *UbuntuJavaInstall) Options(t *templates.Template) {
}

func (tpl *UbuntuJavaInstall) Render(p urknall.Package) {
	p.AddTemplate("java", &UbuntuJavaInstallTemplate{})
}

func (tpl *UbuntuJavaInstall) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target, tpl, inputs)
}

type UbuntuJavaInstallTemplate struct {
}

func (m *UbuntuJavaInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("java-install",
    Shell(ppa),
    Shell("sudo apt-get -y update"),
		InstallPackages("openjdk-8-jdk"),
	)
}
