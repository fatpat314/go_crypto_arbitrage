package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// sushi_swap()
	uni_swap()
}

func uni_swap() {
	var uni_pair_list [2]string
	uni_pair_list[0] = "https://info.uniswap.org/#/pools/0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8"
	uni_pair_list[1] = "https://info.uniswap.org/#/pools/0x6c6bc977e13df9b0de53b251522280bb72383700"

	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})
	// <div class="sc-chPdSV goKJOd css-1qjlrio" style="white-space: nowrap;">1 DAI =  1.0004 USDC</div>
	c.OnHTML("div.sc-jKjlTe", func(e *colly.HTMLElement) {
		fmt.Println(string(e.Text))

	})

	for i := 0; i < len(uni_pair_list); i++ {
		c.Visit(uni_pair_list[i])
	}
}

func sushi_swap() {
	var sushi_pair_list [2]string
	sushi_pair_list[0] = "https://analytics.sushi.com/pairs/0x795065dcc9f64b5614c407a6efdc400da6221fb0"
	sushi_pair_list[1] = "https://analytics.sushi.com/pairs/0x397ff1542f962076d0bfe58ea045ffa2d347aca0"

	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div.MuiGrid-root p", func(e *colly.HTMLElement) {
		fmt.Println(string(e.Text))

	})

	for i := 0; i < len(sushi_pair_list); i++ {
		c.Visit(sushi_pair_list[i])
	}
}
