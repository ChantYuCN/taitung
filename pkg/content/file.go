package content

import (
	"io/fs"
	"log"
	"path/filepath"
)

// func visit(root string, di fs.DirEntry, err error) error {
// 	log.Printf("Visited: %v\n", root)
// 	return nil
// }

func SearchTheFile(sn string, folder string) string {
	found := false
	var abspath string
	//err := filepath.WalkDir(folder, visit)
	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// handle possible path err, just in case...
			return err
		}
		if d.IsDir() && d.Name() == sn {
			log.Printf("File found: %v", path)
			abspath = path
			found = true
			return fs.SkipDir
		}
		return nil
	})
	if err != nil {
		log.Print(err)
		return ""
	}

	if !found {
		log.Print("Folder not found")
		return ""
	}
	return abspath
}
