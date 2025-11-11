package main

import (
	"crypto/md5"
	"encoding/hex"
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

	urlDb[id]=Url{
		Id: id,
		OrignalUrl: OrignalURL,
		ShortUrl: shortUrl,
		CreatedAt: time.Now(),
	}

	return shortUrl;

}



func main(){
	fmt.Println("intial")

}