// Package launchpad contains the Launchpad bridge implementation
package launchpad

import (
	"time"

	"github.com/MichaelMure/git-bug/bridge/core"
)

const (
	target = "launchpad-preview"

	metaKeyLaunchpadID    = "launchpad-id"
	metaKeyLaunchpadLogin = "launchpad-login"

	keyProject = "project"

	defaultTimeout = 60 * time.Second
)

var _ core.BridgeImpl = &Launchpad{}

type Launchpad struct{}

func (*Launchpad) Target() string {
	return "launchpad-preview"
}

func (l *Launchpad) LoginMetaKey() string {
	return metaKeyLaunchpadLogin
}

func (*Launchpad) NewImporter() core.Importer {
	return &launchpadImporter{}
}

func (*Launchpad) NewExporter() core.Exporter {
	return nil
}
