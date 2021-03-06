package function

import "gopkg.in/src-d/go-mysql-server.v0/sql"

// Functions for gitbase queries.
var Functions = []sql.Function{
	sql.Function1{Name: "is_tag", Fn: NewIsTag},
	sql.Function1{Name: "is_remote", Fn: NewIsRemote},
	sql.FunctionN{Name: "language", Fn: NewLanguage},
	sql.FunctionN{Name: "uast", Fn: NewUAST},
	sql.Function3{Name: "uast_mode", Fn: NewUASTMode},
	sql.Function2{Name: "uast_xpath", Fn: NewUASTXPath},
	sql.Function2{Name: "uast_extract", Fn: NewUASTExtract},
	sql.Function1{Name: "uast_children", Fn: NewUASTChildren},
}
