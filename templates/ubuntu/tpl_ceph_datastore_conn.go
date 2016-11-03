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

const (
	Ceph_User = "megdc"
  Uid =`uuidgen`

Xml=`<secret ephemeral='no' private='no'>
  <uuid>%v</uuid>
  <usage type='ceph'>
          <name>client.libvirt secret</name>
  </usage>
</secret>`
Setval=`sudo virsh secret-set-value --secret %v --base64 $(cat client.libvirt.key)`
Echo =`echo '%v'`
)

var ubuntucephdatastore *UbuntuCephDatastore

func init() {
	ubuntucephdatastore = &UbuntuCephDatastore{}
	templates.Register("UbuntuCephDatastore", ubuntucephdatastore)
}

type UbuntuCephDatastore struct {
		uuid string
		cephuser string
		poolname string
}

func (tpl *UbuntuCephDatastore) Options(t *templates.Template) {
	if uuid, ok := t.Options[CLUSTERID]; ok {
		tpl.uuid = uuid
	}
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
	if poolname, ok := t.Options[POOLNAME]; ok {
		tpl.poolname = poolname
	}
}

func (tpl *UbuntuCephDatastore) Render(p urknall.Package) {
	p.AddTemplate("cephds", &UbuntuCephDatastoreTemplate{
		uuid: tpl.uuid,
		cephuser: tpl.cephuser,
	})
}

func (tpl *UbuntuCephDatastore) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target,tpl,inputs)
}

type UbuntuCephDatastoreTemplate struct {
	uuid string
	cephuser string
	poolname string
}

func (m *UbuntuCephDatastoreTemplate) Render(pkg urknall.Package) {
	var UserHome string
    Uid := m.uuid
		if m.cephuser == "root" {
	    UserHome = "/" +  m.cephuser
	  } else {
			UserHome = UserHomePrefix + m.cephuser
		}
		if m.poolname != "" {
			poolname = m.poolname
		} else {
			poolname = DefaultPoolname
		}

		pkg.AddCommands("cephdatastore",
		Shell("mkdir -p "+UserHome+"/ceph-cluster"),
		Shell("cd "+UserHome+"/ceph-cluster;ceph auth get-or-create client.libvirt mon 'allow r' osd 'allow class-read object_prefix rbd_children, allow rwx pool="+poolname+"'"),
		Shell("cd "+UserHome+"/ceph-cluster;ceph auth get-key client.libvirt | tee client.libvirt.key"),
		Shell("cd "+UserHome+"/ceph-cluster;ceph auth get client.libvirt -o ceph.client.libvirt.keyring"),
		Shell("cd "+UserHome+"/ceph-cluster;cp ceph.client.* /etc/ceph"),
		Shell("cd "+UserHome+"/ceph-cluster; "+fmt.Sprintf(Echo,Uid)+" >uid"),
		WriteFile(UserHome + "/ceph-cluster" + "/secret.xml",fmt.Sprintf(Xml,Uid),"root",644),
		InstallPackages("libvirt-bin"),
		Shell("cd "+UserHome+"/ceph-cluster;sudo virsh secret-define secret.xml"),
		Shell("cd "+UserHome+"/ceph-cluster;"+ fmt.Sprintf(Setval,Uid)),
	)
}
