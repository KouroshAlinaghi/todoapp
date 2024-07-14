package todolist

import (
	"errors"
	"strconv"
	"strings"
)

type errNotFound struct {
	id      int
	content string
}

func (err errNotFound) Error() string {
	if err.id == 0 {
		return "Could not find todo item with content " + err.content
	} else {
		return "Could not find todo item with id " + strconv.Itoa(err.id)
	}
}

var (
	errDuplicateContent = errors.New("This todo is already present in the list.")
)

type List struct {
	Items  []*Item
	nextId int
}

type Item struct {
	Id       int
	Done     bool
	Selected bool
	Content  string
}

func NewList() *List {
	return &List{nextId: 1, Items: make([]*Item, 0)}
}

func (l *List) fetchById(id int) (int, *Item, error) {
	for index, item := range l.Items {
		if item.Id == id {
			return index, item, nil
		}
	}

	err := errNotFound{id: id, content: ""}
	return 0, nil, err
}

func (l *List) fetchByContent(content string) (*Item, error) {
	for _, item := range l.Items {
		if item.Content == content {
			return item, nil
		}
	}

	err := errNotFound{id: 0, content: content}
	return nil, err
}

func newItem(content string, id int) *Item {
	return &Item{Id: id, Done: false, Selected: false, Content: content}
}

func (l *List) Add(content string) error {
	_, err := l.fetchByContent(content)
	if err == nil {
		return errDuplicateContent
	}

	newItem := newItem(content, l.nextId)
	l.nextId++

	l.Items = append(l.Items, newItem)
	return nil
}

func (l *List) Remove(id int) error {
	index, _, err := l.fetchById(id)
	if err != nil {
		return err
	}

	l.Items = append(l.Items[:index], l.Items[index+1:]...)
	return nil
}

func (l *List) RemoveAll() {
	l.Items = make([]*Item, 0)
}

func (l *List) ToggleDoneSelecteds() {
	for _, item := range l.Items {
		if item.Selected {
			item.Done = !item.Done
		}
	}
}

func (l *List) RemoveSelecteds() {
	newItems := make([]*Item, 0)
	for _, item := range l.Items {
		if !item.Selected {
			newItems = append(newItems, item)
		}
	}
	l.Items = newItems
}

func (l *List) ToggleSelect(id int) error {
	_, item, err := l.fetchById(id)
	if err != nil {
		return err
	}

	item.Selected = !item.Selected
	return nil
}

func (l *List) ToggleDone(id int) error {
	_, item, err := l.fetchById(id)
	if err != nil {
		return err
	}

	item.Done = !item.Done
	return nil
}

func (i Item) String() string {
	doneStr := "✕"
	if i.Done {
		doneStr = "✓"
	}

	selStr := "⚪"
	if i.Selected {
		selStr = "⚫"
	}
	return selStr + " " + strconv.Itoa(i.Id) + ": " + i.Content + " " + doneStr
}

func (l List) String() string {
	result := "ToDo List:\n"
	maxLen := 0
	for _, item := range l.Items {
		if len(item.String()) > maxLen {
			maxLen = len(item.String())
		}
		result += item.String() + "\n"
	}
	result += strings.Repeat("-", maxLen)
	return result
}
