package spiders

// where i defined spider
import (
	"ants/http"
	"log"
)

const (
	BASE_PARSE_NAME     = "base"
	SPIDERS_STATUS_INIT = iota
	SPIDERS_STATUS_RUNNING
	SPIDERS_STATUS_STOP
	SPIDERS_BASIC_COOKIE
)

/*
what a spider do
*	make start request
* 	define basic parse func
*/
type Spider struct {
	Status    int //  init runnign or stop
	Name      string
	StartUrls []string
	ParseMap  map[string]func(response *http.Response) ([]*http.Request, error)
}

func (this *Spider) MakeStartRequests() []*http.Request {
	startRequestSlice := make([]*http.Request, len(this.StartUrls))
	for index, url := range this.StartUrls {
		request, err := http.NewRequest("GET", url, this.Name, BASE_PARSE_NAME, nil, 0)
		if err != nil {
			log.Println(err)
			continue
		}
		startRequestSlice[index] = request
	}
	return startRequestSlice
}