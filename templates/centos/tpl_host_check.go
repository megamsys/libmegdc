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

var centoshostcheck *CentosHostCheck

func init() {
	centoshostcheck = &CentosHostCheck{}
	templates.Register("CentosHostCheck", centoshostcheck)
}

type CentosHostCheck struct{}

func (tpl *CentosHostCheck) Render(p urknall.Package) {
	p.AddTemplate("hostcheck", &CentosHostCheckTemplate{})
}

func (tpl *CentosHostCheck) Options(t *templates.Template) {
}

func (tpl *CentosHostCheck) Run(target urknall.Target) error {
	return urknall.Run(target, &CentosHostCheck{})
}

type CentosHostCheckTemplate struct{}

func (m *CentosHostCheckTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("vtenable",
		Shell("grep -E 'svm|vmx' /proc/cpuinfo"),
	)
}
