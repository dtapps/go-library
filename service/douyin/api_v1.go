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

type AnalysisV1Response struct {
	AwemeDetail struct {
		Anchors interface{} `json:"anchors"`
		Author  struct {
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			CfList          interface{} `json:"cf_list"`
			CloseFriendType int         `json:"close_friend_type"`
			ContactsStatus  int         `json:"contacts_status"`
			ContrailList    interface{} `json:"contrail_list"`
			CoverUrl        []struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_url"`
			CreateTime                             int         `json:"create_time"`
			CustomVerify                           string      `json:"custom_verify"`
			DataLabelList                          interface{} `json:"data_label_list"`
			EndorsementInfoList                    interface{} `json:"endorsement_info_list"`
			EnterpriseVerifyReason                 string      `json:"enterprise_verify_reason"`
			FavoritingCount                        int         `json:"favoriting_count"`
			FollowStatus                           int         `json:"follow_status"`
			FollowerCount                          int         `json:"follower_count"`
			FollowerListSecondaryInformationStruct interface{} `json:"follower_list_secondary_information_struct"`
			FollowerStatus                         int         `json:"follower_status"`
			FollowingCount                         int         `json:"following_count"`
			ImRoleIds                              interface{} `json:"im_role_ids"`
			IsAdFake                               bool        `json:"is_ad_fake"`
			IsBlockedV2                            bool        `json:"is_blocked_v2"`
			IsBlockingV2                           bool        `json:"is_blocking_v2"`
			IsCf                                   int         `json:"is_cf"`
			MaxFollowerCount                       int         `json:"max_follower_count"`
			Nickname                               string      `json:"nickname"`
			NotSeenItemIdList                      interface{} `json:"not_seen_item_id_list"`
			NotSeenItemIdListV2                    interface{} `json:"not_seen_item_id_list_v2"`
			OfflineInfoList                        interface{} `json:"offline_info_list"`
			PersonalTagList                        interface{} `json:"personal_tag_list"`
			PreventDownload                        bool        `json:"prevent_download"`
			RiskNoticeText                         string      `json:"risk_notice_text"`
			SecUid                                 string      `json:"sec_uid"`
			Secret                                 int         `json:"secret"`
			ShareInfo                              struct {
				ShareDesc      string `json:"share_desc"`
				ShareDescInfo  string `json:"share_desc_info"`
				ShareQrcodeUrl struct {
					Height  int      `json:"height"`
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
				} `json:"share_qrcode_url"`
				ShareTitle       string `json:"share_title"`
				ShareTitleMyself string `json:"share_title_myself"`
				ShareTitleOther  string `json:"share_title_other"`
				ShareUrl         string `json:"share_url"`
				ShareWeiboDesc   string `json:"share_weibo_desc"`
			} `json:"share_info"`
			ShortId             string      `json:"short_id"`
			Signature           string      `json:"signature"`
			SignatureExtra      interface{} `json:"signature_extra"`
			SpecialPeopleLabels interface{} `json:"special_people_labels"`
			Status              int         `json:"status"`
			TextExtra           interface{} `json:"text_extra"`
			TotalFavorited      int         `json:"total_favorited"`
			Uid                 string      `json:"uid"`
			UniqueId            string      `json:"unique_id"`
			UserAge             int         `json:"user_age"`
			UserCanceled        bool        `json:"user_canceled"`
			UserPermissions     interface{} `json:"user_permissions"`
			VerificationType    int         `json:"verification_type"`
		} `json:"author"`
		AuthorMaskTag int   `json:"author_mask_tag"`
		AuthorUserId  int64 `json:"author_user_id"`
		AwemeControl  struct {
			CanComment     bool `json:"can_comment"`
			CanForward     bool `json:"can_forward"`
			CanShare       bool `json:"can_share"`
			CanShowComment bool `json:"can_show_comment"`
		} `json:"aweme_control"`
		AwemeId               string      `json:"aweme_id"`
		AwemeType             int         `json:"aweme_type"`
		ChallengePosition     interface{} `json:"challenge_position"`
		ChapterList           interface{} `json:"chapter_list"`
		CollectStat           int         `json:"collect_stat"`
		CommentGid            int64       `json:"comment_gid"`
		CommentList           interface{} `json:"comment_list"`
		CommentPermissionInfo struct {
			CanComment              bool `json:"can_comment"`
			CommentPermissionStatus int  `json:"comment_permission_status"`
			ItemDetailEntry         bool `json:"item_detail_entry"`
			PressEntry              bool `json:"press_entry"`
			ToastGuide              bool `json:"toast_guide"`
		} `json:"comment_permission_info"`
		CommerceConfigData interface{} `json:"commerce_config_data"`
		CommonBarInfo      string      `json:"common_bar_info"`
		ComponentInfoV2    string      `json:"component_info_v2"`
		CoverLabels        interface{} `json:"cover_labels"`
		CreateTime         int         `json:"create_time"`
		Desc               string      `json:"desc"`
		DiggLottie         struct {
			CanBomb  int    `json:"can_bomb"`
			LottieId string `json:"lottie_id"`
		} `json:"digg_lottie"`
		DisableRelationBar      int           `json:"disable_relation_bar"`
		DislikeDimensionList    interface{}   `json:"dislike_dimension_list"`
		DuetAggregateInMusicTab bool          `json:"duet_aggregate_in_music_tab"`
		Duration                int           `json:"duration"`
		Geofencing              []interface{} `json:"geofencing"`
		GeofencingRegions       interface{}   `json:"geofencing_regions"`
		GroupId                 string        `json:"group_id"`
		HybridLabel             interface{}   `json:"hybrid_label"`
		ImageAlbumMusicInfo     struct {
			BeginTime int `json:"begin_time"`
			EndTime   int `json:"end_time"`
			Volume    int `json:"volume"`
		} `json:"image_album_music_info"`
		ImageInfos     interface{} `json:"image_infos"`
		ImageList      interface{} `json:"image_list"`
		Images         interface{} `json:"images"`
		ImgBitrate     interface{} `json:"img_bitrate"`
		ImpressionData struct {
			GroupIdListA   []int64     `json:"group_id_list_a"`
			GroupIdListB   []int64     `json:"group_id_list_b"`
			SimilarIdListA interface{} `json:"similar_id_list_a"`
			SimilarIdListB interface{} `json:"similar_id_list_b"`
		} `json:"impression_data"`
		InteractionStickers  interface{} `json:"interaction_stickers"`
		IsAds                bool        `json:"is_ads"`
		IsCollectsSelected   int         `json:"is_collects_selected"`
		IsDuetSing           bool        `json:"is_duet_sing"`
		IsImageBeat          bool        `json:"is_image_beat"`
		IsLifeItem           bool        `json:"is_life_item"`
		IsStory              int         `json:"is_story"`
		IsTop                int         `json:"is_top"`
		ItemWarnNotification struct {
			Content string `json:"content"`
			Show    bool   `json:"show"`
			Type    int    `json:"type"`
		} `json:"item_warn_notification"`
		LabelTopText interface{} `json:"label_top_text"`
		LongVideo    interface{} `json:"long_video"`
		Music        struct {
			Album            string        `json:"album"`
			ArtistUserInfos  interface{}   `json:"artist_user_infos"`
			Artists          []interface{} `json:"artists"`
			AuditionDuration int           `json:"audition_duration"`
			Author           string        `json:"author"`
			AuthorDeleted    bool          `json:"author_deleted"`
			AuthorPosition   interface{}   `json:"author_position"`
			AuthorStatus     int           `json:"author_status"`
			AvatarLarge      struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_large"`
			AvatarMedium struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_medium"`
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			BindedChallengeId int  `json:"binded_challenge_id"`
			CanBackgroundPlay bool `json:"can_background_play"`
			CollectStat       int  `json:"collect_stat"`
			CoverHd           struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_large"`
			CoverMedium struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_thumb"`
			DmvAutoShow          bool          `json:"dmv_auto_show"`
			DspStatus            int           `json:"dsp_status"`
			Duration             int           `json:"duration"`
			EndTime              int           `json:"end_time"`
			ExternalSongInfo     []interface{} `json:"external_song_info"`
			Extra                string        `json:"extra"`
			Id                   int64         `json:"id"`
			IdStr                string        `json:"id_str"`
			IsAudioUrlWithCookie bool          `json:"is_audio_url_with_cookie"`
			IsCommerceMusic      bool          `json:"is_commerce_music"`
			IsDelVideo           bool          `json:"is_del_video"`
			IsMatchedMetadata    bool          `json:"is_matched_metadata"`
			IsOriginal           bool          `json:"is_original"`
			IsOriginalSound      bool          `json:"is_original_sound"`
			IsPgc                bool          `json:"is_pgc"`
			IsRestricted         bool          `json:"is_restricted"`
			IsVideoSelfSee       bool          `json:"is_video_self_see"`
			LunaInfo             struct {
				IsLunaUser bool `json:"is_luna_user"`
			} `json:"luna_info"`
			LyricShortPosition interface{} `json:"lyric_short_position"`
			Mid                string      `json:"mid"`
			MusicChartRanks    interface{} `json:"music_chart_ranks"`
			MusicStatus        int         `json:"music_status"`
			MusicianUserInfos  interface{} `json:"musician_user_infos"`
			MuteShare          bool        `json:"mute_share"`
			OfflineDesc        string      `json:"offline_desc"`
			OwnerHandle        string      `json:"owner_handle"`
			OwnerId            string      `json:"owner_id"`
			OwnerNickname      string      `json:"owner_nickname"`
			PgcMusicType       int         `json:"pgc_music_type"`
			PlayUrl            struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlKey  string   `json:"url_key"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"play_url"`
			Position                  interface{} `json:"position"`
			PreventDownload           bool        `json:"prevent_download"`
			PreventItemDownloadStatus int         `json:"prevent_item_download_status"`
			PreviewEndTime            int         `json:"preview_end_time"`
			PreviewStartTime          int         `json:"preview_start_time"`
			ReasonType                int         `json:"reason_type"`
			Redirect                  bool        `json:"redirect"`
			SchemaUrl                 string      `json:"schema_url"`
			SearchImpr                struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			SecUid         string `json:"sec_uid"`
			ShootDuration  int    `json:"shoot_duration"`
			SourcePlatform int    `json:"source_platform"`
			StartTime      int    `json:"start_time"`
			Status         int    `json:"status"`
			StrongBeatUrl  struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"strong_beat_url"`
			TagList           interface{} `json:"tag_list"`
			Title             string      `json:"title"`
			UnshelveCountries interface{} `json:"unshelve_countries"`
			UserCount         int         `json:"user_count"`
			VideoDuration     int         `json:"video_duration"`
		} `json:"music"`
		NicknamePosition    interface{}   `json:"nickname_position"`
		OriginCommentIds    interface{}   `json:"origin_comment_ids"`
		OriginTextExtra     []interface{} `json:"origin_text_extra"`
		OriginalImages      interface{}   `json:"original_images"`
		PackedClips         interface{}   `json:"packed_clips"`
		PhotoSearchEntrance struct {
			EcomType int `json:"ecom_type"`
		} `json:"photo_search_entrance"`
		Position           interface{}   `json:"position"`
		PreviewTitle       string        `json:"preview_title"`
		PreviewVideoStatus int           `json:"preview_video_status"`
		Promotions         []interface{} `json:"promotions"`
		Rate               int           `json:"rate"`
		Region             string        `json:"region"`
		RelationLabels     interface{}   `json:"relation_labels"`
		SearchImpr         struct {
			EntityId   string `json:"entity_id"`
			EntityType string `json:"entity_type"`
		} `json:"search_impr"`
		SeriesPaidInfo struct {
			ItemPrice        int `json:"item_price"`
			SeriesPaidStatus int `json:"series_paid_status"`
		} `json:"series_paid_info"`
		ShareInfo struct {
			ShareDesc     string `json:"share_desc"`
			ShareDescInfo string `json:"share_desc_info"`
			ShareLinkDesc string `json:"share_link_desc"`
			ShareUrl      string `json:"share_url"`
		} `json:"share_info"`
		ShareUrl           string `json:"share_url"`
		ShouldOpenAdReport bool   `json:"should_open_ad_report"`
		ShowFollowButton   struct {
		} `json:"show_follow_button"`
		SocialTagList       interface{} `json:"social_tag_list"`
		StandardBarInfoList interface{} `json:"standard_bar_info_list"`
		Statistics          struct {
			AdmireCount  int    `json:"admire_count"`
			AwemeId      string `json:"aweme_id"`
			CollectCount int    `json:"collect_count"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		Status struct {
			AllowShare        bool   `json:"allow_share"`
			AwemeId           string `json:"aweme_id"`
			InReviewing       bool   `json:"in_reviewing"`
			IsDelete          bool   `json:"is_delete"`
			IsProhibited      bool   `json:"is_prohibited"`
			ListenVideoStatus int    `json:"listen_video_status"`
			PartSee           int    `json:"part_see"`
			PrivateStatus     int    `json:"private_status"`
			ReviewResult      struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
		} `json:"status"`
		TextExtra []struct {
			End         int    `json:"end"`
			HashtagId   string `json:"hashtag_id,omitempty"`
			HashtagName string `json:"hashtag_name,omitempty"`
			IsCommerce  bool   `json:"is_commerce,omitempty"`
			Start       int    `json:"start"`
			Type        int    `json:"type"`
			SecUid      string `json:"sec_uid,omitempty"`
			UserId      string `json:"user_id,omitempty"`
		} `json:"text_extra"`
		UniqidPosition interface{} `json:"uniqid_position"`
		UserDigged     int         `json:"user_digged"`
		Video          struct {
			BigThumbs []struct {
				Duration float64 `json:"duration"`
				Fext     string  `json:"fext"`
				ImgNum   int     `json:"img_num"`
				ImgUrl   string  `json:"img_url"`
				ImgXLen  int     `json:"img_x_len"`
				ImgXSize int     `json:"img_x_size"`
				ImgYLen  int     `json:"img_y_len"`
				ImgYSize int     `json:"img_y_size"`
				Interval int     `json:"interval"`
				Uri      string  `json:"uri"`
			} `json:"big_thumbs"`
			BitRate []struct {
				FPS       int    `json:"FPS"`
				HDRBit    string `json:"HDR_bit"`
				HDRType   string `json:"HDR_type"`
				BitRate   int    `json:"bit_rate"`
				GearName  string `json:"gear_name"`
				IsBytevc1 int    `json:"is_bytevc1"`
				IsH265    int    `json:"is_h265"`
				PlayAddr  struct {
					DataSize int      `json:"data_size"`
					FileCs   string   `json:"file_cs"`
					FileHash string   `json:"file_hash"`
					Height   int      `json:"height"`
					Uri      string   `json:"uri"`
					UrlKey   string   `json:"url_key"`
					UrlList  []string `json:"url_list"`
					Width    int      `json:"width"`
				} `json:"play_addr"`
				QualityType int `json:"quality_type"`
			} `json:"bit_rate"`
			Cover struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover"`
			CoverOriginalScale struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_original_scale"`
			Duration     int `json:"duration"`
			DynamicCover struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"dynamic_cover"`
			Height      int    `json:"height"`
			IsH265      int    `json:"is_h265"`
			IsSourceHDR int    `json:"is_source_HDR"`
			Meta        string `json:"meta"`
			OriginCover struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"origin_cover"`
			PlayAddr struct {
				DataSize int      `json:"data_size"`
				FileCs   string   `json:"file_cs"`
				FileHash string   `json:"file_hash"`
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
			} `json:"play_addr"`
			Ratio string `json:"ratio"`
			Width int    `json:"width"`
		} `json:"video"`
		VideoLabels interface{} `json:"video_labels"`
		VideoTag    []struct {
			Level   int    `json:"level"`
			TagId   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
		} `json:"video_tag"`
		VideoText []interface{} `json:"video_text"`
		WannaTag  struct {
		} `json:"wanna_tag"`
	} `json:"aweme_detail"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	StatusCode int `json:"status_code"`
}

type AnalysisV1Result struct {
	Result AnalysisV1Response // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newAnalysisV1Result(result AnalysisV1Response, body []byte, http gorequest.Response) *AnalysisV1Result {
	return &AnalysisV1Result{Result: result, Body: body, Http: http}
}

// AnalysisV1 抖音解析 https://github.com/iawia002/lux/issues/1184
func (c *Client) AnalysisV1(ctx context.Context, content string) (*AnalysisV1Result, error) {

	// 提取url
	var url string
	if strings.Contains(content, "douyin.com") {
		url = xurls.Relaxed.FindString(content)
	} else if strings.Contains(content, "iesdouyin.com") {
		url = xurls.Relaxed.FindString(content)
	} else {
		return newAnalysisV1Result(AnalysisV1Response{}, nil, gorequest.Response{}), errors.New("url为空")
	}

	// 重定向信息
	request302, err := c.request302(url)
	if err != nil {
		return newAnalysisV1Result(AnalysisV1Response{}, nil, gorequest.Response{}), err
	}

	// 提取编号
	itemIds := regexp.MustCompile(`\d+`).FindStringSubmatch(request302)
	if len(itemIds) < 1 {
		return newAnalysisV1Result(AnalysisV1Response{}, nil, gorequest.Response{}), errors.New("参数错误")
	}

	// 请求
	request, err := c.request(ctx, "https://www.iesdouyin.com/aweme/v1/web/aweme/detail/?aweme_id="+itemIds[0], nil, http.MethodGet)
	if err != nil {
		return newAnalysisV1Result(AnalysisV1Response{}, request.ResponseBody, request), err
	}

	// 定义
	var response AnalysisV1Response
	err = json.Unmarshal(request.ResponseBody, &response)
	return newAnalysisV1Result(response, request.ResponseBody, request), err
}
