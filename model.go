package douyin

type ExplosiveSentenceVideo struct {
	Desc         string `json:"desc"`
	DiggCount    int64  `json:"digg_count"`
	CommentCount int64  `json:"comment_count"`
	PlayCount    int64  `json:"play_count"`
	AwemeId      string `json:"aweme_id" `
	VideoUrl     string `json:"video_url"`
	Time         int64  `json:"time"`
	SecUid       string `json:"sec_uid"`
	DouyinId     string `json:"douyin_id"`
	NickName     string `json:"nick_name"`
	ReleaseTime  int64  `json:"release_time"`
	CoverImage   string `gjson:"cover_image"`
	Comments     string `json:"comments" `
}
type VideoInfo struct {
	AnchorUserID  string `json:"anchor_user_id"`
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
