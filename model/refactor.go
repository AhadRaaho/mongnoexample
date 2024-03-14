package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shipment struct {
	ID                   primitive.ObjectID `json:"_id" bson:"_id"`
	LoadingPointCallDone bool               `json:"loadingPointCallDone"`
	TelcoTracking        bool               `json:"telcoTracking"`
	Indent               int64              `json:"indent"`
	PcLinked             []PcLinked         `json:"pcLinked"`
	IsPcLinked           bool               `json:"isPcLinked"`
	UpdatedAt            time.Time          `json:"updatedAt"`
	CreatedAt            time.Time          `json:"createdAt"`
	Version              int                `json:"__v"`
	HasPlacement         bool               `json:"hasPlacement"`
	IsTruckerAccepted    bool               `json:"isTruckerAccepted"`
	Placement            Placement          `json:"placement"`
	ReportingTime        string             `json:"reportingTime"`
	State                string             `json:"state"`
	TruckerUserID        string             `json:"truckerUserId"`
	IsIpLinked           bool               `json:"isIpLinked"`
	Places               Places             `json:"places"`
}

type PcLinked struct {
	CompanyUUID    int64  `json:"companyUUID"`
	Name           string `json:"name"`
	Comments       string `json:"comments,omitempty"`
	CompanyUUIDStr string `json:"companyUUIDStr"`
}

type Placement struct {
	DocStoreID          string       `json:"docStoreId"`
	Indent              string       `json:"indent"`
	Partner             string       `json:"partner"`
	Price               int64        `json:"price"`
	Tat                 string       `json:"tat"`
	VerifiedPartner     bool         `json:"verifiedPartner"`
	DriverName          string       `json:"driverName"`
	DriverMobile        string       `json:"driverMobile"`
	DriverID            string       `json:"driverId"`
	VehicleID           int64        `json:"vehicleId"`
	VehicleNumber       string       `json:"vehicleNumber"`
	VehicleDocs         *VehicleDocs `json:"vehicleDocs"`
	VehicleType         string       `json:"vehicleType"`
	PlacedByUser        string       `json:"placedByUser"`
	PlacedByUserID      string       `json:"placedByUserId"`
	PlacedByUserPhone   string       `json:"placedByUserPhone"`
	PlacementTime       string       `json:"placementTime"`
	SupplierCompanyUUID string       `json:"suppliercompany_uuid"`
	SupplyManagerUUID   int64        `json:"supplymanager_uuid"`
	AdvancePercentage   int64        `json:"advancePercentage"`
	Item                string       `json:"item"`
	LoadingAddress      string       `json:"loadingAddress"`
	Origin              string       `json:"origin"`
	Destination         string       `json:"destination"`
	RegNo               string       `json:"regNo"`
	DemandMgrPhone      string       `json:"demandMgrPhone"`
	Tds                 int64        `json:"tds"`
	PricingType         string       `json:"pricing_type"`
	Mgt                 int64        `json:"mgt"`
	PptSp               int64        `json:"ppt_sp"`
	PptSh               int64        `json:"ppt_sh"`
	Meta                Meta         `json:"meta"`
}

type VehicleDocs struct {
	Rc   []*Doc `json:"rc"`
	Ins  []*Doc `json:"ins"`
	Prmt []*Doc `json:"prmt"`
}

type Doc struct {
	Url      string `json:"url"`
	Verified bool   `json:"verified"`
}

type Meta struct {
	IncentiveInventory bool   `json:"incentiveInventory"`
	Remarks            string `json:"remarks"`
	SupplierPriority   int    `json:"supplierPriority"`
	PartnerStage       string `json:"partner_stage"`
}

type Places struct {
	Origin      Place1 `json:"origin"`
	Destination Place  `json:"destination"`
}

type Place struct {
	ID             int64   `json:"id"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Address        string  `json:"address"`
	PlacesID       string  `json:"placesId"`
	Verified       bool    `json:"verified"`
	Rejected       bool    `json:"rejected"`
	CityPlaceID    string  `json:"cityPlaceId"`
	CityPlaceLabel string  `json:"cityPlaceLabel"`
	Landmark       string  `json:"landmark"`
	ShipperName    string  `json:"shipperName"`
	ShipperUUID    string  `json:"shipperUUID"`
	ShipperUUIDStr string  `json:"shipperUUIDStr"`
	Zone           string  `json:"zone"`
	StateZone      string  `json:"stateZone"`
	PlacesIDStr    string  `json:"placesIdStr"`
	Label          string  `json:"label"`
}
type Place1 struct {
	ID             int64   `json:"id"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Address        string  `json:"address"`
	PlacesID       int64   `json:"placesId"`
	Verified       bool    `json:"verified"`
	Rejected       bool    `json:"rejected"`
	CityPlaceID    string  `json:"cityPlaceId"`
	CityPlaceLabel string  `json:"cityPlaceLabel"`
	Landmark       string  `json:"landmark"`
	ShipperName    string  `json:"shipperName"`
	ShipperUUID    string  `json:"shipperUUID"`
	ShipperUUIDStr string  `json:"shipperUUIDStr"`
	Zone           string  `json:"zone"`
	StateZone      string  `json:"stateZone"`
	PlacesIDStr    string  `json:"placesIdStr"`
	Label          string  `json:"label"`
}
