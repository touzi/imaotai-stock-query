package purchase

import "strconv"

type PurchaseBody struct {
	Hot      bool   `json:"hot"`
	ItemCode string `json:"itemCode"`
	Lng      string `json:"lng"`
	Lat      string `json:"lat"`
	Province string `json:"province"`
	City     string `json:"city"`
	ShopId   string `json:"shopId"`
	Jt       string `json:"jt"`
}

var (
	ShopIdList = []int{144440424002, 144440400004, 144440400005, 144440400006, 144440400007, 2244042300001, 2344550400001, 2344040000001}
	//ShopIdList = []int{144440100021, 144440100022, 144440100023, 144440100024, 144440100025, 144440100026}
)

func New(shopId int) PurchaseBody {
	return PurchaseBody{
		Hot:      true,
		ItemCode: "10193",
		//Lng:      "113.536131",
		//Lat:      "22.205244",
		//Province: "440000000",
		//City:     "440400000",
		ShopId: strconv.Itoa(shopId),
		Jt:     "qgvGRFMSm%2BG%2FMabe1V7zPG4UCHGMn1s2Alm86u3RTn1O4ZcL5qk1zNBSazULSif%2F6IKBeVa6yDUL2B8EZqhmUNURJ4QKKV6DCupd7V1riBh%2FcFte90eQTsPLmODYCJID4fkE0kd%2BlmKVWozZw%2B4w%2BnfAZkJ6LBZAIUCzFqFbPf3YzBofHcrLW5aYwHaOw%2FrAxThaESflw2lIH%2FG%2F%2FyZiNsQVTqpi18Gkv%2Fg3FWdjPEtTRmHt%2FlO4MKvWQyiMAmIS9WuadL%2F3NgnUOqiAc5Rj30dHaP6IP%2Fv%2BH7wHDLDao4l3yXxKSWzbeTwZL5JfCU%2FwRdu5zZVIFJuPQ5iuQGQQrgybdF0JGMOyELzxm4%2BJ3Ocv357bzTRlUsXm2g93J3Ws%7CNf5KL1XrsYEOHtACIRnfukn04EEaedO%2FQ7cdj%2FkzYUg%3D%7C10%7C0af74766ac9418a035f79f7be344006e",
	}
}

func GetShopPurchaseBodyList() []PurchaseBody {
	shopList := make([]PurchaseBody, 0)
	for _, shopId := range ShopIdList {
		shopList = append(shopList, New(shopId))
	}
	return shopList
}
