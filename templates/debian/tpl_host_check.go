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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
  u "github.com/megamsys/libmegdc/templates/ubuntu"
)


var debianhostcheck *DebianHostCheck

func init() {
	debianhostcheck = &DebianHostCheck{}
	templates.Register("DebianHostCheck", debianhostcheck)
}

type DebianHostCheck struct{}

func (tpl *DebianHostCheck) Render(p urknall.Package) {
	p.AddTemplate("hostinfo", &DebianHostCheckTemplate{})
}

func (tpl *DebianHostCheck) Options(t *templates.Template) {}

func (tpl *DebianHostCheck) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &DebianHostCheck{},inputs)
}

type DebianHostCheckTemplate struct{}

func (m *DebianHostCheckTemplate) Render(pkg urknall.Package) {
  pkg.AddCommands("update",
		u.Shell("apt-get update -y"),
	)
  pkg.AddCommands("cpu-checker",
    u.Shell("apt-get install -y cpu-checker"),
  )
  pkg.AddCommands("kvm-ok",
		u.Shell("kvm-ok"),
	)

}
