package douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"

	"github.com/guonaihong/gout"
	"github.com/kirinlabs/HttpRequest"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 只能通过cookie中的sessionid查看自己的相关信息
func GetUserNickName(sessionid string, options Options) (nickName string, err error) {
	douYinUrl := "https://creator.douyin.com/aweme/v1/creator/user/info"
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"content-type": "application/json;charset=UTF-8",
				"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":      "https://creator.douyin.com/creator-micro/certification",
				"cookie":       fmt.Sprintf("sessionid=%s", sessionid),
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"content-type": "application/json;charset=UTF-8",
				"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":      "https://creator.douyin.com/creator-micro/certification",
				"cookie":       fmt.Sprintf("sessionid=%s", sessionid),
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	nickName = gjson.Get(res, "douyin_user_verify_info.nick_name").String()
	return
}

func GetCommentsVideos(sessionid string, cursor int, options Options) (result []VideoInfo, err error) {
	douYinUrl := fmt.Sprintf("https://creator.douyin.com/aweme/v1/creator/item/list/?aid=2906&app_name=aweme_creator_platform&device_platform=web&cursor=%d&count=20", cursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	data := gjson.Get(res, "item_info_list").String()
	json.Unmarshal([]byte(data), &result)
	return
}

func GetCommentsList(sessionid string, cursor int, videoItemid string, options Options) (result []CommentInfo, err error) {
	douYinUrl := fmt.Sprintf("https://creator.douyin.com/aweme/v1/creator/comment/list/?aid=2906&app_name=aweme_creator_platform&device_platform=web&item_id=%s&cursor=%d&count=20&sort=HOT", url.QueryEscape(videoItemid), cursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}

	if err != nil {
		return
	}
	infoList := gjson.Get(res, "comment_info_list").String()
	json.Unmarshal([]byte(infoList), &result)
	return
}

func GetCommentsListReply(sessionid string, cursor int, comment_id string, options Options) (result []CommentInfo, err error) {
	douYinUrl := fmt.Sprintf("https://creator.douyin.com/aweme/v1/creator/comment/reply/list/?aid=2906&app_name=aweme_creator_platform&device_platform=web&comment_id=%s&cursor=%d&count=20&sort=HOT", url.QueryEscape(comment_id), cursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()

	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	infoList := gjson.Get(res, "comment_info_list").String()
	json.Unmarshal([]byte(infoList), &result)
	return
}
func DeleteCommentsList(sessionid string, commentId string, options Options) (err error) {
	douYinUrl := fmt.Sprintf("https://creator.douyin.com/aweme/v1/creator/comment/delete/?aid=2906&app_name=aweme_creator_platform&device_platform=web")
	var res string
	if options.Address == "" {
		err = gout.POST(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).SetJSON(gout.H{"comment_Id": commentId}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).POST(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).SetJSON(gout.H{"comment_Id": commentId}).BindBody(&res).Do()
	}

	return

}

func GetMyVideos(minCursor, maxCursor int64, sessionid string, options Options) (result []ExplosiveSentenceVideo, err error) {
	douYinUrl := fmt.Sprintf("https://creator.douyin.com/web/api/media/aweme/post/?scene=star_atlas&status=0&count=12&min_cursor=%d&max_cursor=%d", minCursor, maxCursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/content/manage",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://creator.douyin.com/creator-micro/content/manage",
				"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	r := gjson.Get(res, "aweme_list").Array()

	for i, _ := range r {
		var video ExplosiveSentenceVideo

		video.AwemeId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.aweme_id", i)).String()
		video.Desc = gjson.Get(res, fmt.Sprintf("aweme_list.%d.desc", i)).String()
		video.DouyinId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.short_id", i)).String()
		video.SecUid = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.sec_uid", i)).String()
		video.UniqueId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.unique_id", i)).String()
		video.NickName = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.nickname", i)).String()

		video.Time = time.Now().Unix()
		video.CommentCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.comment_count", i)).Int()
		video.DiggCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.digg_count", i)).Int()
		video.ShareCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.share_count", i)).Int()
		video.PlayCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.play_count", i)).Int()
		video.VideoUrl = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.play_addr.url_list.0", i)).String()
		video.Duration = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.duration", i)).Int()

		video.CoverImage = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.cover.url_list.0", i)).String()
		originUrl := gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.origin_cover.uri", i)).String()
		ReleaseTimeStrArray := strings.Split(originUrl, "_")
		parseInt, err := strconv.ParseInt(ReleaseTimeStrArray[len(ReleaseTimeStrArray)-1], 10, 64)
		if err == nil {
			video.ReleaseTime = parseInt
		}
		result = append(result, video)
	}
	return
}

// 通用接口
// 分享链接获取sec_uid
func GetSecUidBySharedUrl(sharedUrl string, options Options) (secUid string, err error) {
	var resp *HttpRequest.Response
	req := HttpRequest.NewRequest()

	req.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36",
	})
	req.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse /* 不进入重定向 */
	})

	if options.Address == "" {
		resp, err = req.Get(sharedUrl)
	} else {
		proxy, err := url.Parse(options.Address)
		if err != nil {
			fmt.Println(err)
		}
		resp, err = HttpRequest.Proxy(http.ProxyURL(proxy)).Get(sharedUrl)
	}
	if err != nil {
		return
	}
	defer resp.Close()
	if err != nil {
		return
	}
	if resp.StatusCode() != 302 {
		err = errors.New("get error")
		return
	}
	if len(resp.Headers().Values("location")) <= 0 {
		err = errors.New("接口可能已经改变")
		return
	}
	location := resp.Headers().Values("location")[0]
	regNew := regexp.MustCompile(`(?:did=)[a-z,A-Z，0-9, _, -]+`)
	if len(regNew.FindAllString(location, -1)) <= 0 {
		err = errors.New("接口可能已经改变")
		return
	}
	secUid = strings.Replace(regNew.FindAllString(location, -1)[0], "did=", "", 1)
	return
}

func GetOthersVideoByTimeStamp(secUid string, begin, end int64, options Options) (result []ExplosiveSentenceVideo, hasMore bool, minCursor, maxCursor int64, err error) {
	var resp *HttpRequest.Response
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36",
	})
	if options.Address == "" {
		resp, err = req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&_signature=PtCNCgAAXljWCq93QOKsFT7QjR",
			secUid, begin, end))
	} else {
		proxy, err1 := url.Parse(options.Address)
		if err1 != nil {
			err = err1
			return
		}
		resp, err = HttpRequest.Proxy(http.ProxyURL(proxy)).Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&_signature=PtCNCgAAXljWCq93QOKsFT7QjR",
			secUid, begin, end))
	}
	if err != nil {
		return
	}
	defer resp.Close()
	if err != nil {
		return
	}
	body, err := resp.Body()
	res := string(body)
	r := gjson.Get(res, "aweme_list").Array()
	hasMore = gjson.Get(res, "has_more").Bool()
	minCursor = gjson.Get(res, "min_cursor").Int()
	maxCursor = gjson.Get(res, "max_cursor").Int()

	for i, _ := range r {
		var video ExplosiveSentenceVideo

		video.AwemeId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.aweme_id", i)).String()
		video.Desc = gjson.Get(res, fmt.Sprintf("aweme_list.%d.desc", i)).String()
		video.VideoUrl = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.play_addr.url_list.0", i)).String()
		video.DouyinId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.short_id", i)).String()
		video.SecUid = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.sec_uid", i)).String()
		video.UniqueId = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.unique_id", i)).String()

		video.NickName = gjson.Get(res, fmt.Sprintf("aweme_list.%d.author.nickname", i)).String()
		video.Time = time.Now().Unix()
		video.CommentCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.comment_count", i)).Int()
		video.DiggCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.digg_count", i)).Int()
		video.ShareCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.share_count", i)).Int()

		video.PlayCount = gjson.Get(res, fmt.Sprintf("aweme_list.%d.statistics.play_count", i)).Int()
		video.Duration = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.duration", i)).Int()

		video.CoverImage = gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.cover.url_list.0", i)).String()
		originUrl := gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.origin_cover.uri", i)).String()
		ReleaseTimeStrArray := strings.Split(originUrl, "_")
		parseInt, err := strconv.ParseInt(ReleaseTimeStrArray[len(ReleaseTimeStrArray)-1], 10, 64)
		if err == nil {
			video.ReleaseTime = parseInt
		}
		if video.ReleaseTime == 0 {
			originUrlVideo := gjson.Get(res, fmt.Sprintf("aweme_list.%d.video.origin_cover.url_list.0", i)).String()
			u, err1 := url.Parse(originUrlVideo)
			if err1 == nil {
				params := u.Query()
				timeStr := params.Get("l")
				if len(timeStr) > 14 {
					tt, _ := time.ParseInLocation("20060102150405", timeStr[0:14], time.Local)

					video.ReleaseTime = tt.Unix()
				}
			}
		}
		if video.ReleaseTime == 0 {
			video.ReleaseTime = video.Time
		}

		result = append(result, video)
	}
	return
}

func GetOthersCommentsByAwemeId(awemeid string, cursor int, options Options) (result []OtherCommentInfo, hasMore bool, err error) {

	douYinUrl := fmt.Sprintf("https://www.douyin.com/aweme/v1/web/comment/list/?aweme_id=%s&cursor=%d&count=50", awemeid, cursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	r := gjson.Get(res, "comments").String()
	if gjson.Get(res, "has_more").Int() > 0 {
		hasMore = true
	}
	json.Unmarshal([]byte(r), &result)
	return
}

func GetOthersUserInfo(secUid string, options Options) (result OtherUserInfo, err error) {
	douYinUrl := fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/user/info/?sec_uid=%s", secUid)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()

	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	if res == "" {
		err = errors.New("没有返回值,稍后继续重试")
		return
	}
	r := gjson.Get(res, "user_info").String()
	json.Unmarshal([]byte(r), &result)
	return
}

// 最大二十
func GetVideosInfoByAwemeId(awemeIdList []string, options Options) (result []ExplosiveSentenceVideo, err error) {
	if len(awemeIdList) > 20 {
		err = errors.New("最大同时支持20个视频")
		return
	}
	awemeIdListStr := strings.Join(awemeIdList, ",")
	douYinUrl := fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s", awemeIdListStr)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()

	}
	if err != nil {
		return
	}
	if res == "" {
		errors.New("请稍后重试")
		return
	}
	r := gjson.Get(res, "item_list").Array()

	for i, _ := range r {
		var video ExplosiveSentenceVideo

		video.AwemeId = gjson.Get(res, fmt.Sprintf("item_list.%d.aweme_id", i)).String()
		video.Desc = gjson.Get(res, fmt.Sprintf("item_list.%d.desc", i)).String()
		video.DouyinId = gjson.Get(res, fmt.Sprintf("item_list.%d.author.short_id", i)).String()
		video.SecUid = gjson.Get(res, fmt.Sprintf("item_list.%d.author.sec_uid", i)).String()
		video.NickName = gjson.Get(res, fmt.Sprintf("item_list.%d.author.nickname", i)).String()
		video.UniqueId = gjson.Get(res, fmt.Sprintf("item_list.%d.author.unique_id", i)).String()

		video.Time = time.Now().Unix()
		video.CommentCount = gjson.Get(res, fmt.Sprintf("item_list.%d.statistics.comment_count", i)).Int()
		video.DiggCount = gjson.Get(res, fmt.Sprintf("item_list.%d.statistics.digg_count", i)).Int()
		video.ShareCount = gjson.Get(res, fmt.Sprintf("item_list.%d.statistics.share_count", i)).Int()

		video.PlayCount = gjson.Get(res, fmt.Sprintf("item_list.%d.statistics.play_count", i)).Int()
		video.Duration = gjson.Get(res, fmt.Sprintf("item_list.%d.video.duration", i)).Int()
		video.VideoUrl = gjson.Get(res, fmt.Sprintf("item_list.%d.video.play_addr.url_list.0", i)).String()

		video.CoverImage = gjson.Get(res, fmt.Sprintf("item_list.%d.video.cover.url_list.0", i)).String()
		video.ReleaseTime = gjson.Get(res, fmt.Sprintf("item_list.%d.create_time", i)).Int()
		result = append(result, video)
	}
	return
}

func GetOthersUserInfoByHtml(secUid string, options Options) (result UserInfoFromHtml, err error) {
	userUrl := fmt.Sprintf("https://www.douyin.com/user/%s", secUid)
	var res *http.Response
	if options.Address == "" {
		res, err = http.Get(userUrl)
		if err != nil {
			return UserInfoFromHtml{}, err
		}

	} else {
		uri, _ := url.Parse(options.Address)

		c := http.Client{
			Transport: &http.Transport{
				// 设置代理
				Proxy: http.ProxyURL(uri),
			},
		}
		res, err = c.Get(userUrl)
		if err != nil {
			return UserInfoFromHtml{}, err
		}

	}
	//str, _ := io.ReadAll(res.Body)
	//fmt.Println(string(str))
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	jsonStr := doc.Find("#RENDER_DATA").Text()
	enEscapeUrl, _ := url.QueryUnescape(jsonStr)

	r := gjson.Get(enEscapeUrl, "37.user.user").String()
	json.Unmarshal([]byte(r), &result)
	if err != nil {
		return
	}
	return
}

func GetFirstPageCommentBySessionId(awemeId string, sessionId string, options Options) (result []OtherCommentInfo, err error) {
	url := fmt.Sprintf("https://www.douyin.com/aweme/v1/web/comment/list/?device_platform=webapp&aid=6383&channel=channel_pc_web&aweme_id=%s&cursor=0&count=20&item_type=0&insert_ids=&rcFT=&pc_client_type=1&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=2560&screen_height=1440&browser_language=zh-CN&browser_platform=Win32&browser_name=Chrome&browser_version=108.0.0.0&browser_online=true&engine_name=Blink&engine_version=108.0.0.0&os_name=Windows&os_version=10&cpu_core_num=12&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=0", awemeId)
	var res string

	if options.Address == "" {
		err = gout.GET(url).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://www.douyin.com/user",
				"cookie":          "sessionid=" + sessionId,
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(url).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"referer":         "https://www.douyin.com/user",
				"cookie":          "sessionid=" + sessionId,
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
	}
	if err != nil {
		return
	}
	r := gjson.Get(res, "comments").String()
	json.Unmarshal([]byte(r), &result)
	return
}
