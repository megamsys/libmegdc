package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
  Scripts = "Scripts"
)

var ubunturuncustomscripts *UbuntuRunCustomScripts

func init() {
	ubunturuncustomscripts = &UbuntuRunCustomScripts{}
	templates.Register("UbuntuRunCustomScripts", ubunturuncustomscripts)
}

type UbuntuRunCustomScripts struct{
  scripts []string
}

func (tpl *UbuntuRunCustomScripts) Render(p urknall.Package) {
	p.AddTemplate("runscripts", &UbuntuRunCustomScriptsTemplate{
    scripts: tpl.scripts,
  })
}

func (tpl *UbuntuRunCustomScripts) Options(t *templates.Template) {
  if scripts, ok := t.Maps[Scripts]; ok {
    tpl.scripts = scripts
  }
}

func (tpl *UbuntuRunCustomScripts) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuRunCustomScripts{
      scripts: tpl.scripts,
    },inputs)
}

type UbuntuRunCustomScriptsTemplate struct{
    scripts []string
}

func (m *UbuntuRunCustomScriptsTemplate) Render(pkg urknall.Package) {
  path := "/var/lib/urknall/runscripts.sh"
  writeScripts(m.scripts,path)
  pkg.AddCommands("shell-scripts",
    Shell("chmod 755 " + path),
    Shell(path),
  )

}
