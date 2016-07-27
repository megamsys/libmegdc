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

package debian

import (
   "github.com/megamsys/libmegdc/templates"
		u "github.com/megamsys/libmegdc/templates/ubuntu"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	InfoDriver     = "InfoDriver"
  Vm    = "Vm"
  HostName  = "HostName"
  Networking = "Networking"
	)

var debianattachonehost *DebianAttachOneHost

func init() {
	debianattachonehost = &DebianAttachOneHost{}
	templates.Register("DebianAttachOneHost", debianattachonehost)
}

type DebianAttachOneHost struct {

	infodriver     string
  vm       string
  hostname   string
  network    string
	}

func (tpl *DebianAttachOneHost) Options(t *templates.Template) {
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

func (tpl *DebianAttachOneHost) Render(p urknall.Package) {
	p.AddTemplate("attachonehost", &DebianAttachOneHostTemplate{
		infodriver:     tpl.infodriver,
    vm:    tpl.vm,
    hostname:   tpl.hostname,
    network: tpl.network,
		})
}

func (tpl *DebianAttachOneHost) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &DebianAttachOneHost{
		infodriver:     tpl.infodriver,
    vm:     tpl.vm,
    hostname:    tpl.hostname,
    network:   tpl.network,
	},inputs)
}

type DebianAttachOneHostTemplate struct {
  infodriver     string
  vm    string
  hostname   string
  network   string
}

func (m *DebianAttachOneHostTemplate) Render(pkg urknall.Package) {
	infodriver := m.infodriver
  vm := m.vm
  hostname := m.hostname
  network := m.network
	 pkg.AddCommands("create-host",
 	u.Shell(" onehost create "+hostname+" --im  "+infodriver+" --vm "+vm+" --net "+network+""),
 	)
  pkg.AddCommands("list",
  u.Shell("onehost list"),
  )
}
