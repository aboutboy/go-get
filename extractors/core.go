package extractors

import (
	"../utils"
	"github.com/PuerkitoBio/goquery"
	simplejson "github.com/bitly/go-simplejson"
	"log"
	//"strings"
	"time"
)

type VideoInfo struct {
	title        string                 `json:"title"`
	url          string                 `json:"url"`
	duration     int64                  `json:"duration"`
	downloadInfo map[string]interface{} `json:"downloadInfo"`
	createTime   int64                  `json:"createTime"`
	site         string                 `json:"site"`
}

type Core interface {
	GetVideoInfo(url string) (info VideoInfo, err error)
	GetHtml(url string) (html string, err error)
	Obj() (obj interface{})
	MatchUrl(url string) bool
}

func (self *VideoInfo) dumps() (info map[string]interface{}) {
	info = make(map[string]interface{})
	info["title"] = self.title
	info["url"] = self.url
	info["duration"] = self.duration
	info["downloadInfo"] = self.downloadInfo
	if self.createTime == 0 {
		self.createTime = utils.GetCurrentMilliseconds()
	}
	info["createTime"] = self.createTime
	return info
}

func (self *VideoInfo) Dumps() (info map[string]interface{}) {
	info = self.dumps()
	info["desc"] = "normal 表示标清，hd1 表示高清，hd2 表示超清，hd3 表示720p hd4 表示1080p"
	return
}

func (self *VideoInfo) Urls(hd string) []string {

	return self.downloadInfo[hd].(map[string]interface{})["urls"].([]string)
}

func (self *VideoInfo) DownloadInfo() map[string]interface{} {

	return self.downloadInfo
}

//实例基类
type Base struct {
	Name            string
	_VIDEO_PATTERNS []string
	Hd              map[string]string
}

func (base *Base) CurrentTime() (ts int64) {
	return time.Now().Unix()
}

func (base *Base) GetVideoInfo(url string) (info VideoInfo, err error) {
	return VideoInfo{}, nil
}

func (base *Base) GetHtml(url string) (html string, err error) {
	log.Println("request url ", url)
	return url + "html", nil
}

func (base *Base) Obj() (obj interface{}) {
	return base
}

func (self *Base) MatchUrl(url string) bool {
	if len(utils.R1Of(self._VIDEO_PATTERNS, url)) > 1 {
		return true
	}
	return false
}

func (base *Base) BuildDoc(url string) (doc *goquery.Document, err error) {
	log.Println("build doc ", url)
	doc, err = goquery.NewDocument(url)
	return
}

func (base *Base) BuildJson(url string) (json *simplejson.Json, err error) {
	video_html, err := utils.GetContent(url, nil)
	bjson := []byte(video_html)
	json, err = simplejson.NewJson(bjson)
	return
}

var (
	Spiders = make(map[string]Core)
)

func init() {
	YouKuRegister()
	QQRegister()
	IQiyiRegister()
	SohuRegister()
	LeTvRegister()
	BiLiBiLiRegister()
	TouTiaoRegister()
	PPTVRegister()
	AcFunRegister()
	DouBanRegister()
	PearVideoRegister()
	YinYueTaiRegister()
	SinaRegister()
	WeiBoRegister()
	HuYaRegister()
	CCTVRegister()
	Open163Register()
	ZuiYouRegister()
}

func GetExtractor(url string) (key string, spider Core) {
	//log.Println(url)
	for a, b := range Spiders {
		if b.MatchUrl(url) {
			key = a
			spider = b
			break
		}
	}
	return
}
