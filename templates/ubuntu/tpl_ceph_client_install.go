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

type UbuntuCephClientInstall struct{}

func (tpl *UbuntuCephClientInstall) Render(p urknall.Package) {
	p.AddTemplate("cephclient", &UbuntuCephClientInstallTemplate{})
}

func (tpl *UbuntuCephClientInstall) Options(t *templates.Template) {}

func (tpl *UbuntuCephClientInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuCephClientInstall{},inputs)
}

type UbuntuCephClientInstallTemplate struct{}

func (m *UbuntuCephClientInstallTemplate) Render(pkg urknall.Package) {

  pkg.AddCommands("update",
    Shell("apt-get update"),
  )
  // pkg.AddCommands("ceph-common",
	// 	Shell("apt-get install ceph-common"),
	// )

}
