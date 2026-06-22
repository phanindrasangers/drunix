package global

import (
	"github.com/onsi/ginkgo/v2/internal"
)

var (
	Suite       *internal.Suite
	Failer      *internal.Failer
	backupSuite *internal.Suite
)

func init() {
	InitializeGlobals()
}

func InitializeGlobals() {
	Failer = internal.NewFailer()
	Suite = internal.NewSuite()
}

func PushClone() error {
	var err error
	backupSuite, err = Suite.Clone()
	return err
}

func PopClone() {
	Suite = backupSuite
}
