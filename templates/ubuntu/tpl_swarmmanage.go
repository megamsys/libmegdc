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
Swarmip = "swarmip"
Token ="token"

Swarmconf = `description "Swarm manage for baremetal (https://console.megam.io)"
author "Megam Systems(https://www.megam.io)"
# When to start the servicee
start on runlevel [2345]
start on (started networking)
start on (local-filesystems)
# When to stop the service
stop on runlevel [016]
stop on (stopping networking)
stop on (stopped swarmmanage)
# Automatically restart process if crashed. Tries 0 times every 60 seconds
respawn# start the cibd seed
script
 echo "[$(date -u +%Y-%m-%dT%T.%3NZ)] (sys) swarmmanage starting" >> /var/log/megam/swarmmanage.log
 exec  docker run -t -p 2375:2375 -t swarm manage -H :2375 --addr %s:2375  token://%s
end script
post-start script
   PID=$(status swarmmanage | egrep -oi '([0-9]+)$' | head -n1)
   echo $PID > /var/run/megam/swarmmanage.pid
end script`

)

var ubuntuswarmmanageinstall *UbuntuSwarmManageInstall

func init() {
	ubuntuswarmmanageinstall = &UbuntuSwarmManageInstall{}
	templates.Register("UbuntuSwarmManageInstall", ubuntuswarmmanageinstall)
}

type UbuntuSwarmManageInstall struct {
swarmip      string
token string

}

func (tpl *UbuntuSwarmManageInstall) Options(t *templates.Template) {
if swarmip, ok := t.Options[Swarmip]; ok {
		tpl.swarmip = swarmip
	}
	if token, ok := t.Options[Token]; ok {
		tpl.token = token
	}
}

func (tpl *UbuntuSwarmManageInstall) Render(p urknall.Package) {
	p.AddTemplate("swarmmanage", &UbuntuSwarmManageInstallTemplate{
    swarmip : tpl.swarmip,
    token   : tpl.token,
	})
}

func (tpl *UbuntuSwarmManageInstall) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target,tpl,inputs)
}

type UbuntuSwarmManageInstallTemplate struct {
swarmip string
  token string
}

func (m *UbuntuSwarmManageInstallTemplate) Render(pkg urknall.Package) {
	swarmip  := m.swarmip
token :=m.token

   pkg.AddCommands("swarmmanage",
 		WriteFile("/etc/init/swarmmanage.conf", fmt.Sprintf(Swarmconf, swarmip, token),Root, 0644),
   	)

    pkg.AddCommands("start",
  Shell("start swarmmanage"),
  )

}
