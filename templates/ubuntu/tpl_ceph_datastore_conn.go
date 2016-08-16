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
  Poolname = "one"
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
}

func (tpl *UbuntuCephDatastore) Options(t *templates.Template) {
	if uuid, ok := t.Options[CLUSTERID]; ok {
		tpl.uuid = uuid
	}
}

func (tpl *UbuntuCephDatastore) Render(p urknall.Package) {
	p.AddTemplate("cephds", &UbuntuCephDatastoreTemplate{
		uuid: tpl.uuid,
	})
}

func (tpl *UbuntuCephDatastore) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuCephDatastore{
		uuid: tpl.uuid,
		},inputs)
}

type UbuntuCephDatastoreTemplate struct {
	uuid string
}

func (m *UbuntuCephDatastoreTemplate) Render(pkg urknall.Package) {
    Uid := m.uuid
		pkg.AddCommands("cephdatastore",
  	AsUser(Ceph_User,Shell("ceph osd pool create "+Poolname+" 128")),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;ceph auth get-or-create client.libvirt mon 'allow r' osd 'allow class-read object_prefix rbd_children, allow rwx pool="+Poolname+"'"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;ceph auth get-key client.libvirt | tee client.libvirt.key"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;ceph auth get client.libvirt -o ceph.client.libvirt.keyring"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;cp ceph.client.* /etc/ceph"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster; "+fmt.Sprintf(Echo,Uid)+" >uid"),
		WriteFile(UserHomePrefix + Ceph_User + "/ceph-cluster" + "/secret.xml",fmt.Sprintf(Xml,Uid),"root",644),
		InstallPackages("libvirt-bin"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;sudo virsh secret-define secret.xml"),
		Shell("cd "+UserHomePrefix + Ceph_User+"/ceph-cluster;"+ fmt.Sprintf(Setval,Uid)),
	)
}
