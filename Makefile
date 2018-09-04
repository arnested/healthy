.PHONY: doc check-doc

doc: README.md

README.md: *.go .godocdown.tmpl
	godocdown --output=README.md
