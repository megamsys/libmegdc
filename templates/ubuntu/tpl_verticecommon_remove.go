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

var ubuntuverticecommonremove *UbuntuMegamCommonRemove

func init() {
	ubuntuverticecommonremove = &UbuntuMegamCommonRemove{}
	templates.Register("UbuntuMegamCommonRemove", ubuntuverticecommonremove)
}

type UbuntuMegamCommonRemove struct{}

func (tpl *UbuntuMegamCommonRemove) Render(p urknall.Package) {
	p.AddTemplate("common", &UbuntuMegamCommonRemoveTemplate{})
}

func (tpl *UbuntuMegamCommonRemove) Options(t *templates.Template) {
}

func (tpl *UbuntuMegamCommonRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuMegamCommonRemove{},inputs)
}

type UbuntuMegamCommonRemoveTemplate struct{}

func (m *UbuntuMegamCommonRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("verticecommon",
		RemovePackage("verticecommon"),
		RemovePackages(""),
		PurgePackages("verticecommon"),
		Shell("dpkg --get-selections megam*"),
	)
	pkg.AddCommands("common-clean",
		Shell("rm -r /var/lib/urknall/common*"),
	)
}
