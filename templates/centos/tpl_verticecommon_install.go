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


var centosverticecommoninstall *CentosMegamCommonInstall

func init() {
	centosverticecommoninstall = &CentosMegamCommonInstall{}
	templates.Register("CentosMegamCommonInstall", centosverticecommoninstall)
}

type CentosMegamCommonInstall struct{}

func (tpl *CentosMegamCommonInstall) Render(p urknall.Package) {
	p.AddTemplate("common", &CentosMegamCommonInstallTemplate{})
}

func (tpl *CentosMegamCommonInstall) Options(t *templates.Template) {
}

func (tpl *CentosMegamCommonInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosMegamCommonInstall{},inputs)
}

type CentosMegamCommonInstallTemplate struct{}

func (m *CentosMegamCommonInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("repository",
		Shell("echo 'deb [arch=amd64] " + DefaultMegamRepo + "' > " + ListFilePath),
		UpdatePackagesOmitError(),
	)
	pkg.AddCommands("verticecommon",
		InstallPackages("verticecommon"),

	)
}
