package Services

import (
	"log"
	Db "Rentmatics_Appartment/Common/DB/Mysql"
	Model "Rentmatics_Appartment/Model"
	
	"fmt"
	"encoding/json"
	"net/http"
)



//Get All Home Detials without service id
func GetAllAppartments(w http.ResponseWriter, r *http.Request) {


	Data := Db.GetAllAppartmentDetails()
	fmt.Println("Datas",Data)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Println("Error on send home details", err)
	}
	w.WriteHeader(http.StatusOK)

	fmt.Println("header", w.Header())
	w.Write(Senddata)

}
func InsertContactProperty(w http.ResponseWriter, r *http.Request) {

	var GetContact Model.ContactProperty
	err := json.NewDecoder(r.Body).Decode(&GetContact)
	if err != nil {
		log.Println("Error on Get particular details", err)
	}
	//Connecting Database
	Data := Db.InsertConatct_DB(GetContact)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}