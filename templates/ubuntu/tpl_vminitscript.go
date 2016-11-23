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

const (
	InitScripts = "InitScripts"
)

var ubuntuvminitscript *UbuntuVmInitScript

func init() {
	ubuntuvminitscript = &UbuntuVmInitScript{}
	templates.Register("UbuntuVmInitScript", ubuntuvminitscript)
}

type UbuntuVmInitScript struct {
	initscripts string
}

func (tpl *UbuntuVmInitScript) Options(t *templates.Template) {
	if initscripts, ok := t.Options[InitScripts]; ok {
		tpl.initscripts = initscripts
	}
}

func (tpl *UbuntuVmInitScript) Render(p urknall.Package) {
	p.AddTemplate("vminitscript", &UbuntuVmInitScriptTemplate{
		initscripts: tpl.initscripts,
	})
}

func (tpl *UbuntuVmInitScript) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target, &UbuntuVmInitScript{
		initscripts: tpl.initscripts,
	}, inputs)
}

type UbuntuVmInitScriptTemplate struct {
	initscripts string
}

func (m *UbuntuVmInitScriptTemplate) Render(pkg urknall.Package) {
	initscripts := m.initscripts
	pkg.AddCommands("WriteScript",
		Shell("mkdir -p /megam"),
	WriteFile("/megam/init.sh", initscripts, "root", 0755),
	)
}
