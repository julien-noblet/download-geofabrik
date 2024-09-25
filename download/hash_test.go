package download_test

import (
	"os"
	"testing"

	"github.com/julien-noblet/download-geofabrik/download"
)

func Test_hashFileMD5(t *testing.T) {
	t.Parallel()

	type args struct {
		filePath string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Check with LICENSE file",
			args:    args{filePath: "../LICENSE"},
			want:    "65d26fcc2f35ea6a181ac777e42db1ea",
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := download.ComputeMD5Hash(thisTest.args.filePath)
			if err != nil != thisTest.wantErr {
				t.Errorf("hashFileMD5(%v) error = %v, wantErr %v", thisTest.args.filePath, err, thisTest.wantErr)

				return
			}

			if got != thisTest.want {
				t.Errorf("hashFileMD5() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Benchmark_hashFileMD5_LICENSE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := download.ComputeMD5Hash("./LICENSE"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_controlHash_LICENSE(b *testing.B) {
	hash, _ := download.ComputeMD5Hash("./LICENSE")
	hashfile := "/tmp/download-geofabrik-test.hash"

	if err := os.WriteFile(hashfile, []byte(hash), 0o600); err != nil {
		b.Errorf("Can't write file %s err: %v", hashfile, err)
	}

	for n := 0; n < b.N; n++ {
		if _, err := download.CheckFileHash(hashfile, hash); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_controlHash(t *testing.T) {
	t.Parallel()

	type args struct {
		hashfile string
		hash     string
	}

	tests := []struct {
		args       args
		name       string
		fileToHash string
		want       bool
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name:       "Check with LICENSE file",
			fileToHash: "../LICENSE",
			args:       args{hashfile: "../LICENCE.md5", hash: "65d26fcc2f35ea6a181ac777e42db1ea"},
			want:       true,
			wantErr:    false,
		},
		{
			name:       "Check with LICENSE file wrong hash",
			fileToHash: "../LICENSE",
			args:       args{hashfile: "../LICENCE.md5", hash: "65d26fcc2f35ea6a181ac777e42db1eb"},
			want:       false,
			wantErr:    false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		hash, _ := download.ComputeMD5Hash(thisTest.fileToHash)

		hashfull := hash + " " + thisTest.fileToHash

		if err := os.WriteFile(thisTest.args.hashfile, []byte(hashfull), 0o600); err != nil {
			t.Errorf("can't write file %s err: %v", thisTest.args.hashfile, err)
		}

		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := download.CheckFileHash(thisTest.args.hashfile, thisTest.args.hash)
			if err != nil != thisTest.wantErr {
				t.Errorf("controlHash() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}

			if got != thisTest.want {
				t.Errorf("controlHash() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Test_VerifyFileChecksum(t *testing.T) {
	type args struct {
		outputPath string
		hashfile   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check with LICENSE file",
			args: args{outputPath: "../LICENSE", hashfile: "../LICENCE.md5"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := download.VerifyFileChecksum(tt.args.outputPath, tt.args.hashfile); got != tt.want {
				t.Errorf("VerifyChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
