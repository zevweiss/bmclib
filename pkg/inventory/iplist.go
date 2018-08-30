package inventory

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/bmc-toolbox/bmcbutler/pkg/asset"
	"github.com/bmc-toolbox/bmcbutler/pkg/config"
)

// A inventory source is required to have a type with these fields
type IpList struct {
	Log       *logrus.Logger
	BatchSize int                  //number of inventory assets to return per iteration
	Channel   chan<- []asset.Asset //the channel to send inventory assets over
	Config    *config.Params       //bmcbutler config
}

type IpListAsset struct {
	BmcAddress string `csv:"bmcaddress"`
	Serial     string `csv:"serial"` //optional
	Vendor     string `csv:"vendor"` //optional
	Type       string `csv:"type"`   //optional
}

func (i *IpList) AssetIter(ipList string) {

	ips := strings.Split(ipList, ",")

	assets := make([]asset.Asset, 0)
	for _, ip := range ips {
		assets = append(assets, asset.Asset{IpAddress: ip})
	}

	//pass the asset to the channel
	i.Channel <- assets
	close(i.Channel)

}
