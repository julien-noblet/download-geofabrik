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
	readErrorMsg         = "can't read %s: %w"
	openErrorMsg         = "can't open %s: %w"
	copyErrorMsg         = "can't copy %s: %w"
	closeErrorMsg        = "can't close file: %w"
	hashFileNotFoundMsg  = "Hash file %s not found"
	hashFileReadErrorMsg = "Can't read hash file %s"
	hashMismatchMsg      = "Checksum MISMATCH for %s"
	hashMatchMsg         = "Checksum OK for %s"
	hashingFileMsg       = "Hashing %s"
	md5HashMsg           = "MD5 : %s"
	checksumErrorMsg     = "checksum error"
	hashFileErrorMsg     = "can't hash file"
)

// CheckFileHash checks if the hash of a file matches the provided hash.
func CheckFileHash(hashfile, expectedHash string) (bool, error) {
	if !FileExist(hashfile) {
		log.Infof(hashFileNotFoundMsg, hashfile)
		return false, nil
	}

	fileContent, err := os.ReadFile(hashfile)
	if err != nil {
		log.Infof(hashFileReadErrorMsg, hashfile)
		return false, fmt.Errorf(readErrorMsg, hashfile, err)
	}

	fileHash := strings.Split(string(fileContent), " ")[0]
	log.Infof("Hash from file: %s", fileHash)

	return strings.EqualFold(expectedHash, fileHash), nil
}

// ComputeMD5Hash computes the MD5 hash of a file.
func ComputeMD5Hash(filePath string) (string, error) {
	if !FileExist(filePath) {
		return "", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(openErrorMsg, filePath, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.WithError(err).Errorf(fmt.Errorf(closeErrorMsg, err).Error())
		}
	}()

	hash := md5.New() //nolint:gosec // MD5 is used to control with md5sum files
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf(copyErrorMsg, filePath, err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// VerifyFileChecksum verifies the checksum of a file.
func VerifyFileChecksum(file, hashfile string) bool {
	log.Infof(hashingFileMsg, file)

	hashed, err := ComputeMD5Hash(file)
	if err != nil {
		log.WithError(err).Fatal(hashFileErrorMsg)
	}

	log.Debugf(md5HashMsg, hashed)

	ret, err := CheckFileHash(hashfile, hashed)
	if err != nil {
		log.WithError(err).Error(checksumErrorMsg)
	}

	if ret {
		log.Infof(hashMatchMsg, file)
	} else {
		log.Infof(hashMismatchMsg, file)
	}

	return ret
}
