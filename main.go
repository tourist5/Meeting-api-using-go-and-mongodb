package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Connection mongoDB with helper class
collection := helper.ConnectDB()
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/meetings", getMeetings).Methods("GET")
	r.HandleFunc("/api/meetings/{id}", getMeeting).Methods("GET")
	r.HandleFunc("/api/meetings", createMeeting).Methods("POST")
	

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

	// Raw string representation of the MongoDB doc _id
idStr := "5d2399ef96fb765873a24bae"

// The MongoDB BSON ObjectID is essentially a byte slice with a length of 12
hexByte, err := hex.DecodeString(idStr)
fmt.Println("ID hexByte len:", len(hexByte))
fmt.Println("ID hexByte type:", reflect.TypeOf(hexByte))

// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
docID, err := primitive.ObjectIDFromHex(idStr)
fmt.Println("\nID ObjectIDFromHex:", docID)
fmt.Println("ID ObjectIDFromHex err:", err)
fmt.Println("ID hexByte type:", reflect.TypeOf(docID))

// Declare a struct instance of the MongoDB fields that will contain the document returned
result := Meeting{}
fmt.Println("\nresult type:", reflect.TypeOf(result))
fmt.Println("result BEFORE:", result)

// call the collection's Find() method and return Cursor object into result
fmt.Println(`bson.M{"_id": docID}:`, bson.M{"_id": docID})
err = col.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

// Check for any errors returned by MongoDB FindOne() method call
if err != nil {
fmt.Println("FindOne() ObjectIDFromHex ERROR:", err)
os.Exit(1)
} else {
// Print out data from the document result
fmt.Println("result AFTER:", result, "\n")

// Struct instances are objects with MongoDB fields that can be accessed as attributes
fmt.Println("FindOne() result:", result)
fmt.Printf("result doc ID: %v\n", result.ID)
fmt.Println("result.Title:", result.Title)
fmt.Println("result.Participants", result.Participants)
fmt.Println("result.Start time", result.Start time)
fmt.Println("result.End time", result.End time)
fmt.Println("result.Timestamp", result.Timestamp)
}

Meeting.inventory.find( {
	 tags: { $all: [ "mohan"  ] } 
  } )
}


func getMeetingss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var meetings []models.Meeting

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var meeting models.Meeting
		// & character returns the memory address of the following variable.
		err := cur.Decode(&meeting) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		meetings = append(meetings, meeting)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(meetings) // encode similar to serialize process.
}
func createMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&meeting)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
