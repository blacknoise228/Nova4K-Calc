androidbuild:
	fyne package -os android -appID com.blacknoise228.NovaStarResolutionCalc

windowsbuild:
	fyne-cross windows -arch amd64 -app-id com.blacknoise228.NovaStarResolutionCalc

build: 
	go build -o NovaStarResolutionCalc

.PHONY: androidbuild windowsbuild build