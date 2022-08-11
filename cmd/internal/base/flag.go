package base

import (
	"go/build"
	"strings"
)

var (
	BuildContext     = defaultContext()
	BuildMod         string // -mod flag
	BuildModExplicit bool   // whether -mod was set explicitly
	BuildX           bool   // -x flag
	BuildV           bool   // -v flag
	BuildSSA         bool   // -ssa flag
	DebugSSATrace    bool   // -ssa-trace flag
)

func defaultContext() build.Context {
	return build.Default
}

type BuildFlagMask int

const (
	DefaultBuildFlags BuildFlagMask = 0
	OmitModFlag       BuildFlagMask = 1 << iota
	OmitVFlag
	OmitSSAFlag
	OmitSSATraceFlag
)

// AddBuildFlags adds the flags common to the build, run, and test commands.
func AddBuildFlags(cmd *Command, mask BuildFlagMask) {
	cmd.Flag.BoolVar(&BuildX, "x", false, "print the commands.")
	if mask&OmitVFlag != 0 {
		cmd.Flag.BoolVar(&BuildV, "v", false, "print the names of packages as they are compiled.")
	}
	if mask&OmitModFlag != 0 {
		cmd.Flag.Var(explicitStringFlag{value: &BuildMod, explicit: &BuildModExplicit}, "mod", "module download mode to use: readonly, vendor, or mod.")
	}
	if mask&OmitSSAFlag != 0 {
		cmd.Flag.BoolVar(&BuildSSA, "ssa", false, "print SSA instruction code")
	}
	if mask&OmitSSATraceFlag != 0 {
		cmd.Flag.BoolVar(&DebugSSATrace, "ssa-trace", false, "trace SSA interpreter code")
	}
	cmd.Flag.Var((*tagsFlag)(&BuildContext.BuildTags), "tags", "a comma-separated list of build tags to consider satisfied during the build")
}

// tagsFlag is the implementation of the -tags flag.
type tagsFlag []string

func (v *tagsFlag) Set(s string) error {
	// Split on commas, ignore empty strings.
	*v = []string{}
	for _, s := range strings.Split(s, ",") {
		if s != "" {
			*v = append(*v, s)
		}
	}
	return nil
}

func (v *tagsFlag) String() string {
	return "<TagsFlag>"
}

// explicitStringFlag is like a regular string flag, but it also tracks whether
// the string was set explicitly to a non-empty value.
type explicitStringFlag struct {
	value    *string
	explicit *bool
}

func (f explicitStringFlag) String() string {
	if f.value == nil {
		return ""
	}
	return *f.value
}

func (f explicitStringFlag) Set(v string) error {
	*f.value = v
	if v != "" {
		*f.explicit = true
	}
	return nil
}
