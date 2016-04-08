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
	Bridgename     = "Bridgename"
  Port    = "Port"
	)

var ubuntucreatebridge *UbuntuCreateBridge

func init() {
	ubuntucreatebridge = &UbuntuCreateBridge{}
	templates.Register("UbuntuCreateBridge", ubuntucreatebridge)
}

type UbuntuCreateBridge struct {
	bridgename      string
  port       string
	}

func (tpl *UbuntuCreateBridge) Options(t *templates.Template) {
	if bridgename, ok := t.Options[Bridgename]; ok {
		tpl.bridgename = bridgename
	}
  if port, ok := t.Options[Port]; ok {
		tpl.port = port
	}
}

func (tpl *UbuntuCreateBridge) Render(p urknall.Package) {
	p.AddTemplate("bridge", &UbuntuCreateBridgeTemplate{
		bridgename:     tpl.bridgename,
    port:    tpl.port,
		})
}

func (tpl *UbuntuCreateBridge) Run(target urknall.Target) error {
	return urknall.Run(target, &UbuntuCreateBridge{
		bridgename:     tpl.bridgename,
    port:     tpl.port,
	})
}

type UbuntuCreateBridgeTemplate struct {
  bridgename     string
  port    string
}

func (m *UbuntuCreateBridgeTemplate) Render(pkg urknall.Package) {
	bridgename := m.bridgename
  port := m.port

	pkg.AddCommands("bridgeutils",
		 Shell("apt-get install -y bridge-utils"),
		 )
	 pkg.AddCommands("create-bridge",
 	 Shell("brctl addbr "+bridgename+""),
 	)
 	pkg.AddCommands("add-port",
 		Shell("brctl addif "+bridgename+" "+port+""),
 	)

}
