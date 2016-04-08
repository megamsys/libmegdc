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
	"github.com/megamsys/urknall"
    u "github.com/megamsys/libmegdc/templates/ubuntu"
	//"github.com/megamsys/libgo/cmd"
)

const (
	Bridgename     = "Bridgename"
  Port    = "Port"
	)

var debiancreatebridge *DebianCreateBridge

func init() {
	debiancreatebridge = &DebianCreateBridge{}
	templates.Register("DebianCreateBridge", debiancreatebridge)
}

type DebianCreateBridge struct {
	bridgename      string
  port       string
	}

func (tpl *DebianCreateBridge) Options(t *templates.Template) {
	if bridgename, ok := t.Options[Bridgename]; ok {
		tpl.bridgename = bridgename
	}
  if port, ok := t.Options[Port]; ok {
		tpl.port = port
	}
}

func (tpl *DebianCreateBridge) Render(p urknall.Package) {
	p.AddTemplate("bridge", &DebianCreateBridgeTemplate{
		bridgename:     tpl.bridgename,
    port:    tpl.port,
		})
}

func (tpl *DebianCreateBridge) Run(target urknall.Target) error {
	return urknall.Run(target, &DebianCreateBridge{
		bridgename:     tpl.bridgename,
    port:     tpl.port,
	})
}

type DebianCreateBridgeTemplate struct {
  bridgename     string
  port    string
}

func (m *DebianCreateBridgeTemplate) Render(pkg urknall.Package) {
	bridgename := m.bridgename
  port := m.port

	pkg.AddCommands("bridgeutils",
		  u.Shell("apt-get install -y bridge-utils"),
		 )
	 pkg.AddCommands("create-bridge",
 	 u.Shell("brctl addbr "+bridgename+""),
 	)
 	pkg.AddCommands("add-port",
 		u.Shell("brctl addif "+bridgename+" "+port+""),
 	)

}
