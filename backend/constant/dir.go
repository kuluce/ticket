package constant

import (
	"path"
)

var (
	DataDir        = ""
	ResourceDir    = path.Join(DataDir, "resource")
	AppResourceDir = path.Join(ResourceDir, "apps")
	AppInstallDir  = path.Join(DataDir, "apps")
)
