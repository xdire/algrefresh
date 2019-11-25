package util

import (
	"os"
	"testing"
	"time"
)



func TestCombinePattern(t *testing.T) {
	psPattern, err := CreatePrefixSuffixNamePattern([]string{"test."}, []string{".in", ".out"})
	psPattern2, err := CreatePrefixSuffixNamePattern([]string{"quests.", "test."}, []string{".out"})
	if err != nil {
		t.Error(err)
	}
	type args struct {
		files []File
		re    CompiledPattern
	}
	tests := []struct {
		name 			string
		args 			args
		prefixEquals1 	string
		prefixEquals2	string
		suffixEquals1	string
		suffixEquals2	string
	}{
		{
			"Prefix Suffix pattern combination 1",
			args{
				files: []File{
					{
						Path: "here1",
						Info: FileInfoFake{FakeName: "test.idx1.in"},
					},
					{
						Path: "here2",
						Info: FileInfoFake{FakeName: "test.idx1.out"},
					},
					{
						Path: "here3",
						Info: FileInfoFake{FakeName: "test.idx2.out"},
					},
					{
						Path: "here4",
						Info: FileInfoFake{FakeName: "test.idx2.in"},
					},
					{
						Path: "nothere5",
						Info: FileInfoFake{FakeName: "test.idx2.an"},
					},
					{
						Path: "nothere6",
						Info: FileInfoFake{FakeName: "quests.idx2.in"},
					},
				},
				re:    psPattern,
			},
			"test.",
			"",
			".in",
			".out",
		},
		{
			"Prefix Suffix pattern combination 2",
			args{
				files: []File{
					{
						Path: "here1",
						Info: FileInfoFake{FakeName: "test.idx1.in"},
					},
					{
						Path: "here2",
						Info: FileInfoFake{FakeName: "test.idx1.out"},
					},
					{
						Path: "here3",
						Info: FileInfoFake{FakeName: "test.idx2.out"},
					},
					{
						Path: "here4",
						Info: FileInfoFake{FakeName: "test.idx2.in"},
					},
					{
						Path: "nothere5",
						Info: FileInfoFake{FakeName: "test.idx2.an"},
					},
					{
						Path: "nothere6",
						Info: FileInfoFake{FakeName: "quests.idx2.out"},
					},
				},
				re:    psPattern2,
			},
			"test.",
			"quests.",
			".in",
			".out",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CombinePattern(tt.args.files, tt.args.re)
			t.Logf("\n%+v\n%+v", res["idx1"], res["idx2"])
			if len(res) != 2 {
				t.Errorf("Length should be equal 4 items for thest to pass")
			}
			for _, v := range res {
				for _, fileComb := range v.Files {
					if fileComb.PrefixSelector != tt.prefixEquals1 && fileComb.PrefixSelector != tt.prefixEquals2 {
						t.Errorf("Prefix should be propertly determined for test")
					}
					if fileComb.SuffixSelector != tt.suffixEquals1 && fileComb.SuffixSelector != tt.suffixEquals2 {
						t.Errorf("Suffix should be properly determined for test")
					}
				}
			}
		})
	}
}

type FileInfoFake struct {
	FakeName string
}

func (f FileInfoFake) Name() string {
	return f.FakeName
}

func (f FileInfoFake) Size() int64 {
	panic("implement me")
}

func (f FileInfoFake) Mode() os.FileMode {
	panic("implement me")
}

func (f FileInfoFake) ModTime() time.Time {
	panic("implement me")
}

func (f FileInfoFake) IsDir() bool {
	panic("implement me")
}

func (f FileInfoFake) Sys() interface{} {
	panic("implement me")
}