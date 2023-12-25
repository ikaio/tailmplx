Go + HTMX + Templ + Tailwind

# Setup
```sh
go install github.com/cosmtrek/air@latest
go get

npm install -g pnpm
cd styles
pnpm i

cd ..
air
```

# Build
```sh
mkdir build
cp -r public build
cp -r database build
cp -r tmp/main build/main
cp -r .env build/.env
zip -r build build
rm -rf build
```

## TailwindCSS IntelliSense
Getting it working by mapping it to HTML.
```js
// settings.json
"tailwindCSS.includeLanguages": {
    "templ": "html"
}
```