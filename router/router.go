package router

import (
	"BackendLinklyMedia/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/seller-dashboard/billboard/mylisting", controller.GetAllbillboards).Methods("GET")
	router.HandleFunc("/seller-dashboard/billboard", controller.Createbillboard).Methods("POST")
	//router.HandleFunc("/seller-dashboard/billboard/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/seller-dashboard/updatebillboard/{id}", controller.UpdateBillboard).Methods("PUT")
	router.HandleFunc("/seller-dashboard/billboard/{id}", controller.DeleteABillboard).Methods("DELETE")
	router.HandleFunc("/seller-dashboard/deleteallbillboards", controller.DeleteAllBillboard).Methods("DELETE")
	router.HandleFunc("/seller-dashboard/influencer", controller.CreateInfluencer).Methods("POST")
	router.HandleFunc("/seller-dashboard/updateinfluncer/{id}", controller.UpdateInfluencer).Methods("PUT")
	router.HandleFunc("/seller-dashboard/influencer/mylisting", controller.GetAllInfluencers).Methods("GET")
	router.HandleFunc("/seller-dashboard/influencer/agencyid/{id}", controller.GetInfluencersByAgencyID).Methods("GET")
	return router
}
