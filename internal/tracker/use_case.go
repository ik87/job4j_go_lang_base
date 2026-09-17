package tracker

import (
	"strings"

	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.GetItems() {
		out.Out(item.toString())
	}
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("find: ")
	find := in.Get()
	for _, item := range tracker.GetItems() {
		if strings.Contains(strings.ToLower(item.toString()), strings.ToLower(find)) {
			out.Out(item.toString())
		}
	}
}

type DelUsecase struct{}

func (u DelUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid: ")
	uuid := in.Get()
	tracker.DelItem(uuid)
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter uuid: ")
	uuid := in.Get()
	out.Out("enter new name: ")
	name := in.Get()
	tracker.UpdateItem(Item{Name: name, ID: uuid})
}
