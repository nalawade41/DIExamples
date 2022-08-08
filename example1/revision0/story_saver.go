package exxample1

import "DIExamples/example1/storageService"

// Write a programe to save the story for a story  writer in document like notebook
type Story struct {
	content []byte
	storage *storageService.StoryNotebook
}

func NewStory() *Story {
	return &Story{
		storage: storageService.NewStoryNotebook(),
	}
}
func (p *Story) Load(title string) {
	p.content = p.storage.Load(title)
}
func (p *Story) Save(title string) {
	p.storage.Save(title, p.content)
}

// 1. Change requests -> Make it so story can be saved in napkin storage or in a word document or any other storage
