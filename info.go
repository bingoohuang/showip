package ip

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// nolint lll
// https://topic.alibabacloud.com/a/go-combat-golang-get-public-ip-view-intranet-ip-detect-ip-type-verify-ip-range-ip-address-string-and-int-conversion-judge-by-ip_1_38_10267608.html

// Info ...
type Info struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}

// IP ...
type IP struct {
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	Isp       string `json:"isp"`
}

// TabaoAPI ...
func TabaoAPI(ip string) *Info {
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*10)
	defer cncl()

	addr := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
	if err != nil {
		logrus.Warnf("failed http.Get(%s), × err: %v", addr, err)

		return nil
	}

	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("failed ioutil.ReadAll, × err: %v", err)
		return nil
	}

	var result Info

	if err := json.Unmarshal(out, &result); err != nil {
		logrus.Warnf("failed json.Unmarshal %s, × err: %v", string(out), err)
		return nil
	}

	return &result
}
