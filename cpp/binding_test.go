package cpp_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	bytes, err := os.ReadFile("/Users/lihua.llh/Downloads/cat-565.h")
	assert.NoError(err)
	n, err := sitter.ParseCtx(context.Background(), bytes, cpp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(translation_unit (declaration type: (primitive_type) declarator: (init_declarator declarator: (identifier) value: (number_literal))))",
		n.String(),
	)
	n2, err := sitter.ParseCtx(context.Background(), []byte(`
class MY_DLL_API MyInterface {
public:
    virtual ~MyInterface() {}
    virtual void DoSomething() = 0; 
};
struct DLL_EXPORT MyStruct {
    int field1;
    wchar_t* name;
};
`), cpp.GetLanguage())
	assert.NoError(err)
	fmt.Println(n2.String())
}
