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

package ubuntu

import (
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	LvName  = "LvName"
	PvName  = "PvName"
)

var ubuntulvmremove *UbuntuLvmRemove

func init() {
	ubuntulvmremove = &UbuntuLvmRemove{}
	templates.Register("UbuntuLvmRemove", ubuntulvmremove)
}

type UbuntuLvmRemove struct {
	lvname []string
	vgname string
  pvname string
}

func (tpl *UbuntuLvmRemove) Options(t *templates.Template) {
	if lvname, ok := t.Maps[LvName]; ok {
		tpl.lvname = lvname
	}
	if vgname, ok := t.Options[VgName]; ok {
		tpl.vgname = vgname
	}
  if pvname, ok := t.Options[PvName]; ok {
		tpl.pvname = pvname
	}
}

func (tpl *UbuntuLvmRemove) Render(p urknall.Package) {
	p.AddTemplate("lvmremove", &UbuntuLvmRemoveTemplate{
		lvname:     tpl.lvname,
		vgname:     tpl.vgname,
    pvname:     tpl.pvname,
	})
}

func (tpl *UbuntuLvmRemove) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuLvmRemove{
		lvname:     tpl.lvname,
		vgname:     tpl.vgname,
    pvname:     tpl.pvname,
	},inputs)
}

type UbuntuLvmRemoveTemplate struct {
  lvname []string
	vgname string
  pvname string
}

func (m *UbuntuLvmRemoveTemplate) Render(pkg urknall.Package) {
  lvname := ArraytoString("","",m.lvname)
	vg := m.vgname
  pv := m.pvname
 	pkg.AddCommands("lvm-remove",
		Shell("lvremove -f " + lvname),
		Shell("vgremove "+vg),
    Shell("pvremove "+pv),
	)
}
