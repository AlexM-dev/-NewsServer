package services

import (
	"GoNews/pkg/rss"
	"GoNews/pkg/storage"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type jsonEx struct {
	Rss            []string
	Request_period int
}

func Update(db storage.Interface, errCh chan<- error) {
	var d jsonEx
	ch := make(chan storage.Post, 100)
	plan, err := ioutil.ReadFile("..\\GoNews\\cmd\\goNews\\config.json")
	if err != nil {
		errCh <- err
	}

	err = json.Unmarshal(plan, &d)
	if err != nil {
		errCh <- err
	}
	log.Println(d)

	go dbAdd(db, ch, errCh)
	go updatePosts(d, ch, errCh)
}

func dbAdd(db storage.Interface, ch <-chan storage.Post, errCh chan<- error) {
	m := make(map[string]bool)
	for {
		d := <-ch
		if _, ok := m[d.Link]; !ok {
			err := db.AddPost(d)
			if err != nil {
				errCh <- err
			}
			log.Println(d.Link, "yes")
			m[d.Link] = true
		} else {
			log.Println(d.Link, "no")
		}
	}
}

func updatePosts(d jsonEx, ch chan<- storage.Post, errCh chan<- error) {
	for {
		for _, i := range d.Rss {
			go rss.PostsFrom(i, ch, errCh)
		}
		time.Sleep(time.Duration(d.Request_period) * time.Minute)
		log.Println("ttttime")
	}
}

func CatchErr(errCh <-chan error) {
	for {
		log.Fatal(<-errCh)
	}
}
