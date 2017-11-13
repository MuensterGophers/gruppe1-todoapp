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
	db, err := sql.Open("postgres", "postgres://postgres:foo123@localhost:5432/todo_db?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	ret := &Controller{
		db: db,
	}

	return ret
}

// Pagination einbauen? (zum Demonstrieren von slices)
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	rows, err := c.db.Query("SELECT id, message FROM todo")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	ret := make([]Model, 0, 10);

	for rows.Next() {
		model := Model{}

		err = rows.Scan(&model.Id, &model.Message)

		if err != nil {
			fmt.Println(err)
			continue
		}

		ret = append(ret, model)
	}

	//warum pointer? wegen interface?
	response, _ := json.Marshal(&ret)
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
