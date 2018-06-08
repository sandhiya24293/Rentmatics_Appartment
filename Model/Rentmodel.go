package Model


type Appartment struct{

    Aid                         string
	Apartmentname	            string
	Streetname	                string
	Areaname	                string
	City	                    string
	Pin	                        string
	District	                string
	State	                    string
	Contactnumber	            string
	KeyFeatures	                string
	PropertyHighlights	        string
	ApartmentAmenities          string
	CommunityFeatures	        string
	LocalSchools	            string
	NearbyNeighborhoods	        string
	AvailableFlats              int
	PropertyDetails             string
	OfficeHours	                string
	DrivingDirections           string
	ProfessionallyManagedBy     string
	
}
type Appartmentimages struct{
	Imageid int
	Apparmentid int
	Imageurl string
}
type AppartmentRating struct{
	RatingID int
	Appartmentid int
	Rating int
	Date string
	Comments string
	
}
type Appartment_BHK1 struct{
	Bhk1id int
	ApparmentID int
	Bed	int
	Bath	 int
	Minimumprice int
	Maximuxprice int
	Availability int 
	SquareFeet string
	Deposit int	
}
type Appartment_BHK2 struct{
	Bhk2id int
	ApparmentID int
	Bed	int
	Bath	 int
	Minimumprice int
	Maximuxprice int
	Availability int 
	SquareFeet string
	Deposit int	
}
type Appartment_BHK3 struct{
	Bhk3id int
	ApparmentID int
	Bed	int
	Bath	 int
	Minimumprice int
	Maximuxprice int
	Availability int 
	SquareFeet string
	Deposit int	
}

type AppartmentResponse struct{
	Apparment_Details Appartment
	Appartment_image []Appartmentimages
	Apparment_ratings []AppartmentRating
	Appartment1BHk []Appartment_BHK1
	Appartment2BHk []Appartment_BHK2
	Appartment3BHk []Appartment_BHK3
	
	
}

type ContactProperty struct{
		
		Appartmetid	          int
		Firstname	          string
		LastName	          string
		Emailid	              string
		Contactnumber	      string
		MoveDate	          string
		Bed	                  int
		Bath                  int
		Comments              string
	
}