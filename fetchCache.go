package main

import lookupcache "github.com/nishantd01/eyeota_assignment/lookupcache"

func (cache DataMap) GetSegmentForOrgAndKey(orgKey string, paramKey string) []lookupcache.SegmentConfig {

	resultString := cache.dataMap[orgKey][paramKey][""]
	result := []lookupcache.SegmentConfig{}

	for _, seg := range resultString {
		result = append(result, lookupcache.SegmentConfig{Id: seg})
	}
	return result

}

func (cache DataMap) GetSegmentForOrgAndKeyAndVal(orgKey string, paramKey string, paramVal string) []lookupcache.SegmentConfig {
	resultString := cache.dataMap[orgKey][paramKey][paramVal]
	result := []lookupcache.SegmentConfig{}

	for _, seg := range resultString {
		result = append(result, lookupcache.SegmentConfig{Id: seg})
	}
	return result
}
