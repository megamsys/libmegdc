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

package centos

import (
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

const (
	BRIDGE_NAME = "Bridge"
	PHY_DEV     = "PhyDev"
)

var centoscreatenetwork *CentosCreateNetwork

func init() {
	centoscreatenetwork = &CentosCreateNetwork{}
	templates.Register("CentosCreateNetwork", centoscreatenetwork)
}

type CentosCreateNetwork struct {
	BridgeName string
	PhyDev     string
}

func (tpl *CentosCreateNetwork) Options(t *templates.Template) {
	if bg, ok := t.Options[BRIDGE_NAME]; ok {
		tpl.BridgeName = bg
	}

	if ph, ok := t.Options[PHY_DEV]; ok {
		tpl.PhyDev = ph
	}
}

func (tpl *CentosCreateNetwork) Render(p urknall.Package) {
	p.AddTemplate("createnetwork", &CentosCreateNetworkTemplate{
		BridgeName: tpl.BridgeName,
		PhyDev:     tpl.PhyDev,
	})
}

func (tpl *CentosCreateNetwork) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosCreateNetwork{
		BridgeName: tpl.BridgeName,
		PhyDev:     tpl.PhyDev,
	},inputs)
}

type CentosCreateNetworkTemplate struct {
	BridgeName string
	PhyDev     string
}

func (m *CentosCreateNetworkTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("ovs-createnetwork",
		Shell("sudo echo '"+"%"+"oneadmin ALL=(root) NOPASSWD: /usr/bin/ovs-vsctl' > //etc/sudoers.d/openvswitch"),
		Shell("sudo echo '"+"%"+"oneadmin ALL=(root) NOPASSWD: /usr/bin/ovs-ofctl' >> //etc/sudoers.d/openvswitch"),
		Shell("sudo ovs-vsctl add-br "+m.BridgeName),
		Shell("sudo echo 'auto "+m.BridgeName+"' >> /etc/network/interfaces"),
		Shell("sudo ovs-vsctl add-port "+m.BridgeName+" "+m.PhyDev+""),
		UpdatePackagesOmitError(),
	)

}
