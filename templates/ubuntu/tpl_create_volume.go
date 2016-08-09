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
	Disk     = "Disk"
	)

var ubuntucreatevolume *UbuntuCreateVolume

func init() {
	ubuntucreatevolume = &UbuntuCreateVolume{}
	templates.Register("UbuntuCreateVolume", ubuntucreatevolume)
}

type UbuntuCreateVolume struct {
	disk      string
	}

func (tpl *UbuntuCreateVolume) Options(t *templates.Template) {
	if disk, ok := t.Options[Disk]; ok {
		tpl.disk = disk
	}
}

func (tpl *UbuntuCreateVolume) Render(p urknall.Package) {
	p.AddTemplate("volume", &UbuntuCreateVolumeTemplate{
		disk:     tpl.disk,
		})
}

func (tpl *UbuntuCreateVolume) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuCreateVolume{
		disk:     tpl.disk,
	},inputs)
}

type UbuntuCreateVolumeTemplate struct {
  disk     string
}

func (m *UbuntuCreateVolumeTemplate) Render(pkg urknall.Package) {
	disk := m.disk

	pkg.AddCommands("lvm",
		 Shell("apt-get install -y lvm2"),
		 )
	 pkg.AddCommands("physicalvolume",
 	 Shell("pvcreate /dev/"+disk+""),
 	)
 	pkg.AddCommands("volumegroup",
 		Shell("vgcreate vgpool /dev/"+disk+""),
 	)
 	pkg.AddCommands("logicalvolume",
 		Shell("lvcreate -L 3G -n lvstuff vgpool  "),
 	)
}
