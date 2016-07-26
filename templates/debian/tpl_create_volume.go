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

package debian

import (
  "github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
  u "github.com/megamsys/libmegdc/templates/ubuntu"
)

const (
	Disk     = "Disk"
	)

var debiancreatevolume *DebianCreateVolume

func init() {
	debiancreatevolume = &DebianCreateVolume{}
	templates.Register("DebianCreateVolume", debiancreatevolume)
}

type DebianCreateVolume struct {
	disk      string
	}

func (tpl *DebianCreateVolume) Options(t *templates.Template) {
	if disk, ok := t.Options[Disk]; ok {
		tpl.disk = disk
	}
}

func (tpl *DebianCreateVolume) Render(p urknall.Package) {
	p.AddTemplate("volume", &DebianCreateVolumeTemplate{
		disk:     tpl.disk,
		})
}

func (tpl *DebianCreateVolume) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &DebianCreateVolume{
		disk:     tpl.disk,
	},inputs)
}

type DebianCreateVolumeTemplate struct {
  disk     string
}

func (m *DebianCreateVolumeTemplate) Render(pkg urknall.Package) {
	disk := m.disk

	pkg.AddCommands("lvm",
		 u.Shell("apt-get install -y lvm2"),
		 )
	 pkg.AddCommands("physicalvolume",
 	u.Shell("pvcreate /dev/"+disk+""),
 	)
 	pkg.AddCommands("volumegroup",
 		u.Shell("vgcreate vgpool /dev/"+disk+""),
 	)
 	pkg.AddCommands("logicalvolume",
 		u.Shell("lvcreate -L 3G -n lvstuff vgpool  "),
 	)
}
