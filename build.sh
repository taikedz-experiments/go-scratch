mkdir -p bin
echo '*' > bin/.gitignore
go build -o bin/filter-map filter-map.go
