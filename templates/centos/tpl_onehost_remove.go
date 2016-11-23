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

var centosonehostremove *CentosOneHostRemove

func init() {
	centosonehostremove = &CentosOneHostRemove{}
	templates.Register("CentosOneHostRemove", centosonehostremove)
}

type CentosOneHostRemove struct{}

func (tpl *CentosOneHostRemove) Render(p urknall.Package) {
	p.AddTemplate("onehost", &CentosOneHostRemoveTemplate{})
}

func (tpl *CentosOneHostRemove) Options(t *templates.Template) {
}

func (tpl *CentosOneHostRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosOneHostRemove{},inputs)
}

type CentosOneHostRemoveTemplate struct{}

func (m *CentosOneHostRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("onehost",
		RemovePackage("opennebula-node openvswitch-common openvswitch-switch bridge-utils sshpass"),
		RemovePackages(""),

	)
	pkg.AddCommands("onehost-clean",
		Shell("rm -rf /var/lib/urknall/onehost*"),
	)
}
