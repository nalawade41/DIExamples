package revisiondi

import "fmt"

// A `Story` contains some story content and an abstract storage reference.
type Story struct {
	content []byte
	storage StoryStorage
}

// `StoryStorage` is just an interface that defines the behavior of a Story storage.
// This is all that `Story` knows (and needs to know) about storing and retrieving Storys.
type StoryStorage interface {
	Type() string        // Return a string describing the storage type.
	Load(string) []byte  // Load a Story by name.
	Save(string, []byte) // Save a Story by name.
}

// `NewStory` constructs a `Story` object. We use this constructor to inject an object
// that satisfies the `StoryStorage` interface.
func NewStory(ps StoryStorage) *Story {
	return &Story{
		content: []byte("I am a Story from a " + ps.Type() + "."),
		storage: ps, // Storage created elsewhere and injected here
	}
}

// `Save` simply calls `Save` on the interface type. The `Story` object neither knows
// nor cares about which actual storage object receives this method call.
func (p *Story) Save(name string) {
	p.storage.Save(name, p.content)
}

// `Load` also invokes the injected storage object without knowing it.
func (p *Story) Load(name string) {
	p.content = p.storage.Load(name)
}

// `String` makes Story a Stringer, allowing us to drop it anywhere a string would be
// expected.
func (p *Story) String() string {
	return string(p.content)
}

// #### The notebook

// A `Notebook` is the classic storage device of a poet.
type Notebook struct {
	Storys map[string][]byte
}

func NewNotebook() *Notebook {
	return &Notebook{
		Storys: map[string][]byte{},
	}
}

// After adding `Save` and `Load`, `Notebook` implicitly satisfies `StoryStorage`.
func (n *Notebook) Save(name string, contents []byte) {
	n.Storys[name] = contents
}

func (n *Notebook) Load(name string) []byte {
	return n.Storys[name]
}

// `Type` returns an informal description of the storage type.
func (n *Notebook) Type() string {
	return "Notebook"
}

// A `Napkin` is the emergency storage device of a poet.
// It can store only one Story.
type Napkin struct {
	Story []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		Story: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.Story = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.Story
}

func (n *Napkin) Type() string {
	return "Napkin"
}

// Create and connect objects, then save and load a few Storys from different storage objects.
func storySaverProvider() {
	notebook := NewNotebook()
	napkin := NewNapkin()

	// First, write a Story into a notebook.
	// `NewStory()` injects the dependency.
	Story := NewStory(notebook)
	Story.Save("My first Story")

	// Create a new Story object to prove that the notebook storage works.
	Story = NewStory(notebook)
	Story.Load("My first Story")
	fmt.Println(Story)

	// Now we do the same with a napkin as storage.
	Story = NewStory(napkin)
	// Note the Story still just uses `Save` and `Load`. "Notebook? Napkin? I don't care."
	Story.Save("My second Story")
	Story = NewStory(napkin)
	Story.Load("My second Story")
	fmt.Println(Story)
}
