package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"time"
)

var header = map[string]string{
	"Host":            `book.douban.com`,
	"Connection":      `keep-alive`,
	"Cache-Control":   `max-age=0`,
	"User-Agent":      `Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36`,
	"Accept":          `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
	"Accept-Language": `zh-CN,zh;q=0.9,en;q=0.8`,
	"Referer":         `https://book.douban.com/`, // MUST
}

const pageSize = 15

// 将 HTML 转为 goquery.Document
func getDocument(url string) *goquery.Document {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))

	if resp.StatusCode != 200 {
		fmt.Println("Request StatusCode != 200", url)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title" `
	Pic   string `json:"pic" `
}

// 将 document 解析为 book
func parseDocument(url string, list *[]Book) bool {
	doc := getDocument(url)
	count := 0
	doc.Find("ul.interest-list li.subject-item").Each(func(i int, s *goquery.Selection) {
		a := s.Find("div.info h2 a")
		title, _ := a.Attr("title")
		href, _ := a.Attr("href")
		pic, _ := s.Find("div.pic a img").Attr("src")
		book := Book{
			Id:    href[32 : len(href)-1],
			Title: title,
			Pic:   pic,
		}
		*list = append(*list, book)
		count++
	})
	return count == pageSize
}

func getBooks(username string) map[string]*[]Book {
	types := []string{"wish", "do", "collect"}

	books := make(map[string]*[]Book)
	for _, t := range types {
		list := make([]Book, 0, 100)
		books[t] = &list
		for i := 0; ; i++ {
			url := fmt.Sprintf("https://book.douban.com/people/%v/%v?start=%d", username, t, i*pageSize)
			hasNext := parseDocument(url, &list)
			if !hasNext {
				break
			}
		}
	}

	return books
}

func writeJson(books map[string]*[]Book, path string) {
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(string(b))
	if err != nil {
		panic(err)
	}
}

func main() {

	var username string
	var path string

	flag.StringVar(&username, "u", "", "豆瓣用户名，默认为空")
	flag.StringVar(&path, "p", "", "json 文件存储路径，默认为空")
	flag.Parse()

	fmt.Printf("username=%v path=%v\n", username, path)
	fmt.Println("start ", time.Now().Format("2006-01-02 15:04:05"))
	books := getBooks(username)
	writeJson(books, path)
	fmt.Println("done ", time.Now().Format("2006-01-02 15:04:05"))
}
