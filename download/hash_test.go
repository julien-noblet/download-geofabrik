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
			args:    args{filePath: "./LICENSE"},
			want:    "65d26fcc2f35ea6a181ac777e42db1ea",
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := download.HashFileMD5(thisTest.args.filePath)
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
		if _, err := download.HashFileMD5("./LICENSE"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_controlHash_LICENSE(b *testing.B) {
	hash, _ := download.HashFileMD5("./LICENSE")
	hashfile := "/tmp/download-geofabrik-test.hash"

	if err := os.WriteFile(hashfile, []byte(hash), 0o600); err != nil {
		b.Errorf("Can't write file %s err: %v", hashfile, err)
	}

	for n := 0; n < b.N; n++ {
		if _, err := download.ControlHash(hashfile, hash); err != nil {
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
			args:       args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1ea"},
			want:       true,
			wantErr:    false,
		},
		{
			name:       "Check with LICENSE file wrong hash",
			fileToHash: "../LICENSE",
			args:       args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1eb"},
			want:       false,
			wantErr:    false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		hash, _ := download.HashFileMD5(thisTest.fileToHash)

		if err := os.WriteFile(thisTest.args.hashfile, []byte(hash), 0o600); err != nil {
			t.Errorf("can't write file %s err: %v", thisTest.args.hashfile, err)
		}

		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := download.ControlHash(thisTest.args.hashfile, thisTest.args.hash)
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
