module giao/pkg/e_book

go 1.19

require github.com/PuerkitoBio/goquery v1.8.1

require giao/pkg/util v0.0.0

replace giao/pkg/util => ../util

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	golang.org/x/net v0.7.0 // indirect
)
