package subcmds

import (
	"github.com/caas-one/kubectl-switch/server"
	"github.com/caas-one/kubectl-switch/server/fileutil"
)

// RemoveAll remove all cluster name
type RemoveAll struct{}

// register cmd
func init() {
	server.RegisterSubCmd((*Remove)(nil))

}

var _ server.SubCommand = &Remove{}

// Exec exec removell
func (r *RemoveAll) Exec(cmd *server.CmdShim) error {
	path := fileutil.GetBase()
	exist, _ := fileutil.PathStatus(path)
	if exist {
		return fileutil.DelDir(path)
	}
	return nil
}
