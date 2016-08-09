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

var centosgatewayinstall *CentosGatewayInstall

func init() {
	centosgatewayinstall = &CentosGatewayInstall{}
	templates.Register("CentosGatewayInstall", centosgatewayinstall)
}

type CentosGatewayInstall struct{}

func (tpl *CentosGatewayInstall) Render(p urknall.Package) {
	p.AddTemplate("gateway", &CentosGatewayInstallTemplate{})
}

func (tpl *CentosGatewayInstall) Options(t *templates.Template) {
}

func (tpl *CentosGatewayInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosGatewayInstall{},inputs)
}

type CentosGatewayInstallTemplate struct{}

func (m *CentosGatewayInstallTemplate) Render(pkg urknall.Package) {
	//fail on Java -version (1.8 check)
	pkg.AddCommands("repository",
		Shell("echo 'deb [arch=amd64] "+DefaultMegamRepo+"' > "+ListFilePath),
		UpdatePackagesOmitError(),
	)

	pkg.AddCommands("verticegateway",
		InstallPackages("verticegateway"),
	)
}
