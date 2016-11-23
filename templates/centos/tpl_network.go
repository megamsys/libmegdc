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
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Bridge     = "Bridge"
  Iptype    = "Iptype"
	Ip   = "Ip"
  Size = "Size"
	Networkmask  = "Networkmask"
	Gatewayip   = "Gatewayip"
  Dns1  = "Dns1"
  Dns2  = "Dns2"
  Opennetwork = `NAME   = "open-lvm"
TYPE   = FIXED
BRIDGE = %s
AR = [ TYPE = "%s", IP = "%s", SIZE = "%s" ]
DNS = "%s %s"
GATEWAY    = "%s"
NETWORK_MASK = "%s"`
	)

var centoscreatenetworkopennebula *CentosCreateNetworkOpennebula

func init() {
	centoscreatenetworkopennebula = &CentosCreateNetworkOpennebula{}
	templates.Register("CentosCreateNetworkOpennebula", centoscreatenetworkopennebula)
}

type CentosCreateNetworkOpennebula struct {
	bridge      string
  iptype       string
	ip      string
  size    string
	networkmask      string
	gateway      string
	dns1     string
  dns2    string
	}

func (tpl *CentosCreateNetworkOpennebula) Options(t *templates.Template) {
	if bridge, ok := t.Options[Bridge]; ok {
		tpl.bridge = bridge
	}
  if iptype, ok := t.Options[Iptype]; ok {
		tpl.iptype = iptype
	}
	if ip, ok := t.Options[Ip]; ok {
		tpl.ip = ip
	}
  if size, ok := t.Options[Size]; ok {
    tpl.size = size
  }
	if networkmask, ok := t.Options[Networkmask]; ok {
		tpl.networkmask = networkmask
	}

	if gateway, ok := t.Options[Gatewayip]; ok {
		tpl.gateway = gateway
	}
	if dns1, ok := t.Options[Dns1]; ok {
		tpl.dns1 = dns1
	}
  if dns2, ok := t.Options[Dns2]; ok {
		tpl.dns2 = dns2
	}
}

func (tpl *CentosCreateNetworkOpennebula) Render(p urknall.Package) {
	p.AddTemplate("createnetworkopennebula", &CentosCreateNetworkOpennebulaTemplate{
		bridge:     tpl.bridge,
    iptype:    tpl.iptype,
		ip:   tpl.ip,
    size: tpl.size,
		networkmask:   tpl.networkmask,
		gateway:   tpl.gateway,
		dns1: tpl.dns1,
    dns2: tpl.dns2,
		})
}

func (tpl *CentosCreateNetworkOpennebula) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &CentosCreateNetworkOpennebula{
		bridge:     tpl.bridge,
    iptype:     tpl.iptype,
		ip:   tpl.ip,
    size:  tpl.size,
		networkmask:   tpl.networkmask,
		gateway:   tpl.gateway,
		dns1:    tpl.dns1,
    dns2: tpl.dns2,
	},inputs)
}

type CentosCreateNetworkOpennebulaTemplate struct {
  bridge     string
  iptype    string
	ip   string
  size string
	networkmask string
	gateway string
	dns1  string
  dns2 string
}

func (m *CentosCreateNetworkOpennebulaTemplate) Render(pkg urknall.Package) {

	bridge := m.bridge
   iptype := m.iptype
   ip    := m.ip
   size := m.size
	networkmask := m.networkmask
	gateway := m.gateway
	dns1 := m.dns1
  dns2 := m.dns2
  pkg.AddCommands("create-network",
 WriteFile("/var/lib/open-networkconf",fmt.Sprintf(Opennetwork, bridge, iptype, ip, size, dns1, dns2, gateway, networkmask),  "root" , 0644),
  Shell(" onevnet create /var/lib/open-networkconf"),
 )
 pkg.AddCommands("list",
 Shell("onevnet list"),
 )
}
