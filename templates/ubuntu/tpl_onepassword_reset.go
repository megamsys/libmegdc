package ubuntu

import (
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

const (
	ONEUSER     = "OneUser"
  ONEPASSWORD = "OnePassword"
	)

var ubuntuonepasswordreset *UbuntuOnePasswordReset

func init() {
	ubuntuonepasswordreset = &UbuntuOnePasswordReset{}
	templates.Register("UbuntuOnePasswordReset", ubuntuonepasswordreset)
}
type UbuntuOnePasswordReset struct{
oneuser string
onepassword string
}

func (tpl *UbuntuOnePasswordReset) Options(t *templates.Template) {
	if oneuser, ok := t.Options[ONEUSER]; ok {
		tpl.oneuser = oneuser
	}
  if onepassword, ok := t.Options[ONEPASSWORD]; ok {
		tpl.onepassword = onepassword
	}

}

type UbuntuOnePasswordResetTemplate struct {
  oneuser     string
  onepassword    string
}

func (tpl *UbuntuOnePasswordReset) Render(p urknall.Package) {
	p.AddTemplate("resetonepassword", &UbuntuOnePasswordResetTemplate{
		oneuser:     tpl.oneuser,
    onepassword:    tpl.onepassword,
		})
}

func (tpl *UbuntuOnePasswordReset) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuOnePasswordResetTemplate{
		oneuser:     tpl.oneuser,
  onepassword:     tpl.onepassword,
	},inputs)
}

func (m *UbuntuOnePasswordResetTemplate) Render(pkg urknall.Package) {
	oneuser := m.oneuser
  onepassword := m.onepassword

	 pkg.AddCommands("resetonepassword",
     Shell("echo '"+oneuser+":"+onepassword+"' >/var/lib/one/.one/one_auth"),
 	)
	pkg.AddCommands("one-restart",
		Shell("service opennebula restart"),
		Shell("service opennebula-sunstone restart")
 )
}
