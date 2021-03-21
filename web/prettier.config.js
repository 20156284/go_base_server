// prettier.config.js
module.exports = {
    printWidth: 80,             // 超过最大值换行
    tabWidth: 2,                // 缩进字节数
    useTabs: false,             // 缩进不使用tab，使用空格
    semi: false,                // 句尾添加分号
    singleQuote: true,          // 使用单引号代替双引号
    trailingComma: 'none',      // 在对象或数组最后一个元素后面是否加逗号
    bracketSpacing: true,       // 在对象，数组括号与文字之间加空格 "{ foo: bar }"
    jsxBracketSameLine: true,   // 在jsx中把'>' 是否单独放一行
    arrowParens: 'avoid',       // (x) => {} 箭头函数参数只有一个时是否要有小括号。avoid：省略括号
    insertPragma: false,        // Prettier可以在文件的顶部插入一个 @format的特殊注释，以表明改文件已经被Prettier格式化过了
    endOfLine: 'auto',          // 结尾是 \n \r \n\r auto
}
