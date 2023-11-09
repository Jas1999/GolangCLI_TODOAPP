package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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
		CreatedAt:   time.Now(),
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

func (t *Todos) Print() {
	//for i, item := range *t {
	//	fmt.Printf("%d - %s \n", i, item.Task)
	//	i++
	// }

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for i, item := range *t {
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
		i++
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Tasks here"},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
