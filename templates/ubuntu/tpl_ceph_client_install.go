package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)


var ubuntucephclientinstall *UbuntuCephClientInstall

func init() {
	ubuntucephclientinstall = &UbuntuCephClientInstall{}
	templates.Register("UbuntuCephClientInstall", ubuntucephclientinstall)
}

type UbuntuCephClientInstall struct{
	cephuser string
}

func (tpl *UbuntuCephClientInstall) Render(p urknall.Package) {
	p.AddTemplate("cephclient", &UbuntuCephClientInstallTemplate{
		cephuser: tpl.cephuser,
	})
}

func (tpl *UbuntuCephClientInstall) Options(t *templates.Template) {
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
}

func (tpl *UbuntuCephClientInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, tpl,inputs)
}

type UbuntuCephClientInstallTemplate struct{
	cephuser string
}

func (m *UbuntuCephClientInstallTemplate) Render(pkg urknall.Package) {
	var CephHome string
	if m.cephuser == "root" {
		CephHome = "/root"
	} else {
		CephHome = UserHomePrefix + m.cephuser
	}
  pkg.AddCommands("install",
		Shell("apt-get update -y"),
		InstallPackages("ceph-common ceph ceph-mds ca-certificates apt-transport-https"),
		Mkdir(CephHome+"/.ssh",m.cephuser,0700),
	)
}
