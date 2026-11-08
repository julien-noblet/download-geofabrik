package downloader

import (
	"crypto/md5" //nolint:gosec // MD5 is used to control with md5sum files
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	readErrorMsg = "cannot read %s: %w"
	openErrorMsg = "cannot open %s: %w"
	copyErrorMsg = "cannot copy %s: %w"
)

// ErrChecksumMismatch indicates that a file's calculated hash differs from its published checksum.
var ErrChecksumMismatch = errors.New("checksum mismatch")

// CheckFileHash checks if the hash of a file matches the provided hash.
func CheckFileHash(hashfile, expectedHash string) (bool, error) {
	if !FileExists(hashfile) {
		slog.Warn("Hash file not found", "file", hashfile)

		return false, nil
	}

	fileContent, err := os.ReadFile(hashfile)
	if err != nil {
		return false, fmt.Errorf(readErrorMsg, hashfile, err)
	}

	fields := strings.Fields(string(fileContent))
	if len(fields) == 0 {
		slog.Warn("Hash file is empty", "file", hashfile)

		return false, nil
	}

	fileHash := fields[0]
	slog.Info("Hash from file", "hash", fileHash)

	if !strings.EqualFold(expectedHash, fileHash) {
		return false, fmt.Errorf("%w: %s", ErrChecksumMismatch, hashfile)
	}

	return true, nil
}

// ComputeMD5Hash computes the hexadecimal MD5 checksum of the file at filePath.
// If the file does not exist, it returns ("", nil) without error.
func ComputeMD5Hash(filePath string) (string, error) {
	if !FileExists(filePath) {
		return "", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(openErrorMsg, filePath, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("Failed to close file", "file", filePath, "error", err)
		}
	}()

	bufPtr, ok := copyBufPool.Get().(*[]byte)
	if !ok {
		buf := make([]byte, streamBufferSize)
		bufPtr = &buf
	}

	defer copyBufPool.Put(bufPtr)

	hash := md5.New() //nolint:gosec // MD5 is used to control with md5sum files
	if _, err := io.CopyBuffer(hash, file, *bufPtr); err != nil {
		return "", fmt.Errorf(copyErrorMsg, filePath, err)
	}

	var digest [md5.Size]byte

	sum := hash.Sum(digest[:0])

	return hex.EncodeToString(sum), nil
}

// VerifyFileChecksum verifies the checksum of a file.
func VerifyFileChecksum(file, hashfile string) (bool, error) {
	slog.Info("Hashing file", "file", file)

	hashed, err := ComputeMD5Hash(file)
	if err != nil {
		return false, fmt.Errorf("hashing file %s: %w", file, err)
	}

	slog.Debug("MD5 Hash", "hash", hashed)

	ret, err := CheckFileHash(hashfile, hashed)
	if err != nil {
		return false, fmt.Errorf("checking checksum for %s: %w", file, err)
	}

	if ret {
		slog.Info("Checksum OK", "file", file)
	}

	return ret, nil
}
