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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)


var ubuntuhostcheck *UbuntuHostCheck

func init() {
	ubuntuhostcheck = &UbuntuHostCheck{}
	templates.Register("UbuntuHostCheck", ubuntuhostcheck)
}

type UbuntuHostCheck struct{}

func (tpl *UbuntuHostCheck) Render(p urknall.Package) {
	p.AddTemplate("hostcheck", &UbuntuHostCheckTemplate{})
}

func (tpl *UbuntuHostCheck) Options(t *templates.Template) {}

func (tpl *UbuntuHostCheck) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuHostCheck{},inputs)
}

type UbuntuHostCheckTemplate struct{}

func (m *UbuntuHostCheckTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("updating",
		Shell("apt-get update -y"),
	)
  pkg.AddCommands("cpu-checker",
    Shell("apt-get install -y qemu-system-x86 qemu-kvm cpu-checker"),
  )
	 pkg.AddCommands("kvm-ok",
		   If("-f /dev/kvm"	,Shell("kvm-ok  | grep \"KVM acceleration can be used\"")),
	)
}
