# Covid Tracker

## Go API

First install these 3 libraries:

```
"github.com/gorilla/mux"
"github.com/rs/cors"
"github.com/gocolly/colly"
```

To run the api type: `go run api.go`

## Web App

This project uses [nano-react-app](https://github.com/adrianmcli/nano-react-app).

- `npm start` — This will spawn a development server with a default port of `1234`.
- `npm run build` — This will output a production build in the `dist` directory.

### Custom port

You can use the `-p` flag to specify a port for development. To do this, you can either run `npm start` with an additional flag:

```
npm start -- -p 3000
```

Or edit the `start` script directly:

```
parcel index.html -p 3000
```