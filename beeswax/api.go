package beeswax

import (
	"fmt"
	"github.com/motive-interactive/beeswax-api/beeswax/adgroup"
	"github.com/motive-interactive/beeswax-api/beeswax/augment"
	"github.com/motive-interactive/beeswax-api/beeswax/base"
	"github.com/motive-interactive/beeswax-api/beeswax/bid"
	"github.com/motive-interactive/beeswax-api/beeswax/currency"
	"github.com/motive-interactive/beeswax-api/beeswax/logs"
)

func testPkg() {
	fmt.Print(adgroup.GhostBidding{})
	fmt.Print(augment.AugmentorRequest{})
	fmt.Print(base.AdGroupId{})
	fmt.Print(bid.Adcandidate{})
	fmt.Print(currency.Currency{})
	fmt.Print(logs.ActivityLogMessage{})
}