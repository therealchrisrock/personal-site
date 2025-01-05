tsc & \
npx tailwind \
  -i './static/css/input.css' \
  -o './static/css/output.css' \
  --watch & \
templ generate --watch & \
air

