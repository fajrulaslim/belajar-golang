menjalankan semua file unit test, masuk ke folder yang dimaksud
go test

verbose untuk melihat detail
go test -v

go test -v -run TestNamaFunction
go test -v -run=TestNamaFunction

menjalankan test di root utama
go test ./...

go get github.com/stretchr/testify

go mod tidy

go test -v -run=TestNamaFunction/SubTestFunction
go test -v -run /SubTestFunction

Menjalankan benchmark
go test -v -bench.
go test -v -run=NotMatchUnitTest -bench=.
go test -v -bench. ./...
go test -v -run=NotMatchUnitTest -bench=TestNamaFunction/SubTestFunction