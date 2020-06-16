package main

import (
	adGroup "./beeswax/adgroup"
	augment "./beeswax/augment"
	base "./beeswax/base"
	bid "./beeswax/bid"
	logs "./beeswax/logs"
	openrtb "./beeswax/openrtb"

)

func main() {
	_ = adGroup.GhostBidding{}
	_ = augment.AugmentorRequest{}
	_ = base.AdGroupId{}
	_ = bid.Adcandidate{}
	_ = logs.ActivityLogMessage{}
	_ = openrtb.AdvertiserInfo{}

}