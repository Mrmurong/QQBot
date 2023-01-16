package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/", func(context *gin.Context) {
		dataReader := context.Request.Body
		rawData, _ := ioutil.ReadAll(dataReader)
		postType := gjson.Get(string(rawData), "post_type").String()
		if postType == "message" {
			message := gjson.Get(string(rawData), "message").String()
			if message == "小爱同学" {
				context.JSON(http.StatusOK, gin.H{
					"reply": "我在呢,主人",
				})
			}
			if message == "日常" {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetMsg(),
				})
			}
			if string([]rune(message)[0:3]) == "阵眼 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetZhenYan(string([]rune(message)[3:])),
				})
			}
			if string([]rune(message)[0:2]) == "宏 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetQiXue(string([]rune(message)[2:])),
				})
			}
			if message == "骚话" {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetSaoHua(),
				})
			}
			if string([]rune(message)[0:3]) == "攻略 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": "[CQ:image,file=" + GetGongLve(string([]rune(message)[3:])) + "]",
				})
			}
			if string([]rune(message)[0:3]) == "前置 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": "[CQ:image,file=" + GetQianZhi(string([]rune(message)[3:])) + "]",
				})
			}
			if string([]rune(message)[0:4]) == "刷马点 " {
				name, url := GetShuaMa(string([]rune(message)[4:]))
				context.JSON(http.StatusOK, gin.H{
					"reply": "刷新马驹：" + name + "\r\n 刷新地点:\r\n[CQ:image,file=" + url + "]",
				})
			}
			if string([]rune(message)[0:5]) == "马驹位置 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetMaJuAddress(string([]rune(message)[5:])),
				})
			}
			if string([]rune(message)[0:3]) == "配装 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetPeiZhuang(string([]rune(message)[3:])),
				})
			}
			if string([]rune(message)[0:3]) == "小药 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetXiaoYao(string([]rune(message)[3:])),
				})
			}
			if message == "小药" {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetAllXiaoYao(),
				})
			}
			if string([]rune(message)[0:4]) == "器物谱 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetQiWuPu(string([]rune(message)[4:])),
				})
			}
			if string([]rune(message)[0:3]) == "家具 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetJiaJu(string([]rune(message)[3:])),
				})
			}
			if message == "开服" {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetKaiFu("天鹅坪"),
				})
			}
			if string([]rune(message)[0:3]) == "日常 " {
				context.JSON(http.StatusOK, gin.H{
					"reply": GetFetureRiChang(string([]rune(message)[3:])),
				})
			}
			if message == "使用说明"{
				context.JSON(http.StatusOK,gin.H{
					"reply":"小爱的使用说明地址在这呢！https://github.com/Mrmurong/-/wiki/%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E",
				})
			}

		}

	})

	_ = r.Run()
}

func GetMsg() string {
	prestige := ""
	lucky := ""
	team_public := ""
	team_mijing := ""
	team_tuandui := ""
	response, err := http.Get("https://www.jx3api.com/data/active/current")
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}
	data := gjson.Get(string(body), "data").String()
	prestiges := gjson.Get(data, "prestige").Array()
	for i := 0; i < len(prestiges); i++ {
		prestige += prestiges[i].String() + "、"
	}
	luckys := gjson.Get(data, "lucky").Array()
	for i := 0; i < len(luckys); i++ {
		lucky += luckys[i].String() + "、"
	}
	teams := gjson.Get(data, "team").Array()
	team_public = teams[0].String()
	team_mijing = teams[1].String()
	team_tuandui = teams[2].String()
	prestige = prestige[:len(prestige)-1]
	lucky = lucky[:len(lucky)-1]
	return "当前时间：" + gjson.Get(data, "date").String() + "\t 星期" + gjson.Get(data, "week").String() + "\r\n今日大战：" + gjson.Get(data, "war").String() + "\r\n今日战场：" + gjson.Get(data, "battle").String() + "\r\n阵营任务：" + gjson.Get(data, "camp").String() + "\r\n门派事件：" + gjson.Get(data, "school").String() + "\r\n援驰任务：" + gjson.Get(data, "relief").String() + "\r\n福缘宠物：" + lucky + "\r\n家园声望：" + prestige + "\r\n公共任务：" + team_public + "\r\n秘境任务：" + team_mijing + "\r\n团队秘境：" + team_tuandui
}
func GetFetureRiChang(day string) string {
	prestige := ""
	result := ""
	days, err := strconv.Atoi(day)
	if err != nil {
		return "查询时请输入当前日期后推多少天！"
	}
	if days < 7 {
		time_use := time.Now().AddDate(0, 0, days)
		day_use := time_use.Day()
		response, err := http.Get("https://www.jx3api.com/data/active/calculate?num=7")
		if err != nil {

		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {

		}
		datatmp := gjson.Get(string(body), "data").String()
		data := gjson.Get(datatmp, "data").Array()
		for i := 0; i < len(data); i++ {
			if gjson.Get(data[i].String(), "day").String() == strconv.Itoa(day_use) {
				prestiges := gjson.Get(data[i].String(), "prestige").Array()
				for i := 0; i < len(prestiges); i++ {
					prestige += prestiges[i].String() + "、"
				}
				prestige = prestige[:len(prestige)-1]
				result = "查询时间：" + gjson.Get(data[i].String(), "date").String() + "\t 星期" + gjson.Get(data[i].String(), "week").String() + "\r\n今日大战：" + gjson.Get(data[i].String(), "war").String() + "\r\n今日战场：" + gjson.Get(data[i].String(), "battle").String() + "\r\n阵营任务：" + gjson.Get(data[i].String(), "camp").String() + "\r\n援驰任务：" + gjson.Get(data[i].String(), "relief").String() + "\r\n家园声望：" + prestige
			}
		}
	} else {
		time_use := time.Now().AddDate(0, 0, days)
		day_use := time_use.Day()
		response, err := http.Get("https://www.jx3api.com/data/active/calculate?num=" + day)
		if err != nil {

		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {

		}
		datatmp := gjson.Get(string(body), "data").String()
		data := gjson.Get(datatmp, "data").Array()
		for i := 0; i < len(data); i++ {
			if gjson.Get(data[i].String(), "day").String() == strconv.Itoa(day_use) {
				prestiges := gjson.Get(data[i].String(), "prestige").Array()
				for i := 0; i < len(prestiges); i++ {
					prestige += prestiges[i].String() + "、"
				}
				prestige = prestige[:len(prestige)-1]
				result = "查询时间：" + gjson.Get(data[i].String(), "date").String() + "\t 星期" + gjson.Get(data[i].String(), "week").String() + "\r\n今日大战：" + gjson.Get(data[i].String(), "war").String() + "\r\n今日战场：" + gjson.Get(data[i].String(), "battle").String() + "\r\n阵营任务：" + gjson.Get(data[i].String(), "camp").String() + "\r\n援驰任务：" + gjson.Get(data[i].String(), "relief").String() + "\r\n家园声望：" + prestige
			}
		}
	}
	return result
}
func GetZhenYan(menpai string) string {
	value := ""
	response, err := http.Get("https://www.jx3api.com/data/school/matrix?name=" + menpai)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}
	data := gjson.Get(string(body), "data").String()
	desc := gjson.Get(data, "descs").Array()
	value = menpai + ":\t【" + gjson.Get(data, "skillName").String() + "】\r\n" + gjson.Get(desc[0].String(), "name").String() + ":\t" + gjson.Get(desc[0].String(), "desc").String() + "\r\n" + gjson.Get(desc[1].String(), "name").String() + ":\t" + gjson.Get(desc[1].String(), "desc").String() + "\r\n" + gjson.Get(desc[2].String(), "name").String() + ":\t" + gjson.Get(desc[2].String(), "desc").String() + "\r\n" + gjson.Get(desc[3].String(), "name").String() + ":\t" + gjson.Get(desc[3].String(), "desc").String() + "\r\n" + gjson.Get(desc[4].String(), "name").String() + ":\t" + gjson.Get(desc[4].String(), "desc").String() + "\r\n" + gjson.Get(desc[5].String(), "name").String() + ":\t" + gjson.Get(desc[5].String(), "desc").String() + "\r\n" + gjson.Get(desc[6].String(), "name").String() + ":\t" + gjson.Get(desc[6].String(), "desc").String() + "\r\n"
	return value
}
func GetQiXue(menpai string) string {
	response, err := http.Get("https://www.jx3api.com/data/school/macro?name=" + menpai)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	qixue := gjson.Get(string(data), "qixue").String()
	hong := gjson.Get(data, "macro").String()
	shijian := gjson.Get(data, "time").String()
	return menpai + "\t更新时间：" + shijian + "\r\n" + hong + "\r\n奇穴：" + qixue
}
func GetSaoHua() string {
	response, err := http.Get("https://www.jx3api.com/data/chat/random")
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return gjson.Get(data, "text").String()
}
func GetGongLve(qiyu string) string {
	response, err := http.Get("https://www.jx3api.com/data/lucky/sub/strategy?name=" + qiyu)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return gjson.Get(data, "url").String()
}
func GetQianZhi(qiyu string) string {
	response, err := http.Get("https://www.jx3api.com/data/lucky/sub/require?name=" + qiyu)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return gjson.Get(data, "url").String()
}
func GetShuaMa(address string) (string, string) {
	name := ""
	url := ""
	response, err := http.Get("https://www.jx3api.com/data/useless/refresh?name=" + address)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	datatemp := gjson.Get(string(body), "data").String()
	data := gjson.Get(string(datatemp), "data").Array()
	for i := 0; i < len(data); i++ {
		name += gjson.Get(data[i].String(), "name").String() + ","
		url = gjson.Get(data[i].String(), "url").String()
	}
	name = name[:len(name)-1]
	return name, url
}
func GetMaJuAddress(maju string) string {
	result := ""
	response, err := http.Get("https://www.jx3api.com/data/useless/refresh?name=" + maju)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	datatemp := gjson.Get(string(body), "data").String()
	data := gjson.Get(string(datatemp), "data").Array()
	for i := 0; i < len(data); i++ {
		result += maju + "刷新地点：" + gjson.Get(data[i].String(), "map").String() + "[CQ:image,file=" + gjson.Get(data[i].String(), "url").String() + "]\r\n"
	}

	return result
}
func GetPeiZhuang(menpai string) string {
	response, err := http.Get("https://www.jx3api.com/data/school/equip?name=" + menpai)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return "PVE:\r\n [CQ:image,file=" + gjson.Get(data, "pve").String() + "]\r\n PVP:\r\n [CQ:image,file=" + gjson.Get(data, "pvp").String() + "]"
}
func GetXiaoYao(menpai string) string {
	response, err := http.Get("https://www.jx3api.com/data/school/snacks?name=" + menpai)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return "[CQ:image,file=" + gjson.Get(data, "url").String() + "]"
}
func GetAllXiaoYao() string {
	response, err := http.Get("https://www.jx3api.com/data/school/snacks")
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return "[CQ:image,file=" + gjson.Get(data, "url").String() + "]"
}
func GetQiWuPu(address string) string {
	result := address + "的家具产出有：\r\n"
	response, err := http.Get("https://www.jx3api.com/data/home/travel?name=" + address)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").Array()
	for i := 0; i < len(data); i++ {
		result += gjson.Get(data[i].String(), "name").String() + "[CQ:image,file=" + gjson.Get(data[i].String(), "image_path").String() + "]\r\n" + gjson.Get(data[i].String(), "tip").String() + "\r\n"
	}
	return result
}
func GetJiaJu(jiaju string) string {
	response, err := http.Get("https://www.jx3api.com/data/home/furniture?name=" + jiaju)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	return jiaju + "的获取方式为：\r\n" + gjson.Get(data, "source").String() + " [CQ:image,file=" + gjson.Get(data, "image_path").String() + "]\r\n" + gjson.Get(data, "tip").String()
}
func GetKaiFu(fu string) string {
	response, err := http.Get("https://www.jx3api.com/data/server/check?server=" + fu)
	if err != nil {

	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	data := gjson.Get(string(body), "data").String()
	if gjson.Get(data, "status").Bool() {
		return fu + "已经开服了,快点上线领奇遇！！！"
	} else {
		return "忘了跟你说，GWW带着小姨子跑了,剑三倒闭了！"
	}
}
