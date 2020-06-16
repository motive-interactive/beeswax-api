package beeswax

import (
	"beeswax/adgroup"
	"beeswax/augment"
	"beeswax/base"
	"beeswax/bid"
	"beeswax/currency"
	"beeswax/logs"
	"beeswax/openrtb"
	"fmt"
)

func testPkg() {
	fmt.Print(adgroup.GhostBidding{})
	fmt.Print(augment.AugmentorRequest{})
	fmt.Print(base.AdGroupId{})
	fmt.Print(bid.Adcandidate{})
	fmt.Print(currency.Currency{})
	fmt.Print(logs.ActivityLogMessage{})
	fmt.Print(openrtb.AdvertiserInfo{})
}