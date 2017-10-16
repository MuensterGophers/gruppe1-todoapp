package todo

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

/**
 * - Wofür benötigen wir dieses struct?
 * - Ist das jetzt ein Singleton?
 */
type Controller struct {
	db *sql.DB
}

func NewController() *Controller {
	db, err := sql.Open("postgres", "postgres://postgres:foo123@10.0.1.13:5432/todo?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	ret := &Controller{
		db: db,
	}

	return ret
}

/**
 * Warum c *Controller?
 */
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	// error nicht ignorieren
	rows, err := c.db.Query("SELECT id, message FROM todo")

	fmt.Println(err)

	defer rows.Close()

	// Todo: Über alle Ergebnisse iterieren
	rows.Next()

	var id int
	var message string

	rows.Scan(&id, &message)

	//c.repo.get
	model := Model{Id: id, Message: message}
	//warum pointer? wegen interface?
	response, _ := json.Marshal(&model)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	m := Model{}

	// bufferovrflow möglich? bei
	// r.Body().Read()
	body, _ := ioutil.ReadAll(r.Body)
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
