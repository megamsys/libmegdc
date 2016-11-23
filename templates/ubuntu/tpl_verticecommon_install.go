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


var ubuntuverticecommoninstall *UbuntuMegamCommonInstall

func init() {
	ubuntuverticecommoninstall = &UbuntuMegamCommonInstall{}
	templates.Register("UbuntuMegamCommonInstall", ubuntuverticecommoninstall)
}

type UbuntuMegamCommonInstall struct{}

func (tpl *UbuntuMegamCommonInstall) Render(p urknall.Package) {
	p.AddTemplate("common", &UbuntuMegamCommonInstallTemplate{})
}

func (tpl *UbuntuMegamCommonInstall) Options(t *templates.Template) {
}

func (tpl *UbuntuMegamCommonInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuMegamCommonInstall{},inputs)
}

type UbuntuMegamCommonInstallTemplate struct{}

func (m *UbuntuMegamCommonInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("verticecommon",
		InstallPackages("verticecommon"),

	)
}
