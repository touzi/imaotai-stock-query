package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maotai/model/purchase"
	"maotai/utils"
	"os"
	"sync"
	"time"
)

const (
	HOST = "https://h5.moutai519.com.cn/xhr/front/mall/item/purchaseInfo"
)

var (
	startTime  = time.Now()
	pushTime   = startTime
	cookie     string
	httpClient *utils.Request
)

func init() {
	_, cookie = readCookie()
	if cookie == "" {
		os.Exit(0)
	}

	httpClient = utils.NewHttpClient()
}

func main() {
	for {
		contentChan := make(chan purchase.Purchase, 100)
		wg := sync.WaitGroup{}
		for i, v := range purchase.GetShopPurchaseBodyList() {
			wg.Add(1)
			go func(i int, v purchase.PurchaseBody, contentChan chan purchase.Purchase) {
				err, res := httpClient.WithHeaders(map[string]interface{}{
					"Cookie": cookie,
				}).PostJson(HOST, v)

				if err != nil {
					fmt.Errorf("请求异常： %s", err.Error())
				}
				p := purchase.Purchase{}
				json.Unmarshal(res, &p)
				wg.Done()
				contentChan <- p
				//fmt.Printf("%+v \n", p)
			}(i, v, contentChan)
		}
		wg.Wait()
		close(contentChan)

		var content string
		var i int
		var stockSum int
		for p := range contentChan {
			i++
			msg := fmt.Sprintf("- 门店%d：%s \t 库存：%d \t \n", i, p.Data.ShopInfo.ShopName, p.Data.PurchaseInfo.Inventory)
			fmt.Print(msg)
			content += msg
			stockSum += p.Data.PurchaseInfo.Inventory
		}
		if stockSum > 0 {
			push(content)
		}
		time.Sleep(time.Second * 5)
	}
}

func readCookie() (err error, cookie string) {
	file, err := os.Open("cookie.txt")
	if err != nil {
		fmt.Errorf("open cookie file error: %s", err.Error())
		return
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("read cookie file error: %s", err.Error())
		return
	}
	cookie = string(fileContent)
	return
}

func push(content string) {
	if pushTime.Add(time.Hour).Before(time.Now()) {
		utils.New().Send(content)
		pushTime = time.Now()
	}
}
