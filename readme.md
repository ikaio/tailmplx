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
rm -rf public/upload
mkdir public/upload
cp -r public build
cp -r tmp/main build/main
cp -r .env.dev build/.env.dev
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