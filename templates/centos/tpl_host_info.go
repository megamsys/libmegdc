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

package centos

import (
	"github.com/megamsys/libmegdc/templates"

	"github.com/megamsys/urknall"

)

var centoshostinfo *CentosHostInfo

func init() {
	centoshostinfo = &CentosHostInfo{}
	templates.Register("CentosHostInfo", centoshostinfo)
}

type CentosHostInfo struct{}

func (tpl *CentosHostInfo) Render(p urknall.Package) {
	p.AddTemplate("hostinfo", &CentosHostInfoTemplate{})
}

func (tpl *CentosHostInfo) Options(t *templates.Template) {
}

func (tpl *CentosHostInfo) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosHostInfo{},inputs)
}

type CentosHostInfoTemplate struct{}

func (m *CentosHostInfoTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("disk",
		Shell("df -h"),
	)
	pkg.AddCommands("memory",
		Shell("free -m"),
	)
	pkg.AddCommands("blockdevices",
		Shell("lsblk"),
	)
	pkg.AddCommands("cpu",
		Shell("lscpu"),
	)
	pkg.AddCommands("hostname",
		Shell("hostname"),
	)
	pkg.AddCommands("dnsserver",
		Shell("cat /etc/resolv.conf"),
	)
	pkg.AddCommands("ipaddress",
		Shell("yum install -y net-tools"),
		Shell("ifconfig"),
	)
	pkg.AddCommands("bridge",
 		 Shell("if /sbin/brctl ; then brctl show; else echo 'no bridge is available'; fi"),
 	 )
}
