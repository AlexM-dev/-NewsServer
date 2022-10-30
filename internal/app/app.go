package app

import (
	"GoNews/internal/services"
	"GoNews/pkg/storage"
	"sync"
)

func Run(db storage.Interface) {
	var wg sync.WaitGroup
	errCh := make(chan error)
	wg.Add(2)
	go services.Update(db, errCh)
	go services.CatchErr(errCh)
	wg.Wait()
}
