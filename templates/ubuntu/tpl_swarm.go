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
	//"github.com/megamsys/libgo/cmd"
)
var ubuntuswarminstall *UbuntuSwarmInstall

func init() {
	ubuntuswarminstall = &UbuntuSwarmInstall{}
	templates.Register("UbuntuSwarmInstall", ubuntuswarminstall)
}

type UbuntuSwarmInstall struct {}

func (tpl *UbuntuSwarmInstall) Options(t *templates.Template) {}

func (tpl *UbuntuSwarmInstall) Render(p urknall.Package) {
	p.AddTemplate("swarm", &UbuntuSwarmInstallTemplate{})
}

func (tpl *UbuntuSwarmInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuSwarmInstall{},inputs)
}

type UbuntuSwarmInstallTemplate struct {}

func (m *UbuntuSwarmInstallTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("swarminstall",
		 Shell(" curl -fsSL https://get.docker.com/ | sh"),
	 )
	 
	pkg.AddCommands("Run",
		Shell("docker run" ),
	)

	pkg.AddCommands("create ",
		Shell("--rm swarm create"),
	)

}
