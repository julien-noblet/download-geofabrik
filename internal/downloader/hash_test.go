package download_test

import (
	"os"
	"path/filepath"
	"testing"

	download "github.com/julien-noblet/download-geofabrik/internal/downloader"
)

func Test_hashFileMD5(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	emptyFile := filepath.Join(tmpDir, "empty.txt")
	requireNoError(t, os.WriteFile(emptyFile, []byte(""), 0o600))

	customFile := filepath.Join(tmpDir, "hello.txt")
	requireNoError(t, os.WriteFile(customFile, []byte("hello world"), 0o600))

	tests := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{
			name:     "Check with LICENSE file",
			filePath: "../../LICENSE",
			want:     "65d26fcc2f35ea6a181ac777e42db1ea",
			wantErr:  false,
		},
		{
			name:     "Check empty file MD5",
			filePath: emptyFile,
			want:     "d41d8cd98f00b204e9800998ecf8427e",
			wantErr:  false,
		},
		{
			name:     "Check custom content MD5",
			filePath: customFile,
			want:     "5eb63bbbe01eeed093cb22bb8f5acdc3",
			wantErr:  false,
		},
		{
			name:     "Check non existent file",
			filePath: filepath.Join(tmpDir, "non_existent.txt"),
			want:     "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := download.ComputeMD5Hash(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeMD5Hash(%v) error = %v, wantErr %v", tt.filePath, err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("ComputeMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Benchmark_hashFileMD5_LICENSE(b *testing.B) {
	for range b.N {
		if _, err := download.ComputeMD5Hash("../../LICENSE"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_controlHash_LICENSE(b *testing.B) {
	hash, _ := download.ComputeMD5Hash("../../LICENSE")
	tmpDir := b.TempDir()
	hashfile := filepath.Join(tmpDir, "test.hash")

	if err := os.WriteFile(hashfile, []byte(hash), 0o600); err != nil {
		b.Errorf("Can't write file %s err: %v", hashfile, err)
	}

	for range b.N {
		if _, err := download.CheckFileHash(hashfile, hash); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_controlHash(t *testing.T) {
	tmpDir := t.TempDir()
	validHashFile := filepath.Join(tmpDir, "valid.md5")
	emptyHashFile := filepath.Join(tmpDir, "empty.md5")
	whitespaceHashFile := filepath.Join(tmpDir, "space.md5")

	requireNoError(t, os.WriteFile(emptyHashFile, []byte(""), 0o600))
	requireNoError(t, os.WriteFile(whitespaceHashFile, []byte("   \n\t  "), 0o600))

	hash, _ := download.ComputeMD5Hash("../../LICENSE")
	hashfull := hash + "  ../../LICENSE\n"
	requireNoError(t, os.WriteFile(validHashFile, []byte(hashfull), 0o600))

	tests := []struct {
		name     string
		hashfile string
		hash     string
		want     bool
		wantErr  bool
	}{
		{
			name:     "Check with LICENSE file valid hash",
			hashfile: validHashFile,
			hash:     "65d26fcc2f35ea6a181ac777e42db1ea",
			want:     true,
			wantErr:  false,
		},
		{
			name:     "Check case insensitive matching",
			hashfile: validHashFile,
			hash:     "65D26FCC2F35EA6A181AC777E42DB1EA",
			want:     true,
			wantErr:  false,
		},
		{
			name:     "Check with wrong hash",
			hashfile: validHashFile,
			hash:     "65d26fcc2f35ea6a181ac777e42db1eb",
			want:     false,
			wantErr:  false,
		},
		{
			name:     "Check non existent hashfile",
			hashfile: filepath.Join(tmpDir, "missing.md5"),
			hash:     "65d26fcc2f35ea6a181ac777e42db1ea",
			want:     false,
			wantErr:  false,
		},
		{
			name:     "Check empty hashfile",
			hashfile: emptyHashFile,
			hash:     "65d26fcc2f35ea6a181ac777e42db1ea",
			want:     false,
			wantErr:  false,
		},
		{
			name:     "Check whitespace hashfile",
			hashfile: whitespaceHashFile,
			hash:     "65d26fcc2f35ea6a181ac777e42db1ea",
			want:     false,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := download.CheckFileHash(tt.hashfile, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckFileHash() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("CheckFileHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_VerifyFileChecksum(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	dataFile := filepath.Join(tmpDir, "data.bin")
	hashFile := filepath.Join(tmpDir, "data.bin.md5")
	wrongHashFile := filepath.Join(tmpDir, "wrong.bin.md5")

	requireNoError(t, os.WriteFile(dataFile, []byte("data to verify"), 0o600))
	hash, err := download.ComputeMD5Hash(dataFile)
	requireNoError(t, err)

	requireNoError(t, os.WriteFile(hashFile, []byte(hash+"  data.bin\n"), 0o600))
	requireNoError(t, os.WriteFile(wrongHashFile, []byte("00000000000000000000000000000000  data.bin\n"), 0o600))

	tests := []struct {
		name       string
		outputPath string
		hashfile   string
		want       bool
	}{
		{
			name:       "Valid checksum",
			outputPath: dataFile,
			hashfile:   hashFile,
			want:       true,
		},
		{
			name:       "Mismatching checksum",
			outputPath: dataFile,
			hashfile:   wrongHashFile,
			want:       false,
		},
		{
			name:       "Missing hashfile",
			outputPath: dataFile,
			hashfile:   filepath.Join(tmpDir, "missing.md5"),
			want:       false,
		},
		{
			name:       "Missing data file",
			outputPath: filepath.Join(tmpDir, "missing.bin"),
			hashfile:   hashFile,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := download.VerifyFileChecksum(tt.outputPath, tt.hashfile); got != tt.want {
				t.Errorf("VerifyFileChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
