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
Nodeip = "nodeip"
Swarmtoken ="swarmtoken"

DockernodejoinConf = `description "Docker join node for baremetal (https://console.megam.io)"
author "Megam Systems(https://www.megam.io)"
# When to start the servicee
start on runlevel [2345]
start on (started networking)
start on (local-filesystems)
# When to stop the service
stop on runlevel [016]
stop on (stopping networking)
stop on (stopped dockernodejoin)
# Automatically restart process if crashed. Tries 0 times every 60 seconds
respawn# start the cibd seed
script
 echo "[$(date -u +%Y-%m-%dT%T.%3NZ)] (sys) dockernodejoin starting" >> /var/log/megam/dockernodejoin.log
 exec  docker run -d swarm join --advertise=%s:2375 token:// %s  >> /var/log/megam/dockernodejoin.log 2>&1
end script
post-start script
   PID=$(status dockernodejoin | egrep -oi '([0-9]+)$' | head -n1)
   echo $PID > /var/run/megam/dockernode.pid
end script`

)

var ubuntudockernodejoininstall *UbuntuDockerNodeJoinInstall

func init() {
	ubuntudockernodejoininstall = &UbuntuDockerNodeJoinInstall{}
	templates.Register("UbuntuDockerNodeJoinInstall", ubuntudockernodejoininstall)
}

type UbuntuDockerNodeJoinInstall struct {
nodeip      string
swarmtoken string

}

func (tpl *UbuntuDockerNodeJoinInstall) Options(t *templates.Template) {
if nodeip, ok := t.Options[Nodeip]; ok {
		tpl.nodeip = nodeip
	}
	if swarmtoken, ok := t.Options[Swarmtoken]; ok {
		tpl.swarmtoken = swarmtoken
	}
}

func (tpl *UbuntuDockerNodeJoinInstall) Render(p urknall.Package) {
	p.AddTemplate("docker", &UbuntuDockerNodeJoinInstallTemplate{
  nodeip : tpl.nodeip,
    swarmtoken : tpl.swarmtoken,
	})
}

func (tpl *UbuntuDockerNodeJoinInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuDockerNodeJoinInstall{
    swarmtoken : tpl.swarmtoken,
  nodeip : tpl.nodeip,
	},inputs)
}

type UbuntuDockerNodeJoinInstallTemplate struct {
nodeip string
  swarmtoken string
}

func (m *UbuntuDockerNodeJoinInstallTemplate) Render(pkg urknall.Package) {
	nodeip  := m.nodeip
swarmtoken :=m.swarmtoken

   pkg.AddCommands("dockernodejoin",
 		WriteFile("/etc/init/dockernodejoin.conf", fmt.Sprintf(DockernodejoinConf, nodeip, swarmtoken),Root, 0644),
   	)
  pkg.AddCommands("start swarm",
		 Shell(" start dockernodejoin"),
	 )
}
