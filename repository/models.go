// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

type PlmnTraffic struct {
	Date              string  `json:"date"`
	Moentity          string  `json:"moentity"`
	RanSiteCity       string  `json:"ran_site_city"`
	RanSiteRegion     string  `json:"ran_site_region"`
	RanSiteLongitude  float64 `json:"ran_site_longitude"`
	RanSiteLatitude   float64 `json:"ran_site_latitude"`
	RanSiteSiteVendor string  `json:"ran_site_site_vendor"`
	Traffic           float64 `json:"traffic"`
	TrafficType       string  `json:"traffic_type"`
}
