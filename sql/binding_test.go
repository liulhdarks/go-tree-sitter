package tree_sitter_sql_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/smacker/go-tree-sitter/sql"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)
	n, err := sitter.ParseCtx(context.Background(), []byte(`
create table table_name(
id interger,
name text
);

select * from table_name where id>0;
`), tree_sitter_sql.GetLanguage())
	assert.NoError(err)
	print(n.String())
	assert.Equal(
		"(program (statement (create_table (keyword_create) (keyword_table) (object_reference name: (identifier)) (column_definitions (column_definition name: (identifier) custom_type: (object_reference name: (identifier))) (column_definition name: (identifier) type: (keyword_text))))) (statement (select (keyword_select) (select_expression (term value: (all_fields)))) (from (keyword_from) (relation (object_reference name: (identifier))) (where (keyword_where) predicate: (binary_expression left: (field name: (identifier)) right: (literal))))))",
		n.String(),
	)
}
