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
  ppa = `sudo apt-add-repository ppa:brightbox/ruby-ng
  `
)
var ubunturubyinstall *UbuntuRubyInstall

func init() {
	ubunturubyinstall = &UbuntuRubyInstall{}
	templates.Register("UbuntuRubyInstall", ubunturubyinstall)
}

type UbuntuRubyInstall struct {
}

func (tpl *UbuntuRubyInstall) Options(t *templates.Template) {

}

func (tpl *UbuntuRubyInstall) Render(p urknall.Package) {
	p.AddTemplate("ruby", &UbuntuRubyInstallTemplate{})
}

func (tpl *UbuntuRubyInstall) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target, tpl, inputs)
}

type UbuntuRubyInstallTemplate struct {
}

func (m *UbuntuRubyInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("ruby-install",
    Shell(ppa),
		InstallPackages("ruby2.3 ruby2.3-dev"),
	)
}
