package ubuntu

import (
  "os"
  "fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

var ubuntucephclusterinstall *UbuntuCephClusterInstall

const (
  CEPHUSER = "CephUser"
  CEPHPASSWORD = "CephPassword"
  CLUSTERID = "ClusterId"
  OSDs       = "Osds"
  CLIENTHOST = "ClientHostName"
  CLIENTIP = "ClientIP"
  CLIENTUSER = "ClientUser"
  CLIENTPASSWORD = "ClientPassword"
)


func init() {
	ubuntucephclusterinstall = &UbuntuCephClusterInstall{}
	templates.Register("UbuntuCephClusterInstall", ubuntucephclusterinstall)
}

type UbuntuCephClusterInstall struct {
	cephuser string
}

func (tpl *UbuntuCephClusterInstall) Options(t *templates.Template) {
	if cephuser, ok := t.Options[CephUser]; ok {
		tpl.cephuser = cephuser
	}
}

func (tpl *UbuntuCephClusterInstall) Render(p urknall.Package) {
	p.AddTemplate("cephcluster", &UbuntuCephClusterInstallTemplate{
		cephuser: tpl.cephuser,
		cephhome: UserHomePrefix + tpl.cephuser,
	})
}

func (tpl *UbuntuCephClusterInstall) Run(target urknall.Target, inputs []string) error {
	return urknall.Run(target, &UbuntuCephClusterInstall{}, inputs)
}

type UbuntuCephClusterInstallTemplate struct {
	cephuser string
	cephhome string
}

func (m *UbuntuCephClusterInstallTemplate) Render(pkg urknall.Package) {
	host, _ := os.Hostname()
	ip := findIps()
	CephUser := m.cephuser
	CephHome := m.cephhome

	pkg.AddCommands("install-depends",
		InstallPackages("apt-transport-https  sudo openssh-server ntp sshpass"),
	)
	pkg.AddCommands("install-ceph",
		InstallPackages("ceph-deploy ceph-common ceph-mds dnsmasq ceph"),
	)

	// pkg.AddCommands("cephuser_add",
	//  AddUser(CephUser,false),
	// )
	//
	// pkg.AddCommands("cephuser_sudoer",
	//   Shell("echo '"+CephUser+" ALL = (root) NOPASSWD:ALL' | sudo tee /etc/sudoers.d/"+CephUser+""),
	// )
	//
	// pkg.AddCommands("chmod_sudoer",
	//   Shell("sudo chmod 0440 /etc/sudoers.d/"+CephUser+""),
	// )

	pkg.AddCommands("etchost",
		Shell("echo '"+ip+" "+host+"' >> /etc/hosts"),
	)

	if _, err := os.Stat(CephHome + "/.ssh/id_rsa"); err != nil {
		pkg.AddCommands("ssh-keygen",
			Mkdir(CephHome+"/.ssh", CephUser, 0700),
			AsUser(CephUser, Shell("ssh-keygen -N '' -t rsa -f "+CephHome+"/.ssh/id_rsa")),
			AsUser(CephUser, Shell("cp "+CephHome+"/.ssh/id_rsa.pub "+CephHome+"/.ssh/authorized_keys")),
		)
	}

	pkg.AddCommands("ssh_known_hosts",
		WriteFile(CephHome+"/.ssh/ssh_config", StrictHostKey, CephUser, 0755),
		WriteFile(CephHome+"/.ssh/config", fmt.Sprintf(SSHHostConfig, host, host, CephUser), CephUser, 0755),
	)

	pkg.AddCommands("new-cluster",
		AsUser(CephUser, Shell("mkdir "+CephHome+"/ceph-cluster")),
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy new "+host+" ")),
	)
	pkg.AddCommands("write_cephconf",
		AsUser(CephUser, Shell("echo 'osd_pool_default_size = 2' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd crush chooseleaf type = 1' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'mon_pg_warn_max_per_osd = 0' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd max object name len = 256' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd max object namespace len = 64' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'rbd default features = 1' >> "+CephHome+"/ceph-cluster/ceph.conf")),
	)
	pkg.AddCommands("mon-init",
		AsUser(CephUser, Shell("ceph-deploy mon create-initial")),
	)
}
