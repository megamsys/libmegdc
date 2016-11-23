
package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)


var ubuntuparseexistlvm *UbuntuParseExistLvm

func init() {
	ubuntuparseexistlvm = &UbuntuParseExistLvm{}
	templates.Register("UbuntuParseExistLvm", ubuntuparseexistlvm)
}

type UbuntuParseExistLvm struct{}

func (tpl *UbuntuParseExistLvm) Render(p urknall.Package) {
	p.AddTemplate("parselvms", &UbuntuParseExistLvmTemplate{})
}

func (tpl *UbuntuParseExistLvm) Options(t *templates.Template) {}

func (tpl *UbuntuParseExistLvm) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuParseExistLvm{},inputs)
}

type UbuntuParseExistLvmTemplate struct{}

func (m *UbuntuParseExistLvmTemplate) Render(pkg urknall.Package) {

  pkg.AddCommands("list-pv",
    Shell("pvdisplay | grep \"PV Name\""),
  )
  pkg.AddCommands("list-vg",
    Shell("vgdisplay | grep \"VG Name\""),
  )
  pkg.AddCommands("list-lv",
    Shell("lvdisplay | grep \"LV Path\""),
  )
}
