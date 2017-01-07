package store

import "github.com/murlokswarm/flux"

var (
	Nav      = &navStore{}
	Articles = &articleStore{}
)

func init() {
	flux.Register(Nav)
	flux.Register(Articles)
}
