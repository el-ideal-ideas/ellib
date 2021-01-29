// an alias of strings package.

package elstr

import "strings"

type SBuilder strings.Builder
type SReader strings.Reader
type SReplacer strings.Replacer

var SCompare = strings.Compare
var SNewReader = strings.NewReader
var SNewReplacer = strings.NewReplacer
var SCount = strings.Count
var SContains = strings.Contains
var SContainsAny = strings.ContainsAny
var SContainsRune = strings.ContainsRune
var SLastIndex = strings.LastIndex
var SIndexByte = strings.IndexByte
var SIndexRune = strings.IndexRune
var SIndexAny = strings.IndexAny
var SLastIndexAny = strings.LastIndexAny
var SLastIndexByte = strings.LastIndexByte
var SSplitN = strings.SplitN
var SSplitAfterN = strings.SplitAfterN
var SSplit = strings.Split
var SSplitAfter = strings.SplitAfter
var SFields = strings.Fields
var SFieldsFunc = strings.FieldsFunc
var SJoin = strings.Join
var SHasPrefix = strings.HasPrefix
var SHasSuffix = strings.HasSuffix
var SMap = strings.Map
var SRepeat = strings.Repeat
var SToUpper = strings.ToUpper
var SToLower = strings.ToLower
var SToTitle = strings.ToTitle
var SToUpperSpecial = strings.ToUpperSpecial
var SToLowerSpecial = strings.ToLowerSpecial
var SToTitleSpecial = strings.ToTitleSpecial
var SToValidUTF8 = strings.ToValidUTF8
var STitle = strings.Title
var STrimLeftFunc = strings.TrimLeftFunc
var STrimRightFunc = strings.TrimRightFunc
var STrimFunc = strings.TrimFunc
var SIndexFunc = strings.IndexFunc
var SLastIndexFunc = strings.LastIndexFunc
var STrim = strings.Trim
var STrimLeft = strings.TrimLeft
var STrimRight = strings.TrimRight
var STrimSpace = strings.TrimSpace
var STrimPrefix = strings.TrimPrefix
var STrimSuffix = strings.TrimSuffix
var SReplace = strings.Replace
var SReplaceAll = strings.ReplaceAll
var SEqualFold = strings.EqualFold
var SIndex = strings.Index
