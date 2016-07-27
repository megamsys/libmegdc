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


var debianhostinfo *DebianHostInfo

func init() {
	debianhostinfo = &DebianHostInfo{}
	templates.Register("DebianHostInfo", debianhostinfo)
}

type DebianHostInfo struct{}

func (tpl *DebianHostInfo) Render(p urknall.Package) {
	p.AddTemplate("hostinfo", &DebianHostInfoTemplate{})
}

func (tpl *DebianHostInfo) Options(t *templates.Template) {}

func (tpl *DebianHostInfo) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &DebianHostInfo{},inputs)
}

type DebianHostInfoTemplate struct{}

func (m *DebianHostInfoTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("disk",
		u.Shell("df -h"),
	)
  pkg.AddCommands("memory",
    u.Shell("free -m"),
  )
  pkg.AddCommands("blockdevices",
	  u.Shell("lsblk"),
	)
  pkg.AddCommands("cpu",
		u.Shell("lscpu"),
	)
  pkg.AddCommands("hostname",
		u.Shell("hostname"),
	)
	pkg.AddCommands("dnsserver",
		u.Shell("cat /etc/resolv.conf"),
	)
  pkg.AddCommands("ipaddress",
		u.Shell("ifconfig"),
	)
	pkg.AddCommands("bridge",
   u.Shell("if /sbin/brctl ; then brctl show; else echo 'no bridge is available'; fi"),
		)

}
