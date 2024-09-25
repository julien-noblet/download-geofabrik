package download

import (
	"crypto/md5" //nolint:gosec // MD5 is used to control with md5sum files
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/apex/log"
)

const (
	readErrorMsg  = "can't read %s: %w"
	openErrorMsg  = "can't open %s: %w"
	copyErrorMsg  = "can't copy %s: %w"
	closeErrorMsg = "can't close file: %w"
)

// ControlHash checks if the hash of a file matches the provided hash.
func ControlHash(hashfile, expectedHash string) (bool, error) {
	if !FileExist(hashfile) {
		log.Infof("Hash file %s not found", hashfile)

		return false, nil
	}

	fileContent, err := os.ReadFile(hashfile)
	if err != nil {
		log.Infof("Can't read hash file %s", hashfile)

		return false, fmt.Errorf(readErrorMsg, hashfile, err)
	}

	fileHash := strings.Split(string(fileContent), " ")[0]
	log.Infof("Hash from file: %s", fileHash)

	return strings.EqualFold(expectedHash, fileHash), nil
}

// HashFileMD5 computes the MD5 hash of a file.
func HashFileMD5(filePath string) (string, error) {
	if !FileExist(filePath) {
		return "", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(openErrorMsg, filePath, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			e := fmt.Errorf(closeErrorMsg, err)
			log.WithError(err).Errorf(e.Error()) //nolint:govet // Need to use different log library to avoid this
		}
	}()

	hash := md5.New() //nolint:gosec // MD5 is used to control with md5sum files
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf(copyErrorMsg, filePath, err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// VerifyChecksum verifies the checksum of a file.
func VerifyChecksum(file, hashfile string) bool {
	log.Infof("Hashing %s", file)

	hashed, err := HashFileMD5(file)
	if err != nil {
		log.WithError(err).Fatal("can't hash file")
	}

	log.Debugf("MD5 : %s", hashed)

	ret, err := ControlHash(hashfile, hashed)
	if err != nil {
		log.WithError(err).Error("checksum error")
	}

	if ret {
		log.Infof("Checksum OK for %s", file)
	} else {
		log.Infof("Checksum MISMATCH for %s", file)
	}

	return ret
}
