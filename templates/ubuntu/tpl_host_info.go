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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)


var ubuntuhostinfo *UbuntuHostInfo

func init() {
	ubuntuhostinfo = &UbuntuHostInfo{}
	templates.Register("UbuntuHostInfo", ubuntuhostinfo)
}

type UbuntuHostInfo struct{}

func (tpl *UbuntuHostInfo) Render(p urknall.Package) {
	p.AddTemplate("hostinfo", &UbuntuHostInfoTemplate{})
}

func (tpl *UbuntuHostInfo) Options(t *templates.Template) {}

func (tpl *UbuntuHostInfo) Run(target urknall.Target) error {
	return urknall.Run(target, &UbuntuHostInfo{})
}

type UbuntuHostInfoTemplate struct{}

func (m *UbuntuHostInfoTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("disk",
		Shell("df -h | grep \"/dev/sd\""),
	)
  pkg.AddCommands("memory",
    Shell("free -m"),
  )
  pkg.AddCommands("blockdevices",
		Shell("lsblk | grep sd"),
	)
  pkg.AddCommands("cpu",
		Shell("lscpu | grep \"CPU(s):\" "),
	)
  pkg.AddCommands("hostname",
		Shell("hostname"),
	)
	pkg.AddCommands("dnsserver",
		Shell("cat /etc/resolv.conf | grep nameserver"),
	)
  pkg.AddCommands("ipaddress",
		Shell("ifconfig | grep addr:"),
	)
	pkg.AddCommands("bridge",
  	Shell("if [ -f /sbin/brctl ] ; then brctl show; else echo 'brctl not installed'; fi"),
		)
  pkg.AddCommands("os",
  	Shell("grep PRETTY_NAME /etc/*-release | awk -F '=\"' '{print $2}'"),
		)

}



