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
	KnownHostsList = `
	ConnectTimeout 5
	Host *
	StrictHostKeyChecking no
	`
)

var centossshpass *CentosSshPass

func init() {
	centossshpass = &CentosSshPass{}
	templates.Register("CentosSshPass", centossshpass)
}

type CentosSshPass struct{
	Host string
}


func (tpl *CentosSshPass) Render(p urknall.Package) {
	p.AddTemplate("sshpass", &CentosSshPassTemplate{
		Host: tpl.Host,
	})
}

func (tpl *CentosSshPass) Options(t *templates.Template) {

if hs, ok := t.Options["HOST"]; ok {
	tpl.Host = hs
}
}


func (tpl *CentosSshPass) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosSshPass{
		Host: tpl.Host,
	},inputs)
}

type CentosSshPassTemplate struct{
	Host string
}

func (m *CentosSshPassTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("install-sshpass",
	  InstallPackages("sshpass"),
	)
	pkg.AddCommands("SSHPass",
		AsUser("oneadmin", Shell("sshpass -p 'oneadmin' scp -o StrictHostKeyChecking=no /var/lib/one/.ssh/id_dsa.pub oneadmin@"+ m.Host +":/var/lib/one/.ssh/authorized_keys")),
    WriteFile("/var/lib/one/.ssh/config",KnownHostsList,"oneadmin", 0755),
	)
}
