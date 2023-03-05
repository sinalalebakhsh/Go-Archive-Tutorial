package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct { 
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	// fmt.Println("", "Name", "\t\t", "Price", "\t\t", "URLs-Link")
	// fmt.Println("", "----", "\t\t", "-----", "\t\t", "------------------------------------")

	var items []item

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item {
			Name: h.ChildText("h2.product-title"),
			Price: h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		// fmt.Println("", item.Name, "\t", item.Price, "\t", item.ImgUrl)

		items = append(items, item)
	
	})

	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("http://j2store.net/demo/index.php/shop")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("products.json", content, 0644)
}

/************************  RESULT   ****************************************

go run .

[{"name":"Simple","price":"$80.00","imgurl":"http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_01.png"},{"name":"Variable","price":"$40.00","imgurl":"http://j2store.net/demo/images/themeparrot/leggings/leggins_01.png"},{"name":"Configurable","price":"$50.00","imgurl":"http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_01.png"},{"name":"Downloadable","price":"$30.00","imgurl":"http://j2store.net/demo/images/themeparrot/product_image_07.png"},{"name":"Blender1","price":"$110.00","imgurl":"http://j2store.net/demo/images/themeparrot/blenders/blenders_01.png"},{"name":"Blender2","price":"$150.00","imgurl":"http://j2store.net/demo/images/themeparrot/blenders/blenders_02.png"},{"name":"Blender3","price":"$79.00","imgurl":"http://j2store.net/demo/images/themeparrot/blenders/blenders_03.png"},{"name":"Blender4","price":"$110.00","imgurl":"http://j2store.net/demo/images/themeparrot/blenders/blenders_05.png"},{"name":"T-Shirt1","price":"$95.00","imgurl":"http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_05.png"},{"name":"T-Shirt2","price":"$80.00","imgurl":"http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_01.png"},{"name":"T-Shirt3","price":"$110.00","imgurl":"http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_02.png"},{"name":"Shirt1","price":"$150.00","imgurl":"http://j2store.net/demo/images/themeparrot/shirt/shirt_04.png"},{"name":"T-Shirt4","price":"$170.00","imgurl":"http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_04.png"},{"name":"Shirt4","price":"$85.00","imgurl":"http://j2store.net/demo/images/themeparrot/shirt/shirt_03.png"},{"name":"Shirt3","price":"$80.00","imgurl":"http://j2store.net/demo/images/themeparrot/shirt/shirt_05.png"},{"name":"Leggings1","price":"$120.00","imgurl":"http://j2store.net/demo/images/themeparrot/leggings/leggins_01.png"},{"name":"Leggings2","price":"$130.00","imgurl":"http://j2store.net/demo/images/themeparrot/leggings/leggins_02.png"},{"name":"Leggings3","price":"$120.00","imgurl":"http://j2store.net/demo/images/themeparrot/leggings/leggins_03.png"},{"name":"Leggings4","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/leggings/leggins_04.png"},{"name":"Shirt2","price":"$120.00","imgurl":"http://j2store.net/demo/images/themeparrot/shirt/shirt_01.png"},{"name":"Kettle1","price":"$50.00","imgurl":"http://j2store.net/demo/images/themeparrot/kettle/kettle_01.png"},{"name":"Kettle2","price":"$45.00","imgurl":"http://j2store.net/demo/images/themeparrot/kettle/kettle_02.png"},{"name":"Kettle3","price":"$65.00","imgurl":"http://j2store.net/demo/images/themeparrot/kettle/kettle_03.png"},{"name":"Kettle4","price":"$50.00","imgurl":"http://j2store.net/demo/images/themeparrot/kettle/kettle_04.png"},{"name":"Knife1","price":"$40.00","imgurl":"http://j2store.net/demo/images/themeparrot/knife/knife_01.png"},{"name":"Knife2","price":"$30.00","imgurl":"http://j2store.net/demo/images/themeparrot/knife/knife_02.png"},{"name":"Knife3","price":"$130.00","imgurl":"http://j2store.net/demo/images/themeparrot/knife/knife_03.png"},{"name":"Knife4","price":"$80.00","imgurl":"http://j2store.net/demo/images/themeparrot/knife/knife_04.png"},{"name":"Toaster1","price":"$200.00","imgurl":"http://j2store.net/demo/images/themeparrot/toaster/toaster_01.png"},{"name":"Toaster2","price":"$110.00","imgurl":"http://j2store.net/demo/images/themeparrot/toaster/toaster_02.png"},{"name":"Toaster3","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/toaster/toaster_03.png"},{"name":"Toaster4","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/toaster/toaster_04.png"},{"name":"Pots\u0026pans1","price":"$150.00","imgurl":"http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_01.png"},{"name":"Pots\u0026pans2","price":"$90.00","imgurl":"http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_02.png"},{"name":"Pots\u0026pans3","price":"$80.00","imgurl":"http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_03.png"},{"name":"Pots\u0026pans4","price":"$75.00","imgurl":"http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_04.png"},{"name":"Dressing table1","price":"$150.00","imgurl":"http://j2store.net/demo/images/themeparrot/dressingtable/dt1.png"},{"name":"Dressing table2","price":"$55.00","imgurl":"http://j2store.net/demo/images/themeparrot/dressingtable/dt2.png"},{"name":"Sofa4","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/sofa/sofa1.png"},{"name":"Sofa3","price":"$60.00","imgurl":"http://j2store.net/demo/images/themeparrot/sofa/sofa2.png"},{"name":"Sofa2","price":"$50.00","imgurl":"http://j2store.net/demo/images/themeparrot/sofa/sofa3.png"},{"name":"Sofa1","price":"$75.00","imgurl":"http://j2store.net/demo/images/themeparrot/sofa/sofa4.png"},{"name":"Dressing table3","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/dressingtable/dt3.jpg"},{"name":"Dressing table4","price":"$200.00","imgurl":"http://j2store.net/demo/images/themeparrot/dressingtable/dt4.jpg"},{"name":"TV-Stand1","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/tvstand/tv2.png"},{"name":"TV-Stand2","price":"$150.00","imgurl":"http://j2store.net/demo/images/themeparrot/tvstand/tv1.png"},{"name":"Bed2","price":"$130.00","imgurl":"http://j2store.net/demo/images/themeparrot/bed/bed1.png"},{"name":"Bed4","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/bed/bed5.jpg"},{"name":"Bed3","price":"$200.00","imgurl":"http://j2store.net/demo/images/themeparrot/bed/bed2.png"},{"name":"Bed1","price":"$100.00","imgurl":"http://j2store.net/demo/images/themeparrot/bed/bed3.png"},{"name":"TV-Stand3","price":"$8.00","imgurl":"http://j2store.net/demo/images/themeparrot/tvstand/tv3.png"},{"name":"Top1","price":"$25.00","imgurl":"http://j2store.net/demo/images/themeparrot/tops/tops_01.png"},{"name":"Hat1","price":"$13.00","imgurl":"http://j2store.net/demo/images/themeparrot/hats/hats_01.png"},{"name":"TV-Stand4","price":"$5.00","imgurl":"http://j2store.net/demo/images/themeparrot/tvstand/tv4.png"},{"name":"Top2","price":"$15.00","imgurl":"http://j2store.net/demo/images/themeparrot/tops/tops_02.png"},{"name":"Study table1","price":"$10.00","imgurl":"http://j2store.net/demo/images/themeparrot/studytable/study1.png"},{"name":"Hat2","price":"$15.00","imgurl":"http://j2store.net/demo/images/themeparrot/hats/hats_02.png"},{"name":"Study table4","price":"$20.00","imgurl":"http://j2store.net/demo/images/themeparrot/studytable/study1.png"},{"name":"Hat4","price":"$30.00","imgurl":"http://j2store.net/demo/images/themeparrot/hats/hats_03.png"},{"name":"Top4","price":"$50.00","imgurl":"http://j2store.net/demo/images/themeparrot/tops/tops_04.png"},{"name":"Study table3","price":"$15.00","imgurl":"http://j2store.net/demo/images/themeparrot/studytable/study9.jpg"},{"name":"Hat3","price":"$25.00","imgurl":"http://j2store.net/demo/images/themeparrot/hats/hats_03.png"},{"name":"Top3","price":"$35.00","imgurl":"http://j2store.net/demo/images/themeparrot/tops/tops_03.png"},{"name":"Study table2","price":"$17.00","imgurl":"http://j2store.net/demo/images/themeparrot/studytable/study2.png"}]
 ****************************************************************************/ 