package rss

import (
	"GoNews/pkg/storage"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func Update(db storage.Interface) {
	type jsonEx struct {
		Rss            []string
		Request_period int
	}
	var d jsonEx
	ch := make(chan storage.Post, 50)
	//var wg sync.WaitGroup
	plan, err := ioutil.ReadFile("C:\\Users\\Александр\\Documents\\Go\\GoNews\\pkg\\storage\\rss\\config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(plan, &d)
	log.Println(d)
	go func() {
		for {
			d := <-ch
			err = db.AddPost(d)
			log.Println("Add post")
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	for {
		for _, i := range d.Rss {
			go postsFrom(i, ch)
		}
		time.Sleep(time.Duration(d.Request_period) * time.Minute)
	}
}
func postsFrom(s string, ch chan<- storage.Post) {
	response, err := http.Get(s)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	xmlData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	rss := new(Rss)

	buffer := bytes.NewBuffer(xmlData)

	decoded := xml.NewDecoder(buffer)

	err = decoded.Decode(rss)

	if err != nil {
		log.Fatal(err)
	}

	total := len(rss.Channel.Items)
	log.Println(total)
	for i := 0; i < total; i++ {
		log.Println(i, total)
		ch <- *storage.CreatePost(rss.Channel.Items[i].Title, rss.Channel.Items[i].Description, rss.Channel.Items[i].Link)
		log.Println("Read post")
	}
}
