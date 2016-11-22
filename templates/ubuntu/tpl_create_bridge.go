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
	"fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Bridgename = "Bridgename"
	Dnsnames   = "Dnsnames"
	PhyDev     = "PhyDev"
	Network    = "Network"
	Netmask    = "Netmask"
	Gateway    = "Gateway"
	Host       = "Host"
	Interface  = `auto lo
iface lo inet loopback

auto %s
auto %s
iface %s inet static
  address %s
	network %s
  netmask %s
  gateway %s
bridge_ports %s
dns-nameservers %s
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
	dnsnames   string
	host       string
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
	if dnsnames, ok := t.Options[Dnsnames]; ok {
		tpl.dnsnames = dnsnames
	}
	if host, ok := t.Options[Host]; ok {
		tpl.host = host
	}
}

func (tpl *UbuntuCreateBridge) Render(p urknall.Package) {
	p.AddTemplate("kvm_network", &UbuntuCreateBridgeTemplate{
		bridgename: tpl.bridgename,
		phydev:     tpl.phydev,
		network:    tpl.network,
		netmask:    tpl.netmask,
		gateway:    tpl.gateway,
		dnsnames:   tpl.dnsnames,
		host:       tpl.host,
	})
}

func (tpl *UbuntuCreateBridge) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, tpl, inputs)
}

type UbuntuCreateBridgeTemplate struct {
	bridgename string
	phydev     string
	network    string
	netmask    string
	gateway    string
	dnsnames   string
	host       string
}

func (m *UbuntuCreateBridgeTemplate) Render(pkg urknall.Package) {
	var dnsnames string
	ip := m.host
	bridgename := m.bridgename
	phydev := m.phydev
	network := m.network
	netmask := m.netmask
	gateway := m.gateway

	if m.phydev == "" {
		phydev = "eth0"
	}

	if m.bridgename == "" {
		bridgename = "one"
	}
	if m.dnsnames == "" {
		dnsnames = "8.8.8.8 8.8.4.4"
	} else {
		dnsnames = m.dnsnames
	}

	pkg.AddCommands("configure",
		Shell("apt-get install -y bridge-utils"),
		Shell("cp /etc/network/interfaces /etc/network/bkinterfaces"),
		WriteFile("/etc/network/interfaces", fmt.Sprintf(Interface, phydev, bridgename, bridgename, ip, network, netmask, gateway, phydev, dnsnames), "root", 0644),
	)
	pkg.AddCommands("create-bridge",
		Shell("brctl addbr "+bridgename+""),
		Shell("brctl show"),
	)

}
