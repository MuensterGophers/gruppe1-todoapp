package todo

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

/**
 * - Wofür benötigen wir dieses struct?
 * - Ist das jetzt ein Singleton?
 */
type Controller struct {
	//repo
}

/**
 * Warum c *Controller?
 */
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	//c.repo.get
	model := Model{Id:123, Message:"Milch kaufen"}
	//warum pointer? wegen interface?
	response, _ := json.Marshal(&model)
	w.Header().Add("Content-Type","application/json")
	w.Write(response)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	m := Model{}

	// bufferovrflow möglich? bei
	//r.Body().Read()
	body , _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &m)

	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	m.Id = 123
	b, err := json.Marshal(&m)

	w.WriteHeader(201)
	w.Write(b)
}
