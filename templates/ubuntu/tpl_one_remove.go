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

var ubuntuoneremove *UbuntuOneRemove

func init() {
	ubuntuoneremove = &UbuntuOneRemove{}
	templates.Register("UbuntuOneRemove", ubuntuoneremove)
}

type UbuntuOneRemove struct{}

func (tpl *UbuntuOneRemove) Render(p urknall.Package) {
	p.AddTemplate("one", &UbuntuOneRemoveTemplate{})
}

func (tpl *UbuntuOneRemove) Options(t *templates.Template) {
}

func (tpl *UbuntuOneRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuOneRemove{},inputs)
}

type UbuntuOneRemoveTemplate struct{}

func (m *UbuntuOneRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("one",
		RemovePackage("opennebula opennebula-sunstone"),
		RemovePackages(""),
		PurgePackages("opennebula opennebula-sunstone"),
	)
	pkg.AddCommands("one-clean",
		Shell("rm -r /var/lib/urknall/one.*"),
	)
}
