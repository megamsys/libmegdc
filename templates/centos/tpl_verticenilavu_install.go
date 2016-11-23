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

const (
	// DefaultMegamRepo is the default megam repository to install if its not provided.
	DefaultMegamRepo = "http://get.megam.io/1.0/ubuntu/14.04/ trusty nightly"

	ListFilePath = "/etc/apt/sources.list.d/megam.list"


	Strict = `
	ConnectTimeout 5
	Host *
	StrictHostKeyChecking no
	`
)

var centosnilavuinstall *CentosNilavuInstall

func init() {
	centosnilavuinstall = &CentosNilavuInstall{}
	templates.Register("CentosNilavuInstall", centosnilavuinstall)
}

type CentosNilavuInstall struct{}

func (tpl *CentosNilavuInstall) Render(p urknall.Package) {
	p.AddTemplate("nilavu", &CentosNilavuInstallTemplate{})
}

func (tpl *CentosNilavuInstall) Options(t *templates.Template) {
}

func (tpl *CentosNilavuInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosNilavuInstall{},inputs)
}

type CentosNilavuInstallTemplate struct{}

func (m *CentosNilavuInstallTemplate) Render(pkg urknall.Package) {
  //fail on ruby2.0 < check

	pkg.AddCommands("repository",
		Shell("echo 'deb [arch=amd64] "+DefaultMegamRepo+"' > "+ListFilePath),
		UpdatePackagesOmitError(),
	)

	pkg.AddCommands("verticenilavu",
		InstallPackages("verticenilavu"),
	)
}
