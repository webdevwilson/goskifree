deps:
	go get github.com/gopherjs/gopherjs/js
	go get honnef.co/go/js/dom

clean:
	rm goskifree.js goskifree.js.map

build: deps
	gopherjs build -o goskifree.js

serve: deps
	gopherjs serve