package rss

import (
	"GoNews/pkg/storage"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func PostsFrom(s string, ch chan<- storage.Post, errCh chan<- error) {
	response, err := http.Get(s)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	xmlData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errCh <- err
	}

	rss := new(Rss)

	buffer := bytes.NewBuffer(xmlData)

	decoded := xml.NewDecoder(buffer)

	err = decoded.Decode(rss)

	if err != nil {
		errCh <- err
	}

	total := len(rss.Channel.Items)
	log.Println("total", total)
	for i := 0; i < total; i++ {
		log.Println(i, total)
		temp := *storage.CreatePost(rss.Channel.Items[i].Title, rss.Channel.Items[i].Description, rss.Channel.Items[i].Link)
		ch <- temp
		log.Println(temp.ID)
	}
	log.Println("VSE")
}
