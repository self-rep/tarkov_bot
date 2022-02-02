mkdir ~/bin
go build -o tarkov_bot -ldflags="-s -w" main.go # LDFLAGS will shorten the size of the binary
mv tarkov_bot ~/bin
mv ./assets ~/bin # move ammo data and settings to same folder as the bot binary
