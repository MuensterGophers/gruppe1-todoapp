package todo

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

type Controller struct {
	db *gorm.DB
}

func NewController() *Controller {
	db, err := gorm.Open("postgres", "postgres://postgres:foo123@localhost:5432/todo_db?sslmode=disable")

	// Sollten wir db.Close() aufrufen?

	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&Model{}).Error

	if err != nil {
		log.Println(err)
	}

	ret := &Controller{
		db: db,
	}

	return ret
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	ret := make([]Model, 0, 10);

	 err := c.db.Find(&ret).Error;

	if err != nil {
		fmt.Println(err)
		return
	}

	response, _ := json.Marshal(&ret)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	m := Model{}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &m)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = c.db.Create(&m).Error;

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	b, err := json.Marshal(&m)

	w.WriteHeader(201)
	w.Write(b)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := &Model{}
	m.ID = uint(id)

	err = c.db.Delete(&m).Error

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := Model{}

	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &m)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if uint(id) != m.ID {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ids don't match")
		return
	}

	err = c.db.Debug().Save(&m).Error

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}