package blockchain

import "os"

type FileDownloader func(src string) (*os.File, error)
