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

var centosnilavuremove *CentosNilavuRemove

func init() {
	centosnilavuremove = &CentosNilavuRemove{}
	templates.Register("CentosNilavuRemove", centosnilavuremove)
}

type CentosNilavuRemove struct{}

func (tpl *CentosNilavuRemove) Render(p urknall.Package) {
	p.AddTemplate("nilavu", &CentosNilavuRemoveTemplate{})
}

func (tpl *CentosNilavuRemove) Options(t *templates.Template) {
}

func (tpl *CentosNilavuRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosNilavuRemove{},inputs)
}

type CentosNilavuRemoveTemplate struct{}

func (m *CentosNilavuRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("verticenilavu",
		RemovePackage("verticenilavu"),
		RemovePackages(""),
		Shell("dpkg --get-selections megam*"),
	)
	pkg.AddCommands("nilavu-clean",
		Shell("rm -rf /var/lib/urknall/nilavu*"),
	)
}
