package v201809

import "encoding/xml"

type CustomerSyncService struct {
	Auth
}

func NewCustomerSyncService(auth *Auth) *CustomerSyncService {
	return &CustomerSyncService{Auth: *auth}
}

type AdGroupChangeData struct {
	AdGroupID                         int64   `xml:"adGroupId"`
	AdGroupChangeStatus               string  `xml:"adGroupChangeStatus"`
	ChangedAds                        []int64 `xml:"changedAds"`
	ChangedCriteria                   []int64 `xml:"changedCriteria"`
	RemovedCriteria                   []int64 `xml:"removedCriteria"`
	ChangedFeeds                      []int64 `xml:"changedFeeds"`
	RemovedFeeds                      []int64 `xml:"removedFeeds"`
	ChangedAdGroupBidModifierCriteria []int64 `xml:"changedAdGroupBidModifierCriteria"`
	RemovedAdGroupBidModifierCriteria []int64 `xml:"removedAdGroupBidModifierCriteria"`
}

type CampaignChangeData struct {
	CampaignID              int64             `xml:"campaignId"`
	CampaignChangeStatus    string            `xml:"campaignChangeStatus"`
	ChangedAdGroups         AdGroupChangeData `xml:"changedAdGroups"`
	AddedCampaignCriteria   []int64           `xml:"addedCampaignCriteria"`
	RemovedCampaignCriteria []int64           `xml:"removedCampaignCriteria"`
	ChangedFeeds            []int64           `xml:"changedFeeds"`
	RemovedFeeds            []int64           `xml:"removedFeeds"`
}

type FeedChangeData struct {
	//todo feed
}

type CustomerChangeData struct {
	ChangedCampaigns    *[]CampaignChangeData `xml:"rval>changedCampaigns"`
	ChangedFeeds        *[]FeedChangeData     `xml:"rval>changedFeeds"`
	LastChangeTimestamp string                `xml:"rval>lastChangeTimestamp"`
}

type CustomerSyncSelector struct {
	XMLName       xml.Name
	DateTimeRange DateRange `xml:"dateTimeRange"`
	CampaignIds   *[]int64  `xml:"campaignIds,omitempty"`
	FeedIds       *[]int64  `xml:"feedIds,omitempty"`
}

func (s *CustomerSyncService) Get(selector CustomerSyncSelector) (changeData CustomerChangeData, err error) {
	selector.XMLName = xml.Name{baseSyncUrl, "selector"}

	respBody, err := s.Auth.request(
		customerSyncServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     CustomerSyncSelector
		}{
			XMLName: xml.Name{
				Space: baseSyncUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return changeData, err
	}
	getResp := CustomerChangeData{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return changeData, err
	}
	return getResp, err
}
