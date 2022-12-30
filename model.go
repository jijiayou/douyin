package douyin

//type ExplosiveSentenceVideo struct {
//	Desc         string `json:"desc"`
//	AwemeId      string `json:"aweme_id" `
//	Time         int64  `json:"time"`
//	SecUid       string `json:"sec_uid"`
//	UniqueId     string `json:"unique_id"`
//	DouyinId     string `json:"douyin_id"`
//	NickName     string `json:"nick_name"`
//	VideoUrl     string `json:"video_url"`
//	DiggCount    int64  `json:"digg_count"`
//	CommentCount int64  `json:"comment_count"`
//	ShareCount   int64  `json:"share_count"`
//	PlayCount    int64  `json:"play_count"`
//	CollectCount int64  `json:"collect_count"`
//	ReleaseTime  int64  `json:"release_time"`
//	CoverImage   string `json:"cover_image"`
//	Duration     int64  `json:"duration"`
//	Comments     string `json:"comments" `
//}

type ExplosiveSentenceVideo struct {
	ID           int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Desc         string `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`
	AwemeId      string `gorm:"column:aweme_id" db:"aweme_id" json:"aweme_id" form:"aweme_id"`
	Time         int64  `gorm:"column:time" db:"time" json:"time" form:"time"`
	SecUid       string `gorm:"column:sec_uid" db:"sec_uid" json:"sec_uid" form:"sec_uid"`
	UniqueId     string `gorm:"column:unique_id" db:"unique_id" json:"unique_id" form:"unique_id"`
	DouyinId     string `gorm:"column:douyin_id" db:"douyin_id" json:"douyin_id" form:"douyin_id"`
	VideoUrl     string `gorm:"column:video_url" db:"video_url" json:"video_url" form:"video_url"`
	CommentCount int64  `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"`
	DiggCount    int64  `gorm:"column:digg_count" db:"digg_count" json:"digg_count" form:"digg_count"`
	ShareCount   int64  `gorm:"column:share_count" db:"share_count" json:"share_count" form:"share_count"`
	PlayCount    int64  `gorm:"column:play_count" db:"play_count" json:"play_count" form:"play_count"`
	NickName     string `gorm:"column:nick_name" db:"nick_name" json:"nick_name" form:"nick_name"`
	CollectCount int64  `gorm:"column:collect_count" db:"collect_count" json:"collect_count" form:"collect_count"`
	ReleaseTime  int64  `gorm:"column:release_time" db:"release_time" json:"release_time" form:"release_time"`
	CoverImage   string `gorm:"column:cover_image" db:"cover_image" json:"cover_image" form:"cover_image"`
	Comments     string `gorm:"column:comments" db:"comments" json:"comments" form:"comments"`
	Duration     int64  `gorm:"column:duration" db:"duration" json:"duration" form:"duration"`
}

func (ExplosiveSentenceVideo) TableName() string {
	return "explosive_sentence_video"
}

type VideoInfo struct {
	AnchorUserID  string `json:"anchor_user_id"`
	AwemeId       string `json:"aweme_id"`
	CommentCount  int    `json:"comment_count"`
	CoverImageURL string `json:"cover_image_url"`
	CreateTime    string `json:"create_time"`
	Duration      int    `json:"duration"`
	ItemID        string `json:"item_id"`
	ItemLink      string `json:"item_link"`
	MediaType     int    `json:"media_type"`
	Title         string `json:"title"`
}

type CommentInfo struct {
	CommentID       string          `json:"comment_id"`
	CreateTime      string          `json:"create_time"`
	DiggCount       string          `json:"digg_count"`
	Followed        bool            `json:"followed"`
	Following       bool            `json:"following"`
	IsAuthor        bool            `json:"is_author"`
	Level           int             `json:"level"`
	ReplyCount      string          `json:"reply_count"`
	ReplyToUserInfo ReplyToUserInfo `json:"reply_to_user_info"`
	Status          int             `json:"status"`
	Text            string          `json:"text"`
	UserDigg        bool            `json:"user_digg"`
	UserInfo        UserInfo        `json:"user_info"`
}
type ReplyToUserInfo struct {
	AvatarURL  string `json:"avatar_url"`
	ScreenName string `json:"screen_name"`
	UserID     string `json:"user_id"`
}
type UserInfo struct {
	AvatarURL  string `json:"avatar_url"`
	ScreenName string `json:"screen_name"`
	UserID     string `json:"user_id"`
}

type OtherCommentInfo struct {
	AwemeID           string         `json:"aweme_id"`
	Cid               string         `json:"cid"`
	CreateTime        int            `json:"create_time"`
	DiggCount         int            `json:"digg_count"`
	IsAuthorDigged    bool           `json:"is_author_digged"`
	LabelText         string         `json:"label_text"`
	LabelType         int            `json:"label_type"`
	ReplyComment      []ReplyComment `json:"reply_comment"`
	ReplyCommentTotal int            `json:"reply_comment_total"`
	ReplyID           string         `json:"reply_id"`
	ReplyToReplyID    string         `json:"reply_to_reply_id"`
	Status            int            `json:"status"`
	StickPosition     int            `json:"stick_position"`
	Text              string         `json:"text"`
	User              User           `json:"user"`
	UserBuried        bool           `json:"user_buried"`
	UserDigged        int            `json:"user_digged"`
}

type ReplyComment struct {
	AwemeID        string        `json:"aweme_id"`
	CanShare       bool          `json:"can_share"`
	Cid            string        `json:"cid"`
	CreateTime     int           `json:"create_time"`
	DiggCount      int           `json:"digg_count"`
	ImageList      interface{}   `json:"image_list"`
	IsAuthorDigged bool          `json:"is_author_digged"`
	IsHot          bool          `json:"is_hot"`
	IsNoteComment  int           `json:"is_note_comment"`
	LabelList      interface{}   `json:"label_list"`
	LabelText      string        `json:"label_text"`
	LabelType      int           `json:"label_type"`
	ReplyComment   interface{}   `json:"reply_comment"`
	ReplyID        string        `json:"reply_id"`
	ReplyToReplyID string        `json:"reply_to_reply_id"`
	Status         int           `json:"status"`
	Text           string        `json:"text"`
	TextExtra      []interface{} `json:"text_extra"`
	TextMusicInfo  interface{}   `json:"text_music_info"`
	User           User          `json:"user"`
	UserBuried     bool          `json:"user_buried"`
	UserDigged     int           `json:"user_digged"`
}
type User struct {
	SecUID   string `json:"sec_uid"`
	Nickname string `json:"nickname"`
	ShortID  string `json:"short_id"`
}
type Options struct {
	Address string `json:"address"`
}

// 通用用户接口
type OtherUserInfo struct {
	AvatarLarger            AvatarLarger  `json:"avatar_larger"`
	AvatarMedium            AvatarMedium  `json:"avatar_medium"`
	UniqueID                string        `json:"unique_id"`
	IsGovMediaVip           bool          `json:"is_gov_media_vip"`
	FollowStatus            int           `json:"follow_status"`
	FavoritingCount         int           `json:"favoriting_count"`
	SecUID                  string        `json:"sec_uid"`
	Secret                  int           `json:"secret"`
	IsMixUser               bool          `json:"is_mix_user"`
	UID                     string        `json:"uid"`
	TotalFavorited          string        `json:"total_favorited"`
	EnterpriseVerifyReason  string        `json:"enterprise_verify_reason"`
	MplatformFollowersCount int           `json:"mplatform_followers_count"`
	FollowingCount          int           `json:"following_count"`
	FollowerCount           int           `json:"follower_count"`
	ShowFavoriteList        bool          `json:"show_favorite_list"`
	MixInfo                 []MixInfo     `json:"mix_info"`
	AwemeCount              int           `json:"aweme_count"`
	MixCount                int           `json:"mix_count"`
	Nickname                string        `json:"nickname"`
	VerificationType        int           `json:"verification_type"`
	PlatformSyncInfo        []interface{} `json:"platform_sync_info"`
	ShortID                 string        `json:"short_id"`
	Signature               string        `json:"signature"`
	AvatarThumb             AvatarThumb   `json:"avatar_thumb"`
}
type AvatarLarger struct {
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
}
type AvatarMedium struct {
	URLList []string `json:"url_list"`
	URI     string   `json:"uri"`
}
type MixInfo struct {
	MixID   string `json:"mix_id"`
	MixName string `json:"mix_name"`
}

type AvatarThumb struct {
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
}

type UserInfoFromHtml struct {
	UID                     string      `json:"uid"`
	SecUID                  string      `json:"secUid"`
	ShortID                 string      `json:"shortId"`
	Nickname                string      `json:"nickname"`
	Desc                    string      `json:"desc"`
	Gender                  int         `json:"gender"`
	AvatarURL               string      `json:"avatarUrl"`
	Avatar300URL            string      `json:"avatar300Url"`
	FollowStatus            int         `json:"followStatus"`
	FollowerStatus          int         `json:"followerStatus"`
	AwemeCount              int         `json:"awemeCount"`
	FollowingCount          int         `json:"followingCount"`
	FollowerCount           int         `json:"followerCount"`
	MplatformFollowersCount int         `json:"mplatformFollowersCount"`
	FavoritingCount         int         `json:"favoritingCount"`
	TotalFavorited          int         `json:"totalFavorited"`
	UniqueID                string      `json:"uniqueId"`
	Age                     int         `json:"age"`
	Country                 string      `json:"country"`
	Province                string      `json:"province"`
	City                    string      `json:"city"`
	District                string      `json:"district"`
	School                  interface{} `json:"school"`
	EnterpriseVerifyReason  string      `json:"enterpriseVerifyReason"`
	Secret                  int         `json:"secret"`
	UserCanceled            bool        `json:"userCanceled"`
	ShareQrcodeURL          string      `json:"shareQrcodeUrl"`
	ShareInfo               struct {
		BoolPersist   int    `json:"boolPersist"`
		ShareDesc     string `json:"shareDesc"`
		ShareImageURL struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
		} `json:"shareImageUrl"`
		ShareQrcodeURL struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
		} `json:"shareQrcodeUrl"`
		ShareURL       string `json:"shareUrl"`
		ShareWeiboDesc string `json:"shareWeiboDesc"`
	} `json:"shareInfo"`
	RoomID                       int         `json:"roomId"`
	IsBlocked                    bool        `json:"isBlocked"`
	IsBlock                      bool        `json:"isBlock"`
	IsBan                        bool        `json:"isBan"`
	FavoritePermission           interface{} `json:"favoritePermission"`
	ShowFavoriteList             bool        `json:"showFavoriteList"`
	ViewHistoryPermission        bool        `json:"viewHistoryPermission"`
	IPLocation                   string      `json:"ipLocation"`
	IsGovMediaVip                bool        `json:"isGovMediaVip"`
	IsStar                       bool        `json:"isStar"`
	NeedSpecialShowFollowerCount bool        `json:"needSpecialShowFollowerCount"`
	IsNotShow                    bool        `json:"isNotShow"`
	IsOverFollower               bool        `json:"isOverFollower"`
}
