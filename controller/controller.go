package controller

import (
	"BackendLinklyMedia/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://linklydb:linklydb@cluster0.ysi5axm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "linklymedia"
const (
	CollBillboards  = "billboards"
	CollInfluencers = "influencers"
)

// MOST IMPORTANT
var collection *mongo.Collection

// Connect MongoDB
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection Success")
	collection = client.Database(dbName).Collection(CollBillboards)
	fmt.Println("Collection reference is ready")
}
func insertOneBillboard(billboard model.Billboard) {
	inserted, err := collection.InsertOne(context.Background(), billboard)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 billboard in DB with id: ", inserted.InsertedID)
}

// DELETE  1 RECORD
func deleteOneBillboard(billboardID string) {
	id, _ := primitive.ObjectIDFromHex(billboardID)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count is ", deletecount)

}

// Delete many records
func deleteAllBillboards() int64 {
	filter := bson.D{{}}
	deleteresult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No. of billboards deleted id ", deleteresult.DeletedCount)
	return deleteresult.DeletedCount
}

// get all billboards
func getAllBillboards() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var billboards []primitive.M
	for cur.Next(context.Background()) {
		var billboard bson.M
		err := cur.Decode(&billboard)
		if err != nil {
			log.Fatal(err)
		}
		billboards = append(billboards, billboard)
	}
	defer cur.Close(context.Background())
	return billboards
}

// Actual controller-file
func GetAllbillboards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	allbillboards := getAllBillboards()
	json.NewEncoder(w).Encode(allbillboards)
}
func Createbillboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var billboard model.Billboard
	_ = json.NewDecoder(r.Body).Decode(&billboard)
	insertOneBillboard(billboard)
	json.NewEncoder(w).Encode(billboard)
}
func DeleteABillboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneBillboard(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllBillboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllBillboards()
	json.NewEncoder(w).Encode(count)

}

func UpdateBillboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedBillboard model.Billboard
	if err := json.NewDecoder(r.Body).Decode(&updatedBillboard); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"billboardname": updatedBillboard.Name,
		"priceperday":   updatedBillboard.PricePerDay,
		"dailyviews":    updatedBillboard.DailyViews,
		"dimensions":    updatedBillboard.Dimensions,
		"vacantfrom":    updatedBillboard.VacantFrom,
		"vacanttill":    updatedBillboard.VacantTill,
		"ownerdetails":  updatedBillboard.OwnerDetails,
		"description":   updatedBillboard.Description,
		"location":      updatedBillboard.LocationLink,
		"landmarks":     updatedBillboard.LandMarks,
		"discounts":     updatedBillboard.Discounts,
		"stars":         updatedBillboard.Stars,
	}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// ---------------------------------------------------INFLUENCERS-------------------------------------------------------------
func CreateInfluencer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-control-Allow-Methods", "POST")
	var influencer model.Influencers
	if err := json.NewDecoder(r.Body).Decode(&influencer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	inserted, err := collection.InsertOne(context.Background(), influencer)
	if err != nil {
		http.Error(w, "Failed to insert influencer", http.StatusInternalServerError)
		return
	}
	fmt.Println("Inserted 1 Influencer in DB with id: ", inserted.InsertedID)

}
func UpdateInfluencer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var updatedInfluencer *model.Influencers
	if err := json.NewDecoder(r.Body).Decode(&updatedInfluencer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"personal details": updatedInfluencer.PersonalDetails,
		"servicesoffered":  updatedInfluencer.Services,
		"platforms":        updatedInfluencer.Platforms,
		"ownerdetails":     updatedInfluencer.OwnerDetails,
		"description":      updatedInfluencer.Description,
		"discounts":        updatedInfluencer.Discounts,
		"stars":            updatedInfluencer.Stars,
	}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)

}
