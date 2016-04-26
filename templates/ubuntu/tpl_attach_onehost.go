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
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	InfoDriver     = "InfoDriver"
  Vm    = "Vm"
  HostName  = "HostName"
  Networking = "Networking"
	)

var ubuntuattachonehost *UbuntuAttachOneHost

func init() {
	ubuntuattachonehost = &UbuntuAttachOneHost{}
	templates.Register("UbuntuAttachOneHost", ubuntuattachonehost)
}

type UbuntuAttachOneHost struct {

	infodriver     string
  vm       string
  hostname   string
  network    string
	}

func (tpl *UbuntuAttachOneHost) Options(t *templates.Template) {
	if infodriver, ok := t.Options[InfoDriver]; ok {
		tpl.infodriver = infodriver
	}
  if vm, ok := t.Options[Vm]; ok {
		tpl.vm = vm
	}
  if hostname, ok := t.Options[HostName]; ok {
    tpl.hostname = hostname
  }
  if network, ok := t.Options[Networking]; ok {
		tpl.network = network
	}

}

func (tpl *UbuntuAttachOneHost) Render(p urknall.Package) {
	p.AddTemplate("attachonehost", &UbuntuAttachOneHostTemplate{
		infodriver:     tpl.infodriver,
    vm:    tpl.vm,
    hostname:   tpl.hostname,
    network: tpl.network,
		})
}

func (tpl *UbuntuAttachOneHost) Run(target urknall.Target) error {
	return urknall.Run(target, &UbuntuAttachOneHost{
		infodriver:     tpl.infodriver,
    vm:     tpl.vm,
    hostname:    tpl.hostname,
    network:   tpl.network,
	})
}

type UbuntuAttachOneHostTemplate struct {
  infodriver     string
  vm    string
  hostname   string
  network   string
}

func (m *UbuntuAttachOneHostTemplate) Render(pkg urknall.Package) {
	infodriver := m.infodriver
  vm := m.vm
  hostname := m.hostname
  network := m.network

	 pkg.AddCommands("create-host",
 	 Shell(" onehost create "+hostname+" --im  "+infodriver+" --vm "+vm+" --net "+network+""),
 	)
	pkg.AddCommands("list",
	Shell("onehost list"),
	)
}
