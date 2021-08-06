package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func apipath(base string) string {
	return filepath.Join(os.Getenv("GOROOT"), "api", base)
}

//pkg syscall (windows-386), const CERT_E_CN_NO_MATCH = 2148204815
var reSym = regexp.MustCompile(`^pkg (\S+)\s?(.*)?, (?:(var|func|type|const|method))\s?(\((\*?[A-Z]\w*)\))? ([A-Z]\w*)`)

// pkg reflect, type Type interface, Align
var reInterface = regexp.MustCompile(`^pkg (\S+)\s?(.*)?, type ([A-Z]\w*) interface, ([A-Z]\w*)`)

// var reMethod = regexp.MustCompile(`^pkg (\S+)\s?(.*)?, method (\((\*?[A-Z]\w*)\))? ([A-Z]\w*)`) // (\(\*?[A-Z]\w*\)) ([A-Z]\w*)`)
// var reNum = regexp.MustCompile(`^\-?[0-9]+$`)

type ApiInfo struct {
	Name       string   // name key
	Pkg        string   // pkgpath
	Kind       string   // var|const|func|type|method|
	MethodType string   // type or empty
	MethodName string   // name or empty
	MethodPtr  string   // * or emtpy
	Tags       []string // (os-arch-cgo)
	Ver        string   // go1.15 | go1.16 or empty
}

func (i ApiInfo) String() string {
	info := fmt.Sprintf("%v %v.%v", i.Kind, i.Pkg, i.Name)
	if i.Kind == "method" {
		info = fmt.Sprintf("%v (%v%v.%v).%v", i.Kind, i.MethodPtr, i.Pkg, i.MethodType, i.MethodName)
	}
	if i.Ver != "" {
		info += " (" + i.Ver + ")"
	}
	if len(i.Tags) > 1 {
		info += " " + strings.Join(i.Tags[1:], ",")
	}
	return info
}

type GoApi struct {
	Keys map[string]*ApiInfo // name - os-arch-cgo
	Ver  string
}

func (c *GoApi) List() []string {
	var list []string
	for k, _ := range c.Keys {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func LoadApi(ver string, tagVer bool) (*GoApi, error) {
	f, err := os.Open(apipath(ver + ".txt"))
	if err != nil {
		return nil, err
	}
	var verInfo string
	if tagVer {
		verInfo = ver
	}
	sc := bufio.NewScanner(f)
	keys := make(map[string]*ApiInfo)
	toTag := func(tag string) string {
		if tag == "" {
			return "-"
		}
		return tag
	}
	for sc.Scan() {
		l := sc.Text()
		has := func(v string) bool { return strings.Contains(l, v) }
		if has("struct, ") {
			continue
		}
		if has("interface, ") {
			if m := reInterface.FindStringSubmatch(l); m != nil {
				// m[1] pkg
				// m[2] os-arch-cgo
				// m[3] name
				// m[4] method
				pkg := m[1]
				tag := m[2]
				mtype := m[3]
				mname := m[4]
				key := pkg + "." + mtype
				if _, ok := keys[key]; !ok {
					keys[key] = &ApiInfo{
						Pkg:  pkg,
						Name: mtype,
						Kind: "type",
						Tags: []string{toTag(tag)},
						Ver:  verInfo,
					}
				}
				key = pkg + "." + mtype + "." + mname
				keys[key] = &ApiInfo{
					Pkg:        pkg,
					Name:       mtype + "." + mname,
					Kind:       "method",
					MethodType: mtype,
					MethodName: mname,
					Tags:       []string{toTag(tag)},
					Ver:        verInfo,
				}
			}
			continue
		}
		if m := reSym.FindStringSubmatch(l); m != nil {
			// m[1] pkg
			// m[2] os-arch-cgo
			// m[3] var|func|type|const|method
			// m[4] (*?Type)
			// m[5] *?Type
			// m[6] name
			pkg := m[1]
			tag := m[2]
			if tag != "" {
				tag = tag[1 : len(tag)-1]
			}
			kind := m[3]
			mtype := m[5]
			name := m[6]
			var mptr string
			var mname string
			if kind == "method" {
				mname = name
				if mtype[0] == '*' {
					mtype = mtype[1:]
					mptr = "*"
				}
				name = mtype + "." + name
			}
			key := pkg + "." + name
			if api, ok := keys[key]; ok {
				if api.Tags[0] != "-" {
					api.Tags = append(api.Tags, toTag(tag))
				}
				continue
			}
			keys[key] = &ApiInfo{
				Pkg:        pkg,
				Name:       name,
				Kind:       kind,
				MethodType: mtype,
				MethodName: mname,
				MethodPtr:  mptr,
				Tags:       []string{toTag(tag)},
				Ver:        verInfo,
			}
		}
	}
	return &GoApi{Ver: ver, Keys: keys}, nil
}

type ApiCheck struct {
	Base map[string]*ApiInfo
	Apis []*GoApi
}

func NewApiCheck() *ApiCheck {
	ac := &ApiCheck{}
	ac.Base = make(map[string]*ApiInfo)
	return ac
}

func (ac *ApiCheck) LoadBase(vers ...string) error {
	for _, ver := range vers {
		api, err := LoadApi(ver, false)
		if err != nil {
			return err
		}
		for k, v := range api.Keys {
			if info, ok := ac.Base[k]; ok {
				if info.Tags[0] != "-" {
					info.Tags = append(info.Tags, v.Tags...)
				}
			} else {
				ac.Base[k] = v
			}
		}
	}
	return nil
}

func (ac *ApiCheck) LoadApi(vers ...string) error {
	for _, ver := range vers {
		api, err := LoadApi(ver, true)
		if err != nil {
			return err
		}
		for k, v := range api.Keys {
			if info, ok := ac.Base[k]; ok {
				if info.Tags[0] != "-" {
					info.Tags = append(info.Tags, v.Tags...)
				}
				delete(api.Keys, k)
			}
		}
		ac.Apis = append(ac.Apis, api)
	}
	return nil
}

func (ac *ApiCheck) FindApis(name string) (vers string, info *ApiInfo) {
	for _, api := range ac.Apis {
		if info, ok := api.Keys[name]; ok {
			return api.Ver, info
		}
	}
	return
}

func (ac *ApiCheck) ApiVers() (vers []string) {
	for _, api := range ac.Apis {
		vers = append(vers, api.Ver)
	}
	return
}

func (ac *ApiCheck) LoadPkg(pkg string) (infos []*ApiInfo) {
	for _, api := range ac.Apis {
		for _, v := range api.Keys {
			if v.Pkg == pkg {
				infos = append(infos, v)
			}
		}
	}
	for _, v := range ac.Base {
		if v.Pkg == pkg {
			infos = append(infos, v)
		}
	}
	sort.Slice(infos, func(i, j int) bool {
		n := strings.Compare(infos[i].Kind, infos[j].Kind)
		if n == 0 {
			return infos[i].Name < infos[j].Name
		}
		return n < 0
	})
	return
}

type Style struct {
	Ver string
	Tag string
}

type LineData struct {
	Ver  string
	Tag  []string
	Info string
}

var (
	osMap = map[string]string{
		"windows-386 windows-amd64":                    "windows",
		"darwin-386-cgo darwin-amd64 darwin-amd64-cgo": "darwin",
		"darwin-386 darwin-386-cgo darwin-386-cgo darwin-amd64 darwin-amd64 darwin-amd64-cgo darwin-amd64-cgo":            "darwin",
		"linux-386 linux-386-cgo linux-amd64 linux-amd64-cgo linux-arm linux-arm-cgo":                                     "linux",
		"linux-386-cgo linux-amd64 linux-amd64-cgo linux-arm linux-arm-cgo":                                               "linux",
		"linux-386 linux-386-cgo linux-amd64 linux-amd64-cgo linux-arm":                                                   "linux",
		"freebsd-386 freebsd-386-cgo freebsd-amd64 freebsd-amd64-cgo freebsd-arm freebsd-arm-cgo":                         "freebsd",
		"freebsd-386-cgo freebsd-amd64 freebsd-amd64-cgo freebsd-arm freebsd-arm-cgo":                                     "freebsd",
		"netbsd-386 netbsd-386-cgo netbsd-amd64 netbsd-amd64-cgo netbsd-arm netbsd-arm-cgo netbsd-arm64 netbsd-arm64-cgo": "netbsd",
		"openbsd-386 openbsd-386-cgo openbsd-amd64 openbsd-amd64-cgo":                                                     "openbsd",
	}
	osReplacer *strings.Replacer
)

func init() {
	var oldnews []string
	for k, v := range osMap {
		oldnews = append(oldnews, k, v)
	}
	osReplacer = strings.NewReplacer(oldnews...)
}

var (
	tagsName = make(map[string]string)
	tagIndex int
)

func (d LineData) TagName() (name string, tags []string) {
	if d.Ver != "" {
		name += "_" + strings.Replace(d.Ver, ".", "", -1)
		tags = append(tags, "// +build "+d.Ver)
	}
	switch len(d.Tag) {
	case 0:
	case 1:
		name += "_" + strings.Replace(d.Tag[0], "-", "_", -1)
	default:
		otags := strings.Join(d.Tag, " ")
		ntags := osReplacer.Replace(otags)
		if strings.Index(ntags, " ") == -1 {
			name += "_" + ntags
		} else {
			tags = append(tags, "// +build "+strings.Replace(ntags, "-", ",", -1))
			if tag, ok := tagsName[ntags]; ok {
				name += "_" + tag
			} else {
				h := fnv.New32()
				h.Write([]byte(ntags))
				name += fmt.Sprintf("_%v", h.Sum32())
			}
		}
	}
	return
}

type File struct {
	Id      string
	Name    string
	Tags    []string
	ExtList []string
	TypList []string
}

func (ac *ApiCheck) Export(pkgPath string) (map[string]*File, error) {
	extList, typList, ok := ac.ExportData(pkgPath)
	if !ok {
		return nil, fmt.Errorf("empty export pkg %v", pkgPath)
	}
	fileMap := make(map[string]*File)
	for _, v := range extList {
		name, tags := v.TagName()
		f, ok := fileMap[name]
		if !ok {
			f = &File{Name: name, Tags: tags}
			fileMap[name] = f
		}
		f.ExtList = append(f.ExtList, v.Info)
	}
	for _, v := range typList {
		name, tags := v.TagName()
		f, ok := fileMap[name]
		if !ok {
			f = &File{Name: name, Tags: tags}
			fileMap[name] = f
		}
		f.TypList = append(f.TypList, v.Info)
	}
	return fileMap, nil
}

func cleanTags(tags []string) []string {
	sort.Strings(tags)
	var i, j int
	for {
		if i >= len(tags)-1 {
			break
		}
		for j = i + 1; j < len(tags) && tags[i] == tags[j]; j++ {
		}
		tags = append(tags[:i+1], tags[j:]...)
		i++
	}
	return tags
}

// linux-386-cgo linux-amd64 linux-amd64-cgo linux-arm linux-arm-cgo
// darwin-386-cgo darwin-amd64 darwin-amd64-cgo
// freebsd-386 freebsd-386-cgo freebsd-amd64 freebsd-amd64-cgo freebsd-arm freebsd-arm-cgo
// netbsd-386 netbsd-386-cgo netbsd-amd64 netbsd-amd64-cgo netbsd-arm netbsd-arm-cgo netbsd-arm64 netbsd-arm64-cgo
// openbsd-386 openbsd-386-cgo openbsd-amd64 openbsd-amd64-cgo
// windows-386 windows-amd64
func (ac *ApiCheck) ExportData(pkgPath string) (extList []LineData, typList []LineData, ok bool) {
	infos := ac.LoadPkg(pkgPath)
	if len(infos) == 0 {
		return
	}
	pkgList := strings.Split(pkgPath, "/")
	pkgName := pkgList[len(pkgList)-1]
	for _, t := range infos {
		tags := cleanTags(t.Tags)[1:]
		switch t.Kind {
		case "var":
			extList = append(extList, LineData{
				t.Ver,
				tags,
				fmt.Sprintf("%q : &%v", pkgPath+"."+t.Name, pkgName+"."+t.Name),
			})
		case "func":
			extList = append(extList, LineData{
				t.Ver,
				tags,
				fmt.Sprintf("%q : %v", pkgPath+"."+t.Name, pkgName+"."+t.Name),
			})
		case "type":
			typList = append(typList, LineData{
				t.Ver,
				tags,
				fmt.Sprintf("(*%v.%v)(nil)", pkgName, t.Name),
			})
			for _, v := range infos {
				if v.Kind == "method" {
					if v.MethodType == t.Name {
						tags := v.Tags[1:]
						sort.Strings(tags)
						extList = append(extList, LineData{
							v.Ver,
							tags,
							fmt.Sprintf("\"(%v%v.%v).%v\" : (%v%v.%v).%v",
								v.MethodPtr, pkgPath, v.MethodType, v.MethodName,
								v.MethodPtr, pkgName, v.MethodType, v.MethodName),
						})
					}
				}
			}
		}
	}
	ok = len(extList) > 0 || len(typList) > 0
	return
}
