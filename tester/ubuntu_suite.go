package tester

import (
  "io"
   "os"
   "bytes"
   "fmt"
    "github.com/megamsys/megdc/handler"
    "github.com/megamsys/libgo/cmd"
    "testing"
    "errors"
  	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(&HostSuite{})

type HostSuite struct {
}


type Temp struct {
	All              bool        `json:"all"`
	LvmInstall       bool        `json:"lvminstall"`
	FormatPartitions bool        `json:"formatfartiontions"`
	DeletePartitions bool        `json:"deletepartitions"`
  NetworkInfo     bool        `json:"networkinfo"`
	ParseExistLvm    bool        `json:"parseexistlvm"`
	RemoveLvm        bool        `json:"removelvm"`
	Host             string      `json:"ipaddress"`
	Disks            cmd.MapFlag `json:"osds"`
	Mount            cmd.MapFlag `json:"mount"`
	LvPaths          cmd.MapFlag `json:"lvpaths"`
	VgName           string      `json:"vgname"`
	PvName           string      `json:"pvname"`
	Username         string      `json:"username"`
	Password         string      `json:"password"`
	UserMail         string      `json:"email"`
}

func (s *HostSuite) TestSetConf(c *check.C) {
}

func runner(packages []string, i *handler.WrappedParms) string{
  var outBuffer bytes.Buffer

 inputs := []string{"email=info@megam.io"}

 writer := io.MultiWriter(&outBuffer, os.Stdout)
  i.IfNoneAddPackages(packages)
  if h, err := handler.NewHandler(i); err != nil {
    return err.Error()
  } else if err := h.Run(writer,inputs); err != nil {

    fmt.Println(err)
    return err.Error()
  }

  s := outBuffer.String()
  return s
}

func (s *HostSuite) TestRunners(c *check.C) {
  fmt.Println("**************************")
  z := Temp{ NetworkInfo: true, Host: "192.168.0.143", Username: "megam", Password: "megam"}
  f := handler.NewWrap(&z)
  ss := runner([]string{"NetworkInfo"}, f)

   fmt.Println(ss)

   err := errors.New("testing")
   fmt.Println(err)
	c.Assert(nil, check.NotNil)
}
// */
