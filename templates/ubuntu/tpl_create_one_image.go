/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package ubuntu

import (
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
	//"github.com/megamsys/libgo/cmd"
)

const (
	DataStoreId = "DataStoreId"
	Name        = "Name"
	ImageUrl    = "ImageUrl"
	Type        = "Type"
)

var ubuntucreateoneimage *UbuntuCreateOneImage

func init() {
	ubuntucreateoneimage = &UbuntuCreateOneImage{}
	templates.Register("UbuntuCreateOneImage", ubuntucreateoneimage)
}

type UbuntuCreateOneImage struct {
	datastoreid string
	name        string
	imageurl    string
	imagetype   string
	imagesize   string
}

func (tpl *UbuntuCreateOneImage) Options(t *templates.Template) {
	if datastoreid, ok := t.Options[DataStoreId]; ok {
		tpl.datastoreid = datastoreid
	}
	if name, ok := t.Options[Name]; ok {
		tpl.name = name
	}
	if imageurl, ok := t.Options[ImageUrl]; ok {
		tpl.imageurl = imageurl
	}
	if imagetype, ok := t.Options[Type]; ok {
		tpl.imagetype = imagetype
	}
	if imagesize, ok := t.Options[Size]; ok {
		tpl.imagesize = imagesize
	}

}

func (tpl *UbuntuCreateOneImage) Render(p urknall.Package) {
	p.AddTemplate("createoneimage", &UbuntuCreateOneImageTemplate{
		datastoreid: tpl.datastoreid,
		name:        tpl.name,
		imageurl:    tpl.imageurl,
		imagetype:   tpl.imagetype,
		imagesize:   tpl.imagesize,
	})
}

func (tpl *UbuntuCreateOneImage) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target, &UbuntuCreateOneImage{
		datastoreid: tpl.datastoreid,
		name:        tpl.name,
		imageurl:    tpl.imageurl,
		imagetype:   tpl.imagetype,
		imagesize:   tpl.imagesize,
	}, inputs)
}

type UbuntuCreateOneImageTemplate struct {
	datastoreid string
	name        string
	imageurl    string
	imagetype   string
	imagesize   string
}

func (m *UbuntuCreateOneImageTemplate) Render(pkg urknall.Package) {
	datastoreid := m.datastoreid
	name := m.name
	imageurl := m.imageurl
	imagetype := m.imagetype
	imagesize := m.imagesize

	pkg.AddCommands("downloadImage",
		Shell("mkdir -p /var/lib/megam/images"),
		Shell("cd /var/lib/megam/images; wget "+imageurl+" -O "+name+".tar.gz"),
	)
	pkg.AddCommands("UntarImage",
		Shell("mkdir /var/lib/megam/images/tmp"),
		Shell("cd /var/lib/megam/images;tar xvf "+name+".tar.gz -C /var/lib/megam/images/tmp"),
	)

	pkg.AddCommands("GetImage",
		Shell("cd /var/lib/megam/images/tmp;first_file=$(ls | sort -n | head -1 );mv $first_file "+name+".img"),
		Shell("mv /var/lib/megam/images/tmp/"+name+".img /var/lib/megam/images"),
		Shell("rm -rf /var/lib/megam/images/tmp"),
	)
	if imagetype == "OS" {
		pkg.AddCommands("oneimagecreate",
			Shell("oneimage create --datastore "+datastoreid+" --name "+name+" --type "+imagetype+" --path /var/lib/megam/images/"+name+".img"),
		)
	} else {
		pkg.AddCommands("oneimagecreate",
			Shell("oneimage create --datastore "+datastoreid+" --name "+name+" --type "+imagetype+" --size "+imagesize+" --path /var/lib/megam/images/"+name+".img"),
		)
	}
	pkg.AddCommands("clearCaches",
		RemoveAllCaches("/var/lib/urknall/createoneimage.*"),
	)
}
