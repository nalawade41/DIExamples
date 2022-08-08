package storageService

type StoryNotebook struct{}

func NewStoryNotebook() *StoryNotebook {
	return &StoryNotebook{}
}
func (n *StoryNotebook) Save(name string, contents []byte) {
}

func (n *StoryNotebook) Load(name string) []byte {
	return []byte{}
}
