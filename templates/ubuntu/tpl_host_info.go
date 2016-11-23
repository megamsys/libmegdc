
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

func (tpl *UbuntuHostInfo) Run(target urknall.Target,inputs map[string]string) error {
	return urknall.Run(target, &UbuntuHostInfo{},inputs)
}

type UbuntuHostInfoTemplate struct{}

func (m *UbuntuHostInfoTemplate) Render(pkg urknall.Package) {
  pkg.AddCommands("get_config",
    Shell("free -m"),
		Shell("lsblk"),
		Shell("lscpu | grep \"CPU(s):\" "),
		Shell("hostname"),
		Shell("grep PRETTY_NAME /etc/*-release | awk -F '=\"' '{print $2}'"),
  )
	pkg.AddCommands("kvm-check",
				Shell("if [ -c /dev/kvm ]; then echo 'KVM acceleration can be used'; else echo 'KVM acceleration can not be used'; fi;"),
	)

}
