# fap [![CLA assistant](https://cla-assistant.io/readme/badge/kirillDanshin/fap)](https://cla-assistant.io/kirillDanshin/fap) [![ReviewNinja](https://app.review.ninja/62991967/badge)](https://app.review.ninja/kirillDanshin/fap)

Go fap

## **WARNING!** EARLY BETA!

Please send pull requests and send bugs if you catch one.

Feauture requests are welcome!

# Install

`go get -u github.com/kirillDanshin/fap`

# Use

## Init from web template

Create a project directory, `cd` there and run

`fap init from web`

It'll initialize git repository with readme, .gitignore and code from web template.

### Build

`fap build`

Builds your package with qtc templates and runs `go install`

### Debug build

`fap build -d`

Builds your package with qtc templates and runs `go install -tags "debug"`

### Release

`fap release "Some message"`

Adds all changes to git, commit them with "Some message" and pushes it to origin/master

# TODO

- [x] `fap init`
- [x] `fap build` builds qtc templates and runs go install
- [x] `fap release` (alias for `git add .; git commit; git push`)
- [ ] fap/web helper package
- [ ] fapfile support
- [ ] fap template API
- [ ] `fap deploy`
- [ ] `fap test`
