package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
type Billboard struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"billboardname" bson:"billboardname" validate:"required,min=2, max=100"`
	PricePerDay  int                `json:"priceperday" bson:"priceperday" validate:"required"`
	DailyViews   int                `json:"dailyviews" bson:"dailyviews" validate:"required"`
	Dimensions   string             `json:"dimensions" bson:"dimensions" validate:"required"`
	VacantFrom   time.Time          `json:"vacant" bson:"vacantfrom" validate:"required"`
	VacantTill   time.Time          `json:"vacanttill" bson:"vacanttill" validate:"required"`
	OwnerDetails *OwnerDetails      `json:"ownerdetails" bson:"ownerdetails"` //take details of user->seller
	Description  string             `json:"description" bson:"description" validate:"required,min=2, max=500"`
	LocationLink string             `json:"location" bson:"location" validate:"required"`
	LandMarks    *LandMarks         `json:"landmarks" bson:"landmarks"`
	Discounts    int                `json:"discount" bson:"discounts"`
	Stars        int                `json:"stars" bson:"stars" validate:"required,max=5"`
}

type OwnerDetails struct {
	AgencyID            int    `json:"agencyid" bson:"agencyid"`
	AgencyName          string `json:"agencyname" bson:"agencyname" validate:"required,min=2,max=100"`
	AgencyContactNumber int    `json:"agencycontactnumber" bson:"agencycontactnumber" validate:"required,min"`
}
type LandMarks struct {
	Name string `josn:"name" bson:"name"`
}
type Influencers struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	PersonalDetails *PersonalDetails   `json:"personal_details" bson:"personal details" validate:"required"`
	Services        *Services          `json:"servicesoffered" bson:"servicesoffered" validate:"required"`
	Platforms       *Platforms         `json:"platforms" bson:"platforms" validate:"required"`
	Dimensions      string             `json:"dimensions" bson:"dimensions" validate:"required"`
	OwnerDetails    *OwnerDetails      `json:"ownerdetails" bson:"ownerdetails"` //take details of user->seller
	Description     string             `json:"description" bson:"description" validate:"required,min=2, max=500"`
	LocationLink    string             `json:"location" bson:"location" validate:"required"`
	Discounts       int                `json:"discount" bson:"discounts"`
	Stars           int                `json:"stars" bson:"stars" validate:"required,max=5"`
}
type Service struct {
	Name  string `json:"name" bson:"name"`
	Views string `json:"views" bson:"views"`
	Price int    `json:"price" bson:"price"`
}
type Platform struct {
	Name  string `json:"name" bson:"name"`   // e.g., "youtube", "instagram", "customplatform"
	Reach int    `json:"reach" bson:"reach"` // Number of followers or audience size
	Since string `json:"since" bson:"since"` // Date or duration, e.g., "2018-06-15" or "3 years"
}

type Platforms struct {
	YouTube        *Platform  `json:"youtube,omitempty" bson:"youtube,omitempty"`
	Instagram      *Platform  `json:"instagram,omitempty" bson:"instagram,omitempty"`
	Snapchat       *Platform  `json:"snapchat,omitempty" bson:"snapchat,omitempty"`
	X              *Platform  `json:"X,omitempty" bson:"X,omitempty"`
	Facebook       *Platform  `json:"facebook,omitempty" bson:"facebook,omitempty"`
	LinkedIn       *Platform  `json:"linkedin,omitempty" bson:"linkedin,omitempty"`
	OtherPlatforms []Platform `json:"other_platforms,omitempty" bson:"other_platforms,omitempty"` // Custom or unknown platforms
}

type Services struct {
	InstaReel     *Service  `json:"instareel,omitempty" bson:"instareel,omitempty"`
	InstaPost     *Service  `json:"instapost,omitempty" bson:"instapost,omitempty"`
	InstaStory    *Service  `json:"instastory,omitempty" bson:"instastory,omitempty"`
	YtPodcast     *Service  `json:"ytpodcast,omitempty" bson:"ytpodcast,omitempty"`
	YtShorts      *Service  `json:"ytshorts,omitempty" bson:"ytshorts,omitempty"`
	YtVideo       *Service  `json:"ytvideo,omitempty" bson:"ytvideo,omitempty"`
	OtherServices []Service `json:"other_services,omitempty" bson:"other_services,omitempty"`
}
type PersonalDetails struct {
	Name         string `json:"influencername" bson:"influencername" validate:"required,min=2, max=100"`
	Sex          string `json:"sex" bson:"sex" validate:"required"`
	Age          int    `json:"age" bson:"age" validate:"required"`
	LocationLink string `json:"location" bson:"location" validate:"required"`
}
