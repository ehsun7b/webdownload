package save

// Saver interface
type Saver interface {
	Save([]byte, string) error
}
