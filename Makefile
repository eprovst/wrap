FLAGS = -ldflags '-s -w'

.PHONY: test clean distclean

test:
	go test -v ./...

clean:
	rm -rf build

distclean: clean
	rm -rf dist

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

prime/bash-complete.sh prime/zsh-complete.sh:
	go run scripts/generate_completions.go
	mkdir -p prime
	mv *-complete.sh prime

build: cmd/wrap
	go build ./$<

build/windows/wrap.exe: cmd/wrap
	go generate ./$<
	GOARCH=amd64 GOOS=windows go build -o $@ $(FLAGS) ./$<
	rm $</resource.syso

dist/Wrap_Win64_nightly.exe: build/windows/wrap.exe
	mkdir -p dist
	makensis -V2 -DARCH=x64 scripts/installer.nsi

build/darwin/wrap: cmd/wrap
	GOARCH=amd64 GOOS=darwin go build -o $@ $(FLAGS) ./$<

dist/Wrap_macOS_nightly.zip: build/darwin/wrap
	mkdir -p ./build/macOS/dist/wrap.app/Contents/MacOS
	mkdir -p ./build/macOS/dist/wrap.app/Contents/Resources
	cp scripts/Info.plist \
		build/macOS/dist/wrap.app/Contents/
	cp build/darwin/wrap \
		build/macOS/dist/wrap.app/Contents/MacOS/
	cp assets/wrap/wrap.icns \
		build/macOS/dist/wrap.app/Contents/Resources/wrapApp.icns
	cp assets/filetypes/wrap/wrap.icns \
		build/macOS/dist/wrap.app/Contents/Resources/WRAP.icns
	cp assets/filetypes/fountain/fountain.icns \
		build/macOS/dist/wrap.app/Contents/Resources/FOUNTAIN.icns
	mkdir -p dist
	zip -q -r $@ build/macOS/dist/wrap.app

nightly: prime/bash-complete.sh prime/zsh-complete.sh \
	dist/Wrap_Win64_nightly.exe dist/Wrap_macOS_nightly.zip

dist/Wrap_Win64.exe: dist/Wrap_Win64_nightly.exe
	mv $< $@

dist/Wrap_macOS.zip: dist/Wrap_macOS_nightly.zip
	mv $< $@

release: nightly dist/Wrap_Win64.exe dist/Wrap_macOS.zip
