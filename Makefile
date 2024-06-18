templ:
	templ generate --watch -proxy=http://localhost:3000
tailwind:
	npx tailwindcss -i ./ui/assets/input.css -o ./ui/assets/main.css --watch
