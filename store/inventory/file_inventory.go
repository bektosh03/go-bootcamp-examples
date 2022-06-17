package inventory

import "os"

type FileInventory struct {
	db *os.File
}
