package Mysql

import (
	"log"
	Model "Rentmatics_Appartment/Model"
		InitDb "Rentmatics_Appartment/Common/DB/Mysql/InitializeDB"
	 "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

	

var OpenConnection = make(map[string]*sql.DB)

func init() {
	OpenConnection = InitDb.Ret()
}


func GetAllAppartmentDetails() (Apres  []Model.AppartmentResponse ){
	var temp_Appartment Model.Appartment
	var temp_AppartRes Model.AppartmentResponse
	
	

	rows, err := OpenConnection["Rentmatics"].Query("Select * from Apartment")
	fmt.Println(err)

	for rows.Next() {
	rows.Scan(
			
	&temp_Appartment.Aid,                        
	&temp_Appartment.Apartmentname,	            
	&temp_Appartment.Streetname,	                
	&temp_Appartment.Areaname,	                
	&temp_Appartment.City,	                    
	&temp_Appartment.Pin,	                        
	&temp_Appartment.District,	                
	&temp_Appartment.State,	                    
	&temp_Appartment.Contactnumber,	            
	&temp_Appartment.KeyFeatures,	                
	&temp_Appartment.PropertyHighlights,	        
	&temp_Appartment.ApartmentAmenities,          
	&temp_Appartment.CommunityFeatures,	        
	&temp_Appartment.LocalSchools,	            
	&temp_Appartment.NearbyNeighborhoods,	        
	&temp_Appartment.AvailableFlats,             
	&temp_Appartment.PropertyDetails,             
	&temp_Appartment.OfficeHours,	                
	&temp_Appartment.DrivingDirections,           
	&temp_Appartment.ProfessionallyManagedBy,     
	)
	
	//Retrieve slice of Appartment images
	rows, err := OpenConnection["Rentmatics"].Query("select * from Apartmentimages where Apartmentid=?", temp_Appartment.Aid)
	fmt.Println(err)
	var Rentimgarray []Model.Appartmentimages
	for rows.Next() {
		var Rentimage Model.Appartmentimages
			rows.Scan(

					&Rentimage.Imageid,
					&Rentimage.Apparmentid,
					&Rentimage.Imageurl)
					
				Rentimgarray = append(Rentimgarray, Rentimage)
			}
			
		//Retrieve Ratings 
			row1, err := OpenConnection["Rentmatics"].Query("select * from Apartmentratings where Apartmentid=?", temp_Appartment.Aid)
			
			var Rentratingarray []Model.AppartmentRating

			for row1.Next() {
				var Rentrating Model.AppartmentRating

				row1.Scan(

					&Rentrating.RatingID,
					&Rentrating.Appartmentid,
					&Rentrating.Rating,
					&Rentrating.Date,
					&Rentrating.Comments,
					
					
					)

				Rentratingarray = append(Rentratingarray, Rentrating)
		}
		//Retrieve BHK1 
		row2, err := OpenConnection["Rentmatics"].Query("select * from Apartment_1BHK where Apatmentid=?", temp_Appartment.Aid)
			
			var Rent1bhkarray []Model.Appartment_BHK1

			for row2.Next() {
				var Rent1bhktemp  Model.Appartment_BHK1

				row2.Scan(

					&Rent1bhktemp.Bhk1id,
					&Rent1bhktemp.ApparmentID, 
					&Rent1bhktemp.Bed,	
					&Rent1bhktemp.Bath,	 
					&Rent1bhktemp.Minimumprice, 
					&Rent1bhktemp.Maximuxprice, 
					&Rent1bhktemp.Availability, 
					&Rent1bhktemp.SquareFeet, 
					&Rent1bhktemp.Deposit, 	
					
					
					)

				Rent1bhkarray = append(Rent1bhkarray, Rent1bhktemp)
		}
		
		//Retrieve 2 BHK 
		row3, err := OpenConnection["Rentmatics"].Query("select * from Apartment_2BHK where Apatmentid=?", temp_Appartment.Aid)
			
			var Rent2bhkarray []Model.Appartment_BHK2

			for row3.Next() {
				var Rent2bhktemp  Model.Appartment_BHK2

				row3.Scan(

					&Rent2bhktemp.Bhk2id,
					&Rent2bhktemp.ApparmentID, 
					&Rent2bhktemp.Bed,	
					&Rent2bhktemp.Bath,	 
					&Rent2bhktemp.Minimumprice, 
					&Rent2bhktemp.Maximuxprice, 
					&Rent2bhktemp.Availability, 
					&Rent2bhktemp.SquareFeet, 
					&Rent2bhktemp.Deposit, 	
					
					
					)

				Rent2bhkarray = append(Rent2bhkarray, Rent2bhktemp)
		}
		
		//Retrieve 3 BHK
		row4, err := OpenConnection["Rentmatics"].Query("select * from  Apartment_3BHK where Apatmentid=?", temp_Appartment.Aid)
			
			var Rent3bhkarray []Model.Appartment_BHK3

			for row4.Next() {
				var Rent3bhktemp  Model.Appartment_BHK3

				row4.Scan(

					&Rent3bhktemp.Bhk3id,
					&Rent3bhktemp.ApparmentID, 
					&Rent3bhktemp.Bed,	
					&Rent3bhktemp.Bath,	 
					&Rent3bhktemp.Minimumprice, 
					&Rent3bhktemp.Maximuxprice, 
					&Rent3bhktemp.Availability, 
					&Rent3bhktemp.SquareFeet, 
					&Rent3bhktemp.Deposit, 	
					)

				Rent3bhkarray = append(Rent3bhkarray, Rent3bhktemp)
		}
		
		
			temp_AppartRes.Apparment_Details = temp_Appartment
			temp_AppartRes.Appartment_image = Rentimgarray
			temp_AppartRes.Apparment_ratings = Rentratingarray
			temp_AppartRes.Appartment1BHk = Rent1bhkarray
			temp_AppartRes.Appartment2BHk = Rent2bhkarray
			temp_AppartRes.Appartment3BHk = Rent3bhkarray
			
			Apres = append(Apres,temp_AppartRes)
			}
			return 
			


	
}
func InsertConatct_DB(Contactvar Model.ContactProperty)string{
	rows, err := OpenConnection["Rentmatics"].Exec("insert into Appartment_contactproperty (Appartmetid,Firstname,LastName,Emailid,contactnumber,MoveDate,Bed,Bath,comments) values (?,?,?,?,?,?,?,?,?)", Contactvar.Appartmetid,Contactvar.Firstname,Contactvar.LastName,Contactvar.Emailid,Contactvar.Contactnumber,Contactvar.MoveDate,Contactvar.Bed,Contactvar.Bath,Contactvar.Comments)

	if err != nil {
		log.Println("Error -DB: Executive insert", err,rows)
	}
	return "success"
}
