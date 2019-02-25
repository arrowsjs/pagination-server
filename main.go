package main

import (
	"github.com/arrowsjs/pagination-server/db"
	"github.com/arrowsjs/pagination-server/resource"
	"github.com/efritz/chevron"
	"github.com/efritz/nacelle"
	basehttp "github.com/efritz/nacelle/base/http"
)

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	processes.RegisterInitializer(
		db.NewInitializer(),
		nacelle.WithInitializerName("db"),
	)

	processes.RegisterProcess(
		basehttp.NewServer(chevron.NewInitializer(resource.SetupRoutesFunc)),
		nacelle.WithProcessName("server"),
	)

	return nil
}

func main() {
	nacelle.NewBootstrapper("pagination-server", setup).BootAndExit()
}
