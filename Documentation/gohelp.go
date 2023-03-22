package main

func main() {
	//Doc prints the documentation comments associated with the item identified by its
	//arguments (a package, const, func, type, var, method, or struct field)
	//followed by a one-line summary of each of the first-level items "under"
	//that item (package-level declarations for a package, methods for a type,
	//etc.).
	//
	//Doc accepts zero, one, or two arguments.
	//
	//Given no arguments, that is, when run as
	//
	//        go doc
	//
	//it prints the package documentation for the package in the current directory.

	/*If the package is a command (package main), the exported symbols of the package
	are elided from the presentation unless the -cmd flag is provided.  ???????????????????

	*/

	/*
		When run with one argument, the argument is treated as a Go-syntax-like
		representation of the item to be documented. What the argument selects depends
		on what is installed in GOROOT and GOPATH, as well as the form of the argument,
			which is schematically one of these:

		go doc <pkg>
		go doc <sym>[.<methodOrField>]
		go doc [<pkg>.]<sym>[.<methodOrField>]
		go doc [<pkg>.][<sym>.]<methodOrField>

	*/

	/*
	Examples:
		go doc
		Show documentation for current package.
		go doc Foo
		Show documentation for Foo in the current package.
		(Foo starts with a capital letter so it cannot match
		a package path.)
		go doc encoding/json
		Show documentation for the encoding/json package.
		go doc json
		Shorthand for encoding/json.
		go doc json.Number (or go doc json.number)
		Show documentation and method summary for json.Number.
		go doc json.Number.Int64 (or go doc json.number.int64)
		Show documentation for json.Number's Int64 method.
		go doc cmd/doc
		Show package docs for the doc command.
		go doc -cmd cmd/doc
		Show package docs and exported symbols within the doc command.
		go doc template.new
		Show documentation for html/template's New function.
		(html/template is lexically before text/template)
		go doc text/template.new # One argument
		Show documentation for text/template's New function.
		go doc text/template new # Two arguments
		Show documentation for text/template's New function.

	*/

}
