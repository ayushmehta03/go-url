package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)


type Url struct{
	Id string `json:"id"`
	OrignalUrl string `json:"orignal_url"`
	ShortUrl string `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}

var urlDb=make(map[string] Url)

func generateShortUrl(OrignalURL string) string{


	// intilaize the md5 module
	hahser:=md5.New()


	// pass the orignal url by converting into bytes cause hashing algorithm inly works for bytes
	hahser.Write([] byte(OrignalURL))


	//  give the final bytes
	data:=hahser.Sum(nil)


	// convert into readable hexadecimal string

	hash:=hex.EncodeToString(data);


	 return hash[:8]



}


func createURL( OrignalURL string) string {
	shortUrl:=generateShortUrl(OrignalURL)

	id:=shortUrl


	// save to the map with unqiue id key and value as struct
	urlDb[id]=Url{
		Id: id,
		OrignalUrl: OrignalURL,
		ShortUrl: shortUrl,
		CreatedAt: time.Now(),
	}



	return shortUrl;

}


func getURL(id string) (Url,error){
	url,ok:=urlDb[id]
	if !ok{
		return Url{},errors.New("no such url found")
	}
	return url,nil
}



func shortURLHandler(w http.ResponseWriter, r*http.Request){

 var data struct{
		URL string `json:"url"`
	}


	err:=json.NewDecoder(r.Body).Decode(&data)
	if err!=nil{
		http.Error(w,"Inavlid request body",http.StatusBadRequest)
	}

	shortUrl:=createURL(data.URL)


	response:= struct{
		ShortUrl string	`json:"short_url"`
	}{ShortUrl: shortUrl}


		w.Header().Set("Content-Type","application/json")

		json.NewEncoder(w).Encode(response)





}



func redirectUrlHandler( w http.ResponseWriter, r *http.Request){

	id:=r.URL.Path[len("/redirect/"):]

	url,err:=getURL(id)
	if err!=nil{
		http.Error(w,"Invalid Request",http.StatusBadRequest)


	}

	http.Redirect(w,r,url.OrignalUrl,http.StatusFound)
}


func main(){
 fmt.Println("starting.....")


	 
	http.HandleFunc("/shorten",shortURLHandler)

	http.HandleFunc("/redirect/",redirectUrlHandler)

err:=	http.ListenAndServe(":3000",nil)
if err!=nil{
fmt.Println("error on starting server")

}



}