package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
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


	fmt.Println(hash[:8])


	return "ayush"
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


func main(){
	fmt.Println("intial")

}