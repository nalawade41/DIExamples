package main

func main() {
	// notebook := revisiondi.NewNotebook()
	// napkin := revisiondi.NewNapkin()

	// // For testing purposes we can just create the the mock storage and mock the methods wihtout actually writing it to the storage
	// // We can use any kind of storeage now as long as we implement the storeage interface without making any changes in the existing code

	// // First, write a Story into a notebook.
	// // `NewStory()` injects the dependency.
	// Story := revisiondi.NewStory(notebook)
	// Story.Save("My first Story")

	// // Create a new Story object to prove that the notebook storage works.
	// Story = revisiondi.NewStory(notebook)
	// Story.Load("My first Story")
	// fmt.Println(Story)

	// // Now we do the same with a napkin as storage.
	// Story = revisiondi.NewStory(napkin)
	// // Note the Story still just uses `Save` and `Load`. "Notebook? Napkin? I don't care."
	// Story.Save("My second Story")
	// Story = revisiondi.NewStory(napkin)
	// Story.Load("My second Story")
	// fmt.Println(Story)
}

// Different types of DI
//1. Constructor this example consists one
//2. Using Reciever method (Mehtod injection)
//3. DI Using config
//4. JIT DI
