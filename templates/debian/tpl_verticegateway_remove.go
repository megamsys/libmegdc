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
	"github.com/megamsys/libmegdc/templates"
	u "github.com/megamsys/libmegdc/templates/ubuntu"
	"github.com/megamsys/urknall"
)

var debiangatewayremove *DebianGatewayRemove

func init() {
	debiangatewayremove = &DebianGatewayRemove{}
	templates.Register("DebianGatewayRemove", debiangatewayremove)
}

type DebianGatewayRemove struct{}

func (tpl *DebianGatewayRemove) Render(p urknall.Package) {
	p.AddTemplate("gateway", &DebianGatewayRemoveTemplate{})
}

func (tpl *DebianGatewayRemove) Options(t *templates.Template) {
}

func (tpl *DebianGatewayRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &DebianGatewayRemove{},inputs)
}

type DebianGatewayRemoveTemplate struct{}

func (m *DebianGatewayRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("verticegateway",
		u.RemovePackage("verticegateway"),
		u.RemovePackages(""),
		u.PurgePackages("verticegateway"),
		u.Shell("dpkg --get-selections megam*"),
	)
	pkg.AddCommands("verticegateway-clean",
		u.Shell("rm -r /var/lib/urknall/gateway*"),
	)
}
