package vue_test

import (
    "context"
    "testing"

    sitter "github.com/smacker/go-tree-sitter"
    "github.com/smacker/go-tree-sitter/vue"
    "github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
    assert := assert.New(t)
    n, err := sitter.ParseCtx(context.Background(), []byte(`
<html>
<head>
    <title>Vue Counter Example</title>
    <!-- 引入 Vue.js 的 CDN 链接 -->
    <script src="https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.js"></script>
</head>
<body>
    <!-- 创建一个 id 为 app 的 div，Vue 实例会挂载在这里 -->
    <div id="app">
        <h1>Simple Vue Counter</h1>
        <p>{{ message }}</p>
        <button @click="incrementCounter">Click me to increment!</button>
        <p>Counter: {{ counter }}</p>
    </div>

    <script>
        // 创建一个新的 Vue 实例
        var app = new Vue({
            // Vue 实例将挂载到 #app div
            el: '#app',
            // 实例的数据对象
            data: {
                message: 'Hello Vue!',
                counter: 0
            },
            // 实例的方法对象
            methods: {
                incrementCounter() {
                    this.counter += 1; // 增加 counter 的值
                }
            }
        });
    </script>
</body>
</html>
`), vue.GetLanguage())
    assert.NoError(err)
    print(n.String())
    //assert.Equal(
    //    "(module (expression_statement (call function: (identifier) arguments: (argument_list (integer)))))",
    //    n.String(),
    //)
}
