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

package debian

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
	 u "github.com/megamsys/libmegdc/templates/ubuntu"
)

var debianverticecommoninstall *DebianMegamCommonInstall

func init() {
	debianverticecommoninstall = &DebianMegamCommonInstall{}
	templates.Register("DebianMegamCommonInstall", debianverticecommoninstall)
}

type DebianMegamCommonInstall struct{}

func (tpl *DebianMegamCommonInstall) Render(p urknall.Package) {
	p.AddTemplate("common", &DebianMegamCommonInstallTemplate{})
}

func (tpl *DebianMegamCommonInstall) Options(t *templates.Template) {
}

func (tpl *DebianMegamCommonInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &DebianMegamCommonInstall{},inputs)
}

type DebianMegamCommonInstallTemplate struct{}

func (m *DebianMegamCommonInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("repository",
		u.Shell("echo 'deb [arch=amd64] " + DefaultMegamRepo + "' > " + ListFilePath),
		u.UpdatePackagesOmitError(),
	)
	pkg.AddCommands("verticecommon",
		u.InstallPackages("verticecommon"),

	)
}
