package ubuntu

import (
  "fmt"
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"
)

var ubuntucephclusterinstall *UbuntuCephClusterInstall

const (
  USERNAME  = "Username"
  CLUSTERID = "ClusterId"
  OSDs       = "Osds"
  HOST       = "Host"
  CLIENTHOST = "ClientHostName"
  CLIENTIP = "ClientIP"
  CLIENTUSER = "ClientUser"
  CLIENTPASSWORD = "ClientPassword"
  CLIENTKEY  = "ClientPrivatKey"
  POOLNAME = "PoolName"
  DefaultPoolname = "one"
)

var CephHome, ClientHome,poolname string


func init() {
	ubuntucephclusterinstall = &UbuntuCephClusterInstall{}
	templates.Register("UbuntuCephClusterInstall", ubuntucephclusterinstall)
}

type UbuntuCephClusterInstall struct {
	cephuser string
  host string
  hostname string
  poolname string
}

func (tpl *UbuntuCephClusterInstall) Options(t *templates.Template) {
	if cephuser, ok := t.Options[USERNAME]; ok {
		tpl.cephuser = cephuser
	}
  if host, ok := t.Options[HOST]; ok {
    tpl.host = host
  }
  if hostname, ok := t.Options[CLIENTHOST]; ok {
    tpl.hostname = hostname
  }
  if poolname, ok := t.Options[POOLNAME]; ok {
    tpl.poolname = poolname
  }
}

func (tpl *UbuntuCephClusterInstall) Render(p urknall.Package) {
	p.AddTemplate("cephcluster", &UbuntuCephClusterInstallTemplate{
		cephuser: tpl.cephuser,
    host: tpl.host,
    hostname: tpl.hostname,
    poolname: tpl.poolname,
	})
}

func (tpl *UbuntuCephClusterInstall) Run(target urknall.Target, inputs map[string]string) error {
	return urknall.Run(target,tpl, inputs)
}

type UbuntuCephClusterInstallTemplate struct {
	cephuser string
  host string
  hostname string
  poolname string
}

func (m *UbuntuCephClusterInstallTemplate) Render(pkg urknall.Package) {
  ip := m.host
  host := m.hostname
	CephUser := m.cephuser
  if m.cephuser == "root" {
    CephHome = "/root"
  } else {
    CephHome = UserHomePrefix + m.cephuser
  }
  if m.poolname != "" {
    poolname = m.poolname
  } else {
    poolname = DefaultPoolname
  }

	pkg.AddCommands("depends_install",
		InstallPackages("apt-transport-https  sudo openssh-server ntp sshpass"),
	)
	pkg.AddCommands("install",
		InstallPackages("ceph-deploy ceph-common ceph-mds ceph"),
	)

		pkg.AddCommands("access_prepare",
      Shell("echo '"+ip+" "+host+"' >> /etc/hosts"),
			Mkdir(CephHome+"/.ssh", CephUser, 0700),
			AsUser(CephUser, Shell("ssh-keygen -N '' -t rsa -f "+CephHome+"/.ssh/id_rsa")),
			AsUser(CephUser, Shell("cat "+CephHome+"/.ssh/id_rsa.pub >>"+CephHome+"/.ssh/authorized_keys")),
      WriteFile(CephHome+"/.ssh/ssh_config", StrictHostKey, CephUser, 0755),
  		WriteFile(CephHome+"/.ssh/config", fmt.Sprintf(SSHHostConfig, host, host, CephUser), CephUser, 0755),
		)

	pkg.AddCommands("create_new",
		AsUser(CephUser, Shell("mkdir -p "+CephHome+"/ceph-cluster")),
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy new "+host)),
	)
	pkg.AddCommands("configure",
    AsUser(CephUser, Shell("sed -i \"s/^[ ]*mon_host.*/mon_host = "+ip+"/\" "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd_pool_default_size = 2' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd crush chooseleaf type = 1' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'mon_pg_warn_max_per_osd = 0' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd max object name len = 256' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'osd max object namespace len = 64' >> "+CephHome+"/ceph-cluster/ceph.conf")),
		AsUser(CephUser, Shell("echo 'rbd default features = 1' >> "+CephHome+"/ceph-cluster/ceph.conf")),
	)
	pkg.AddCommands("mon_init",
		AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;ceph-deploy mon create-initial")),
    AsUser(CephUser, Shell("cd "+CephHome+"/ceph-cluster;sudo cp ceph.client.* /etc/ceph/")),
    Shell("sudo chmod +r /etc/ceph/ceph.client.admin.keyring"),
    Shell("sudo chown -R "+CephUser+":"+CephUser+" /etc/ceph/ceph.client.admin.keyring"),
	)
  	pkg.AddCommands("create_pool",
      AsUser(CephUser, Shell("ceph osd pool create "+poolname+" 128")),
    )
}
