// This keeps track of what's going on in the open documents (the state)
// basically like a compiler
// In a proper LSP for a language, you would use some tooling
// for the particular language

package analysis

type State struct {
	// Maps filenames to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}
