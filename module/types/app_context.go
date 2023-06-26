package types

import "go.a5r.dev/services/module/repositories"

type Repositories struct {
	Module *repositories.ModuleRepository
	Auth   *repositories.AuthRepository
}

type AppContext struct {
	Meta         *Meta
	Repositories *Repositories
}
