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
  "github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
	Disk     = "Disk"
	)

var centoscreatevolume *CentosCreateVolume

func init() {
	centoscreatevolume = &CentosCreateVolume{}
	templates.Register("CentosCreateVolume", centoscreatevolume)
}

type CentosCreateVolume struct {
	disk      string
	}

func (tpl *CentosCreateVolume) Options(t *templates.Template) {
  if disk, ok := t.Options[Disk]; ok {
		tpl.disk = disk
	}
}

func (tpl *CentosCreateVolume) Render(p urknall.Package) {
	p.AddTemplate("volume", &CentosCreateVolumeTemplate{
		disk:     tpl.disk,
		})
}

func (tpl *CentosCreateVolume) Run(target urknall.Target) error {
	return urknall.Run(target, &CentosCreateVolume{
		disk:     tpl.disk,
	})
}

type CentosCreateVolumeTemplate struct {
  disk     string
}

func (m *CentosCreateVolumeTemplate) Render(pkg urknall.Package) {
	disk := m.disk

	pkg.AddCommands("lvm",
		 Shell("yum install -y lvm2"),
		 )
	 pkg.AddCommands("physicalvolume",
 	Shell("pvcreate -y  /dev/"+disk+""),
 	)
 	pkg.AddCommands("volumegroup",
 		Shell("vgcreate vgpool /dev/"+disk+""),
 	)
 	pkg.AddCommands("logicalvolume",
 		Shell("lvcreate -L 3G -n lvstuff vgpool  "),
 	)
}
