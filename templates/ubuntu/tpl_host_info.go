
package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)


var ubuntuhostinfo *UbuntuHostInfo

func init() {
	ubuntuhostinfo = &UbuntuHostInfo{}
	templates.Register("UbuntuHostInfo", ubuntuhostinfo)
}

type UbuntuHostInfo struct{}

func (tpl *UbuntuHostInfo) Render(p urknall.Package) {
	p.AddTemplate("hostinfos", &UbuntuHostInfoTemplate{})
}

func (tpl *UbuntuHostInfo) Options(t *templates.Template) {}

func (tpl *UbuntuHostInfo) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuHostInfo{},inputs)
}

type UbuntuHostInfoTemplate struct{}

func (m *UbuntuHostInfoTemplate) Render(pkg urknall.Package) {

  pkg.AddCommands("memory",
    Shell("free -m"),
  )
  pkg.AddCommands("disk",
		Shell("lsblk"),
	)
  pkg.AddCommands("cpu",
		Shell("lscpu | grep \"CPU(s):\" "),
	)
  pkg.AddCommands("hostname",
		Shell("hostname"),
	)
  pkg.AddCommands("os",
  	Shell("grep PRETTY_NAME /etc/*-release | awk -F '=\"' '{print $2}'"),
		)
}
