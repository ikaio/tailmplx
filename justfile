
install:
    go install github.com/cosmtrek/air@latest
    go mod tidy
    npm install -g pnpm
    cd styles && pnpm i

build:
    templ generate 
    cd styles && pnpm run build 
    go build -o ./tmp/main

release: build
    mkdir release
    cp -r public release
    cp -r tmp/main release/main
    cp -r .env.dev release/.env.dev
    zip -r release release
    rm -rf release
