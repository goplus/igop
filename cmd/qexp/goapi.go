package main

import (
	"bufio"
	"fmt"
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
var sym = regexp.MustCompile(`^pkg (\S+)\s?(.*)?, (?:(var|func|type|const|method))\s?(\((\*?[A-Z]\w*)\))? ([A-Z]\w*)`)
var reMethod = regexp.MustCompile(`^pkg (\S+)\s?(.*)?, method (\((\*?[A-Z]\w*)\))? ([A-Z]\w*)`) // (\(\*?[A-Z]\w*\)) ([A-Z]\w*)`)
var num = regexp.MustCompile(`^\-?[0-9]+$`)

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
		if has(", method (") {
			// m[1] pkg
			// m[2] os-arch-cgo
			// m[3] var|func|type|const|method
			// m[4] (*?Type)
			// m[5] *?Type
			// m[6] name
		}
		if has("interface, ") || has("struct, ") {
			continue
		}
		if m := sym.FindStringSubmatch(l); m != nil {
			// m[1] pkg
			// m[2] os-arch-cgo
			// m[3] var|func|type|const|method
			// m[4] (*?Type)
			// m[5] *?Type
			// m[6] name
			pkg := m[1]
			tag := m[2]
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

type Data struct {
	Ver  string
	Tag  []string
	Info string
}

func (ac *ApiCheck) Export(pkgPath string) (extList []Data, typList []Data, ok bool) {
	infos := ac.LoadPkg(pkgPath)
	if len(infos) == 0 {
		return
	}
	pkgList := strings.Split(pkgPath, "/")
	pkgName := pkgList[len(pkgList)-1]
	for _, t := range infos {
		switch t.Kind {
		case "var":
			extList = append(extList, Data{
				t.Ver,
				t.Tags,
				fmt.Sprintf("%q : &%v", pkgPath+"."+t.Name, pkgName+"."+t.Name),
			})
		case "func":
			extList = append(extList, Data{
				t.Ver,
				t.Tags,
				fmt.Sprintf("%q : %v", pkgPath+"."+t.Name, pkgName+"."+t.Name),
			})
		case "type":
			typList = append(typList, Data{
				t.Ver,
				t.Tags,
				fmt.Sprintf("(*%v.%v)(nil)", pkgName, t.Name),
			})
			for _, v := range infos {
				if v.Kind == "method" {
					if v.MethodType == t.Name {
						extList = append(extList, Data{
							v.Ver,
							v.Tags,
							fmt.Sprintf("\"(%v%v.%v).%v\" : (%v%v.%v).%v",
								v.MethodPtr, pkgPath, v.MethodType, v.MethodName,
								v.MethodPtr, pkgPath, v.MethodType, v.MethodName),
						})
					}
				}
			}
		}
	}
	ok = len(extList) > 0 || len(typList) > 0
	return
}
