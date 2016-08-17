package ubuntu

import (
	"github.com/megamsys/urknall"
	"github.com/megamsys/libmegdc/templates"
)

const (
  Repos = "Repos"
)

var ubuntuaddrepos *UbuntuAddRepos

func init() {
	ubuntuaddrepos = &UbuntuAddRepos{}
	templates.Register("UbuntuAddRepos", ubuntuaddrepos)
}

type UbuntuAddRepos struct{
  repos string
}

func (tpl *UbuntuAddRepos) Render(p urknall.Package) {
	p.AddTemplate("addrepos", &UbuntuAddReposTemplate{
    repos: tpl.repos,
  })
}

func (tpl *UbuntuAddRepos) Options(t *templates.Template) {
  if repos, ok := t.Options[Repos]; ok {
    tpl.repos = repos
  }
}

func (tpl *UbuntuAddRepos) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &UbuntuAddRepos{
      repos: tpl.repos,
    },inputs)
}

type UbuntuAddReposTemplate struct{
    repos string
}

func (m *UbuntuAddReposTemplate) Render(pkg urknall.Package) {

  pkg.AddCommands("addrepos",
    Shell(m.repos),
		Shell("apt-get update -y"),
  )

}
