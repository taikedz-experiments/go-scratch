bin/filter-map: src/filter-map.go
	mkdir -p bin
	go build -o bin/filter-map src/filter-map.go

run: bin/filter-map
	./bin/filter-map
