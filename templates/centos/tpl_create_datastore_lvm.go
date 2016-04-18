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
	"github.com/megamsys/libmegdc/templates"
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

var centoscreatedatastorelvm *CentosCreateDatastoreLvm

func init() {
	centoscreatedatastorelvm = &CentosCreateDatastoreLvm{}
	templates.Register("CentosCreateDatastoreLvm", centoscreatedatastorelvm)
}

type CentosCreateDatastoreLvm struct {

	poolname     string
  vgname      string
  hostname   string
	}

func (tpl *CentosCreateDatastoreLvm) Options(t *templates.Template) {
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

func (tpl *CentosCreateDatastoreLvm) Render(p urknall.Package) {
	p.AddTemplate("createdatastorelvm", &CentosCreateDatastoreLvmTemplate{
		poolname:   tpl.poolname,
    vgname:    tpl.vgname,
    hostname:   tpl.hostname,
		})
}

func (tpl *CentosCreateDatastoreLvm) Run(target urknall.Target) error {
	return urknall.Run(target, &CentosCreateDatastoreLvm{
		poolname:     tpl.poolname,
    vgname:     tpl.vgname,
    hostname:    tpl.hostname,
	})
}

type CentosCreateDatastoreLvmTemplate struct {
  poolname    string
  vgname    string
  hostname   string
}

func (m *CentosCreateDatastoreLvmTemplate) Render(pkg urknall.Package) {
	poolname := m.poolname
  vgname := m.vgname
  hostname := m.hostname

	 pkg.AddCommands("create-datastore",
  WriteFile("/var/lib/dslvmconf",fmt.Sprintf(Dslvmconf, poolname, hostname, vgname, hostname),  "root" , 0644),
 	 Shell(" onedatastore create /var/lib/dslvmconf"),
 	)
	pkg.AddCommands("list",
	Shell("onedatastore list"),
	)
}
