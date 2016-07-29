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
	"os"
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Hdd     = "Osd"
	VgName  = "VgName"
)

var ubuntulvminstall *UbuntuLvmInstall

func init() {
	ubuntulvminstall = &UbuntuLvmInstall{}
	templates.Register("UbuntuLvmInstall", ubuntulvminstall)
}

type UbuntuLvmInstall struct {
	osds      []string
	vgname string
}

func (tpl *UbuntuLvmInstall) Options(t *templates.Template) {
	if osds, ok := t.Maps[Hdd]; ok {
		tpl.osds = osds
	}
	if vgname, ok := t.Options[VgName]; ok {
		tpl.vgname = vgname
	}
}

func (tpl *UbuntuLvmInstall) Render(p urknall.Package) {
	p.AddTemplate("lvm", &UbuntuLvmInstallTemplate{
		osds:     tpl.osds,
		phydev:    tpl.phydev,
	})
}

func (tpl *UbuntuLvmInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuLvmInstall{
		osds:     tpl.osds,
		vgname: tpl.vgname,

	},inputs)
}

type UbuntuLvmInstallTemplate struct {
  osds     []string
	vgname string
}

func (m *UbuntuLvmInstallTemplate) Render(pkg urknall.Package) {
  osddir := ArraytoString("/dev/","",m.osds)
	vg := m.vgname
 pkg.AddCommands("lvminstall",
	  UpdatePackagesOmitError(),
		InstallPackages("clvm lvm2 kvm"),
	)
	pkg.AddCommands("vg-setup",
		Shell("pvcreate "+osddir+""),
		Shell("vgcreate "+vg+" "+osddir+""),
	)
}
