package douyin

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/mvdan/xurls"
	"net/http"
	"regexp"
	"strings"
)

type AnalysisV2Response struct {
	StatusCode int `json:"status_code"`
	ItemList   []struct {
		AwemePoiInfo struct {
			Tag  string `json:"tag"`
			Icon struct {
				UrlList []string `json:"url_list"`
				Uri     string   `json:"uri"`
			} `json:"icon"`
			PoiName  string `json:"poi_name"`
			TypeName string `json:"type_name"`
		} `json:"aweme_poi_info"`
		Images interface{} `json:"images"`
		Author struct {
			Geofencing       interface{} `json:"geofencing"`
			CardEntries      interface{} `json:"card_entries"`
			ShortId          string      `json:"short_id"`
			Nickname         string      `json:"nickname"`
			FollowStatus     int         `json:"follow_status"`
			UniqueId         string      `json:"unique_id"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			Uid              string      `json:"uid"`
			AvatarLarger     struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_larger"`
			PolicyVersion interface{} `json:"policy_version"`
			MixInfo       interface{} `json:"mix_info"`
			Signature     string      `json:"signature"`
			AvatarThumb   struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_medium"`
			FollowersDetail interface{} `json:"followers_detail"`
			TypeLabel       interface{} `json:"type_label"`
		} `json:"author"`
		ChaList []struct {
			ViewCount      int    `json:"view_count"`
			HashTagProfile string `json:"hash_tag_profile"`
			Cid            string `json:"cid"`
			CoverItem      struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_item"`
			UserCount    int         `json:"user_count"`
			ConnectMusic interface{} `json:"connect_music"`
			Type         int         `json:"type"`
			IsCommerce   bool        `json:"is_commerce"`
			ChaName      string      `json:"cha_name"`
			Desc         string      `json:"desc"`
		} `json:"cha_list"`
		Duration     int         `json:"duration"`
		LongVideo    interface{} `json:"long_video"`
		Desc         string      `json:"desc"`
		AuthorUserId int64       `json:"author_user_id"`
		LabelTopText interface{} `json:"label_top_text"`
		IsPreview    int         `json:"is_preview"`
		CreateTime   int         `json:"create_time"`
		ShareUrl     string      `json:"share_url"`
		RiskInfos    struct {
			Warn             bool   `json:"warn"`
			Type             int    `json:"type"`
			Content          string `json:"content"`
			ReflowUnplayable int    `json:"reflow_unplayable"`
		} `json:"risk_infos"`
		Promotions interface{} `json:"promotions"`
		Music      struct {
			Duration int    `json:"duration"`
			Id       int64  `json:"id"`
			Mid      string `json:"mid"`
			Title    string `json:"title"`
			CoverHd  struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_thumb"`
			Author  string `json:"author"`
			PlayUrl struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"play_url"`
			Position interface{} `json:"position"`
			Status   int         `json:"status"`
		} `json:"music"`
		CommentList interface{} `json:"comment_list"`
		ForwardId   string      `json:"forward_id"`
		GroupIdStr  string      `json:"group_id_str"`
		Video       struct {
			OriginCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"origin_cover"`
			HasWatermark bool `json:"has_watermark"`
			Duration     int  `json:"duration"`
			Height       int  `json:"height"`
			DynamicCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			Width    int         `json:"width"`
			Ratio    string      `json:"ratio"`
			BitRate  interface{} `json:"bit_rate"`
			Vid      string      `json:"vid"`
			PlayAddr struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"` // 真实去水印地址
			} `json:"play_addr"`
			Cover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover"`
		} `json:"video"`
		TextExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagId   int64  `json:"hashtag_id"`
		} `json:"text_extra"`
		VideoLabels interface{} `json:"video_labels"`
		VideoText   interface{} `json:"video_text"`
		AwemeType   int         `json:"aweme_type"`
		ImageInfos  interface{} `json:"image_infos"`
		AwemeId     string      `json:"aweme_id"`
		Statistics  struct {
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
			AwemeId      string `json:"aweme_id"`
		} `json:"statistics"`
		IsLiveReplay bool `json:"is_live_replay"`
		ShareInfo    struct {
			ShareWeiboDesc string `json:"share_weibo_desc"`
			ShareDesc      string `json:"share_desc"`
			ShareTitle     string `json:"share_title"`
		} `json:"share_info"`
		Geofencing interface{} `json:"geofencing"`
		GroupId    int64       `json:"group_id"`
	} `json:"item_list"`
	FilterList []interface{} `json:"filter_list"`
	Extra      struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
}

type AnalysisV2Result struct {
	Result AnalysisV2Response // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newAnalysisV2Result(result AnalysisV2Response, body []byte, http gorequest.Response) *AnalysisV2Result {
	return &AnalysisV2Result{Result: result, Body: body, Http: http}
}

// AnalysisV2 抖音解析
func (c *Client) AnalysisV2(ctx context.Context, content string) (*AnalysisV2Result, error) {

	// 提取url
	var url string
	if strings.Contains(content, "douyin.com") {
		url = xurls.Relaxed.FindString(content)
	} else if strings.Contains(content, "iesdouyin.com") {
		url = xurls.Relaxed.FindString(content)
	} else {
		return newAnalysisV2Result(AnalysisV2Response{}, nil, gorequest.Response{}), errors.New("url为空")
	}

	// 重定向信息
	request302, err := c.request302(url)
	if err != nil {
		return newAnalysisV2Result(AnalysisV2Response{}, nil, gorequest.Response{}), err
	}

	// 提取编号
	itemIds := regexp.MustCompile(`\d+`).FindStringSubmatch(request302)
	if len(itemIds) < 1 {
		return newAnalysisV2Result(AnalysisV2Response{}, nil, gorequest.Response{}), errors.New("参数错误")
	}

	// 请求
	request, err := c.request(ctx, "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids="+itemIds[0], nil, http.MethodGet)
	if err != nil {
		return newAnalysisV2Result(AnalysisV2Response{}, request.ResponseBody, request), err
	}

	// 定义
	var response AnalysisV2Response
	err = json.Unmarshal(request.ResponseBody, &response)
	return newAnalysisV2Result(response, request.ResponseBody, request), err
}
