package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fahimishraq99/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://FahimIshraq:18301077@cluster0.bux7kvd.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Student"
const colName = "todolist"

// Most Important
var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB helpers - file

//insert 1 record

func insertOneStudent(name model.Student) {
	inserted, err := collection.InsertOne(context.Background(), name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 student in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneStudent(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"studied": true}} //bson.M for shorter and clearer filter declaration

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneStudent(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Student got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllStudent() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of students deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all records from database

func getAllStudents() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var students []primitive.M

	for cur.Next(context.Background()) {
		var student bson.M
		err := cur.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	defer cur.Close(context.Background())
	return students
}

// Actual controller - file

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allStudents := getAllStudents()
	json.NewEncoder(w).Encode(allStudents)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	insertOneStudent(student)
	json.NewEncoder(w).Encode(student)

}

func MarkAsStudied(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneStudent(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneStudent(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllStudent()
	json.NewEncoder(w).Encode(count)
}
