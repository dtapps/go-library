package cloudflare

// List, search, sort, and filter a zones' DNS records.
// https://developers.cloudflare.com/api/resources/dns/subresources/records/methods/list/
type ResponseDnsRecordsGet struct {
	Result []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
		Comment string `json:"comment,omitempty"`
	} `json:"result"`
}
