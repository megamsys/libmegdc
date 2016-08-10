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
  //"fmt"
  )

const (
  Mount = "Mount"
	Partitions = "Partitions"
)

var ubuntuformatpartition *UbuntuFormatPartition

func init() {

	ubuntuformatpartition = &UbuntuFormatPartition{}
	templates.Register("UbuntuFormatPartition", ubuntuformatpartition)
}

type UbuntuFormatPartition struct{
  mount string
  partitions string
}

func (tpl *UbuntuFormatPartition) Options(t *templates.Template) {
  if mount, ok := t.Options[Mount]; ok {
		tpl.mount = mount
	}
  if partitions, ok := t.Options[Partitions]; ok {
		tpl.partitions = partitions
	}
}

func (tpl *UbuntuFormatPartition) Render(p urknall.Package) {
	p.AddTemplate("delete", &UbuntuFormatPartitionTemplate{
    mount : tpl.mount,
    partitions : tpl.partitions,

  })
}

func (tpl *UbuntuFormatPartition) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuFormatPartition{
    mount : tpl.mount,
    partitions : tpl.partitions,
    },inputs)
}

type UbuntuFormatPartitionTemplate struct{
  mount string
  partitions string

}

func (m *UbuntuFormatPartitionTemplate) Render(pkg urknall.Package) {
Mount := m.mount
//Partitions := m.partitions
ZapDisk := "sda"

pkg.AddCommands("del-partition",
  Shell("umount "+Mount+""),
  Shell("mkfs.ext4 "+ZapDisk+""),
  )
}
