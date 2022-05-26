package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maotai/model/purchase"
	"maotai/utils"
	"os"
	"time"
)

const (
	HOST = "https://h5.moutai519.com.cn/xhr/front/mall/item/purchaseInfo"
)

var (
	startTime  = time.Now()
	pushTime   = startTime
	contentMap = make(map[string]struct{})
)

func main() {
	httpClient := utils.NewHttpClient()
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
	cookie := string(fileContent)
	for {
		go func() {
			var content string
			var hasStock bool
			for i, v := range purchase.GetShopPurchaseBodyList() {
				err, res := httpClient.WithHeaders(map[string]interface{}{
					"Cookie": cookie,
				}).PostJson(HOST, v)

				if err != nil {
					fmt.Errorf("请求异常： %s", err.Error())
				}
				p := purchase.Purchase{}
				json.Unmarshal(res, &p)

				//fmt.Printf("%+v \n", p)
				msg := fmt.Sprintf("- 门店%d：%s \t 库存：%d \t \n", i+1, p.Data.ShopInfo.ShopName, p.Data.PurchaseInfo.Inventory)
				fmt.Print(msg)
				content += msg
				if p.Data.PurchaseInfo.Inventory > 0 {
					hasStock = true
				}
			}
			if !hasStock {
				push(content)
			}
		}()
		time.Sleep(time.Second * 10)
	}
}

func push(content string) {
	_, ok := contentMap[content]
	if pushTime.Add(10*time.Second).Before(time.Now()) && !ok {
		utils.New().Send(content)
		pushTime = time.Now()
		contentMap[content] = struct{}{}
	}
}
