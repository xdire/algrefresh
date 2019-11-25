package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type File struct {
	Path string
	Info os.FileInfo
}

type FileCombination struct {
	File
	PrefixSelector	string
	SuffixSelector	string
}

type Combination struct {
	Files 			[]FileCombination
	Matched 		string
}

type CompiledPattern interface {
	Pattern() 		*regexp.Regexp
	PrefixesSize()	int
	SuffixesSize()	int
	Prefixes()		[]string
	Suffixes()		[]string
}

type CommonCompiledPattern struct {
	p *regexp.Regexp
	prefixList []string
	suffixList []string
}

func (c CommonCompiledPattern) Prefixes() []string {
	return c.prefixList
}

func (c CommonCompiledPattern) Suffixes() []string {
	return c.suffixList
}

func (c CommonCompiledPattern) Pattern() *regexp.Regexp {
	return c.p
}

func (c CommonCompiledPattern) PrefixesSize() int {
	if c.prefixList != nil {
		return len(c.prefixList)
	}
	return 0
}

func (c CommonCompiledPattern) SuffixesSize() int {
	if c.suffixList != nil {
		return len(c.suffixList)
	}
	return 0
}

func GetDirectoryFiles(directory string) ([]File, error) {
	files := make([]File, 0)
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, File{
			Path: path,
			Info: info,
		})
		return nil
	})
	if err != nil {
		return files, fmt.Errorf("cannot get directory files, %+v", err)
	}
	return files, nil
}

func CombinePattern(files []File, pattern CompiledPattern) map[string]*Combination {
	combs := make(map[string]*Combination)
	for _, v := range files {
		res := pattern.Pattern().FindStringSubmatch(v.Info.Name())
		if res != nil {
			str := ""
			suffix := ""
			prefix := ""
			if len(res) == 4 {
			// Both suffixes and prefixes were defined, take all of those
				prefix = res[1]
				str = res[2]
				suffix = res[3]
			} else if len(res) == 3 && pattern.PrefixesSize() - 1 > 0 {
			// Situation when its either prefix or suffix, and if prefixes
			// are defined then we taking the prefix as a found param
				prefix = res[1]
				str = res[2]
				if pattern.SuffixesSize() == 1 {
					suffix = pattern.Suffixes()[0]
				}
			} else if len(res) == 3 && pattern.SuffixesSize() - 1 > 0 {
			// Situation when its either prefix or suffix, and if suffixes
			// are defined then we taking the suffix as a found param
				suffix = res[2]
				str = res[1]
				if pattern.PrefixesSize() == 1 {
					prefix = pattern.Prefixes()[0]
				}
			} else if len(res) == 2 {
			// No prefix and suffix were defined for the search
				str = res[1]
			} else {
			// Probably just skip this kind of variant, probably that
			// can be named as undefined behavior
				continue
			}
			// If found populate
			if comb, ok := combs[str]; ok {
				comb.Files = append(comb.Files, FileCombination{
					File:           v,
					PrefixSelector: prefix,
					SuffixSelector: suffix,
				})
				continue
			}
			// Create if not found
			combs[str] = &Combination{
				Files:   []FileCombination{{
					File:           v,
					PrefixSelector: prefix,
					SuffixSelector: suffix,
				}},
				Matched: str,
			}
		}
	}
	return combs
}

func CreatePrefixSuffixNamePattern(prefixes []string, suffixes []string) (CompiledPattern, error) {
	prefixesCombine := make([]string, len(prefixes))
	for i, v := range prefixes {
		prefixesCombine[i] = fmt.Sprintf("(?:%s)", v)
	}
	suffixesCombine := make([]string, len(suffixes))
	for i, v := range suffixes {
		suffixesCombine[i] = fmt.Sprintf("(?:%s)", v)
	}
	p := ""
	s := ""
	// Combine prefixes
	if len(prefixesCombine) > 1 {
		p = "(" + strings.Join(prefixesCombine, "|") + ")"
	} else if len(prefixesCombine) > 0 {
		p = prefixesCombine[0]
	}
	// Combine suffixes
	if len(suffixesCombine) > 1 {
		s = "(" + strings.Join(suffixesCombine, "|") + ")"
	} else if len(suffixesCombine) > 0 {
		s = suffixesCombine[0]
	}
	// Compile
	var r *regexp.Regexp
	var err error
	// Output
	if len(p) > 0 && len(s) > 0{
		r, err = regexp.Compile("" +  p + "(.*)" + s + "")
	} else if len(p) > 0 {
		r, err = regexp.Compile("" +  p + "(.*)")
	} else if len(s) > 0 {
		r, err = regexp.Compile("(.*)" + s + "")
	}
	if err != nil {
		return nil, err
	}
	return CommonCompiledPattern{
		p:        r,
		prefixList: prefixes,
		suffixList: suffixes,
	}, nil
}
