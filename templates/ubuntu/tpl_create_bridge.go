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

package ubuntu

import (
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Bridgename = "Bridgename"
	PhyDev     = "PhyDev"
	Network    = "Network"
	Netmask    = "Netmask"
	Gateway    = "Gateway"
	Dnsname1   = "Dnsname1"
	Dnsname2   = "Dnsname2"
	Interface  = `auto lo
iface lo inet loopback

auto eth0
auto %s
iface %s inet static
  address %s
  network %s
  netmask %s
  gateway %s
bridge_ports %s
dns-nameservers %s %s
source /etc/network/interfaces.d/*.cfg`
)

var ubuntucreatebridge *UbuntuCreateBridge

func init() {
	ubuntucreatebridge = &UbuntuCreateBridge{}
	templates.Register("UbuntuCreateBridge", ubuntucreatebridge)
}

type UbuntuCreateBridge struct {
	bridgename string
	phydev     string
	network    string
	netmask    string
	gateway    string
	dnsname1   string
	dnsname2   string
}

func (tpl *UbuntuCreateBridge) Options(t *templates.Template) {
	if bridgename, ok := t.Options[Bridgename]; ok {
		tpl.bridgename = bridgename
	}
	if phydev, ok := t.Options[PhyDev]; ok {
		tpl.phydev = phydev
	}
	if network, ok := t.Options[Network]; ok {
		tpl.network = network
	}
	if netmask, ok := t.Options[Netmask]; ok {
		tpl.netmask = netmask
	}

	if gateway, ok := t.Options[Gateway]; ok {
		tpl.gateway = gateway
	}
	if dnsname1, ok := t.Options[Dnsname1]; ok {
		tpl.dnsname1 = dnsname1
	}
	if dnsname2, ok := t.Options[Dnsname2]; ok {
		tpl.dnsname2 = dnsname2
	}
}

func (tpl *UbuntuCreateBridge) Render(p urknall.Package) {
	p.AddTemplate("bridge", &UbuntuCreateBridgeTemplate{
		bridgename: tpl.bridgename,
		phydev:     tpl.phydev,
		network:    tpl.network,
		netmask:    tpl.netmask,
		gateway:    tpl.gateway,
		dnsname1:   tpl.dnsname1,
		dnsname2:   tpl.dnsname2,
	})
}

func (tpl *UbuntuCreateBridge) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, &UbuntuCreateBridge{
		bridgename: tpl.bridgename,
		phydev:     tpl.phydev,
		network:    tpl.network,
		netmask:    tpl.netmask,
		gateway:    tpl.gateway,
		dnsname1:   tpl.dnsname1,
		dnsname2:   tpl.dnsname2,
	}, inputs)
}

type UbuntuCreateBridgeTemplate struct {
	bridgename string
	phydev     string
	network    string
	netmask    string
	gateway    string
	dnsname1   string
	dnsname2   string
}

func (m *UbuntuCreateBridgeTemplate) Render(pkg urknall.Package) {
	ip := IP("")
	bridgename := m.bridgename
	phydev := m.phydev
	network := m.network
	netmask := m.netmask
	gateway := m.gateway
	dnsname1 := m.dnsname1
	dnsname2 := m.dnsname2
	pkg.AddCommands("bridgeutils",
		Shell("apt-get install -y bridge-utils"),
	)
	pkg.AddCommands("interfaces",
		WriteFile("/etc/network/interfaces", fmt.Sprintf(Interface, bridgename, bridgename, ip, network, netmask, gateway, phydev, dnsname1, dnsname2), "root", 0644),
	)
	pkg.AddCommands("create-bridge",
		Shell("brctl addbr "+bridgename+""),
	)
	pkg.AddCommands("list-bridge",
		Shell("brctl show"),
	)

}
