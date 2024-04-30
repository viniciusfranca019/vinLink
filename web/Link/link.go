package Link

import "github.com/google/uuid"

type link struct {
	link, id string
}

func New(url string) link {
	id, err := uuid.NewV7()

	if err != nil {
		panic("error on new id")
	}

	return link{link: url, id: id.String()}
}

func (link link) Id() string {
	return link.id
}

