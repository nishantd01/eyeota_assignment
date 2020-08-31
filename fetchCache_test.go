package main

import (
	"testing"
)

func TestGetSegmentForOrgAndKey(t *testing.T) {

	cache := initializeCache("./data/testData.json")

	vals := cache.GetSegmentForOrgAndKey("org1", "paramName1")

	// Check length, it will return an empty SegmentConfig array , if not throw an error
	if len(vals) != 0 {
		t.Errorf("The array is not empty!")
	}

	vals = cache.GetSegmentForOrgAndKey("org1", "testedu")
	if len(vals) != 1 {
		t.Error(" length is invalid for org1,testedu, it should be 1\n")
	}
	//value should be n277
	if vals[0].Id != "n277" {
		t.Errorf(" Invalid value =%v  , Value should have been n277\n", vals[0].Id)
	}

	vals = cache.GetSegmentForOrgAndKey("org1", "gen")

	if len(vals) != 0 {
		t.Errorf("The array is not empty! it should not contain value for org1 , gen = %v", vals[0].Id)
	}

}

func TestGetSegmentForOrgAndKeyAndVal(t *testing.T) {

	cache := initializeCache("./data/testData.json")

	vals := cache.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal1")

	if len(vals) != 1 {
		t.Error(" length is invalid for org1,testedu, it should be 1\n")
	}

	if vals[0].Id != "seg_1234" {
		t.Errorf(" Invalid value =%v  , Value should have been seg_1234\n", vals[0].Id)
	}

	vals = cache.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2")

	if len(vals) != 1 {
		t.Error(" length is invalid for org1,paramName1,paramVal2 it should be 1\n")
	}

	if vals[0].Id != "intr.edu" {
		t.Errorf(" Invalid value =%v  , Value should have been intr.edu\n", vals[0].Id)
	}

	vals = cache.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal6")

	if len(vals) != 3 {
		t.Error(" length is invalid for org1,paramName1,paramVal6 it should be 3\n")
	}

	var visitedVal map[string]bool
	visitedVal = make(map[string]bool)

	for _, val := range vals {
		id := val.Id
		if id != "dem.infg.m" && id != "dem.infg.f" && id != "intr.heal" {
			t.Error(" Invalid value it should be either dem.infg.m or intr.heal or dem.infg.f\n")
		} else if visitedVal[val.Id] == true {
			t.Errorf(" Duplicate Value =%v\n", val.Id)
		} else {
			visitedVal[val.Id] = true
		}
	}

	vals = cache.GetSegmentForOrgAndKeyAndVal("org1", "testedu", "")

	// log.Println(vals)

	if len(vals) != 1 {
		t.Error(" length is invalid for org1,testedu, blank it should be 1\n")
	}

	if vals[0].Id != "n277" {
		t.Errorf(" Invalid value =%v  , Value should have been n277\n", vals[0].Id)
	}

	vals = cache.GetSegmentForOrgAndKeyAndVal("org1", "testedu", "random")

	// log.Println(vals)

	if len(vals) != 0 {
		t.Error(" length is invalid for org1,testedu,random  it should be 0\n")
	}

}
