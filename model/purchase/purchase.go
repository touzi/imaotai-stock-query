package purchase

type Purchase struct {
	Code int `json:"code"`
	Data struct {
		ShopInfo struct {
			ShopId   string `json:"shopId"`
			ShopName string `json:"shopName"`
			Desc     string `json:"desc"`
			PicUrl   string `json:"picUrl"`
			ShopType int    `json:"shopType"`
			ShopTag  string `json:"shopTag"`
		} `json:"shopInfo"`
		PurchaseInfo struct {
			ItemCode         string `json:"itemCode"`
			Inventory        int    `json:"inventory"`
			CanAddCart       bool   `json:"canAddCart"`
			ForbiddenBuyDesc string `json:"forbiddenBuyDesc"`
			LimitCount       int    `json:"limitCount"`
		} `json:"purchaseInfo"`
	} `json:"data"`
}
