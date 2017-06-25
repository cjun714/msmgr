package exec

import (
	"bytes"
	"os/exec"
	"util/log"

	"github.com/pkg/errors"
)

// RunCmd is used to executed shell cmd
func RunCmd(exe string, arg ...string) error {
	cmd := exec.Command(exe, arg...)

	var out bytes.Buffer
	cmd.Stdout = &out
	var er bytes.Buffer
	cmd.Stderr = &er

	e := cmd.Run()
	if e != nil {
		log.E(e.Error())
		log.E(er.String())
		return errors.Wrap(e, er.String())
	}
	log.I(out.String())
	return nil
}
