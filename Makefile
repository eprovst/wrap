.PHONY: test
test:
	@go test -v ./...

.PHONY: clean
clean:
	@rm -rf build

.PHONY: distclean
distclean: clean
	@rm -rf dist

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

prime/bash-complete.sh prime/zsh-complete.sh:
	@go run scripts/generate_completions.go
	@mkdir -p ./prime
	@mv ./bash-complete.sh prime/bash-complete.sh
	@mv ./zsh-complete.sh prime/zsh-complete.sh

build/windows/wrap.exe:
	@go generate # Prepare resource.syso
	@GOARCH=amd64 GOOS=windows go build -o build/windows/wrap.exe -ldflags '-s -w'
	@rm resource.syso # Remove resource.syso

dist/Wrap_Win64.exe: build/windows/wrap.exe
	@mkdir -p ./dist
	@makensis -V2 -DARCH=x64 scripts/installer.nsi

build/darwin/wrap:
	@GOARCH=amd64 GOOS=darwin go build -o build/darwin/wrap -ldflags '-s -w'

dist/Wrap_macOS.zip: build/darwin/wrap
	@mkdir -p ./build/dist/macOS/wrap.app/Contents/MacOS
	@mkdir -p ./build/dist/macOS/wrap.app/Contents/Resources
	@cp ./scripts/Info.plist \
		./build/dist/macOS/wrap.app/Contents/Info.plist
	@cp ./build/darwin/wrap \
		./build/dist/macOS/wrap.app/Contents/MacOS/wrap
	@cp ./assets/wrap/wrap.icns \
		./build/dist/macOS/wrap.app/Contents/Resources/wrapApp.icns
	@cp ./assets/filetypes/wrap/wrap.icns \
		./build/dist/macOS/wrap.app/Contents/Resources/WRAP.icns
	@cp ./assets/filetypes/fountain/fountain.icns \
		./build/dist/macOS/wrap.app/Contents/Resources/FOUNTAIN.icns
	@mkdir -p ./dist
	@zip -q -r ./dist/Wrap_macOS.zip ./build/dist/macOS/wrap.app

.PHONY: release
release: prime/bash-complete.sh prime/zsh-complete.sh \
	dist/Wrap_Win64.exe dist/Wrap_macOS.zip

dist/Wrap_Win64_nightly.exe: dist/Wrap_Win64.exe
	@mv ./dist/Wrap_Win64.exe ./dist/Wrap_Win64_nightly.exe

dist/Wrap_macOS_nightly.zip: dist/Wrap_macOS.zip
	@mv ./dist/Wrap_macOS.zip ./dist/Wrap_macOS_nightly.zip

.PHONY: nightlies
nightlies: prime/bash-complete.sh prime/zsh-complete.sh \
	dist/Wrap_Win64_nightly.exe dist/Wrap_macOS_nightly.zip
