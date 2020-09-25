package reqres

import (
	"github.com/theNorstroem/spectools/pkg/orderedmap"
	"gopkg.in/yaml.v3"
)

// DecodeExtension decodes the extension node["ename"] and stores its data into the value pointed to by v.
//
// returns true if ename was found and could be decoded without error
// conversion of extension YAML into a Go value.
func DecodeExtension(node *orderedmap.OrderedMap, ename string, v interface{}) bool {
	iValue, found := node.Get(ename)
	if found {
		fieldYamlNode := iValue.(*yaml.Node)
		err := fieldYamlNode.Decode(v)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
