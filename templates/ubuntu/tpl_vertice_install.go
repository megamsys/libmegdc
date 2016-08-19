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
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

const (
	apiConf =`sed -i 's/^[ \t]*api = "https:\/\/api.megam.io\/v2".*/    api = "http:\/\/%s:9000\/v2"/' /var/lib/megam/vertice/vertice.conf`
  nsqdConf = `sed -i 's/^[ \t]*nsqd = \["localhost:4150"\].*/    nsqd = \["%s:4150"\]/' /var/lib/megam/vertice/vertice.conf`
  scyllaConf = `sed -i 's/^[ \t]*scylla = \["localhost"\].*/    scylla = \["%s"\]/' /var/lib/megam/vertice/vertice.conf`
)

var ubuntuverticeinstall *UbuntuMegamdInstall

func init() {
	ubuntuverticeinstall = &UbuntuMegamdInstall{}
	templates.Register("UbuntuMegamdInstall", ubuntuverticeinstall)
}

type UbuntuMegamdInstall struct{
	hostip string
}

func (tpl *UbuntuMegamdInstall) Render(p urknall.Package) {
	p.AddTemplate("vertice", &UbuntuMegamdInstallTemplate{
		hostip: tpl.hostip,
	})
}

func (tpl *UbuntuMegamdInstall) Options(t *templates.Template) {
	if host,ok := t.Options[HOST]; ok {
		tpl.hostip = host
	}
}

func (tpl *UbuntuMegamdInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target,tpl,inputs)
}

type UbuntuMegamdInstallTemplate struct{
	hostip string
}

func (m *UbuntuMegamdInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("vertice",
	InstallPackages("vertice"),
	)
	pkg.AddCommands("conf",
	Shell(fmt.Sprintf(apiConf, m.hostip)),
	Shell(fmt.Sprintf(nsqdConf, m.hostip)),
	Shell(fmt.Sprintf(scyllaConf, m.hostip)),
	Shell("sudo restart vertice"),
	)
}
