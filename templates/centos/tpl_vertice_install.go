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
	"github.com/megamsys/libmegdc/templates"

	"github.com/megamsys/urknall"

)

var centosverticeinstall *CentosMegamdInstall

func init() {
	centosverticeinstall = &CentosMegamdInstall{}
	templates.Register("CentosMegamdInstall", centosverticeinstall)
}

type CentosMegamdInstall struct{}

func (tpl *CentosMegamdInstall) Render(p urknall.Package) {
	p.AddTemplate("vertice", &CentosMegamdInstallTemplate{})
}

func (tpl *CentosMegamdInstall) Options(t *templates.Template) {
}

func (tpl *CentosMegamdInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosMegamdInstall{},inputs)
}

type CentosMegamdInstallTemplate struct{}

func (m *CentosMegamdInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("repository",
		Shell("echo 'deb [arch=amd64] "+DefaultMegamRepo+"' > "+ListFilePath),
		UpdatePackagesOmitError(),
	)

	pkg.AddCommands("vertice",
	InstallPackages("vertice"),
	)

}
