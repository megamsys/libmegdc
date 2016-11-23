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
)

var ubuntuformatpartitions *UbuntuFormatPartitions

func init() {

	ubuntuformatpartitions = &UbuntuFormatPartitions{}
	templates.Register("UbuntuFormatPartitions", ubuntuformatpartitions)
}

type UbuntuFormatPartitions struct{
  mount string
  partitions string
}

func (tpl *UbuntuFormatPartitions) Options(t *templates.Template) {
  if mount, ok := t.Options[Mount]; ok {
		tpl.mount = mount
	}
  if partitions, ok := t.Options[Disks]; ok {
		tpl.partitions = partitions
	}
}

func (tpl *UbuntuFormatPartitions) Render(p urknall.Package) {
	p.AddTemplate("delete", &UbuntuFormatPartitionsTemplate{
    mount : tpl.mount,
    partitions : tpl.partitions,

  })
}

func (tpl *UbuntuFormatPartitions) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuFormatPartitions{
    mount : tpl.mount,
    partitions : tpl.partitions,
    },inputs)
}

type UbuntuFormatPartitionsTemplate struct{
  mount string
  partitions string

}

func (m *UbuntuFormatPartitionsTemplate) Render(pkg urknall.Package) {
pkg.AddCommands("del-partition",
  Shell("umount "+m.partitions+""),
  Shell("mkfs.ext4 "+ m.partitions +""),
  )
}
