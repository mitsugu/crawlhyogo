package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
	"time"
)

func main() {
  url := "https://web.pref.hyogo.lg.jp/kf16/coronavirus_data.html"
	doc, _ := htmlquery.LoadURL(url)

	// 年月日
	year := time.Now().Year()
	path := "//*[@id='kensa_new']/following-sibling::text()"
	node, _ := htmlquery.Query(doc, path)
	d := strings.Replace(strings.Replace(strings.Replace(node.Data, "月", "/", 1), "日24時現在）", "", 1), "（", strconv.Itoa(year)+"/", 1)

	// 宿泊療養
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[6]/text()"
	node, _ = htmlquery.Query(doc, path)
	syuku := strings.Replace(node.Data, ",", "", 1)
	// 軽症・中等症
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[4]/text()"
	node, _ = htmlquery.Query(doc, path)
	kei := strings.Replace(node.Data, ",", "", 1)
	// 入院・宿泊療養調整等
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[7]/text()"
	node, _ = htmlquery.Query(doc, path)
	cyou := strings.Replace(node.Data, ",", "", 1)
	// 自宅療養
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[9]/text()"
	node, _ = htmlquery.Query(doc, path)
	home := strings.Replace(node.Data, ",", "", 1)
	// その他医療機関福祉施設等
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[10]/text()"
	node, _ = htmlquery.Query(doc, path)
	fukushi := strings.Replace(node.Data, ",", "", 1)
	// 重症
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[5]/text()"
	node, _ = htmlquery.Query(doc, path)
	jyuu := strings.Replace(node.Data, ",", "", 1)
	// 死亡
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[11]/text()"
	node, _ = htmlquery.Query(doc, path)
	shibou := strings.Replace(node.Data, ",", "", 1)
	// 退院
  path = "//div[@id='tmp_contents']/table[2]/tbody/tr[4]/td[12]/text()"
	node, _ = htmlquery.Query(doc, path)
	taiin := strings.Replace(node.Data, ",", "", 1)

	fmt.Printf("%v, %v, %v, %v, %v, %v, %v,,, %v, %v\n", d, syuku, kei, cyou, home, fukushi, jyuu, shibou, taiin)

}

