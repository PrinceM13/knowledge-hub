package server

import "github.com/PrinceM13/knowledge-hub-api/internal/db/user"

type App struct {
	UserRepo user.Repository
}
