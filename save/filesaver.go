package save

import "io/ioutil"

// FileSaver simple saver to file
type FileSaver struct {
}

// Save method to save byte array into a file
func (fs *FileSaver) Save(content []byte, file string) error {
	err := ioutil.WriteFile(file, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
