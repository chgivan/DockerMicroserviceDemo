package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Message struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}
type Messages []Message

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

func main(){	
	dbSession, err := mgo.Dial("mongo")
    if err != nil {
        panic(err)
    }
    defer dbSession.Close()

    dbSession.SetMode(mgo.Monotonic, true)
    EnsureIndex(dbSession)	
	
	var routes = Routes{
		Route{"Index","GET","/",Index},
		Route{"GetMessages","GET","/messages",GetAllMessages(dbSession)},
		Route{"GetMessage","GET","/messages/{messageId}",GetMessage(dbSession)},
		Route{"CreateMessage","POST","/messages",CreateMessage(dbSession)},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	log.Fatal(http.ListenAndServe(":80",router))
}

func EnsureIndex(s *mgo.Session){
	session := s.Copy()
    defer session.Close()

    c := GetMessageDBCollection(session)

    index := mgo.Index{
        Key:        []string{"id"},
        Unique:     true,
        DropDups:   true,
        Background: true,
        Sparse:     true,
    }

    err := c.EnsureIndex(index)
    if err != nil {
        panic(err)
    }
}

func Index(w http.ResponseWriter, r *http.Request){
	sendJson(w, "Welcome!", http.StatusOK)	
}

func GetAllMessages(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		session := s.Copy()
        defer session.Close()

        c := GetMessageDBCollection(session)

        var messages Messages
        err := c.Find(bson.M{}).All(&messages)
        if err != nil {
            sendJson(w, "Database error", http.StatusInternalServerError)
            log.Println("Failed get all messages: ", err)
            return
        }
	
		if messages != nil {
 			sendJson(w, messages, http.StatusOK)
		}else{
			sendJson(w, "Messages does not exists", http.StatusNotFound)
		}       
	}
}

func CreateMessage(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
return func (w http.ResponseWriter, r *http.Request){
	session := s.Copy()
	defer session.Close()
	
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		sendJson(w, "Incorrect Body", http.StatusBadRequest)
		return
	}

	c := GetMessageDBCollection(session)
	
	err = c.Insert(message)
	if err != nil {
		if mgo.IsDup(err){
			sendJson(w, "Message with this ID already exists", http.StatusBadRequest)
			return
		}

		sendJson(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed insert message: ", err)
		return
	}
	
	sendJson(w, message, http.StatusCreated)
}}

func GetMessage(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
return func (w http.ResponseWriter, r *http.Request){
	session := s.Copy()
	defer session.Close()

	vars := mux.Vars(r)
	messageId := vars["messageId"]
	
	
	c := GetMessageDBCollection(session)
	
	var message Message
	err := c.Find(bson.M{"id": messageId}).One(&message)
	if err != nil {
		if(err.Error() == "not found"){
			sendJson(w, "Message not found", http.StatusNotFound)
			return
		}
		
		sendJson(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed find message: ", err)
		return
	}

	sendJson(w, message, http.StatusOK)
}}

func GetMessageDBCollection (s *mgo.Session) *mgo.Collection {
	return s.DB("messageDB").C("messages")
}

func sendJson(w http.ResponseWriter,body interface{}, status int){
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(body);	
	if  err != nil {
		panic(err)	
	}
}
