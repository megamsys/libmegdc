/*
** Copyright [2013-2015] [Megam Systems]
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

var centosverticeremove *CentosMegamdRemove

func init() {
	centosverticeremove = &CentosMegamdRemove{}
	templates.Register("CentosMegamdRemove", centosverticeremove)
}

type CentosMegamdRemove struct{}

func (tpl *CentosMegamdRemove) Render(p urknall.Package) {
	p.AddTemplate("vertice", &CentosMegamdRemoveTemplate{})
}

func (tpl *CentosMegamdRemove) Options(t *templates.Template) {
}

func (tpl *CentosMegamdRemove) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosMegamdRemove{},inputs)
}

type CentosMegamdRemoveTemplate struct{}

func (m *CentosMegamdRemoveTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("vertice",
		RemovePackage("vertice"),
		RemovePackages(""),
		Shell("dpkg --get-selections megam*"),
	)
	pkg.AddCommands("vertice-clean",
		Shell("rm -rf /var/lib/urknall/vertice*"),
	)
}
