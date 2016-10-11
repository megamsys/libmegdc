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
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	"fmt"
)

var ubuntucephpoolcreate *UbuntuCephPoolCreate

func init() {
	ubuntucephpoolcreate = &UbuntuCephPoolCreate{}
	templates.Register("UbuntuCephPoolCreate", ubuntucephpoolcreate)
}

type UbuntuCephPoolCreate struct {
		poolname string
		cephuser string
}

func (tpl *UbuntuCephPoolCreate) Options(t *templates.Template) {
	if poolname, ok := t.Options[POOLNAME]; ok {
		tpl.poolname = poolname
	}
	if cephuser, ok := t.Options[CEPHUSER]; ok {
		tpl.cephuser = cephuser
	}
}

func (tpl *UbuntuCephPoolCreate) Render(p urknall.Package) {
	p.AddTemplate("pool-create", &UbuntuCephPoolCreateTemplate{
		poolname: tpl.poolname,
		cephuser: tpl.cephuser,
	})
}

func (tpl *UbuntuCephPoolCreate) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target,tpl,inputs)
}

type UbuntuCephPoolCreateTemplate struct {
	poolname string
	cephuser string
}

func (m *UbuntuCephPoolCreateTemplate) Render(pkg urknall.Package) {
	var UserHome string
		if m.cephuser == "root" {
	    UserHome = "/" +  m.cephuser
	  } else {
			UserHome = UserHomePrefix + m.cephuser
		}

		pkg.AddCommands("ceph-pool-create",
  	AsUser(CephUser,Shell("ceph osd pool create "+m.poolname+" 128")),
	)
}
