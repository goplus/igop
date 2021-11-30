// export by github.com/goplus/gossa/cmd/qexp

package parse

import (
	q "text/template/parse"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "parse",
		Path: "text/template/parse",
		Deps: map[string]string{
			"bytes":        "bytes",
			"fmt":          "fmt",
			"runtime":      "runtime",
			"strconv":      "strconv",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ActionNode":     {reflect.TypeOf((*q.ActionNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"BoolNode":       {reflect.TypeOf((*q.BoolNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"BranchNode":     {reflect.TypeOf((*q.BranchNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"ChainNode":      {reflect.TypeOf((*q.ChainNode)(nil)).Elem(), "", "Add,Copy,String,tree,writeTo"},
			"CommandNode":    {reflect.TypeOf((*q.CommandNode)(nil)).Elem(), "", "Copy,String,append,tree,writeTo"},
			"CommentNode":    {reflect.TypeOf((*q.CommentNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"DotNode":        {reflect.TypeOf((*q.DotNode)(nil)).Elem(), "", "Copy,String,Type,tree,writeTo"},
			"FieldNode":      {reflect.TypeOf((*q.FieldNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"IdentifierNode": {reflect.TypeOf((*q.IdentifierNode)(nil)).Elem(), "", "Copy,SetPos,SetTree,String,tree,writeTo"},
			"IfNode":         {reflect.TypeOf((*q.IfNode)(nil)).Elem(), "", "Copy"},
			"ListNode":       {reflect.TypeOf((*q.ListNode)(nil)).Elem(), "", "Copy,CopyList,String,append,tree,writeTo"},
			"Mode":           {reflect.TypeOf((*q.Mode)(nil)).Elem(), "", ""},
			"NilNode":        {reflect.TypeOf((*q.NilNode)(nil)).Elem(), "", "Copy,String,Type,tree,writeTo"},
			"NodeType":       {reflect.TypeOf((*q.NodeType)(nil)).Elem(), "Type", ""},
			"NumberNode":     {reflect.TypeOf((*q.NumberNode)(nil)).Elem(), "", "Copy,String,simplifyComplex,tree,writeTo"},
			"PipeNode":       {reflect.TypeOf((*q.PipeNode)(nil)).Elem(), "", "Copy,CopyPipe,String,append,tree,writeTo"},
			"Pos":            {reflect.TypeOf((*q.Pos)(nil)).Elem(), "Position", ""},
			"RangeNode":      {reflect.TypeOf((*q.RangeNode)(nil)).Elem(), "", "Copy"},
			"StringNode":     {reflect.TypeOf((*q.StringNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"TemplateNode":   {reflect.TypeOf((*q.TemplateNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"TextNode":       {reflect.TypeOf((*q.TextNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"Tree":           {reflect.TypeOf((*q.Tree)(nil)).Elem(), "", "Copy,ErrorContext,Parse,action,add,backup,backup2,backup3,blockControl,checkPipeline,clearActionLine,command,elseControl,endControl,error,errorf,expect,expectOneOf,hasFunction,ifControl,itemList,newAction,newBool,newChain,newCommand,newComment,newDot,newElse,newEnd,newField,newIf,newList,newNil,newNumber,newPipeline,newRange,newString,newTemplate,newText,newVariable,newWith,next,nextNonSpace,operand,parse,parseControl,parseDefinition,parseTemplateName,peek,peekNonSpace,pipeline,popVars,rangeControl,recover,startParse,stopParse,templateControl,term,textOrAction,unexpected,useVar,withControl"},
			"VariableNode":   {reflect.TypeOf((*q.VariableNode)(nil)).Elem(), "", "Copy,String,tree,writeTo"},
			"WithNode":       {reflect.TypeOf((*q.WithNode)(nil)).Elem(), "", "Copy"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsEmptyTree":   reflect.ValueOf(q.IsEmptyTree),
			"New":           reflect.ValueOf(q.New),
			"NewIdentifier": reflect.ValueOf(q.NewIdentifier),
			"Parse":         reflect.ValueOf(q.Parse),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"NodeAction":     {reflect.TypeOf(q.NodeAction), constant.MakeInt64(1)},
			"NodeBool":       {reflect.TypeOf(q.NodeBool), constant.MakeInt64(2)},
			"NodeChain":      {reflect.TypeOf(q.NodeChain), constant.MakeInt64(3)},
			"NodeCommand":    {reflect.TypeOf(q.NodeCommand), constant.MakeInt64(4)},
			"NodeComment":    {reflect.TypeOf(q.NodeComment), constant.MakeInt64(20)},
			"NodeDot":        {reflect.TypeOf(q.NodeDot), constant.MakeInt64(5)},
			"NodeField":      {reflect.TypeOf(q.NodeField), constant.MakeInt64(8)},
			"NodeIdentifier": {reflect.TypeOf(q.NodeIdentifier), constant.MakeInt64(9)},
			"NodeIf":         {reflect.TypeOf(q.NodeIf), constant.MakeInt64(10)},
			"NodeList":       {reflect.TypeOf(q.NodeList), constant.MakeInt64(11)},
			"NodeNil":        {reflect.TypeOf(q.NodeNil), constant.MakeInt64(12)},
			"NodeNumber":     {reflect.TypeOf(q.NodeNumber), constant.MakeInt64(13)},
			"NodePipe":       {reflect.TypeOf(q.NodePipe), constant.MakeInt64(14)},
			"NodeRange":      {reflect.TypeOf(q.NodeRange), constant.MakeInt64(15)},
			"NodeString":     {reflect.TypeOf(q.NodeString), constant.MakeInt64(16)},
			"NodeTemplate":   {reflect.TypeOf(q.NodeTemplate), constant.MakeInt64(17)},
			"NodeText":       {reflect.TypeOf(q.NodeText), constant.MakeInt64(0)},
			"NodeVariable":   {reflect.TypeOf(q.NodeVariable), constant.MakeInt64(18)},
			"NodeWith":       {reflect.TypeOf(q.NodeWith), constant.MakeInt64(19)},
			"ParseComments":  {reflect.TypeOf(q.ParseComments), constant.MakeInt64(1)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
