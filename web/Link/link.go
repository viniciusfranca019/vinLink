package Link

import "github.com/google/uuid"

type link struct {
	link, id string
}

func New(url string) link {
	return link{link: url, id: generateUUID()}
}

func generateUUID() string {
	id, err := uuid.NewV7()

	if err != nil {
		panic("error on new id")
	}

	return id.String()
}

func (link link) Id() string {
	return link.id
}

