package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strconv"
)

type SendingMessageT struct {
	Token   string  `json:"token"`
	Message string  `json:"message"`
	Users   []int   `json:"users"`
}

type SendingPhotoT struct {
	Token   string  `json:"token"`
	Photo string  `json:"photo"`
	Users   []int   `json:"users"`
}

type SendingLocationT struct {
	Token   string  `json:"token"`
	Latitude float32  `json:"latitude"`
	Longitude float32  `json:"longitude"`
	Users   []int   `json:"users"`
}

func main() {

  const baseUrl string = "https://api.telegram.org/bot"	

  //эндпоинт отправки сообщений
  http.HandleFunc("/sendMessage", func(w http.ResponseWriter, r *http.Request) {

	//считываю тело запроса (формат json)
	body, _ := ioutil.ReadAll(r.Body)

	//парсю его в структуру ( сообщение | токен | массив id)
	sending := SendingMessageT{}
	json.Unmarshal(body,&sending)

	//в цикле делаю запросы к телеграму с подставляя id текст и токен
	for i:=0; i < len(sending.Users); i++ {
        go sendMessage(baseUrl, sending.Token, sending.Message, sending.Users[i])
	}

  })
  

  //эндпоинт отправки фото
  http.HandleFunc("/sendPhoto", func(w http.ResponseWriter, r *http.Request) {

	//считываю тело запроса (формат json)
	body, _ := ioutil.ReadAll(r.Body)

	//парсю его в структуру ( сообщение | токен | массив id)
	sending := SendingPhotoT{}
	json.Unmarshal(body,&sending)


	//в цикле делаю запросы к телеграму с подставляя id текст и токен
	for i:=0; i < len(sending.Users); i++ {
        go sendPhoto(baseUrl,sending.Token, sending.Photo, sending.Users[i])
	}
  })

   //эндпоинт отправки местоположения
   http.HandleFunc("/sendLocation", func(w http.ResponseWriter, r *http.Request) {

	//считываю тело запроса (формат json)
	body, _ := ioutil.ReadAll(r.Body)

	//парсю его в структуру ( сообщение | токен | массив id)
	sending := SendingLocationT{}
	json.Unmarshal(body,&sending)


	//в цикле делаю запросы к телеграму с подставляя id текст и токен
	for i:=0; i < len(sending.Users); i++ {
        go sendLocation(baseUrl,sending.Token, sending.Latitude, sending.Longitude, sending.Users[i])
	}
  })

  http.ListenAndServe(":8010", nil)
}

/**
 * Функция для отправки сообщение в телеграм
 */
func sendMessage(baseUrl string, token string, message string, userId int) {
	url := baseUrl + token + "/sendMessage?parse_mode=HTML&chat_id="  + strconv.Itoa(userId)  +  "&text=" + message
    http.Get(url);
}


/**
 * Функция для отправки фото в телеграм
 */
func sendPhoto(baseUrl string, token string, photo string, userId int) {
	url := baseUrl + token + "/sendPhoto?chat_id="  + strconv.Itoa(userId)  +  "&photo=" + photo
	http.Get(url);
}

/**
 * Функция для отправки локации в телеграм
 */
func sendLocation(baseUrl string, token string, latitude float32, longitude float32, userId int){
    url := baseUrl + token + "/sendLocation?chat_id="  + strconv.Itoa(userId)  + "&latitude=" + fmt.Sprint(latitude)  +  "&longitude=" + fmt.Sprint(longitude)
    http.Get(url);
}