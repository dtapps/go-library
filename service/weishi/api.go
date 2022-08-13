package weishi

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/mvdan/xurls"
	"regexp"
	"strings"
)

type AnalysisResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Feeds []struct {
			Id       string `json:"id"`
			Wording  string `json:"wording"`
			Type     int    `json:"type"`
			PosterId string `json:"poster_id"`
			Poster   struct {
				Id                string `json:"id"`
				Type              int    `json:"type"`
				Uid               string `json:"uid"`
				Createtime        int    `json:"createtime"`
				Nick              string `json:"nick"`
				Avatar            string `json:"avatar"`
				Sex               int    `json:"sex"`
				FeedlistTimeId    string `json:"feedlist_time_id"`
				FeedlistHotId     string `json:"feedlist_hot_id"`
				RelatedFeedlistId string `json:"related_feedlist_id"`
				FollowerlistId    string `json:"followerlist_id"`
				InteresterlistId  string `json:"interesterlist_id"`
				ChatlistId        string `json:"chatlist_id"`
				RichFlag          int    `json:"rich_flag"`
				Age               int    `json:"age"`
				Address           string `json:"address"`
				Wealth            struct {
					FlowerNum int `json:"flower_num"`
					Score     int `json:"score"`
				} `json:"wealth"`
				Background        string `json:"background"`
				Status            string `json:"status"`
				FollowStatus      int    `json:"followStatus"`
				ChartScore        int    `json:"chartScore"`
				ChartRank         int    `json:"chartRank"`
				FeedGoldNum       int    `json:"feedGoldNum"`
				AvatarUpdatetime  int    `json:"avatar_updatetime"`
				DescFromOperator  string `json:"desc_from_operator"`
				SyncContent       int    `json:"sync_content"`
				FeedlistPraiseId  string `json:"feedlist_praise_id"`
				Settingmask       int    `json:"settingmask"`
				Originalavatar    string `json:"originalavatar"`
				BlockTime         string `json:"block_time"`
				Grade             int    `json:"grade"`
				Medal             int    `json:"medal"`
				BlockReason       string `json:"block_reason"`
				Qq                int    `json:"qq"`
				RecommendReason   string `json:"recommendReason"`
				LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
				Updateinfo        struct {
					Flag int    `json:"flag"`
					Tip  string `json:"tip"`
					Num  int    `json:"num"`
				} `json:"updateinfo"`
				NickUpdatetime     int64  `json:"nick_updatetime"`
				LastDownloadAvatar string `json:"lastDownloadAvatar"`
				RealName           string `json:"realName"`
				PinyinFirst        string `json:"pinyin_first"`
				CertifDesc         string `json:"certif_desc"`
				PrivateInfo        struct {
					PhoneNum string `json:"phone_num"`
					Name     string `json:"name"`
					IdNum    string `json:"id_num"`
				} `json:"privateInfo"`
				ExternInfo struct {
					MpEx struct {
						DarenPriority        string `json:"daren_priority"`
						LoginResaveOldAvatar string `json:"loginResaveOldAvatar"`
						LoginResaveNewAvatar string `json:"loginResaveNewAvatar"`
						LoginResaveTime      string `json:"loginResaveTime"`
						DarenCompany         string `json:"daren_company"`
						AuditPriority        string `json:"audit_priority"`
						SubPriority          string `json:"sub_priority"`
					} `json:"mpEx"`
					BindAcct  []interface{} `json:"bind_acct"`
					BgPicUrl  string        `json:"bgPicUrl"`
					LevelInfo struct {
						Level           int `json:"level"`
						Score           int `json:"score"`
						PrevUpgradeTime int `json:"prev_upgrade_time"`
					} `json:"level_info"`
					WeishiId            string `json:"weishiId"`
					WeishiidModifyCount string `json:"weishiid_modify_count"`
					WatermarkType       int    `json:"watermark_type"`
					RealNick            string `json:"real_nick"`
					CmtLevel            struct {
						Level           int `json:"level"`
						Cmtscore        int `json:"cmtscore"`
						Dingscore       int `json:"dingscore"`
						PrevUpgradeTime int `json:"prev_upgrade_time"`
					} `json:"cmt_level"`
					FlexibilityFlag int `json:"flexibility_flag"`
					LiveStatus      int `json:"live_status"`
					NowLiveRoomId   int `json:"now_live_room_id"`
					MedalInfo       struct {
						TotalScore int           `json:"total_score"`
						MedalList  []interface{} `json:"medal_list"`
					} `json:"medal_info"`
					H5HasLogin int `json:"h5_has_login"`
				} `json:"extern_info"`
				CertifData struct {
					CertifIcon    string `json:"certif_icon"`
					CertifJumpurl string `json:"certif_jumpurl"`
				} `json:"certifData"`
				IsShowPOI    int `json:"isShowPOI"`
				IsShowGender int `json:"isShowGender"`
				FormatAddr   struct {
					Country  string `json:"country"`
					Province string `json:"province"`
					City     string `json:"city"`
				} `json:"formatAddr"`
				AuthorizeTime int `json:"authorize_time"`
				ActivityInfo  struct {
					InvitePersonid string `json:"invitePersonid"`
				} `json:"activityInfo"`
			} `json:"poster"`
			Video struct {
				FileId       string `json:"file_id"`
				FileSize     int    `json:"file_size"`
				Sha1         string `json:"sha1"`
				PlayIndex    int    `json:"play_index"`
				Duration     int    `json:"duration"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Md5          string `json:"md5"`
				Orientation  int    `json:"orientation"`
				H265Hvc1     int    `json:"h265_hvc1"`
				MaxDb        int    `json:"max_db"`
				VoiceRatio   int    `json:"voice_ratio"`
				Loudnorm     string `json:"loudnorm"`
				MetaLoudnorm struct {
					InputI            string `json:"input_i"`
					InputTp           string `json:"input_tp"`
					InputLra          string `json:"input_lra"`
					InputThresh       string `json:"input_thresh"`
					OutputI           string `json:"output_i"`
					OutputTp          string `json:"output_tp"`
					OutputLra         string `json:"output_lra"`
					OutputThresh      string `json:"output_thresh"`
					NormalizationType string `json:"normalization_type"`
					TargetOffset      string `json:"target_offset"`
					WeishiI           string `json:"weishi_i"`
					WeishiTp          string `json:"weishi_tp"`
					WeishiLra         string `json:"weishi_lra"`
				} `json:"meta_loudnorm"`
			} `json:"video"`
			Images []struct {
				Url          string `json:"url"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Type         int    `json:"type"`
				SpriteWidth  int    `json:"sprite_width"`
				SpriteHeight int    `json:"sprite_height"`
				SpriteSpan   int    `json:"sprite_span"`
			} `json:"images"`
			UgcVideoIds      []interface{} `json:"ugc_video_ids"`
			UgcVideos        []interface{} `json:"ugc_videos"`
			Createtime       int           `json:"createtime"`
			Mask             int           `json:"mask"`
			Score            int           `json:"score"`
			DingCount        int           `json:"ding_count"`
			CommentlistId    string        `json:"commentlist_id"`
			TotalCommentNum  int           `json:"total_comment_num"`
			Comments         []interface{} `json:"comments"`
			MaterialId       string        `json:"material_id"`
			MaterialDesc     string        `json:"material_desc"`
			DingHashId       string        `json:"ding_hash_id"`
			IsDing           int           `json:"is_ding"`
			PlayNum          int           `json:"playNum"`
			CharacterId      string        `json:"character_id"`
			FlowerNum        int           `json:"flower_num"`
			SendFlowerNum    int           `json:"send_flower_num"`
			RichFlag         int           `json:"rich_flag"`
			VideoUrl         string        `json:"video_url"`
			MaterialThumburl string        `json:"material_thumburl"`
			Platform         int           `json:"platform"`
			Reserve          struct {
				Field1 string `json:"2"`
				Field2 string `json:"3"`
				Field3 string `json:"6"`
				Field4 string `json:"36"`
				Field5 string `json:"38"`
				Field6 string `json:"41"`
				Field7 string `json:"47"`
				Field8 string `json:"62"`
			} `json:"reserve"`
			VideoSpecUrls []struct {
				Url           string `json:"url"`
				Size          int    `json:"size"`
				Hardorsoft    int    `json:"hardorsoft"`
				RecommendSpec int    `json:"recommendSpec"`
				HaveWatermark int    `json:"haveWatermark"`
				Width         int    `json:"width"`
				Height        int    `json:"height"`
			} `json:"video_spec_urls"`
			ShareInfo struct {
				JumpUrl string `json:"jump_url"`
				BodyMap []struct {
					Title    string `json:"title"`
					Desc     string `json:"desc"`
					ImageUrl string `json:"image_url"`
					Url      string `json:"url"`
				} `json:"body_map"`
				WxMiniProgram struct {
					WebpageUrl       string `json:"webpageUrl"`
					UserName         string `json:"userName"`
					Path             string `json:"path"`
					HdImageDataURL   string `json:"hdImageDataURL"`
					WithShareTicket  int    `json:"withShareTicket"`
					MiniProgramType  int    `json:"miniProgramType"`
					Appid            string `json:"appid"`
					VideoUserName    string `json:"videoUserName"`
					VideoSource      string `json:"videoSource"`
					VideoCoverWidth  int    `json:"videoCoverWidth"`
					VideoCoverHeight int    `json:"videoCoverHeight"`
					AppThumbUrl      string `json:"appThumbUrl"`
				} `json:"wx_mini_program"`
				SqArkInfo struct {
					ArkData   string `json:"arkData"`
					ShareBody struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"shareBody"`
					CoverProto string `json:"coverProto"`
				} `json:"sq_ark_info"`
				ShareIconUrl   string `json:"share_icon_url"`
				ShareIconTitle string `json:"share_icon_title"`
				BackgroundUrl  string `json:"background_url"`
				ActivityType   int    `json:"activity_type"`
				HaibaoJumpUrl  string `json:"haibao_jump_url"`
				HaibaoBodyMap  struct {
					Field1 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"0"`
					Field2 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"1"`
					Field3 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"2"`
					Field4 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"3"`
					Field5 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"4"`
					Field6 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageUrl string `json:"image_url"`
						Url      string `json:"url"`
					} `json:"5"`
				} `json:"haibao_body_map"`
				BackgroundTitleColor string `json:"background_title_color"`
				HaibaoDesc           string `json:"haibao_desc"`
			} `json:"share_info"`
			FeedGift struct {
				FeedGiftList []interface{} `json:"feedGiftList"`
			} `json:"feedGift"`
			GiftRank []interface{} `json:"giftRank"`
			TopicId  string        `json:"topic_id"`
			Topic    struct {
				Id             string        `json:"id"`
				Name           string        `json:"name"`
				ThumbUrl1      string        `json:"thumbUrl1"`
				ThumbUrl2      string        `json:"thumbUrl2"`
				ThumbUrl3      string        `json:"thumbUrl3"`
				Detail         string        `json:"detail"`
				Createtime     int           `json:"createtime"`
				FeedlistTimeId string        `json:"feedlist_time_id"`
				FeedlistHotId  string        `json:"feedlist_hot_id"`
				MaterialIds    []interface{} `json:"material_ids"`
				Mask           int           `json:"mask"`
				Type           int           `json:"type"`
				Reserve        struct {
				} `json:"reserve"`
				ViewNum    int `json:"view_num"`
				StartTime  int `json:"start_time"`
				EndTime    int `json:"end_time"`
				AppVersion int `json:"appVersion"`
				WorkNum    int `json:"workNum"`
				LikeNum    int `json:"likeNum"`
				Person     struct {
					Id                string `json:"id"`
					Type              int    `json:"type"`
					Uid               string `json:"uid"`
					Createtime        int    `json:"createtime"`
					Nick              string `json:"nick"`
					Avatar            string `json:"avatar"`
					Sex               int    `json:"sex"`
					FeedlistTimeId    string `json:"feedlist_time_id"`
					FeedlistHotId     string `json:"feedlist_hot_id"`
					RelatedFeedlistId string `json:"related_feedlist_id"`
					FollowerlistId    string `json:"followerlist_id"`
					InteresterlistId  string `json:"interesterlist_id"`
					ChatlistId        string `json:"chatlist_id"`
					RichFlag          int    `json:"rich_flag"`
					Age               int    `json:"age"`
					Address           string `json:"address"`
					Wealth            struct {
						FlowerNum int `json:"flower_num"`
						Score     int `json:"score"`
					} `json:"wealth"`
					Background        string `json:"background"`
					Status            string `json:"status"`
					FollowStatus      int    `json:"followStatus"`
					ChartScore        int    `json:"chartScore"`
					ChartRank         int    `json:"chartRank"`
					FeedGoldNum       int    `json:"feedGoldNum"`
					AvatarUpdatetime  int    `json:"avatar_updatetime"`
					DescFromOperator  string `json:"desc_from_operator"`
					SyncContent       int    `json:"sync_content"`
					FeedlistPraiseId  string `json:"feedlist_praise_id"`
					Settingmask       int    `json:"settingmask"`
					Originalavatar    string `json:"originalavatar"`
					BlockTime         string `json:"block_time"`
					Grade             int    `json:"grade"`
					Medal             int    `json:"medal"`
					BlockReason       string `json:"block_reason"`
					Qq                int    `json:"qq"`
					RecommendReason   string `json:"recommendReason"`
					LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
					Updateinfo        struct {
						Flag int    `json:"flag"`
						Tip  string `json:"tip"`
						Num  int    `json:"num"`
					} `json:"updateinfo"`
					NickUpdatetime     int    `json:"nick_updatetime"`
					LastDownloadAvatar string `json:"lastDownloadAvatar"`
					RealName           string `json:"realName"`
					PinyinFirst        string `json:"pinyin_first"`
					CertifDesc         string `json:"certif_desc"`
					PrivateInfo        struct {
						PhoneNum string `json:"phone_num"`
						Name     string `json:"name"`
						IdNum    string `json:"id_num"`
					} `json:"privateInfo"`
					ExternInfo struct {
						MpEx struct {
						} `json:"mpEx"`
						BindAcct  []interface{} `json:"bind_acct"`
						BgPicUrl  string        `json:"bgPicUrl"`
						LevelInfo struct {
							Level           int `json:"level"`
							Score           int `json:"score"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"level_info"`
						WeishiId            string `json:"weishiId"`
						WeishiidModifyCount string `json:"weishiid_modify_count"`
						WatermarkType       int    `json:"watermark_type"`
						RealNick            string `json:"real_nick"`
						CmtLevel            struct {
							Level           int `json:"level"`
							Cmtscore        int `json:"cmtscore"`
							Dingscore       int `json:"dingscore"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"cmt_level"`
						FlexibilityFlag int `json:"flexibility_flag"`
						LiveStatus      int `json:"live_status"`
						NowLiveRoomId   int `json:"now_live_room_id"`
						MedalInfo       struct {
							TotalScore int           `json:"total_score"`
							MedalList  []interface{} `json:"medal_list"`
						} `json:"medal_info"`
						H5HasLogin int `json:"h5_has_login"`
					} `json:"extern_info"`
					CertifData struct {
						CertifIcon    string `json:"certif_icon"`
						CertifJumpurl string `json:"certif_jumpurl"`
					} `json:"certifData"`
					IsShowPOI    int `json:"isShowPOI"`
					IsShowGender int `json:"isShowGender"`
					FormatAddr   struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"formatAddr"`
					AuthorizeTime int `json:"authorize_time"`
					ActivityInfo  struct {
						InvitePersonid string `json:"invitePersonid"`
					} `json:"activityInfo"`
				} `json:"person"`
				FeedId            string `json:"feed_id"`
				PendantMaterialId string `json:"pendant_material_id"`
				MusicMaterialId   string `json:"music_material_id"`
				MusicInfo         struct {
					Id              string        `json:"id"`
					Name            string        `json:"name"`
					Desc            string        `json:"desc"`
					Type            string        `json:"type"`
					ThumbUrl        string        `json:"thumbUrl"`
					Version         int           `json:"version"`
					MiniSptVersion  int           `json:"miniSptVersion"`
					PackageUrl      string        `json:"packageUrl"`
					FeedlistTimeId  string        `json:"feedlist_time_id"`
					FeedlistHotId   string        `json:"feedlist_hot_id"`
					TopicIds        []interface{} `json:"topic_ids"`
					Mask            int           `json:"mask"`
					ShortName       string        `json:"shortName"`
					RichFlag        int           `json:"rich_flag"`
					EffectId        string        `json:"effectId"`
					Rgbcolor        string        `json:"rgbcolor"`
					IsCollected     int           `json:"isCollected"`
					BubbleStartTime int           `json:"bubbleStartTime"`
					BubbleEndTime   int           `json:"bubbleEndTime"`
					CollectTime     int           `json:"collectTime"`
					SdkInfo         struct {
						IsSdk            int `json:"isSdk"`
						SdkMinVersion    int `json:"sdkMinVersion"`
						SdkMaxVersion    int `json:"sdkMaxVersion"`
						SdkMinSptVersion int `json:"sdkMinSptVersion"`
					} `json:"sdkInfo"`
					BigThumbUrl string        `json:"bigThumbUrl"`
					Priority    int           `json:"priority"`
					MusicIDs    []interface{} `json:"musicIDs"`
					Platform    string        `json:"platform"`
					Reserve     struct {
					} `json:"reserve"`
					Category       string        `json:"category"`
					ShootingTips   string        `json:"shooting_tips"`
					VecSubcategory []interface{} `json:"vec_subcategory"`
				} `json:"music_info"`
				PendantMaterialIdIos string `json:"pendant_material_id_ios"`
				MediaMaterialUrl     string `json:"media_material_url"`
				BubbleStartTime      int    `json:"bubble_start_time"`
				BubbleEndTime        int    `json:"bubble_end_time"`
				BubbleCopywrite      string `json:"bubble_copywrite"`
				Rgbcolor             int    `json:"rgbcolor"`
				Lplaynum             int    `json:"lplaynum"`
				QqMusicInfo          struct {
					AlbumInfo struct {
						UiId    int    `json:"uiId"`
						StrMid  string `json:"strMid"`
						StrName string `json:"strName"`
						StrPic  string `json:"strPic"`
					} `json:"albumInfo"`
					SingerInfo struct {
						UiId    int    `json:"uiId"`
						StrMid  string `json:"strMid"`
						StrName string `json:"strName"`
						StrPic  string `json:"strPic"`
					} `json:"singerInfo"`
					SongInfo struct {
						UiId               int    `json:"uiId"`
						StrMid             string `json:"strMid"`
						StrName            string `json:"strName"`
						StrGenre           string `json:"strGenre"`
						IIsOnly            int    `json:"iIsOnly"`
						StrLanguage        string `json:"strLanguage"`
						IPlayable          int    `json:"iPlayable"`
						ITrySize           int    `json:"iTrySize"`
						ITryBegin          int    `json:"iTryBegin"`
						ITryEnd            int    `json:"iTryEnd"`
						IPlayTime          int    `json:"iPlayTime"`
						StrH5Url           string `json:"strH5Url"`
						StrPlayUrl         string `json:"strPlayUrl"`
						StrPlayUrlStandard string `json:"strPlayUrlStandard"`
						StrPlayUrlHq       string `json:"strPlayUrlHq"`
						StrPlayUrlSq       string `json:"strPlayUrlSq"`
						ISize              int    `json:"iSize"`
						ISizeStandard      int    `json:"iSizeStandard"`
						ISizeHq            int    `json:"iSizeHq"`
						ISizeSq            int    `json:"iSizeSq"`
						Copyright          int    `json:"copyright"`
						ISource            int    `json:"iSource"`
					} `json:"songInfo"`
					LyricInfo struct {
						UiSongId   int    `json:"uiSongId"`
						StrSongMid string `json:"strSongMid"`
						StrFormat  string `json:"strFormat"`
						StrLyric   string `json:"strLyric"`
					} `json:"lyricInfo"`
					ConfInfo struct {
						IType               int    `json:"iType"`
						IStartPos           int    `json:"iStartPos"`
						StrLabel            string `json:"strLabel"`
						IsCollected         int    `json:"isCollected"`
						CollectTime         int    `json:"collectTime"`
						Exclusive           int    `json:"exclusive"`
						FollowFeed          string `json:"followFeed"`
						UseCount            int    `json:"useCount"`
						TogetherFeed        string `json:"togetherFeed"`
						TogetherType        int    `json:"togetherType"`
						FeedUseType         int    `json:"feedUseType"`
						DefaultFeedPosition int    `json:"defaultFeedPosition"`
						DefaultTogetherFeed int    `json:"defaultTogetherFeed"`
						BubbleStartTime     int    `json:"bubbleStartTime"`
						BubbleEndTime       int    `json:"bubbleEndTime"`
					} `json:"confInfo"`
					SubtitleInfo struct {
						UiSongId   int    `json:"uiSongId"`
						StrSongMid string `json:"strSongMid"`
						StrFormat  string `json:"strFormat"`
						StrLyric   string `json:"strLyric"`
					} `json:"subtitleInfo"`
					Foreignlyric struct {
						UiSongId   int    `json:"uiSongId"`
						StrSongMid string `json:"strSongMid"`
						StrFormat  string `json:"strFormat"`
						StrLyric   string `json:"strLyric"`
					} `json:"foreignlyric"`
					RecommendInfo struct {
						TraceStr string `json:"traceStr"`
					} `json:"recommendInfo"`
					UnplayableInfo struct {
						UnplayableCode int    `json:"unplayableCode"`
						UnplayableMsg  string `json:"unplayableMsg"`
					} `json:"unplayableInfo"`
				} `json:"qqMusicInfo"`
			} `json:"topic"`
			FlowerNumDb int           `json:"flowerNumDb"`
			FlowerRank  []interface{} `json:"flowerRank"`
			FeedDesc    string        `json:"feed_desc"`
			DescMask    int           `json:"desc_mask"`
			ShieldId    string        `json:"shieldId"`
			VideoCover  struct {
				StaticCover struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"static_cover"`
				DynamicCover struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"dynamic_cover"`
				CoverTime     int           `json:"cover_time"`
				VMetaEffect   []interface{} `json:"vMetaEffect"`
				AnimatedCover struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"animated_cover"`
				SmallAnimatedCover struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"small_animated_cover"`
				CoverWidth      int `json:"cover_width"`
				CoverHeight     int `json:"cover_height"`
				AnimatedCover5F struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"animated_cover_5f"`
				SmallAnimatedCover5F struct {
					Url          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
				} `json:"small_animated_cover_5f"`
			} `json:"video_cover"`
			GeoInfo struct {
				Country   string `json:"country"`
				Province  string `json:"province"`
				City      string `json:"city"`
				Latitude  int    `json:"latitude"`
				Longitude int    `json:"longitude"`
				Altitude  int    `json:"altitude"`
				District  string `json:"district"`
				Name      string `json:"name"`
				Distance  int    `json:"distance"`
				PolyGeoID string `json:"polyGeoID"`
			} `json:"geoInfo"`
			MusicId  string `json:"music_id"`
			VideoBgm struct {
				MusicId    string `json:"music_id"`
				MusicCover string `json:"music_cover"`
				MusicDesc  string `json:"music_desc"`
				Duration   int    `json:"duration"`
				Size       int    `json:"size"`
				FeedId     string `json:"feed_id"`
				MpEx       struct {
					Mp3Id       string `json:"mp3_id"`
					TogetherBgm string `json:"together_bgm"`
				} `json:"mpEx"`
			} `json:"video_bgm"`
			RecgBgm struct {
				MusicId    string `json:"music_id"`
				MusicCover string `json:"music_cover"`
				MusicDesc  string `json:"music_desc"`
				Duration   int    `json:"duration"`
				Size       int    `json:"size"`
				FeedId     string `json:"feed_id"`
				MpEx       struct {
				} `json:"mpEx"`
			} `json:"recg_bgm"`
			EnableRealRcmd      int    `json:"enable_real_rcmd"`
			FeedDescWithat      string `json:"feed_desc_withat"`
			FeedRecommendReason string `json:"feed_recommend_reason"`
			Interaction         struct {
				MpEx struct {
				} `json:"mpEx"`
				Type     int           `json:"type"`
				PersonId string        `json:"person_id"`
				FeedId   string        `json:"feed_id"`
				Score    int           `json:"score"`
				Buttons  []interface{} `json:"buttons"`
			} `json:"interaction"`
			Ornament struct {
				MpEx struct {
				} `json:"mpEx"`
				FilterId    string `json:"filter_id"`
				FilterName  string `json:"filter_name"`
				PendantId   string `json:"pendant_id"`
				PendantCate string `json:"pendant_cate"`
			} `json:"ornament"`
			VideoOrnaments []interface{} `json:"video_ornaments"`
			HaveText       int           `json:"have_text"`
			ExternInfo     struct {
				MpEx struct {
					FeedSource       string `json:"feed_source"`
					FeedCover        string `json:"feed_cover"`
					ShowWxShareIcon  string `json:"show_wx_share_icon"`
					ActivityInfo     string `json:"activity_info"`
					ReportJson       string `json:"report_json"`
					PrepareRecommend string `json:"prepare_recommend"`
					SecurityCheck    string `json:"security_check"`
				} `json:"mpEx"`
				VisibleType       int `json:"visible_type"`
				ActivityShareInfo struct {
				} `json:"activity_share_info"`
				ActTogetherInfo struct {
					ExtInfo struct {
					} `json:"extInfo"`
					AllowTogether int    `json:"allowTogether"`
					TogetherType  int    `json:"togetherType"`
					PolyId        string `json:"polyId"`
					LastFeedId    string `json:"lastFeedId"`
					SrcFeedId     string `json:"srcFeedId"`
					TogetherCount int    `json:"togetherCount"`
					TogetherSpec  struct {
						Field1 int `json:"1"`
						Field2 int `json:"2"`
						Field3 int `json:"3"`
					} `json:"togetherSpec"`
					TogetherJump string `json:"togetherJump"`
					LastPersonId string `json:"lastPersonId"`
					GhostFeed    int    `json:"ghostFeed"`
					SrcBgmId     string `json:"srcBgmId"`
					FeedPosition struct {
					} `json:"feedPosition"`
					DefaultFeedPosition int `json:"defaultFeedPosition"`
					DefaultTogetherFeed int `json:"defaultTogetherFeed"`
				} `json:"actTogetherInfo"`
				DangerMarker int    `json:"danger_marker"`
				Rowkey       string `json:"rowkey"`
				FeedAdsInfo  struct {
					Icon        string `json:"icon"`
					IconWidth   int    `json:"icon_width"`
					IconHeight  int    `json:"icon_height"`
					Schema      string `json:"schema"`
					CommentType int    `json:"comment_type"`
					AdsGoal     int    `json:"ads_goal"`
					AdsType     int    `json:"ads_type"`
					QbossReport struct {
						Qbossid   int    `json:"qbossid"`
						TaskId    int    `json:"task_id"`
						TraceInfo string `json:"trace_info"`
						Position  int    `json:"position"`
					} `json:"qboss_report"`
					Extra struct {
					} `json:"extra"`
					AdsName string `json:"ads_name"`
				} `json:"feedAdsInfo"`
				ClarifyScore int `json:"clarifyScore"`
				ConcernHint  struct {
					EnableHint  int    `json:"enableHint"`
					BeginSecond int    `json:"beginSecond"`
					EndSecond   int    `json:"endSecond"`
					Thumbnail   string `json:"thumbnail"`
					Hint        string `json:"hint"`
				} `json:"concernHint"`
				RewardNum       int    `json:"reward_num"`
				SubtitleFlag    int    `json:"subtitle_flag"`
				UploadLyricFlag int    `json:"upload_lyric_flag"`
				FriendLikeNum   int    `json:"friend_like_num"`
				SafeVisibleType int    `json:"safe_visible_type"`
				RecommendReason string `json:"recommend_reason"`
				NowLiveRoomId   int    `json:"now_live_room_id"`
				ExtraMask       int    `json:"extra_mask"`
				RecommendMore   int    `json:"recommend_more"`
				QuestionList    struct {
					MaxIndex  int           `json:"max_index"`
					Questions []interface{} `json:"questions"`
				} `json:"question_list"`
				InteractConf struct {
					StickerData struct {
						TimeLines []interface{} `json:"time_lines"`
					} `json:"sticker_data"`
					MagicData struct {
						VideoWidth  int           `json:"video_width"`
						VideoHeight int           `json:"video_height"`
						EventList   []interface{} `json:"event_list"`
					} `json:"magic_data"`
					Token           string `json:"token"`
					TemplateTypes   string `json:"template_types"`
					VideoShareCover struct {
						Url          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
					} `json:"video_share_cover"`
					QzoneSkin struct {
						SkinId        string `json:"skin_id"`
						SkinType      int    `json:"skin_type"`
						Picurl        string `json:"picurl"`
						Bgcolor       string `json:"bgcolor"`
						GradientBegin string `json:"gradient_begin"`
						GradientEnd   string `json:"gradient_end"`
						PicurlAnd     string `json:"picurl_and"`
					} `json:"qzone_skin"`
					TemplateName      string `json:"template_name"`
					TemplateId        string `json:"template_id"`
					TemplateBusiness  string `json:"template_business"`
					TemplateTitleSkin struct {
						Url    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"template_title_skin"`
				} `json:"interact_conf"`
				InteractUgcData struct {
					UgcContent       string `json:"ugc_content"`
					HasVote          int    `json:"has_vote"`
					PersonOfficeName string `json:"person_office_name"`
				} `json:"interact_ugc_data"`
				SrcFeedId string        `json:"src_feed_id"`
				VKeyFrame []interface{} `json:"vKeyFrame"`
			} `json:"extern_info"`
			StarRanking struct {
				InRanking   int    `json:"in_ranking"`
				InActivity  int    `json:"in_activity"`
				CallBangImg string `json:"call_bang_img"`
				RankingTips string `json:"ranking_tips"`
			} `json:"starRanking"`
			Tags         []interface{} `json:"tags"`
			CollectionId string        `json:"collectionId"`
			Collection   struct {
				Cid       string `json:"cid"`
				Name      string `json:"name"`
				Cover     string `json:"cover"`
				Desc      string `json:"desc"`
				FeedNum   int    `json:"feedNum"`
				PlayNum   int    `json:"playNum"`
				ShareInfo struct {
					JumpUrl string `json:"jump_url"`
					BodyMap struct {
					} `json:"body_map"`
					WxMiniProgram struct {
						WebpageUrl       string `json:"webpageUrl"`
						UserName         string `json:"userName"`
						Path             string `json:"path"`
						HdImageDataURL   string `json:"hdImageDataURL"`
						WithShareTicket  int    `json:"withShareTicket"`
						MiniProgramType  int    `json:"miniProgramType"`
						Appid            string `json:"appid"`
						VideoUserName    string `json:"videoUserName"`
						VideoSource      string `json:"videoSource"`
						VideoCoverWidth  int    `json:"videoCoverWidth"`
						VideoCoverHeight int    `json:"videoCoverHeight"`
						AppThumbUrl      string `json:"appThumbUrl"`
					} `json:"wx_mini_program"`
					SqArkInfo struct {
						ArkData   string `json:"arkData"`
						ShareBody struct {
							Title    string `json:"title"`
							Desc     string `json:"desc"`
							ImageUrl string `json:"image_url"`
							Url      string `json:"url"`
						} `json:"shareBody"`
						CoverProto string `json:"coverProto"`
					} `json:"sq_ark_info"`
					ShareIconUrl   string `json:"share_icon_url"`
					ShareIconTitle string `json:"share_icon_title"`
					BackgroundUrl  string `json:"background_url"`
					ActivityType   int    `json:"activity_type"`
					HaibaoJumpUrl  string `json:"haibao_jump_url"`
					HaibaoBodyMap  struct {
					} `json:"haibao_body_map"`
					BackgroundTitleColor string `json:"background_title_color"`
					HaibaoDesc           string `json:"haibao_desc"`
				} `json:"shareInfo"`
				AttachInfo string `json:"attach_info"`
				Poster     struct {
					Id                string `json:"id"`
					Type              int    `json:"type"`
					Uid               string `json:"uid"`
					Createtime        int    `json:"createtime"`
					Nick              string `json:"nick"`
					Avatar            string `json:"avatar"`
					Sex               int    `json:"sex"`
					FeedlistTimeId    string `json:"feedlist_time_id"`
					FeedlistHotId     string `json:"feedlist_hot_id"`
					RelatedFeedlistId string `json:"related_feedlist_id"`
					FollowerlistId    string `json:"followerlist_id"`
					InteresterlistId  string `json:"interesterlist_id"`
					ChatlistId        string `json:"chatlist_id"`
					RichFlag          int    `json:"rich_flag"`
					Age               int    `json:"age"`
					Address           string `json:"address"`
					Wealth            struct {
						FlowerNum int `json:"flower_num"`
						Score     int `json:"score"`
					} `json:"wealth"`
					Background        string `json:"background"`
					Status            string `json:"status"`
					FollowStatus      int    `json:"followStatus"`
					ChartScore        int    `json:"chartScore"`
					ChartRank         int    `json:"chartRank"`
					FeedGoldNum       int    `json:"feedGoldNum"`
					AvatarUpdatetime  int    `json:"avatar_updatetime"`
					DescFromOperator  string `json:"desc_from_operator"`
					SyncContent       int    `json:"sync_content"`
					FeedlistPraiseId  string `json:"feedlist_praise_id"`
					Settingmask       int    `json:"settingmask"`
					Originalavatar    string `json:"originalavatar"`
					BlockTime         string `json:"block_time"`
					Grade             int    `json:"grade"`
					Medal             int    `json:"medal"`
					BlockReason       string `json:"block_reason"`
					Qq                int    `json:"qq"`
					RecommendReason   string `json:"recommendReason"`
					LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
					Updateinfo        struct {
						Flag int    `json:"flag"`
						Tip  string `json:"tip"`
						Num  int    `json:"num"`
					} `json:"updateinfo"`
					NickUpdatetime     int    `json:"nick_updatetime"`
					LastDownloadAvatar string `json:"lastDownloadAvatar"`
					RealName           string `json:"realName"`
					PinyinFirst        string `json:"pinyin_first"`
					CertifDesc         string `json:"certif_desc"`
					PrivateInfo        struct {
						PhoneNum string `json:"phone_num"`
						Name     string `json:"name"`
						IdNum    string `json:"id_num"`
					} `json:"privateInfo"`
					ExternInfo struct {
						MpEx struct {
						} `json:"mpEx"`
						BindAcct  []interface{} `json:"bind_acct"`
						BgPicUrl  string        `json:"bgPicUrl"`
						LevelInfo struct {
							Level           int `json:"level"`
							Score           int `json:"score"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"level_info"`
						WeishiId            string `json:"weishiId"`
						WeishiidModifyCount string `json:"weishiid_modify_count"`
						WatermarkType       int    `json:"watermark_type"`
						RealNick            string `json:"real_nick"`
						CmtLevel            struct {
							Level           int `json:"level"`
							Cmtscore        int `json:"cmtscore"`
							Dingscore       int `json:"dingscore"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"cmt_level"`
						FlexibilityFlag int `json:"flexibility_flag"`
						LiveStatus      int `json:"live_status"`
						NowLiveRoomId   int `json:"now_live_room_id"`
						MedalInfo       struct {
							TotalScore int           `json:"total_score"`
							MedalList  []interface{} `json:"medal_list"`
						} `json:"medal_info"`
						H5HasLogin int `json:"h5_has_login"`
					} `json:"extern_info"`
					CertifData struct {
						CertifIcon    string `json:"certif_icon"`
						CertifJumpurl string `json:"certif_jumpurl"`
					} `json:"certifData"`
					IsShowPOI    int `json:"isShowPOI"`
					IsShowGender int `json:"isShowGender"`
					FormatAddr   struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"formatAddr"`
					AuthorizeTime int `json:"authorize_time"`
					ActivityInfo  struct {
						InvitePersonid string `json:"invitePersonid"`
					} `json:"activityInfo"`
				} `json:"poster"`
				UpdateTime    int `json:"updateTime"`
				UpdateFeedNum int `json:"updateFeedNum"`
				IsFollowed    int `json:"isFollowed"`
				LikeNum       int `json:"likeNum"`
			} `json:"collection"`
			MusicBeginTime int `json:"music_begin_time"`
			MusicEndTime   int `json:"music_end_time"`
			MusicInfo      struct {
				AlbumInfo struct {
					UiId    int    `json:"uiId"`
					StrMid  string `json:"strMid"`
					StrName string `json:"strName"`
					StrPic  string `json:"strPic"`
				} `json:"albumInfo"`
				SingerInfo struct {
					UiId    int    `json:"uiId"`
					StrMid  string `json:"strMid"`
					StrName string `json:"strName"`
					StrPic  string `json:"strPic"`
				} `json:"singerInfo"`
				SongInfo struct {
					UiId               int    `json:"uiId"`
					StrMid             string `json:"strMid"`
					StrName            string `json:"strName"`
					StrGenre           string `json:"strGenre"`
					IIsOnly            int    `json:"iIsOnly"`
					StrLanguage        string `json:"strLanguage"`
					IPlayable          int    `json:"iPlayable"`
					ITrySize           int    `json:"iTrySize"`
					ITryBegin          int    `json:"iTryBegin"`
					ITryEnd            int    `json:"iTryEnd"`
					IPlayTime          int    `json:"iPlayTime"`
					StrH5Url           string `json:"strH5Url"`
					StrPlayUrl         string `json:"strPlayUrl"`
					StrPlayUrlStandard string `json:"strPlayUrlStandard"`
					StrPlayUrlHq       string `json:"strPlayUrlHq"`
					StrPlayUrlSq       string `json:"strPlayUrlSq"`
					ISize              int    `json:"iSize"`
					ISizeStandard      int    `json:"iSizeStandard"`
					ISizeHq            int    `json:"iSizeHq"`
					ISizeSq            int    `json:"iSizeSq"`
					Copyright          int    `json:"copyright"`
					ISource            int    `json:"iSource"`
				} `json:"songInfo"`
				LyricInfo struct {
					UiSongId   int    `json:"uiSongId"`
					StrSongMid string `json:"strSongMid"`
					StrFormat  string `json:"strFormat"`
					StrLyric   string `json:"strLyric"`
				} `json:"lyricInfo"`
				ConfInfo struct {
					IType               int    `json:"iType"`
					IStartPos           int    `json:"iStartPos"`
					StrLabel            string `json:"strLabel"`
					IsCollected         int    `json:"isCollected"`
					CollectTime         int    `json:"collectTime"`
					Exclusive           int    `json:"exclusive"`
					FollowFeed          string `json:"followFeed"`
					UseCount            int    `json:"useCount"`
					TogetherFeed        string `json:"togetherFeed"`
					TogetherType        int    `json:"togetherType"`
					FeedUseType         int    `json:"feedUseType"`
					DefaultFeedPosition int    `json:"defaultFeedPosition"`
					DefaultTogetherFeed int    `json:"defaultTogetherFeed"`
					BubbleStartTime     int    `json:"bubbleStartTime"`
					BubbleEndTime       int    `json:"bubbleEndTime"`
				} `json:"confInfo"`
				SubtitleInfo struct {
					UiSongId   int    `json:"uiSongId"`
					StrSongMid string `json:"strSongMid"`
					StrFormat  string `json:"strFormat"`
					StrLyric   string `json:"strLyric"`
				} `json:"subtitleInfo"`
				Foreignlyric struct {
					UiSongId   int    `json:"uiSongId"`
					StrSongMid string `json:"strSongMid"`
					StrFormat  string `json:"strFormat"`
					StrLyric   string `json:"strLyric"`
				} `json:"foreignlyric"`
				RecommendInfo struct {
					TraceStr string `json:"traceStr"`
				} `json:"recommendInfo"`
				UnplayableInfo struct {
					UnplayableCode int    `json:"unplayableCode"`
					UnplayableMsg  string `json:"unplayableMsg"`
				} `json:"unplayableInfo"`
			} `json:"music_info"`
			Header struct {
				Active  int    `json:"active"`
				Type    int    `json:"type"`
				Title   string `json:"title"`
				Jumpurl string `json:"jumpurl"`
			} `json:"header"`
			RewardInfo struct {
				InRewarding int `json:"in_rewarding"`
			} `json:"rewardInfo"`
			NearbyfeedCoverUrls    []interface{} `json:"nearbyfeed_cover_urls"`
			FingerprintCheckStatus int           `json:"fingerprint_check_status"`
		} `json:"feeds"`
		Isdeleted      int           `json:"isdeleted"`
		Recommendfeeds []interface{} `json:"recommendfeeds"`
		Errmsg         string        `json:"errmsg"`
		Idc            string        `json:"_idc"`
	} `json:"data"`
}

type AnalysisResult struct {
	Result AnalysisResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newAnalysisResult(result AnalysisResponse, body []byte, http gorequest.Response, err error) *AnalysisResult {
	return &AnalysisResult{Result: result, Body: body, Http: http, Err: err}
}

// Analysis 微视解析
func (c *Client) Analysis(ctx context.Context, content string) *AnalysisResult {

	// 提取url
	var url string
	if strings.Contains(content, "h5.weishi") {
		url = xurls.Relaxed.FindString(content)
	} else if strings.Contains(content, "isee.weishi") {
		url = xurls.Relaxed.FindString(content)
	} else {
		return newAnalysisResult(AnalysisResponse{}, nil, gorequest.Response{}, errors.New("url为空"))
	}

	// 内容匹配
	var feedid string
	if strings.Contains(url, "h5.weishi") {
		// 重定向信息
		request302, err := c.request302(url)
		if err != nil {
			return newAnalysisResult(AnalysisResponse{}, nil, gorequest.Response{}, err)
		}

		feedid = strings.Split(request302, "/")[3] ///share/video/6734643996347485448/?region=CN&mid=6734637731277851404&u_code=0&titleType=title&utm_source=copy_link&utm_campaign=client_share&utm_medium=android&app=aweme

	} else if strings.Contains(url, "isee.weishi") {
		// 自带参数
		feedid = regexp.MustCompile("id=(.*?)&").FindStringSubmatch(url)[1]
	}

	request, err := c.request(ctx, "https://h5.qzone.qq.com/webapp/json/weishi/WSH5GetPlayPage?feedid="+feedid)

	// 定义
	var response AnalysisResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newAnalysisResult(response, request.ResponseBody, request, err)
}
