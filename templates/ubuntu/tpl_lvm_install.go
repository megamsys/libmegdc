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
)

const (
	Hdd = "Disk"
)

var ubuntulvminstall *UbuntuLvmInstall

func init() {
	ubuntulvminstall = &UbuntuLvmInstall{}
	templates.Register("UbuntuLvmInstall", ubuntulvminstall)
}

type UbuntuLvmInstall struct {
	disks  []string
	vgname string
}

func (tpl *UbuntuLvmInstall) Options(t *templates.Template) {
	if disks, ok := t.Maps[Hdd]; ok {
		tpl.disks = disks
	}
	if vgname, ok := t.Options[VgName]; ok {
		tpl.vgname = vgname
	}
}

func (tpl *UbuntuLvmInstall) Render(p urknall.Package) {
	p.AddTemplate("lvm", &UbuntuLvmInstallTemplate{
		disks:  tpl.disks,
		vgname: tpl.vgname,
	})
}

func (tpl *UbuntuLvmInstall) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, &UbuntuLvmInstall{
		disks:  tpl.disks,
		vgname: tpl.vgname,
	}, inputs)
}

type UbuntuLvmInstallTemplate struct {
	disks  []string
	vgname string
}

func (m *UbuntuLvmInstallTemplate) Render(pkg urknall.Package) {
	diskdir := ArraytoString("/dev/", "", m.disks)
	vg := m.vgname
	pkg.AddCommands("lvminstall",
		UpdatePackagesOmitError(),
		InstallPackages("clvm lvm2 kvm"),
	)
	pkg.AddCommands("vg-setup",
		Shell("pvcreate "+diskdir+""),
		Shell("vgcreate "+vg+" "+diskdir+""),
	)
}
