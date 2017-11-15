// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// make creates the test and example files for the given grammar.
package main

import (
	"bramp.net/antlr4/internal"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const COPYRIGHT = `// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

`

const DOCFILE = `{{template "copyright" .}}
package {{ .PackageName }} // import "bramp.net/antlr4/{{ .PackageName }}"

`

// TESTFILE is the template for a go test file for this grammar.
// It expects to be executed with a pom.
const TESTFILE = `{{template "copyright" .}}
// Package {{ .PackageName }}_test contains tests for the {{ .Project.LongName }} grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by make.go
//
package {{ .PackageName }}_test

import (
	"bramp.net/antlr4/{{ .PackageName }}"
{{ if .Project.HasParser -}}
	"bramp.net/antlr4/internal"
{{- end }}

{{ if .Project.HasParser }}
	"fmt"
{{ end -}}
{{ if or (eq .Project.CaseInsensitiveType "UPPER") (eq .Project.CaseInsensitiveType "lower") -}}
	"strings"
{{ end -}}
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
{{- range $_, $example := .Project.Examples }}
	{{ printf "%q" . }},
{{- end }}
}

{{ if .Project.HasParser }}
type exampleListener struct {
	*{{ .PackageName }}.Base{{ .Project.ListenerName }}
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
{{ end -}}

func Example() {
	{{- if eq .Project.CaseInsensitiveType "UPPER" }}
	// Setup the input (which this parser expects to be uppercased).
	is := antlr.NewInputStream(strings.ToUpper("...some text to parse..."))
	{{ else if eq .Project.CaseInsensitiveType "lower" }}
	// Setup the input (which this parser expects to be lowercased).
	is := antlr.NewInputStream(strings.ToLower("...some text to parse..."))
	{{ else }}
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")
	{{- end }}

	// Create the Lexer
	lexer := {{ .PackageName }}.New{{ .Project.LexerName }}(is)

{{- if .Project.HasParser }}
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := {{ .PackageName }}.New{{ .Project.ParserName }}(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.{{ .Project.EntryPoint | Title }}()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
{{- else }}
	// There is no {{ .PackageName }} Parser so instead use the Lexer to read tokens.
	t := lexer.NextToken()
	for t.GetTokenType() != antlr.TokenEOF {
		// Do something with the token
		t = lexer.NextToken()
	}
{{ end -}}
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("..", filename))
	if err != nil {
		return nil, err
	}

	{{ if eq .Project.CaseInsensitiveType "UPPER" }}
	input = internal.NewCaseChangingStream(input, true)
	{{ else if eq .Project.CaseInsensitiveType "lower" }}
	input = internal.NewCaseChangingStream(input, false)
	{{ end -}}

	return input, nil
}

func Test{{ .Project.LexerName | Title }}(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := {{ .PackageName }}.New{{ .Project.LexerName }}(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i >= MAX_TOKENS {
			t.Errorf("New{{ .Project.LexerName }}(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

{{ if .Project.HasParser }}
func Test{{ .Project.ParserName | Title }}(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := {{ .PackageName }}.New{{ .Project.LexerName }}(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := {{ .PackageName }}.New{{ .Project.ParserName }}(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(internal.NewTestingErrorListener(t, file))

		// Finally test
		p.{{ .Project.EntryPoint | Title }}()

		// TODO(bramp): If there is a "file.tree", then compare the output
		// TODO(bramp): If there is a "file.errors", then check the error
	}
}
{{ end }}
`

type templateData struct {
	PackageName string
	Project     *internal.Project
}

func create(filename string, t *template.Template, data *templateData) error {
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create %q: %s", filename, err)
	}

	if err := t.Execute(out, data); err != nil {
		return fmt.Errorf("failed to generate %q: %s", filename, err)
	}

	if err := out.Close(); err != nil {
		return fmt.Errorf("failed to close file %q: %s", filename, err)
	}

	if err := exec.Command("go", "fmt", filename).Run(); err != nil {
		return fmt.Errorf("failed to `go fmt %q`: %s", filename, err)
	}

	return nil
}

func usage() {
	// TODO Merge makemake.go into this
	fmt.Fprintf(os.Stderr, "Usage: %s [doc|test] ...\n"+
		"  doc <output>\n"+
		"  test <output> <pom.xml> [<grammar.g4> ...]\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	typ := os.Args[1]
	output := os.Args[2]

	if typ != "doc" && typ != "test" {
		log.Fatalf("Type must be one of doc, test, got: %q", typ)
	}

	copyrightTmpl := template.Must(template.New("copyright").Parse(COPYRIGHT))

	data := &templateData{
		PackageName: output,
	}

	var tmpl *template.Template
	var target string

	if typ == "test" {
		if len(os.Args) < 3 {
			usage()
		}

		pom := os.Args[3]
		project, err := internal.ParsePom(pom)
		if err != nil {
			log.Fatalf("Failed to read pom file %q: %s", pom, err)
		}

		//if len(project.Examples) == 0 {
		//	log.Fatalf("Pom.xml contains zero examples: %q", pom)
		//}

		// Ignore all grammars defined in the pom.xml (because sometimes a single pom may span multiple grammars)
		project.Includes = nil
		project.Grammars = nil

		for _, arg := range os.Args[4:] {
			project.AddGrammar(arg)
		}

		data.Project = project

		funcs := template.FuncMap{
			"Join":    strings.Join,
			"ToCamel": strcase.ToCamel, // I'd prefer to use ToCamel, but the go target does't do this yet...
			"Title":   strings.Title,
		}

		tmpl = template.Must(copyrightTmpl.New("test").Funcs(funcs).Parse(TESTFILE))
		target = filepath.Join(output, output+"_test.go")

	} else if typ == "doc" {
		tmpl = template.Must(copyrightTmpl.New("doc").Parse(DOCFILE))
		target = filepath.Join(output, "doc.go")

	} else {
		panic(fmt.Sprintf("Unexpected type %q want doc or test", typ))
	}

	if err := create(target, tmpl, data); err != nil {
		log.Fatalf("%s: %s", typ, err)
	}
}
