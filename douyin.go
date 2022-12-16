package douyin

import (
	"encoding/json"
	"errors"
	"fmt"
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
		nickName = gjson.Get(res, "douyin_user_verify_info.nick_name").String()
		return
	}
	c := &http.Client{}
	err = gout.New(c).GET(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"content-type": "application/json;charset=UTF-8",
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"referer":      "https://creator.douyin.com/creator-micro/certification",
			"cookie":       fmt.Sprintf("sessionid=%s", sessionid),
		}).BindBody(&res).Do()
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
		data := gjson.Get(res, "item_info_list").String()
		json.Unmarshal([]byte(data), &result)
		return
	}
	c := &http.Client{}
	err = gout.New(c).GET(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
			"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
			"accept-language": "zh-CN,zh;q=0.9",
		}).BindBody(&res).Do()
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
		infoList := gjson.Get(res, "comment_info_list").String()
		json.Unmarshal([]byte(infoList), &result)
		return
	}
	c := &http.Client{}
	err = gout.New(c).GET(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
			"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
			"accept-language": "zh-CN,zh;q=0.9",
		}).BindBody(&res).Do()
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
		infoList := gjson.Get(res, "comment_info_list").String()
		json.Unmarshal([]byte(infoList), &result)
		return
	}
	c := &http.Client{}
	err = gout.New(c).GET(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
			"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
			"accept-language": "zh-CN,zh;q=0.9",
		}).BindBody(&res).Do()
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
		return
	}
	c := &http.Client{}
	err = gout.New(c).POST(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"referer":         "https://creator.douyin.com/creator-micro/data/following/comment",
			"cookie":          fmt.Sprintf("sessionid=%s", sessionid),
			"accept-language": "zh-CN,zh;q=0.9",
		}).SetJSON(gout.H{"comment_Id": commentId}).BindBody(&res).Do()
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

	defer resp.Close()
	if err != nil {
		return
	}
	if resp.StatusCode() != 302 {
		err = errors.New("get error")
		return
	}
	location := resp.Headers().Values("location")[0]

	regNew := regexp.MustCompile(`(?:sec_uid=)[a-z,A-Z，0-9, _, -]+`)
	secUid = strings.Replace(regNew.FindAllString(location, -1)[0], "sec_uid=", "", 1)
	return
}

func GetOthersVideoByTimeStamp(secUid string, begin, end int64, options Options) (result []ExplosiveSentenceVideo, hasMore bool, err error) {
	var resp *HttpRequest.Response
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36",
	})
	if options.Address == "" {
		resp, err = req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&_signature=PtCNCgAAXljWCq93QOKsFT7QjR",
			secUid, begin, end))
	} else {
		proxy, err := url.Parse(options.Address)
		if err != nil {
			fmt.Println(err)
		}
		resp, err = HttpRequest.Proxy(http.ProxyURL(proxy)).Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&_signature=PtCNCgAAXljWCq93QOKsFT7QjR",
			secUid, begin, end))
	}

	defer resp.Close()
	if err != nil {
		return
	}
	body, err := resp.Body()
	res := string(body)
	r := gjson.Get(res, "aweme_list").Array()
	hasMore = gjson.Get(res, "has_more").Bool()

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
		result = append(result, video)
	}
	return
}

func GetOthersCommentsByAwemeId(awemeid string, cursor int, options Options) (result []OtherCommentInfo, err error) {

	douYinUrl := fmt.Sprintf("https://www.douyin.com/aweme/v1/web/comment/list/?aweme_id=%s&cursor=%d&count=50", awemeid, cursor)
	var res string
	if options.Address == "" {
		err = gout.GET(douYinUrl).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
		r := gjson.Get(res, "comments").String()
		json.Unmarshal([]byte(r), &result)
		return
	}
	c := &http.Client{}
	err = gout.New(c).GET(douYinUrl).
		SetProxy(options.Address).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"accept-language": "zh-CN,zh;q=0.9",
		}).BindBody(&res).Do()
	r := gjson.Get(res, "comments").String()
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
		r := gjson.Get(res, "user_info").String()
		json.Unmarshal([]byte(r), &result)
		return
	} else {
		c := &http.Client{}
		err = gout.New(c).GET(douYinUrl).
			SetProxy(options.Address).
			SetHeader(gout.H{
				"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
				"accept-language": "zh-CN,zh;q=0.9",
			}).BindBody(&res).Do()
		r := gjson.Get(res, "user_info").String()
		json.Unmarshal([]byte(r), &result)
		return
	}
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
