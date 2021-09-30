package restful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"` //bu işaretleri koymak için Alt Gr+virgüle basıp sonrasında herhangi bi tuşa basmak gerekiyor
	Id        int    `json:"id`
	Title     string `json:"title`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}
