# antlr4-grammars [![Build Status](https://img.shields.io/travis/bramp/antlr4-grammars.svg)](https://travis-ci.org/bramp/antlr4-grammars) [![Coverage](https://img.shields.io/coveralls/bramp/antlr4-grammars.svg)](https://coveralls.io/github/bramp/antlr4-grammars) [![Report card](https://goreportcard.com/badge/github.com/bramp/antlr4-grammars)](https://goreportcard.com/report/github.com/bramp/antlr4-grammars) [![GoDoc](https://godoc.org/github.com/bramp/antlr4-grammars?status.svg)](https://godoc.org/github.com/bramp/antlr4-grammars)

Precompiled Go versions of many of the grammars on [github.com/antlr/grammars-v4](http://github.com/antlr/grammars-v4).

## Example
```go

import (
	"bramp.net/antlr4-grammars/java"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type exampleListener struct {
	*java.BaseJavaListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := java.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := java.NewJavaParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.CompilationUnit()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
}
```

## Supported Languages

| Status | Language     | Notes                                                                       |
| :-: | --------------- | --------------------------------------------------------------------------- |
| ✅  | abnf            |                                                                             |
| ✅  | agc             |                                                                             |
| ✅  | atl             |                                                                             |
| ✅  | bnf             |                                                                             |
| ✅  | brainfuck       |                                                                             |
| ✅  | c               |                                                                             |
| ✅  | clf             |                                                                             |
| ✅  | clif            |                                                                             |
| ✅  | cookie          |                                                                             |
| ✅  | cool            |                                                                             |
| ✅  | creole          |                                                                             |
| ✅  | csv             |                                                                             |
| ✅  | datetime        |                                                                             |
| ✅  | dcm             |                                                                             |
| ✅  | dot             |                                                                             |
| ✅  | emailaddress    |                                                                             |
| ✅  | fasta           |                                                                             |
| ✅  | fol             |                                                                             |
| ✅  | fusion_tables   |                                                                             |
| ✅  | gml             |                                                                             |
| ✅  | graphstream_dgs |                                                                             |
| ✅  | idl             |                                                                             |
| ✅  | iri             |                                                                             |
| ✅  | istc            |                                                                             |
| ✅  | java            |                                                                             |
| ✅  | java8_pt        |                                                                             |
| ✅  | jpa             |                                                                             |
| ✅  | json            |                                                                             |
| ✅  | lambda          |                                                                             |
| ✅  | lcc             |                                                                             |
| ✅  | less            |                                                                             |
| ✅  | mdx             |                                                                             |
| ✅  | memcached_protocol |                                                                             |
| ✅  | metric          |                                                                             |
| ✅  | modelica        |                                                                             |
| ✅  | molecule        |                                                                             |
| ✅  | mps             |                                                                             |
| ✅  | mumath          |                                                                             |
| ✅  | mumps           |                                                                             |
| ✅  | muparser        |                                                                             |
| ✅  | p               |                                                                             |
| ✅  | pcre            |                                                                             |
| ✅  | peoplecode      |                                                                             |
| ✅  | postalcode      |                                                                             |
| ✅  | powerbuilder    |                                                                             |
| ✅  | propcalc        |                                                                             |
| ✅  | properties      |                                                                             |
| ✅  | rcs             |                                                                             |
| ✅  | redcode         |                                                                             |
| ✅  | robotwars       |                                                                             |
| ✅  | romannumerals   |                                                                             |
| ✅  | rpn             |                                                                             |
| ✅  | ruby            |                                                                             |
| ✅  | scss            |                                                                             |
| ✅  | sexpression     |                                                                             |
| ✅  | sharc           |                                                                             |
| ✅  | snobol          |                                                                             |
| ✅  | stacktrace      |                                                                             |
| ✅  | suokif          |                                                                             |
| ✅  | telephone       |                                                                             |
| ✅  | tiny            |                                                                             |
| ✅  | tinyc           |                                                                             |
| ✅  | tnsnames        |                                                                             |
| ✅  | tnt             |                                                                             |
| ✅  | useragent       |                                                                             |
| ✅  | wavefront       |                                                                             |
| ✅  | xml             |                                                                             |
| ✅  | xpath           |                                                                             |
| ❌  | antlr3          | antlr: error(134): ANTLRv3.g4:32:102: symbol action conflicts with generated code in target language or runtime |
| ❌  | antlr4          | antlr: error(134): ANTLRv4Parser.g4:62:5: symbol action conflicts with generated code in target language or runtime |
| ❌  | apex            | antlr: error(134): apex.g4:544:30: symbol type conflicts with generated code in target language or runtime |
| ❌  | arithmetic      |  test: FAIL  bramp.net/antlr4-grammars/arithmetic    0.031s                     |
| ❌  | asm6502         | antlr: error(134): asm6502.g4:71:30: symbol string conflicts with generated code in target language or runtime |
| ❌  | asn             | antlr: error(134): ASN.g4:166:2: symbol type conflicts with generated code in target language or runtime |
| ❌  | basic           | antlr: error(134): jvmBasic.g4:481:8: symbol var conflicts with generated code in target language or runtime |
| ❌  | calculator      | antlr: error(134): calculator.g4:54:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | clojure         | antlr: error(134): Clojure.g4:99:12: symbol map conflicts with generated code in target language or runtime |
| ❌  | cobol85         |  test: FAIL  bramp.net/antlr4-grammars/cobol85 [build failed]                |
| ❌  | cpp             | build: cpp/cpp14_parser.go:28859:6: syntax error: unexpected NewEmptyBaseclauseContext, expecting ( |
| ❌  | csharp          | antlr: error(134): CSharpParser.g4:330:8: symbol type conflicts with generated code in target language or runtime |
| ❌  | csharp          | antlr: error(134): CSharpParser.g4:330:8: symbol type conflicts with generated code in target language or runtime |
| ❌  | css3            | antlr: error(134): css3.g4:196:6: symbol var conflicts with generated code in target language or runtime |
| ❌  | dice            |  test: FAIL  bramp.net/antlr4-grammars/dice  0.032s                           |
| ❌  | ecmascript      | antlr: error(8): ECMAScript.GoTarget.g4:28:8: grammar name ECMAScript and file name ECMAScript.GoTarget.g4 differ |
| ❌  | erlang          | antlr: error(134): Erlang.g4:156:36: symbol type conflicts with generated code in target language or runtime |
| ❌  | fortran77       | antlr: error(134): fortran77.g4:89:6: symbol type conflicts with generated code in target language or runtime |
| ❌  | gff3            | antlr: error(134): gff3.g4:45:28: symbol type conflicts with generated code in target language or runtime |
| ❌  | golang          | antlr: error(134): Golang.g4:151:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | graphql         | antlr: error(134): GraphQL.g4:140:9: symbol type conflicts with generated code in target language or runtime |
| ❌  | gtin            |  test: FAIL  bramp.net/antlr4-grammars/gtin  0.048s                           |
| ❌  | html            |  test: FAIL  bramp.net/antlr4-grammars/html  10.027s                          |
| ❌  | icalendar       | antlr: error(134): ICalendar.g4:305:3: symbol action conflicts with generated code in target language or runtime |
| ❌  | informix        | antlr: error(134): informix.g4:601:266: symbol string conflicts with generated code in target language or runtime |
| ❌  | java8           | antlr: error(134): Java8.g4:73:0: symbol type conflicts with generated code in target language or runtime |
| ❌  | javadoc         | build: javadoc/javadoc_lexer.go:167:10: undefined: _input                   |
| ❌  | kotlin          | antlr: error(134): KotlinParser.g4:56:21: symbol type conflicts with generated code in target language or runtime |
| ❌  | kuka            | antlr: error(134): krl.g4:80:68: symbol type conflicts with generated code in target language or runtime |
| ❌  | logo            | antlr: error(134): logo.g4:128:38: symbol func conflicts with generated code in target language or runtime |
| ❌  | lolcode         | antlr: error(134): lolcode.g4:111:5: symbol func conflicts with generated code in target language or runtime |
| ❌  | lua             | antlr: error(134): Lua.g4:90:15: symbol var conflicts with generated code in target language or runtime |
| ❌  | masm            | build: masm/masm_lexer.go:12:4: syntax error: non-declaration statement outside function body |
| ❌  | modula2pim4     | antlr: error(134): m2pim4.g4:356:22: symbol type conflicts with generated code in target language or runtime |
| ❌  | morsecode       |  test: FAIL  bramp.net/antlr4-grammars/morsecode     0.028s                      |
| ❌  | mysql           | build:       previous declaration at mysql/mysql_parser.go:12215:6               |
| ❌  | oncrpc          |  test: FAIL  bramp.net/antlr4-grammars/oncrpc [build failed]                 |
| ❌  | pascal          | antlr: error(134): pascal.g4:115:23: symbol type conflicts with generated code in target language or runtime |
| ❌  | pddl            | antlr: error(134): Pddl.g4:92:25: symbol type conflicts with generated code in target language or runtime |
| ❌  | pdn             | antlr: error(134): pdn.g4:44:14: symbol string conflicts with generated code in target language or runtime |
| ❌  | pdp7            | antlr: error(134): pdp7.g4:87:5: symbol string conflicts with generated code in target language or runtime |
| ❌  | pgn             | build: pgn/pgn_lexer.go:191:10: undefined: getCharPositionInLine            |
| ❌  | php             | antlr: error(134): PHPParser.g4:478:6: symbol string conflicts with generated code in target language or runtime |
| ❌  | plsql           | antlr: error(134): PlSqlParser.g4:779:41: symbol range conflicts with generated code in target language or runtime |
| ❌  | prolog          | antlr: error(134): prolog.g4:72:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | protobuf3       | antlr: error(134): Protobuf3.g4:138:19: symbol range conflicts with generated code in target language or runtime |
| ❌  | python2         | build: python2/python2_lexer.go:639:14: too many errors                     |
| ❌  | python3         | build: python3/python3_lexer.go:828:13: too many errors                     |
| ❌  | python3_js      | build: python3_js/python3_lexer.go:805:13: too many errors                  |
| ❌  | python3_py      | build: python3_py/python3_lexer.go:774:29: too many errors                  |
| ❌  | python3_ts      | build: python3_ts/python3_lexer.go:14:33: illegal rune literal              |
| ❌  | python3alt      | build: python3alt/altpython3_lexer.go:812:13: too many errors               |
| ❌  | quakemap        | antlr: error(134): quakemap.g4:44:12: symbol string conflicts with generated code in target language or runtime |
| ❌  | r               | build: r/rfilter_parser.go:1283:2: syntax error: unexpected default, expecting } |
| ❌  | scala           | antlr: error(134): Scala.g4:433:17: symbol type conflicts with generated code in target language or runtime |
| ❌  | smalltalk       | antlr: error(134): Smalltalk.g4:50:57: symbol string conflicts with generated code in target language or runtime |
| ❌  | smtlibv2        | antlr: error(134): SMTLIBv2.g4:1089:23: symbol string conflicts with generated code in target language or runtime |
| ❌  | sparql          | antlr: error(134): Sparql.g4:57:45: symbol var conflicts with generated code in target language or runtime |
| ❌  | sqlite          | antlr: error(134): SQLite.g4:34:21: symbol error conflicts with generated code in target language or runtime |
| ❌  | stringtemplate  | build: stringtemplate/st_lexer.go:779:2: syntax error: unexpected default, expecting : |
| ❌  | stringtemplate  | build: stringtemplate/stg_lexer.go:716:2: syntax error: unexpected default, expecting : |
| ❌  | swift2          | antlr: error(134): Swift2.g4:783:34: symbol type conflicts with generated code in target language or runtime |
| ❌  | swift3          | antlr: error(134): Swift3.g4:664:16: symbol type conflicts with generated code in target language or runtime |
| ❌  | tsql            | build:       previous declaration at tsql/tsql_parser.go:123881:8                |
| ❌  | turtle          | build:               want String([]string, antlr.RuleContext) string                    |
| ❌  | ucb_logo        | build: ucb_logo/ucblogo_parser.go:14:14: expected 'STRING', found '.'       |
| ❌  | unicode/graphemes | build:     previous declaration at ./makemake.go:158:6                         |
| ❌  | unicode/unicode16 | build:     previous declaration at ./makemake.go:158:6                         |
| ❌  | url             | antlr: error(134): url.g4:95:18: symbol string conflicts with generated code in target language or runtime |
| ❌  | vb6             | antlr: error(134): VisualBasic6.g4:555:35: symbol type conflicts with generated code in target language or runtime |
| ❌  | vba             | antlr: error(134): vba.g4:569:44: symbol type conflicts with generated code in target language or runtime |
| ❌  | verilog         | antlr: error(134): Verilog2001.g4:1944:24: symbol range conflicts with generated code in target language or runtime |
| ❌  | vhdl            |  test: FAIL  bramp.net/antlr4-grammars/vhdl  10.069s                        |
| ❌  | webidl          | antlr: error(134): WebIDL.g4:584:3: symbol type conflicts with generated code in target language or runtime |
| ❌  | z               | build: z/z_lexer.go:1231:2: syntax error: non-declaration statement outside function body |
| ❌  | z               | build: z/zoperator_parser.go:146:10: escape sequence is invalid Unicode code point |

Less than 50% of the grammars build successfully. This is due to limitations in the ANTLR Go target. The failures are broken down like so:

* **antlr** - ANTLR failed to generate Go code from the grammar.
* **build** - The generated Go code failed to build.
* **test**  - The generated Go code failed the unit tests for that language.

If you wish to help fix the situation then please submit fixes back to the [ANTLR Go target](https://github.com/antlr/antlr4/blob/master/tool/src/org/antlr/v4/codegen/target/GoTarget.java), or the [Gramamrs Github Repo](https://github.com/antlr/grammars-v4).

To generate, build and test a single grammar, just run:

```bash
make <name of grammar>
```

## To update all the grammars

```bash
# Fetch the latest set of grammars
git submodule init
git submodule update

# Run make and all the magic will happen
make all

# Update the table in the README.md with the output
```

## Licence (Apache 2)

*This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.*

```
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
