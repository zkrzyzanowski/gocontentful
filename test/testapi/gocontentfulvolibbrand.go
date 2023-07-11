// Code generated by https://github.com/foomo/gocontentful v1.0.22 - DO NOT EDIT.
package testapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/foomo/contentful"
)

const ContentTypeBrand = "brand"

// ---Brand private methods---

// ---Brand public methods---

func (cc *ContentfulClient) GetAllBrand() (voMap map[string]*CfBrand, err error) {
	if cc == nil {
		return nil, errors.New("GetAllBrand: No client available")
	}
	cc.cacheMutex.sharedDataGcLock.RLock()
	cacheInit := cc.cacheInit
	optimisticPageSize := cc.optimisticPageSize
	cc.cacheMutex.sharedDataGcLock.RUnlock()
	if cacheInit {
		return cc.Cache.entryMaps.brand, nil
	}
	col, err := cc.optimisticPageSizeGetAll("brand", optimisticPageSize)
	if err != nil {
		return nil, err
	}
	allBrand, err := colToCfBrand(col, cc)
	if err != nil {
		return nil, err
	}
	brandMap := map[string]*CfBrand{}
	for _, brand := range allBrand {
		brandMap[brand.Sys.ID] = brand
	}
	return brandMap, nil
}

func (cc *ContentfulClient) GetFilteredBrand(query *contentful.Query) (voMap map[string]*CfBrand, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("getFilteredBrand: No client available")
	}
	col := cc.Client.Entries.List(cc.SpaceID)
	if query != nil {
		col.Query = *query
	}
	col.Query.ContentType("brand").Locale("*").Include(0)
	_, err = col.GetAll()
	if err != nil {
		return nil, errors.New("getFilteredBrand: " + err.Error())
	}
	allBrand, err := colToCfBrand(col, cc)
	if err != nil {
		return nil, errors.New("getFilteredBrand: " + err.Error())
	}
	brandMap := map[string]*CfBrand{}
	for _, brand := range allBrand {
		brandMap[brand.Sys.ID] = brand
	}
	return brandMap, nil
}

func (cc *ContentfulClient) GetBrandByID(id string, forceNoCache ...bool) (vo *CfBrand, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("GetBrandByID: No client available")
	}
	if cc.cacheInit && (len(forceNoCache) == 0 || !forceNoCache[0]) {
		cc.cacheMutex.brandGcLock.RLock()
		defer cc.cacheMutex.brandGcLock.RUnlock()
		vo, ok := cc.Cache.entryMaps.brand[id]
		if ok {
			return vo, nil
		}
		return nil, fmt.Errorf("GetBrandByID: entry '%s' not found in cache", id)
	}
	col := cc.Client.Entries.List(cc.SpaceID)
	col.Query.ContentType("brand").Locale("*").Include(0).Equal("sys.id", id)
	_, err = col.GetAll()
	if err != nil {
		return nil, err
	}
	if len(col.Items) == 0 {
		return nil, fmt.Errorf("GetBrandByID: %s Not found", id)
	}
	vos, err := colToCfBrand(col, cc)
	if err != nil {
		return nil, fmt.Errorf("GetBrandByID: Error converting %s to VO: %w", id, err)
	}
	vo = vos[0]
	return
}

func NewCfBrand(contentfulClient ...*ContentfulClient) (cfBrand *CfBrand) {
	cfBrand = &CfBrand{}
	if len(contentfulClient) != 0 && contentfulClient[0] != nil {
		cfBrand.CC = contentfulClient[0]
	}

	cfBrand.Fields.CompanyName = map[string]string{}

	cfBrand.Fields.Logo = map[string]ContentTypeSys{}

	cfBrand.Fields.CompanyDescription = map[string]string{}

	cfBrand.Fields.Website = map[string]string{}

	cfBrand.Fields.Twitter = map[string]string{}

	cfBrand.Fields.Email = map[string]string{}

	cfBrand.Fields.Phone = map[string][]string{}

	cfBrand.Sys.ContentType.Sys.ID = "brand"
	cfBrand.Sys.ContentType.Sys.Type = FieldTypeLink
	cfBrand.Sys.ContentType.Sys.LinkType = "ContentType"
	return
}
func (vo *CfBrand) GetParents(contentType ...string) (parents []EntryReference, err error) {
	if vo == nil {
		return nil, errors.New("GetParents: Value Object is nil")
	}
	if vo.CC == nil {
		return nil, errors.New("GetParents: Value Object has no Contentful Client set")
	}
	return commonGetParents(vo.CC, vo.Sys.ID, contentType)
}

func (vo *CfBrand) GetPublishingStatus() string {
	if vo == nil {
		return ""
	}
	if vo.Sys.PublishedVersion == 0 {
		return StatusDraft
	}
	if vo.Sys.Version-vo.Sys.PublishedVersion == 1 {
		return StatusPublished
	}
	return StatusChanged
}

// Brand Field getters

func (vo *CfBrand) CompanyName(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockCompanyName.RLock()
	defer vo.Fields.RWLockCompanyName.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyName()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.CompanyName[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyName()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.CompanyName[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyName()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.CompanyName[string(loc)]
}

func (vo *CfBrand) Logo(locale ...Locale) *contentful.AssetNoLocale {
	if vo == nil {
		return nil
	}
	if vo.CC == nil {
		return nil
	}
	vo.Fields.RWLockLogo.RLock()
	defer vo.Fields.RWLockLogo.RUnlock()
	loc := defaultLocale
	reqLoc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		reqLoc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Logo()"}, LogError, ErrLocaleUnsupported)
			}
			return nil
		}
	}
	if _, ok := vo.Fields.Logo[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Logo()"}, LogWarn, ErrNotSet)
			}
			return nil
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Logo[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Logo()"}, LogWarn, ErrNotSetNoFallback)
			}
			return nil
		}
	}
	localizedLogo := vo.Fields.Logo[string(loc)]
	asset, err := vo.CC.GetAssetByID(localizedLogo.Sys.ID)
	if err != nil {
		if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
			vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Logo()"}, LogError, ErrNoTypeOfRefAsset)
		}
		return nil
	}
	tempAsset := &contentful.AssetNoLocale{}
	tempAsset.Sys = asset.Sys
	tempAsset.Fields = &contentful.FileFieldsNoLocale{}
	if _, ok := asset.Fields.Title[string(reqLoc)]; ok {
		tempAsset.Fields.Title = asset.Fields.Title[string(reqLoc)]
	} else {
		tempAsset.Fields.Title = asset.Fields.Title[string(loc)]
	}
	if _, ok := asset.Fields.Description[string(reqLoc)]; ok {
		tempAsset.Fields.Description = asset.Fields.Description[string(reqLoc)]
	} else {
		tempAsset.Fields.Description = asset.Fields.Description[string(loc)]
	}
	if _, ok := asset.Fields.File[string(reqLoc)]; ok {
		tempAsset.Fields.File = asset.Fields.File[string(reqLoc)]
	} else {
		tempAsset.Fields.File = asset.Fields.File[string(loc)]
	}
	return tempAsset
}

func (vo *CfBrand) CompanyDescription(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockCompanyDescription.RLock()
	defer vo.Fields.RWLockCompanyDescription.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyDescription()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.CompanyDescription[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyDescription()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.CompanyDescription[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CompanyDescription()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.CompanyDescription[string(loc)]
}

func (vo *CfBrand) Website(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockWebsite.RLock()
	defer vo.Fields.RWLockWebsite.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Website()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.Website[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Website()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Website[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Website()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.Website[string(loc)]
}

func (vo *CfBrand) Twitter(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockTwitter.RLock()
	defer vo.Fields.RWLockTwitter.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Twitter()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.Twitter[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Twitter()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Twitter[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Twitter()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.Twitter[string(loc)]
}

func (vo *CfBrand) Email(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockEmail.RLock()
	defer vo.Fields.RWLockEmail.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Email()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.Email[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Email()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Email[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Email()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.Email[string(loc)]
}

func (vo *CfBrand) Phone(locale ...Locale) []string {
	if vo == nil {
		return nil
	}
	if vo.CC == nil {
		return nil
	}
	vo.Fields.RWLockPhone.RLock()
	defer vo.Fields.RWLockPhone.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Phone()"}, LogError, ErrLocaleUnsupported)
			}
			return nil
		}
	}
	if _, ok := vo.Fields.Phone[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Phone()"}, LogWarn, ErrNotSet)
			}
			return nil
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Phone[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Phone()"}, LogWarn, ErrNotSetNoFallback)
			}
			return nil
		}
	}
	return vo.Fields.Phone[string(loc)]
}

// Brand Field setters

func (vo *CfBrand) SetCompanyName(companyName string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetCompanyName(companyName: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockCompanyName.Lock()
	defer vo.Fields.RWLockCompanyName.Unlock()
	if vo.Fields.CompanyName == nil {
		vo.Fields.CompanyName = make(map[string]string)
	}
	vo.Fields.CompanyName[string(loc)] = companyName
	return
}

func (vo *CfBrand) SetLogo(logo ContentTypeSys, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetLogo(logo: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockLogo.Lock()
	defer vo.Fields.RWLockLogo.Unlock()
	if vo.Fields.Logo == nil {
		vo.Fields.Logo = make(map[string]ContentTypeSys)
	}
	vo.Fields.Logo[string(loc)] = logo
	return
}

func (vo *CfBrand) SetCompanyDescription(companyDescription string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetCompanyDescription(companyDescription: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockCompanyDescription.Lock()
	defer vo.Fields.RWLockCompanyDescription.Unlock()
	if vo.Fields.CompanyDescription == nil {
		vo.Fields.CompanyDescription = make(map[string]string)
	}
	vo.Fields.CompanyDescription[string(loc)] = companyDescription
	return
}

func (vo *CfBrand) SetWebsite(website string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetWebsite(website: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockWebsite.Lock()
	defer vo.Fields.RWLockWebsite.Unlock()
	if vo.Fields.Website == nil {
		vo.Fields.Website = make(map[string]string)
	}
	vo.Fields.Website[string(loc)] = website
	return
}

func (vo *CfBrand) SetTwitter(twitter string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetTwitter(twitter: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockTwitter.Lock()
	defer vo.Fields.RWLockTwitter.Unlock()
	if vo.Fields.Twitter == nil {
		vo.Fields.Twitter = make(map[string]string)
	}
	vo.Fields.Twitter[string(loc)] = twitter
	return
}

func (vo *CfBrand) SetEmail(email string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetEmail(email: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockEmail.Lock()
	defer vo.Fields.RWLockEmail.Unlock()
	if vo.Fields.Email == nil {
		vo.Fields.Email = make(map[string]string)
	}
	vo.Fields.Email[string(loc)] = email
	return
}

func (vo *CfBrand) SetPhone(phone []string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetPhone(phone: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockPhone.Lock()
	defer vo.Fields.RWLockPhone.Unlock()
	if vo.Fields.Phone == nil {
		vo.Fields.Phone = make(map[string][]string)
	}
	vo.Fields.Phone[string(loc)] = phone
	return
}

func (vo *CfBrand) UpsertEntry() (err error) {
	if vo == nil {
		return errors.New("UpsertEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UpsertEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UpsertEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfBrand UpsertEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfBrand UpsertEntry: Can't unmarshal JSON into CF entry")
	}

	err = vo.CC.Client.Entries.Upsert(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfBrand UpsertEntry: Operation failed: %w", err)
	}
	return
}
func (vo *CfBrand) PublishEntry() (err error) {
	if vo == nil {
		return errors.New("PublishEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("PublishEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("PublishEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfBrand PublishEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfBrand PublishEntry: Can't unmarshal JSON into CF entry")
	}
	err = vo.CC.Client.Entries.Publish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfBrand PublishEntry: publish operation failed: %w", err)
	}
	return
}
func (vo *CfBrand) UnpublishEntry() (err error) {
	if vo == nil {
		return errors.New("UnpublishEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UnpublishEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UnpublishEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfBrand UnpublishEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfBrand UnpublishEntry: Can't unmarshal JSON into CF entry")
	}
	err = vo.CC.Client.Entries.Unpublish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfBrand UnpublishEntry: unpublish operation failed: %w", err)
	}
	return
}
func (vo *CfBrand) UpdateEntry() (err error) {
	if vo == nil {
		return errors.New("UpdateEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UpdateEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UpdateEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfBrand UpdateEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfBrand UpdateEntry: Can't unmarshal JSON into CF entry")
	}

	err = vo.CC.Client.Entries.Upsert(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfBrand UpdateEntry: upsert operation failed: %w", err)
	}
	tmp, errMarshal = json.Marshal(cfEntry)
	if errMarshal != nil {
		return errors.New("CfBrand UpdateEntry: Can't marshal JSON back from CF entry")
	}
	errUnmarshal = json.Unmarshal(tmp, &vo)
	if errUnmarshal != nil {
		return errors.New("CfBrand UpdateEntry: Can't unmarshal JSON back into VO")
	}
	err = vo.CC.Client.Entries.Publish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfBrand UpdateEntry: publish operation failed: %w", err)
	}
	return
}
func (vo *CfBrand) DeleteEntry() (err error) {
	if vo == nil {
		return errors.New("DeleteEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("DeleteEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("DeleteEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfBrand DeleteEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfBrand DeleteEntry: Can't unmarshal JSON into CF entry")
	}
	if cfEntry.Sys.PublishedCounter > 0 {
		errUnpublish := vo.CC.Client.Entries.Unpublish(vo.CC.SpaceID, cfEntry)
		if errUnpublish != nil && !strings.Contains(errUnpublish.Error(), "Not published") {
			return fmt.Errorf("CfBrand DeleteEntry: Unpublish entry failed: %w", errUnpublish)
		}
	}
	errDelete := vo.CC.Client.Entries.Delete(vo.CC.SpaceID, cfEntry.Sys.ID)
	if errDelete != nil {
		return fmt.Errorf("CfBrand DeleteEntry: Delete entry failed: %w", errDelete)
	}
	return nil
}
func (vo *CfBrand) ToReference() (refSys ContentTypeSys) {
	if vo == nil {
		return refSys
	}
	refSys.Sys.ID = vo.Sys.ID
	refSys.Sys.Type = FieldTypeLink
	refSys.Sys.LinkType = FieldLinkTypeEntry
	return
}

func (cc *ContentfulClient) cacheAllBrand(ctx context.Context, resultChan chan<- ContentTypeResult) (vos map[string]*CfBrand, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("cacheAllBrand: No CDA/CPA client available")
	}
	var allBrand []*CfBrand
	col := &contentful.Collection{
		Items: []interface{}{},
	}
	cc.cacheMutex.sharedDataGcLock.RLock()
	defer cc.cacheMutex.sharedDataGcLock.RUnlock()
	if cc.offline {
		for _, entry := range cc.offlineTemp.Entries {
			if entry.Sys.ContentType.Sys.ID == ContentTypeBrand {
				col.Items = append(col.Items, entry)
			}
		}
	} else {
		col, err = cc.optimisticPageSizeGetAll("brand", cc.optimisticPageSize)
		if err != nil {
			return nil, errors.New("optimisticPageSizeGetAll for Brand failed: " + err.Error())
		}
	}
	allBrand, err = colToCfBrand(col, cc)
	if err != nil {
		return nil, errors.New("colToCfBrand failed: " + err.Error())
	}
	brandMap := map[string]*CfBrand{}
	for _, brand := range allBrand {
		if cc.cacheInit {
			existingBrand, err := cc.GetBrandByID(brand.Sys.ID)
			if err == nil && existingBrand != nil && existingBrand.Sys.Version > brand.Sys.Version {
				return nil, fmt.Errorf("cache update canceled because Brand entry %s is newer in cache", brand.Sys.ID)
			}
		}
		brandMap[brand.Sys.ID] = brand
		result := ContentTypeResult{
			EntryID:     brand.Sys.ID,
			ContentType: ContentTypeBrand,
			References:  map[string][]EntryReference{},
		}
		addEntry := func(id string, refs EntryReference) {
			if result.References[id] == nil {
				result.References[id] = []EntryReference{}
			}
			result.References[id] = append(result.References[id], refs)
		}
		_ = addEntry

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		resultChan <- result
	}
	return brandMap, nil
}

func (cc *ContentfulClient) cacheBrandByID(ctx context.Context, id string, entryPayload *contentful.Entry, entryDelete bool) error {
	cc.cacheMutex.brandGcLock.Lock()
	defer cc.cacheMutex.brandGcLock.Unlock()
	cc.cacheMutex.idContentTypeMapGcLock.Lock()
	defer cc.cacheMutex.idContentTypeMapGcLock.Unlock()
	cc.cacheMutex.parentMapGcLock.Lock()
	defer cc.cacheMutex.parentMapGcLock.Unlock()

	var col *contentful.Collection
	if entryPayload != nil {
		col = &contentful.Collection{
			Items: []interface{}{entryPayload},
		}
		id = entryPayload.Sys.ID
	} else {
		if cc.Client == nil {
			return errors.New("cacheBrandByID: No client available")
		}
		if !entryDelete {
			col = cc.Client.Entries.List(cc.SpaceID)
			col.Query.ContentType("brand").Locale("*").Include(0).Equal("sys.id", id)
			_, err := col.GetAll()
			if err != nil {
				return err
			}
		}
	}
	// It was deleted
	if col != nil && len(col.Items) == 0 || entryDelete {
		delete(cc.Cache.entryMaps.brand, id)
		delete(cc.Cache.idContentTypeMap, id)
		// delete as child
		delete(cc.Cache.parentMap, id)
		// delete as parent
		for childID, parents := range cc.Cache.parentMap {
			newParents := []EntryReference{}
			for _, parent := range parents {
				if parent.ID != id {
					newParents = append(newParents, parent)
				}
			}
			cc.Cache.parentMap[childID] = newParents
		}
		return nil
	}
	vos, err := colToCfBrand(col, cc)
	if err != nil {
		return fmt.Errorf("cacheBrandByID: Error converting %s to VO: %w", id, err)
	}
	brand := vos[0]
	if cc.Cache.entryMaps.brand == nil {
		cc.Cache.entryMaps.brand = map[string]*CfBrand{}
	}
	cc.Cache.entryMaps.brand[id] = brand
	cc.Cache.idContentTypeMap[id] = brand.Sys.ContentType.Sys.ID
	allChildrensIds := map[string]bool{}

	_ = allChildrensIds // safety net
	// clean up child-parents that don't exist anymore
	for childID, parents := range cc.Cache.parentMap {
		if _, isCollectedChildID := allChildrensIds[childID]; isCollectedChildID {
			continue
		}
		newParents := []EntryReference{}
		for _, parent := range parents {
			if parent.ID != id {
				newParents = append(newParents, parent)
			}
		}
		cc.Cache.parentMap[childID] = newParents
	}
	return nil
}

func colToCfBrand(col *contentful.Collection, cc *ContentfulClient) (vos []*CfBrand, err error) {
	for _, item := range col.Items {
		var vo CfBrand
		byteArray, _ := json.Marshal(item)
		err = json.NewDecoder(bytes.NewReader(byteArray)).Decode(&vo)
		if err != nil {
			break
		}
		if cc.textJanitor {

			vo.Fields.CompanyName = cleanUpStringField(vo.Fields.CompanyName)

			vo.Fields.CompanyDescription = cleanUpStringField(vo.Fields.CompanyDescription)

			vo.Fields.Website = cleanUpStringField(vo.Fields.Website)

			vo.Fields.Twitter = cleanUpStringField(vo.Fields.Twitter)

			vo.Fields.Email = cleanUpStringField(vo.Fields.Email)

			vo.Fields.Phone = cleanUpStringSliceField(vo.Fields.Phone)

		}
		vo.CC = cc
		vos = append(vos, &vo)
	}
	return vos, err
}
