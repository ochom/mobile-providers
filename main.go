package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	gohttp "github.com/ochom/go-http"
)

// Map ...
type Map map[string]interface{}

// Market ...
type Market struct {
	EventCode int   `json:"EventCode,omitempty"`
	Outcomes  []Map `json:"Outcomes,omitempty"`
}

// GetMarkets ...
func getMarkets(gameID string, row int, markets chan Market, done func()) {

	defer done()

	ctx := context.Background()

	client := gohttp.NewHTTPService(time.Second * 30)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-API-Key":    "9351C2288CA16D34",
	}

	url := fmt.Sprintf("https://sports-stm04.btobet.games/rest/smsbetting/GetMarkets?mobile=0708113456&eventCode=%s&marketCode=%v", gameID, row)

	market := new(Market)
	defer func() {
		markets <- *market
	}()

	res, err := client.Get(ctx, url, headers)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.Unmarshal(res, &market); err != nil {
		log.Println(err)
		return
	}

	if len(market.Outcomes) > 0 {
		fmt.Println("got market", row)
		return
	}
}

func main() {
	provider := MatchMobile("254760113456")
	if provider == nil {
		fmt.Println("no match")
	}

	fmt.Println(*provider)
}

// func main() {
// 	gameID := flag.String("g", "3", "game id")
// 	var wg sync.WaitGroup
// 	max := 1000
// 	markets := make(chan Market, max)
// 	for i := 0; i < max; i++ {
// 		wg.Add(1)
// 		go getMarkets(*gameID, i, markets, func() { wg.Done() })
// 	}

// 	defer close(markets)
// 	wg.Wait()

// 	f, err := os.Create("games.json")
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	defer f.Close()

// 	f.Write([]byte("[\n"))
// 	for v := range markets {
// 		d, err := json.Marshal(v)
// 		if err != nil {
// 			continue
// 		}
// 		f.Write(d)
// 		f.Write([]byte(",\n"))
// 	}
// 	f.Write([]byte("]"))

// }
