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
	"github.com/megamsys/libmegdc/templates"
  	u "github.com/megamsys/libmegdc/templates/ubuntu"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
  "fmt"
)

const (

  VgName    = "VgName"
  PoolName  = "PoolName"
  Hostname   = "Hostname"
  Dslvmconf = `NAME = "lvmds"
DS_MAD = lvm
TM_MAD = lvm
DISK_TYPE = block
POOL_NAME = %s
BRIDGE_LIST = %s
VG_NAME = "%s"
HOST = "%s"
  `
	)

var debiancreatedatastorelvm *DebianCreateDatastoreLvm

func init() {
	debiancreatedatastorelvm = &DebianCreateDatastoreLvm{}
	templates.Register("DebianCreateDatastoreLvm", debiancreatedatastorelvm)
}

type DebianCreateDatastoreLvm struct {

	poolname     string
  vgname      string
  hostname   string
	}

func (tpl *DebianCreateDatastoreLvm) Options(t *templates.Template) {
	if poolname, ok := t.Options[PoolName]; ok {
		tpl.poolname = poolname
	}
  if vgname, ok := t.Options[VgName]; ok {
		tpl.vgname = vgname
	}
  if hostname, ok := t.Options[Hostname]; ok {
    tpl.hostname = hostname
  }
}

func (tpl *DebianCreateDatastoreLvm) Render(p urknall.Package) {
	p.AddTemplate("createdatastorelvm", &DebianCreateDatastoreLvmTemplate{
		poolname:   tpl.poolname,
    vgname:    tpl.vgname,
    hostname:   tpl.hostname,
		})
}

func (tpl *DebianCreateDatastoreLvm) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &DebianCreateDatastoreLvm{
		poolname:     tpl.poolname,
    vgname:     tpl.vgname,
    hostname:    tpl.hostname,
	},inputs)
}

type DebianCreateDatastoreLvmTemplate struct {
  poolname    string
  vgname    string
  hostname   string
}

func (m *DebianCreateDatastoreLvmTemplate) Render(pkg urknall.Package) {
	poolname := m.poolname
  vgname := m.vgname
  hostname := m.hostname

	 pkg.AddCommands("create-datastore",
  u.WriteFile("/var/lib/dslvmconf",fmt.Sprintf(Dslvmconf, poolname, hostname, vgname, hostname),  "root" , 0644),
 	 u.Shell(" onedatastore create /var/lib/dslvmconf"),
 	)
	pkg.AddCommands("list",
	u.Shell("onedatastore list"),
	)
}
