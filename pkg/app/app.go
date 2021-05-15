package app

import (
	"pishikot/pkg/service"
	"pishikot/pkg/store"
)

func Run() {
	store := store.New()
	service := service.NewService(store)
	service.Start()
}
