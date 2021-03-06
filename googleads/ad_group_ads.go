package v201809

import (
	"encoding/xml"
	"fmt"
)

type AdGroupAds []interface{}

func (a1 AdGroupAds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := a1[0]
	e.EncodeToken(start)
	switch a.(type) {
	case TextAd:
		ad := a.(TextAd)
		e.EncodeElement(ad.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
		e.EncodeElement(ad, xml.StartElement{
			xml.Name{"", "ad"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "TextAd"},
			},
		})
		e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{"", "status"}})
		e.EncodeElement(ad.Labels, xml.StartElement{Name: xml.Name{"", "labels"}})
	case ExpandedTextAd:
		ad := a.(ExpandedTextAd)
		e.EncodeElement(ad.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
		e.EncodeElement(ad, xml.StartElement{
			xml.Name{"", "ad"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "ExpandedTextAd"},
			},
		})
		e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{"", "status"}})
		e.EncodeElement(ad.Labels, xml.StartElement{Name: xml.Name{"", "labels"}})
	case Ad:
		ad := a.(Ad)
		e.EncodeElement(ad.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
		e.EncodeElement(ad, xml.StartElement{
			xml.Name{"", "ad"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "Ad"},
			},
		})
		e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{"", "status"}})
	case ImageAd:
		return ERROR_NOT_YET_IMPLEMENTED
	case TemplateAd:
		return ERROR_NOT_YET_IMPLEMENTED
	default:
		return fmt.Errorf("unknown Ad type -> %#v", start)
	}
	e.EncodeToken(start.End())
	return nil
}

func (aga *AdGroupAds) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	typeName := xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"}
	var adGroupId int64
	var status string
	var policySummary *AdGroupAdPolicySummary
	var labels []Label
	var baseCampaignId int64
	var baseAdGroupId int64
	var adStrengthInfo *AdStrengthInfo

	var ad interface{}
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "adGroupId":
				err := dec.DecodeElement(&adGroupId, &start)
				if err != nil {
					return err
				}
			case "ad":
				typeName, err := findAttr(start.Attr, typeName)
				if err != nil {
					return err
				}
				switch typeName {
				case "TextAd":
					a := TextAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ExpandedTextAd":
					a := ExpandedTextAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ImageAd":
					a := ImageAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "TemplateAd":
					a := TemplateAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "DynamicSearchAd":
					a := DynamicSearchAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ExpandedDynamicSearchAd":
					a := ExpandedDynamicSearchAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ResponsiveDisplayAd":
					a := ResponsiveDisplayAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "MultiAssetResponsiveDisplayAd":
					a := MultiAssetResponsiveDisplayAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
				case "ProductAd":
					a := ProductAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "GoalOptimizedShoppingAd":
					a := GoalOptimizedShoppingAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "CallOnlyAd":
					a := CallOnlyAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ResponsiveSearchAd":
					a := ResponsiveSearchAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "UniversalAppAd":
					a := UniversalAppAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ShowcaseAd":
					a := ShowcaseAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "RichMediaAd":
					a := RichMediaAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ThirdPartyRedirectAd":
					a := ThirdPartyRedirectAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "GmailAd":
					a := GmailAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "DeprecatedAd":
					a := DeprecatedAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				default:
					return fmt.Errorf("unknown AdGroupCriterion -> %#v", start)
				}
			case "status":
				err := dec.DecodeElement(&status, &start)
				if err != nil {
					return err
				}
			case "policySummary":
				err := dec.DecodeElement(&policySummary, &start)
				if err != nil {
					return err
				}
			case "labels":
				err := dec.DecodeElement(&labels, &start)
				if err != nil {
					return err
				}
			case "baseCampaignId":
				err := dec.DecodeElement(&baseCampaignId, &start)
				if err != nil {
					return err
				}
			case "baseAdGroupId":
				err := dec.DecodeElement(&baseAdGroupId, &start)
				if err != nil {
					return err
				}
			case "adStrengthInfo":
				err := dec.DecodeElement(&adStrengthInfo, &start)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown AdGroupAd field -> %#v", tag)
			}
		}
	}
	switch a := ad.(type) {
	case TextAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ExpandedTextAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ImageAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case TemplateAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case DynamicSearchAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ResponsiveDisplayAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case MultiAssetResponsiveDisplayAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ExpandedDynamicSearchAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ProductAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case GoalOptimizedShoppingAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case CallOnlyAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ResponsiveSearchAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case UniversalAppAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ShowcaseAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case RichMediaAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case ThirdPartyRedirectAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case GmailAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	case DeprecatedAd:
		a.Status = status
		a.PolicySummary = policySummary
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
		a.AdStrengthInfo = adStrengthInfo
		*aga = append(*aga, a)
	}

	return nil
}
