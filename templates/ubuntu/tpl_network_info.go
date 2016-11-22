
package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
  "fmt"
)

const (
  Grep =`grep %s /etc/network/interfaces | awk -F ' ' '{print $2}'`
	Iface = `ifconfig | grep -B1 "inet addr:%s" | awk '$1!="inet" && $1!="--" {print $1}'`
)

var ubuntunetworkinfo *UbuntuNetworkInfo

func init() {
	ubuntunetworkinfo = &UbuntuNetworkInfo{}
	templates.Register("UbuntuNetworkInfo", ubuntunetworkinfo)
}

type UbuntuNetworkInfo struct{
	host string
}

func (tpl *UbuntuNetworkInfo) Render(p urknall.Package) {
	p.AddTemplate("networkinfos", &UbuntuNetworkInfoTemplate{
		host: tpl.host,
	})
}

func (tpl *UbuntuNetworkInfo) Options(t *templates.Template) {
	if host, ok := t.Options[HOST]; ok {
		tpl.host = host
	}
}

func (tpl *UbuntuNetworkInfo) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, tpl,inputs)
}

type UbuntuNetworkInfoTemplate struct{
	host string
}

func (m *UbuntuNetworkInfoTemplate) Render(pkg urknall.Package) {
  pkg.AddCommands("interface",
    Shell(fmt.Sprintf(Iface, m.host)),
			Shell(fmt.Sprintf(Grep,"network")),
			Shell(fmt.Sprintf(Grep,"gateway")),
			Shell(fmt.Sprintf(Grep,"broadcast")),
			Shell("grep dns-nameservers /etc/network/interfaces | awk -F ' ' '{print $2 \" \" $3}'"),
			Shell(fmt.Sprintf(Grep,"dns-search")),
  )
}
