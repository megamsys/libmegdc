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
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
//	"strings"
)

const (
	DelPartitions = `
#!/bin/bash
Disk=(%s)
for i in ${Disk[@]}
do
echo "sed -e 's/\s*\([\+0-9a-zA-Z]*\).*/\1/' << EOF | fdisk /dev/$i" > delpart
for (( totalPartitions=$(grep -c $i[0-9] /proc/partitions); totalPartitions>0; totalPartitions-- ))
do
	echo "d" >> delpart
	echo $totalPartitions >> delpart
done
echo 'w' >> delpart
echo 'EOF' >> delpart
chmod +x delpart
echo "disk $i zapping unsuccessful" | ./delpart
done
`

)

var ubuntudeletepartitions *UbuntuDeletePartitions

func init() {
	ubuntudeletepartitions = &UbuntuDeletePartitions{}
	templates.Register("UbuntuDeletePartitions", ubuntudeletepartitions)
}

type UbuntuDeletePartitions struct {
	disks []string
}

func (tpl *UbuntuDeletePartitions) Options(t *templates.Template) {
	if disks, ok := t.Maps[Disks]; ok {
		tpl.disks = disks
	}
}

func (tpl *UbuntuDeletePartitions) Render(p urknall.Package) {
	p.AddTemplate("delpart", &UbuntuDeletePartitionsTemplate{
		disks: tpl.disks,
	})
}

func (tpl *UbuntuDeletePartitions) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, &UbuntuDeletePartitions{
		disks: tpl.disks,
	}, inputs)
}

type UbuntuDeletePartitionsTemplate struct {
	disks []string
}

func (m *UbuntuDeletePartitionsTemplate) Render(pkg urknall.Package) {
	Disk := ArraytoString("","",m.disks)
	path := "/var/lib/urknall/zapscripts.sh"
	pkg.AddCommands("write-scripts",
		WriteFile(path, fmt.Sprintf(DelPartitions, Disk), "root", 0755),
  )
	pkg.AddCommands("delete-partitions",
		Shell("cat "+path),
	)
}
