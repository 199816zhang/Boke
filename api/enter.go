package api

import (
	"blogx_server/api/banner_api"
	"blogx_server/api/image_api"
	"blogx_server/api/site_api"
)

type Api struct {
	SiteApi   site_api.SiteApi
	ImageApi  image_api.ImageApi
	BannerApi banner_api.BannerApi
}

var App = Api{}
