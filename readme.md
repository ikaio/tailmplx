Go + Templ + HTMX + Tailwind
# Setup
1. Make sure you have **Go** & **Node NPM**.
2. Make sure you have the **just** command-runner installed: [just](https://github.com/casey/just)
3. Install: **Air**, **PNPM** and project dependencies: `just install`
4. Finally, you're ready to run: `air`
# Releasing
Build a production release by running: `just release`, it will:

1. Create a `/release` folder with `/public` and the binary in it.
	- (gets deleted afterwards so you won't see it)
2. Zip it into `release.zip`.

