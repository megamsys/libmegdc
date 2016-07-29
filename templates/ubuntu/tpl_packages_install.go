package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
  Packages = "Packages"
)

var ubuntucustompackagesinstall *UbuntuCustomPackagesInstall

func init() {
	ubuntucustompackagesinstall = &UbuntuCustomPackagesInstall{}
	templates.Register("UbuntuCustomPackagesInstall", ubuntucustompackagesinstall)
}

type UbuntuCustomPackagesInstall struct{
  packages string
}

func (tpl *UbuntuCustomPackagesInstall) Render(p urknall.Package) {
	p.AddTemplate("packageinstall", &UbuntuCustomPackagesInstallTemplate{
    packages: tpl.packages,
  })
}

func (tpl *UbuntuCustomPackagesInstall) Options(t *templates.Template) {
  if packages, ok := t.Options[Packages]; ok {
    tpl.packages = packages
  }
}

func (tpl *UbuntuCustomPackagesInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuCustomPackagesInstall{
      packages: tpl.packages,
    },inputs)
}

type UbuntuCustomPackagesInstallTemplate struct{
    packages string
}

func (m *UbuntuCustomPackagesInstallTemplate) Render(pkg urknall.Package) {

  pkg.AddCommands("packagesinstall",
    InstallPackages(m.packages),
  )

}
