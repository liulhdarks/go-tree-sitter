package xml_test

import (
    "context"
    "github.com/smacker/go-tree-sitter/xml"
    "testing"

    sitter "github.com/smacker/go-tree-sitter"
    "github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
    assert := assert.New(t)
    n, err := sitter.ParseCtx(context.Background(), []byte(`
<?xml version="1.0" encoding="utf-8"?>
<modification>
  <name>Hello world</name>
  <file path="catalog/controller/common/hello_world.php">
    <operation>
      <search><![CDATA[
if (true) {
]]></search>
      <add position="replace"><![CDATA[
if (false) {
]]>
      </add>
    </operation>
  </file>
</modification>
`), xml.GetLanguage())
    assert.NoError(err)
    print(n.String())
    //assert.Equal(
    //    "(module (expression_statement (call function: (identifier) arguments: (argument_list (integer)))))",
    //    n.String(),
    //)
}
