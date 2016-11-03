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
	"strconv"
)
const(
  OSDCOUNT = "OsdsCount"
  DISKPARTS =  "DiskParts"
)


var ubuntuaddpartitionosds *UbuntuAddPartitionOsds

func init() {
	ubuntuaddpartitionosds = &UbuntuAddPartitionOsds{}
	templates.Register("UbuntuAddPartitionOsds", ubuntuaddpartitionosds)
}

type UbuntuAddPartitionOsds struct {
	osds      []string
	cephuser string
  clienthostname string
  osdno   string
}

func (tpl *UbuntuAddPartitionOsds) Options(t *templates.Template) {
	if osds, ok := t.Maps[DISKPARTS]; ok {
		tpl.osds = osds
	}
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
  if clienthostname, ok := t.Options[CLIENTHOST]; ok {
    tpl.clienthostname = clienthostname
  }
  if osdno, ok := t.Options[CLIENTHOST]; ok {
    tpl.osdno = osdno
  }
}

func (tpl *UbuntuAddPartitionOsds) Render(p urknall.Package) {
	p.AddTemplate("add-osd-part", &UbuntuAddPartitionOsdsTemplate{
		osds:     tpl.osds,
		cephuser: tpl.cephuser,
    clienthostname: tpl.clienthostname,
    osdno: tpl.osdno,
	})
}

func (tpl *UbuntuAddPartitionOsds) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, tpl ,inputs)
}

type UbuntuAddPartitionOsdsTemplate struct {
  osds     []string
	cephuser string
  clienthostname string
  osdno          string
}

func (m *UbuntuAddPartitionOsdsTemplate) Render(pkg urknall.Package) {
	var activate,disks,dirs,mountpoint string
  CephUser := m.cephuser
	CephHome := UserHomePrefix + CephUser
  ClientHostName := m.clienthostname
	if m.cephuser == "root" {
		CephHome = "/root"
	} else {
		CephHome = UserHomePrefix + m.cephuser
	}
	osdhome := "/var/lib/ceph/osd/ceph-"
	osdcount,_ := strconv.Atoi(m.osdno)
	for i,k := range m.osds {
		t := osdcount + i
		no := strconv.Itoa(t)
	  activate = activate + " " + ClientHostName + ":" + osdhome + no
		dirs = dirs + " " + osdhome + no
		mountpoint = mountpoint + " mount " + "/dev/" + k + " "+ osdhome + no + ";"
	}
  disks =  ArraytoString("/dev/","",m.osds)
	pkg.AddCommands("umount",
		 Shell("umount "+ disks ),
		 Shell("mkdir "+ dirs ),
	)
	pkg.AddCommands("mount",
		 Shell(mountpoint),
	)
	pkg.AddCommands("prepare-osds",
    //AsUser("root",Shell("sudo chown -R "+ CephUser +":"+ CephUser +" /etc/ceph/ceph.client.admin.keyring")),
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy --overwrite-conf osd prepare "+ disks )),
	)
	pkg.AddCommands("activate-osds",
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy osd activate "+ disks )),
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy admin "+ ClientHostName )),
		RemoveAllCaches("/var/lib/urknall/aadd-osd-part.*"),
	)
}
