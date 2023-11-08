package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreateAt:    time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(idx int) error {

	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}
	ls[idx-1].CompletedAt = time.Now()
	ls[idx-1].Done = true
	return nil
}

func (t *Todos) Delete(idx int) error {

	ls := *t
	if idx <= 0 || idx > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:idx-1], ls[idx:]...) // get item before and everything after
	return nil
}

func (t *Todos) Load(fileName string) error {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(fileName string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}
