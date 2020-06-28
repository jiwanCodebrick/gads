package v201809

import (
	"encoding/xml"
	"fmt"
)

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.ExtensionSetting
// A setting specifying when and which extensions should serve at a given level (customer, campaign, or ad group).
type ExtensionSetting struct {
	PlatformRestrictions ExtensionSettingPlatform `xml:"platformRestrictions,omitempty"`

	Extensions []Extension `xml:"extensions,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.ExtensionSetting.Platform
// Different levels of platform restrictions
// DESKTOP, MOBILE, NONE
type ExtensionSettingPlatform string

type Extension interface{}

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.CallFeedItem
// Represents a Call extension.
type CallFeedItem struct {
	XMLName xml.Name `json:"-" xml:"extensions"`

	FeedId                  int64                      `xml:"feedId,omitempty"`
	FeedItemId              int64                      `xml:"feedItemId,omitempty"`
	Status                  string                     `xml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty"`
	StartTime               string                     `xml:"startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                     `xml:"endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []FeedItemPolicySummary    `xml:"policySummaries,omitempty"`
	ExtensionFeedItemType   string                     `xml:"ExtensionFeedItem.Type,omitempty"`

	CallPhoneNumber               string             `xml:"callPhoneNumber,omitempty"`
	CallCountryCode               string             `xml:"callCountryCode,omitempty"`
	CallTracking                  bool               `xml:"callTracking,omitempty"`
	CallConversionType            CallConversionType `xml:"callConversionType,omitempty"`
	DisableCallConversionTracking bool               `xml:"disableCallConversionTracking,omitempty"`
}

type SitelinkFeedItem struct {
	FeedId                  int64                      `xml:"feedId,omitempty"`
	FeedItemId              int64                      `xml:"feedItemId,omitempty"`
	Status                  string                     `xml:"status,omitempty"`
	FeedType                *FeedType                  `xml:"feedType,omitempty"`
	StartTime               string                     `xml:"startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                     `xml:"endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        *FeedItemDevicePreference  `xml:"devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"geoTargetingRestriction,omitempty"`
	PolicySummaries         []FeedItemPolicySummary    `xml:"policySummaries,omitempty"`
	ExtensionFeedItemType   string                     `xml:"ExtensionFeedItem.Type,omitempty"`

	SitelinkText                string           `xml:"sitelinkText,omitempty"`
	SitelinkUrl                 string           `xml:"sitelinkUrl,omitempty"`
	SitelinkLine2               string           `xml:"sitelinkLine2,omitempty"`
	SitelinkLine3               string           `xml:"sitelinkLine3,omitempty"`
	SitelinkFinalUrls           UrlList          `xml:"sitelinkFinalUrls,omitempty"`
	SitelinkFinalMobileUrls     UrlList          `xml:"sitelinkFinalMobileUrls,omitempty"`
	SitelinkTrackingUrlTemplate string           `xml:"sitelinkTrackingUrlTemplate,omitempty"`
	SitelinkFinalUrlSuffix      string           `xml:"sitelinkFinalUrlSuffix,omitempty"`
	SitelinkUrlCustomParameters CustomParameters `xml:"sitelinkUrlCustomParameters,omitempty"`
}

func extensionsUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (ext interface{}, err error) {
	extensionsType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return
	}
	switch extensionsType {
	case "CallFeedItem":
		c := CallFeedItem{}
		err = dec.DecodeElement(&c, &start)
		if err != nil {
			return nil, err
		}
		ext = c
	case "SitelinkFeedItem":
		c := SitelinkFeedItem{}
		err = dec.DecodeElement(&c, &start)
		if err != nil {
			return nil, err
		}
		ext = c
	default:
		err = fmt.Errorf("unknown Extensions type %#v", extensionsType)
	}
	return
}

func (s ExtensionSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// e.EncodeToken(start)
	// if s.PlatformRestrictions != "NONE" {
	// 	e.EncodeElement(&s.PlatformRestrictions, xml.StartElement{Name: xml.Name{
	// 		"https://adwords.google.com/api/adwords/cm/v201809",
	// 		"platformRestrictions"}})
	// }
	// switch extType := s.Extensions.(type) {
	// case []CallFeedItem:
	// 	e.EncodeElement(s.Extensions.([]CallFeedItem), xml.StartElement{
	// 		xml.Name{baseUrl, "extensions"},
	// 		[]xml.Attr{
	// 			xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "CallFeedItem"},
	// 		},
	// 	})
	// case []SitelinkFeedItem:
	// 	e.EncodeElement(s.Extensions.([]CallFeedItem), xml.StartElement{
	// 		xml.Name{baseUrl, "extensions"},
	// 		[]xml.Attr{
	// 			xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "SitelinkFeedItem"},
	// 		},
	// 	})
	// default:
	// 	return fmt.Errorf("unknown extension type %#v\n", extType)

	// }

	// e.EncodeToken(start.End())
	// return nil
	return ERROR_NOT_YET_IMPLEMENTED
}

func (s *ExtensionSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	s.Extensions = []Extension{}

	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "platformRestrictions":
				if err := dec.DecodeElement(&s.PlatformRestrictions, &start); err != nil {
					return err
				}
			case "extensions":
				extension, err := extensionsUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				s.Extensions = append(s.Extensions, extension)
			}
		}
	}
	return nil
}
