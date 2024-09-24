package download

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/apex/log"
)

func ControlHash(hashfile, hash string) (bool, error) {
	if FileExist(hashfile) {
		file, err := os.ReadFile(hashfile)
		if err != nil {
			return false, fmt.Errorf("can't read %s : %w", hashfile, err)
		}

		filehash := strings.Split(string(file), " ")[0]

		log.Debugf("Hash from file :%s", filehash)

		return strings.EqualFold(hash, filehash), nil
	}

	return false, nil
}

func HashFileMD5(filePath string) (string, error) {
	var returnMD5String string

	if FileExist(filePath) {
		hash := md5.New() //nolint:gosec // I use md5 to control with md5sum files

		file, err := os.Open(filePath)
		if err != nil {
			return returnMD5String, fmt.Errorf("can't open %s : %w", filePath, err)
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.WithError(err).Fatal("can't save file")
			}
		}()

		if _, err := io.Copy(hash, file); err != nil {
			return returnMD5String, fmt.Errorf("can't copy %s : %w", filePath, err)
		}

		hashInBytes := hash.Sum(nil)[:16]
		returnMD5String = hex.EncodeToString(hashInBytes)

		return returnMD5String, nil
	}

	return returnMD5String, nil
}
