package main

import (
	"os"
	"fmt"
    "time"
    "regexp"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/ChimeraCoder/anaconda"
)

const (
    url = "http://www.topre.co.jp/products/elec/keyboards/index.html"
)

func fetchTopre(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
        fmt.Println(err)
	}
    defer res.Body.Close()

    fmt.Println("fetched", url)

    byteArray, _ := ioutil.ReadAll(res.Body)
    html_text := string(byteArray)

    return html_text, err
}

func tweet(api *anaconda.TwitterApi, tweet_text string) anaconda.Tweet {
    tweeted_text, err := api.PostTweet(tweet_text, nil)
    if err != nil {
        fmt.Println(err)
    }
    return tweeted_text
}

func main() {
    search_words := []string{"スプリット", "左右", "セパレート", "分離"}

    api_key := os.Getenv("TSB_API_KEY")
    api_secret := os.Getenv("TSB_API_SECRET")
    con_key := os.Getenv("TSB_CONSUMER_KEY")
    con_secret := os.Getenv("TSB_CONSUMER_SECRET")

    fmt.Println(os.Environ())

    api := anaconda.NewTwitterApiWithCredentials(api_key, api_secret, con_key, con_secret)

    html_text, _ := fetchTopre(url)

    find_flag := false
    var find_words []string
    for _, word := range search_words {
        r := regexp.MustCompile(word)
        if r.MatchString(html_text) {
            find_words = append(find_words, word)
            find_flag = true
        }
    }

    tweet_text := ""
    if find_flag {
        tweet_text =  time.Now().Format("1月2日15時4分") + "現在、" + url + " に 「" + strings.Join(find_words, ",") + "」が見つかりました。スプリットキーボードが発売されるといいね✨"
    } else if time.Now().Hour() == 23 - 9 {
        tweet_text = time.Now().Format("1月2日") + "はスプリットキーボードが発表されなかったね…また明日に期待💁🏼‍♀️"
    } else {
        tweet_text = ""
    }

    if tweet_text != "" {
        tweeted_text := tweet(api, tweet_text)
        fmt.Print(tweeted_text)
    }
}
