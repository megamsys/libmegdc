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
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

var ubuntusshpass *UbuntuSshPass

func init() {
	ubuntusshpass = &UbuntuSshPass{}
	templates.Register("UbuntuSshPass", ubuntusshpass)
}

type UbuntuSshPass struct{
	Host string
}


func (tpl *UbuntuSshPass) Render(p urknall.Package) {
	p.AddTemplate("sshpass", &UbuntuSshPassTemplate{
		Host: tpl.Host,
	})
}

func (tpl *UbuntuSshPass) Options(t *templates.Template) {

if hs, ok := t.Options[HOSTNODE]; ok {
	tpl.Host = hs
}
}


func (tpl *UbuntuSshPass) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuSshPass{
		Host: tpl.Host,
	},inputs)
}

type UbuntuSshPassTemplate struct{
	Host string
}

func (m *UbuntuSshPassTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("install-sshpass",
	  InstallPackages("sshpass"),
	)
	pkg.AddCommands("SSHPass",
		AsUser("oneadmin", Shell("sshpass -p 'oneadmin' scp -o StrictHostKeyChecking=no /var/lib/one/.ssh/id_rsa.pub oneadmin@"+ m.Host +":/var/lib/one/.ssh/authorized_keys")),

	)
}
