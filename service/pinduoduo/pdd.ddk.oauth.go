package pinduoduo

type PddDdkOauthClient struct {
	client *Client
}

// PddDdkOauth 工具商接口
func (c *Client) PddDdkOauth() *PddDdkOauthClient {
	return &PddDdkOauthClient{
		client: c,
	}
}

type PddDdkOauthCashGiftApi struct {
	client *Client
}

// CashGift 多多礼金接口
func (c *PddDdkOauthClient) CashGift() *PddDdkOauthCashGiftApi {
	return &PddDdkOauthCashGiftApi{
		client: c.client,
	}
}

type PddDdkOauthCmsApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Cms() *PddDdkOauthCmsApi {
	return &PddDdkOauthCmsApi{
		client: c.client,
	}
}

type PddDdkOauthGoodsApi struct {
	client *Client
}

// Goods 商品接口
func (c *PddDdkOauthClient) Goods() *PddDdkOauthGoodsApi {
	return &PddDdkOauthGoodsApi{
		client: c.client,
	}
}

type PddDdkOauthMemberApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Member() *PddDdkOauthMemberApi {
	return &PddDdkOauthMemberApi{
		client: c.client,
	}
}

type PddDdkOauthOrderApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Order() *PddDdkOauthOrderApi {
	return &PddDdkOauthOrderApi{
		client: c.client,
	}
}

type PddDdkOauthPidApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Pid() *PddDdkOauthPidApi {
	return &PddDdkOauthPidApi{
		client: c.client,
	}
}

type PddDdkOauthResourceApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Resource() *PddDdkOauthResourceApi {
	return &PddDdkOauthResourceApi{
		client: c.client,
	}
}

type PddDdkOauthRpApi struct {
	client *Client
}

func (c *PddDdkOauthClient) Rp() *PddDdkOauthRpApi {
	return &PddDdkOauthRpApi{
		client: c.client,
	}
}
