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
Hostip = "Hostip"
Root ="root"
Subnetmask  = "Subnetmask"
DockerConf = `
description "Docker Daemon - node agent on baremetal"
author "Megam Systems(https://www.megam.io)"
# When to start the service
start on runlevel [2345]
start on (started networking)
# When to stop the service
stop on runlevel [016]
stop on (stopping networking)
stop on (stopped dockerdaemon)
# Automatically restart process if crashed. Tries 0 times every 60 seconds
respawn
respawn limit 0 60
# start the  dockerdaemon as docker node agent.
script
 echo "[$(date -u +%Y-%m-%dT%T.%3NZ)] (sys) dockerdaemon starting" >> /var/log/megam/dockerdaemon.log
 exec docker daemon -D -H tcp://%s:2375 --bip %s >> /var/log/megam/dockerdaemon.log 2>&1
end script
post-start script
   PID=$(status dockerdaemon | egrep -oi '([0-9]+)$' | head -n1)
   echo $PID > /var/run/megam/dockerdaemon.pid
end scriptpost-stop script
   rm -f /var/run/megam/dockerdaemon.pid
end script`

)

var ubuntudockerinstall *UbuntuDockerInstall

func init() {
	ubuntudockerinstall = &UbuntuDockerInstall{}
	templates.Register("UbuntuDockerInstall", ubuntudockerinstall)
}

type UbuntuDockerInstall struct {
hostip      string
subnetmask string

}

func (tpl *UbuntuDockerInstall) Options(t *templates.Template) {
if hostip, ok := t.Options[Hostip]; ok {
		tpl.hostip = hostip
	}
	if subnetmask, ok := t.Options[Subnetmask]; ok {
		tpl.subnetmask = subnetmask
	}
}

func (tpl *UbuntuDockerInstall) Render(p urknall.Package) {
	p.AddTemplate("docker", &UbuntuDockerInstallTemplate{

    hostip : tpl.hostip,
    subnetmask : tpl.subnetmask,
	})
}

func (tpl *UbuntuDockerInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuDockerInstall{
    hostip : tpl.hostip,
    subnetmask : tpl.subnetmask,
	},inputs)
}

type UbuntuDockerInstallTemplate struct {
  hostip string
  subnetmask string
}

func (m *UbuntuDockerInstallTemplate) Render(pkg urknall.Package) {
	hostip := m.hostip
subnetmask :=m.subnetmask

	pkg.AddCommands("dockerinstall",
		 Shell(" curl -fsSL https://get.docker.com/ | sh"),
	 )

	pkg.AddCommands("DockerConf",
		WriteFile("/etc/init/dockerdaemon.conf", fmt.Sprintf(DockerConf, hostip, subnetmask), Root, 0644),
	)

	pkg.AddCommands("mkdir",
		Shell("mkdir -p /var/log/megam" ),
    Shell("mkdir -p /var/run/megam"),
	)

	pkg.AddCommands("stop",
		Shell("stop docker"),
	)

  pkg.AddCommands("start dockerdaemon",
  Shell("start dockerdaemon"),
)
}
