package snippy

// Prompt is the interface implemented by an object that
// can interact with the user.
//
// Ask returns the snippet that the user has chosen.
type Prompt interface {
	Ask(snippets string) (Snippet, error)
}
