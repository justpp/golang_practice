module giao/pkg/e_book

go 1.19

require github.com/PuerkitoBio/goquery v1.8.1

require (
	giao/pkg/util v0.0.0
	github.com/spf13/cobra v1.7.0
)

replace giao/pkg/util => ../util

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.7.0 // indirect
)
