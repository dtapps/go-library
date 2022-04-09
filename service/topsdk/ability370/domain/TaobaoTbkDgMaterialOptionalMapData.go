package domain

type TaobaoTbkDgMaterialOptionalMapData struct {
	/*
	   优惠券信息-优惠券开始时间     */
	CouponStartTime *string `json:"coupon_start_time,omitempty" `

	/*
	   优惠券信息-优惠券结束时间     */
	CouponEndTime *string `json:"coupon_end_time,omitempty" `

	/*
	   商品信息-定向计划信息     */
	InfoDxjh *string `json:"info_dxjh,omitempty" `

	/*
	   商品信息-淘客30天推广量     */
	TkTotalSales *string `json:"tk_total_sales,omitempty" `

	/*
	   商品信息-月支出佣金(该字段废弃，请勿再用)     */
	TkTotalCommi *string `json:"tk_total_commi,omitempty" `

	/*
	   优惠券信息-优惠券id     */
	CouponId *string `json:"coupon_id,omitempty" `

	/*
	   商品信息-宝贝id(该字段废弃，请勿再用)     */
	NumIid *int64 `json:"num_iid,omitempty" `

	/*
	   商品信息-商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品信息-商品主图     */
	PictUrl *string `json:"pict_url,omitempty" `

	/*
	   商品信息-商品小图列表     */
	SmallImages *[]string `json:"small_images,omitempty" `

	/*
	   商品信息-商品一口价格     */
	ReservePrice *string `json:"reserve_price,omitempty" `

	/*
	   折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价     */
	ZkFinalPrice *string `json:"zk_final_price,omitempty" `

	/*
	   店铺信息-卖家类型。0表示集市，1表示天猫     */
	UserType *int32 `json:"user_type,omitempty" `

	/*
	   商品信息-宝贝所在地     */
	Provcity *string `json:"provcity,omitempty" `

	/*
	   链接-宝贝地址     */
	ItemUrl *string `json:"item_url,omitempty" `

	/*
	   商品信息-是否包含营销计划     */
	IncludeMkt *string `json:"include_mkt,omitempty" `

	/*
	   商品信息-是否包含定向计划     */
	IncludeDxjh *string `json:"include_dxjh,omitempty" `

	/*
	   商品信息-佣金比率。1550表示15.5%     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   商品信息-30天销量（饿了么卡券信息-总销量）     */
	Volume *int32 `json:"volume,omitempty" `

	/*
	   店铺信息-卖家id     */
	SellerId *int64 `json:"seller_id,omitempty" `

	/*
	   店铺信息-店铺名称     */
	ShopTitle *string `json:"shop_title,omitempty" `

	/*
	   优惠券信息-优惠券总量     */
	CouponTotalCount *int32 `json:"coupon_total_count,omitempty" `

	/*
	   优惠券信息-优惠券剩余量     */
	CouponRemainCount *int32 `json:"coupon_remain_count,omitempty" `

	/*
	   优惠券信息-优惠券满减信息     */
	CouponInfo *string `json:"coupon_info,omitempty" `

	/*
	   商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划     */
	CommissionType *string `json:"commission_type,omitempty" `

	/*
	   店铺信息-店铺dsr评分     */
	ShopDsr *int64 `json:"shop_dsr,omitempty" `

	/*
	   链接-宝贝+券二合一页面链接     */
	CouponShareUrl *string `json:"coupon_share_url,omitempty" `

	/*
	   链接-宝贝推广链接     */
	Url *string `json:"url,omitempty" `

	/*
	   商品信息-一级类目名称     */
	LevelOneCategoryName *string `json:"level_one_category_name,omitempty" `

	/*
	   商品信息-一级类目ID     */
	LevelOneCategoryId *int64 `json:"level_one_category_id,omitempty" `

	/*
	   商品信息-叶子类目名称     */
	CategoryName *string `json:"category_name,omitempty" `

	/*
	   商品信息-叶子类目id     */
	CategoryId *int64 `json:"category_id,omitempty" `

	/*
	   商品信息-商品短标题     */
	ShortTitle *string `json:"short_title,omitempty" `

	/*
	   商品信息-商品白底图     */
	WhiteImage *string `json:"white_image,omitempty" `

	/*
	   拼团专用-拼团结束时间     */
	Oetime *string `json:"oetime,omitempty" `

	/*
	   拼团专用-拼团开始时间     */
	Ostime *string `json:"ostime,omitempty" `

	/*
	   拼团专用-拼团几人团     */
	JddNum *int64 `json:"jdd_num,omitempty" `

	/*
	   拼团专用-拼团拼成价，单位元     */
	JddPrice *string `json:"jdd_price,omitempty" `

	/*
	   预售专用-预售数量     */
	UvSumPreSale *int64 `json:"uv_sum_pre_sale,omitempty" `

	/*
	   链接-物料块id(测试中请勿使用)     */
	XId *string `json:"x_id,omitempty" `

	/*
	   优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元     */
	CouponStartFee *string `json:"coupon_start_fee,omitempty" `

	/*
	   优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用     */
	CouponAmount *string `json:"coupon_amount,omitempty" `

	/*
	   商品信息-宝贝描述(推荐理由)     */
	ItemDescription *string `json:"item_description,omitempty" `

	/*
	   店铺信息-卖家昵称     */
	Nick *string `json:"nick,omitempty" `

	/*
	   拼团专用-拼团一人价（原价)，单位元     */
	OrigPrice *string `json:"orig_price,omitempty" `

	/*
	   拼团专用-拼团库存数量     */
	TotalStock *int32 `json:"total_stock,omitempty" `

	/*
	   拼团专用-拼团已售数量     */
	SellNum *int32 `json:"sell_num,omitempty" `

	/*
	   拼团专用-拼团剩余库存     */
	Stock *int32 `json:"stock,omitempty" `

	/*
	   营销-天猫营销玩法     */
	TmallPlayActivityInfo *string `json:"tmall_play_activity_info,omitempty" `

	/*
	   商品信息-宝贝id     */
	ItemId *int64 `json:"item_id,omitempty" `

	/*
	   商品邮费     */
	RealPostFee *string `json:"real_post_fee,omitempty" `

	/*
	   锁住的佣金率     */
	LockRate *string `json:"lock_rate,omitempty" `

	/*
	   锁佣结束时间     */
	LockRateEndTime *int64 `json:"lock_rate_end_time,omitempty" `

	/*
	   锁佣开始时间     */
	LockRateStartTime *int64 `json:"lock_rate_start_time,omitempty" `

	/*
	   预售商品-优惠     */
	PresaleDiscountFeeText *string `json:"presale_discount_fee_text,omitempty" `

	/*
	   预售商品-付尾款结束时间（毫秒）     */
	PresaleTailEndTime *int64 `json:"presale_tail_end_time,omitempty" `

	/*
	   预售商品-付尾款开始时间（毫秒）     */
	PresaleTailStartTime *int64 `json:"presale_tail_start_time,omitempty" `

	/*
	   预售商品-付定金结束时间（毫秒）     */
	PresaleEndTime *int64 `json:"presale_end_time,omitempty" `

	/*
	   预售商品-付定金开始时间（毫秒）     */
	PresaleStartTime *int64 `json:"presale_start_time,omitempty" `

	/*
	   预售商品-定金（元）     */
	PresaleDeposit *string `json:"presale_deposit,omitempty" `

	/*
	   预售有礼-淘礼金发放时间     */
	YsylTljSendTime *string `json:"ysyl_tlj_send_time,omitempty" `

	/*
	   预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）     */
	YsylCommissionRate *string `json:"ysyl_commission_rate,omitempty" `

	/*
	   预售有礼-预估淘礼金（元）     */
	YsylTljFace *string `json:"ysyl_tlj_face,omitempty" `

	/*
	   预售有礼-推广链接     */
	YsylClickUrl *string `json:"ysyl_click_url,omitempty" `

	/*
	   预售有礼-淘礼金使用结束时间     */
	YsylTljUseEndTime *string `json:"ysyl_tlj_use_end_time,omitempty" `

	/*
	   预售有礼-淘礼金使用开始时间     */
	YsylTljUseStartTime *string `json:"ysyl_tlj_use_start_time,omitempty" `

	/*
	   本地化-销售开始时间     */
	SaleBeginTime *string `json:"sale_begin_time,omitempty" `

	/*
	   本地化-销售结束时间     */
	SaleEndTime *string `json:"sale_end_time,omitempty" `

	/*
	   本地化-到门店距离（米）     */
	Distance *string `json:"distance,omitempty" `

	/*
	   本地化-可用店铺id     */
	UsableShopId *string `json:"usable_shop_id,omitempty" `

	/*
	   本地化-可用店铺名称     */
	UsableShopName *string `json:"usable_shop_name,omitempty" `

	/*
	   活动价     */
	SalePrice *string `json:"sale_price,omitempty" `

	/*
	   跨店满减信息     */
	KuadianPromotionInfo *string `json:"kuadian_promotion_info,omitempty" `

	/*
	   是否品牌精选，0不是，1是     */
	SuperiorBrand *string `json:"superior_brand,omitempty" `

	/*
	   比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0     */
	RewardInfo *int32 `json:"reward_info,omitempty" `

	/*
	   是否品牌快抢，0不是，1是     */
	IsBrandFlashSale *string `json:"is_brand_flash_sale,omitempty" `

	/*
	   本地化-扩展信息     */
	LocalizationExtend *TaobaoTbkDgMaterialOptionalLocalizationMapData `json:"localization_extend,omitempty" `

	/*
	   物料评估-匹配分     */
	MatchScore *string `json:"match_score,omitempty" `

	/*
	   物料评估-收益分     */
	CommiScore *string `json:"commi_score,omitempty" `

	/*
	   是否是热门商品，0不是，1是     */
	HotFlag *string `json:"hot_flag,omitempty" `

	/*
	   前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
	TopnInfo *TaobaoTbkDgMaterialOptionalTopNInfoDTO `json:"topn_info,omitempty" `

	/*
	   百亿补贴信息     */
	BybtInfo *TaobaoTbkDgMaterialOptionalBybtInfoDTO `json:"bybt_info,omitempty" `

	/*
	   商品入驻淘特后产生的所有销量量级，不特指某段具体时间     */
	TtSoldCount *string `json:"tt_sold_count,omitempty" `

	/*
	   猫超买返卡信息     */
	MaifanPromotion *TaobaoTbkDgMaterialOptionalMaifanPromotionDTO `json:"maifan_promotion,omitempty" `

	/*
	   额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=单单奖励(已失效)，1=超级单单奖励(已失效)，2=年货节单单奖励     */
	CpaRewardType *string `json:"cpa_reward_type,omitempty" `

	/*
	   额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
	CpaRewardAmount *string `json:"cpa_reward_amount,omitempty" `

	/*
	   合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据     */
	ActivityId *string `json:"activity_id,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponStartTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponEndTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetInfoDxjh(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.InfoDxjh = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTkTotalSales(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.TkTotalSales = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTkTotalCommi(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.TkTotalCommi = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponId(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetNumIid(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.NumIid = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTitle(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Title = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPictUrl(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.PictUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSmallImages(v []string) *TaobaoTbkDgMaterialOptionalMapData {
	s.SmallImages = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetReservePrice(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ReservePrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetZkFinalPrice(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ZkFinalPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetUserType(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.UserType = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetProvcity(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Provcity = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetItemUrl(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ItemUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetIncludeMkt(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.IncludeMkt = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetIncludeDxjh(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.IncludeDxjh = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCommissionRate(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CommissionRate = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetVolume(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.Volume = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSellerId(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.SellerId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetShopTitle(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ShopTitle = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponTotalCount(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponTotalCount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponRemainCount(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponRemainCount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponInfo(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCommissionType(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CommissionType = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetShopDsr(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.ShopDsr = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponShareUrl(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponShareUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetUrl(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Url = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLevelOneCategoryName(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.LevelOneCategoryName = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.LevelOneCategoryId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCategoryName(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CategoryName = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCategoryId(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.CategoryId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetShortTitle(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ShortTitle = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetWhiteImage(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.WhiteImage = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetOetime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Oetime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetOstime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Ostime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetJddNum(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.JddNum = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetJddPrice(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.JddPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetUvSumPreSale(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.UvSumPreSale = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetXId(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.XId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponStartFee(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponStartFee = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCouponAmount(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CouponAmount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetItemDescription(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ItemDescription = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetNick(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Nick = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetOrigPrice(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.OrigPrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTotalStock(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.TotalStock = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSellNum(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.SellNum = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetStock(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.Stock = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTmallPlayActivityInfo(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.TmallPlayActivityInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetItemId(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.ItemId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetRealPostFee(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.RealPostFee = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLockRate(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.LockRate = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLockRateEndTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.LockRateEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLockRateStartTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.LockRateStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleDiscountFeeText(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleDiscountFeeText = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleTailEndTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleTailEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleTailStartTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleTailStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleEndTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleStartTime(v int64) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetPresaleDeposit(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.PresaleDeposit = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylTljSendTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylTljSendTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylCommissionRate(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylCommissionRate = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylTljFace(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylTljFace = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylClickUrl(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylClickUrl = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylTljUseEndTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylTljUseEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetYsylTljUseStartTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.YsylTljUseStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSaleBeginTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.SaleBeginTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSaleEndTime(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.SaleEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetDistance(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.Distance = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetUsableShopId(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.UsableShopId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetUsableShopName(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.UsableShopName = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSalePrice(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.SalePrice = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetKuadianPromotionInfo(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.KuadianPromotionInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetSuperiorBrand(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.SuperiorBrand = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetRewardInfo(v int32) *TaobaoTbkDgMaterialOptionalMapData {
	s.RewardInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetIsBrandFlashSale(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.IsBrandFlashSale = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetLocalizationExtend(v TaobaoTbkDgMaterialOptionalLocalizationMapData) *TaobaoTbkDgMaterialOptionalMapData {
	s.LocalizationExtend = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetMatchScore(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.MatchScore = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCommiScore(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CommiScore = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetHotFlag(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.HotFlag = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTopnInfo(v TaobaoTbkDgMaterialOptionalTopNInfoDTO) *TaobaoTbkDgMaterialOptionalMapData {
	s.TopnInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetBybtInfo(v TaobaoTbkDgMaterialOptionalBybtInfoDTO) *TaobaoTbkDgMaterialOptionalMapData {
	s.BybtInfo = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetTtSoldCount(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.TtSoldCount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetMaifanPromotion(v TaobaoTbkDgMaterialOptionalMaifanPromotionDTO) *TaobaoTbkDgMaterialOptionalMapData {
	s.MaifanPromotion = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCpaRewardType(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CpaRewardType = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetCpaRewardAmount(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.CpaRewardAmount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalMapData) SetActivityId(v string) *TaobaoTbkDgMaterialOptionalMapData {
	s.ActivityId = &v
	return s
}
