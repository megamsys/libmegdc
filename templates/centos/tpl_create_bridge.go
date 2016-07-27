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
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Bridgename     = "Bridgename"
  PhyDev    = "PhyDev"

	Netmask  = "Netmask"
	Gateway   = "Gateway"

	 Ifcfgeth0 = `
	 DEVICE=%s
ONBOOT=yes
BRIDGE=%s`
 Ifcfgone = `
 DEVICE=%s
TYPE=Bridge
BOOTPROTO=static
IPADDR=%s
NETMASK=%s
GATEWAY=%s
ONBOOT=yes
STP=no`
	)

var centoscreatebridge *CentosCreateBridge

func init() {
	centoscreatebridge = &CentosCreateBridge{}
	templates.Register("CentosCreateBridge", centoscreatebridge)
}

type CentosCreateBridge struct {
	bridgename      string
  phydev       string
	netmask      string
	gateway      string

	}

func (tpl *CentosCreateBridge) Options(t *templates.Template) {
	if bridgename, ok := t.Options[Bridgename]; ok {
		tpl.bridgename = bridgename
	}
  if phydev, ok := t.Options[PhyDev]; ok {
		tpl.phydev = phydev
	}

	if netmask, ok := t.Options[Netmask]; ok {
		tpl.netmask = netmask
	}

	if gateway, ok := t.Options[Gateway]; ok {
		tpl.gateway = gateway
	}

}

func (tpl *CentosCreateBridge) Render(p urknall.Package) {
	p.AddTemplate("bridge", &CentosCreateBridgeTemplate{
		bridgename:     tpl.bridgename,
    phydev:    tpl.phydev,
		netmask:   tpl.netmask,
		gateway:   tpl.gateway,

		})
}

func (tpl *CentosCreateBridge) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &CentosCreateBridge{
		bridgename:     tpl.bridgename,
    phydev:     tpl.phydev,
		netmask:   tpl.netmask,
		gateway:   tpl.gateway,
	},inputs)
}

type CentosCreateBridgeTemplate struct {
  bridgename     string
  phydev    string
	netmask string
	gateway string
}

func (m *CentosCreateBridgeTemplate) Render(pkg urknall.Package) {
	ip := IP("")
	bridgename := m.bridgename
  phydev := m.phydev
	netmask := m.netmask
	gateway := m.gateway

	pkg.AddCommands("bridgeutils",
		 Shell("yum install -y bridge-utils"),
		 )
	 pkg.AddCommands("create-bridge",
 	 Shell("brctl addbr "+bridgename+""),
 	)
	pkg.AddCommands("interfaces",
  WriteFile("/etc/sysconfig/network-scripts/ifcfg-"+phydev+"", fmt.Sprintf(Ifcfgeth0, phydev, bridgename ), "root", 0644),
	WriteFile("/etc/sysconfig/network-scripts/ifcfg-"+bridgename+"", fmt.Sprintf(Ifcfgone, bridgename, ip, netmask, gateway ), "root", 0644),
	)
	pkg.AddCommands("network",
	 Shell("service network restart"),
	)
	pkg.AddCommands("list-bridge",
   Shell("brctl show"),
	)

}
