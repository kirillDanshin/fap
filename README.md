# fap [![CLA assistant](https://cla-assistant.io/readme/badge/kirillDanshin/fap)](https://cla-assistant.io/kirillDanshin/fap) [![ReviewNinja](https://app.review.ninja/62991967/badge)](https://app.review.ninja/kirillDanshin/fap)

Go fap

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

# TODO

- [x] `fap init`
- [x] `fap build` builds qtc templates and runs go install
- [x] `fap release`
- [ ] `fap deploy`
