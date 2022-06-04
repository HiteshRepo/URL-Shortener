## URL SHORTENER API

### ENDPOINTS
1. /shorten_url
```
curl -X POST http://localhost:9091/shorten_url -H "Content-Type: application/json" -d '{"long_url": "https://1very_log_url_to_be_shortened12"}'
```
2. /fetch_url
```
curl -X GET http://localhost:9091/fetch_url -H "Content-Type: application/json" -d '{"short_url": "https://bitly.com/868dG85"}'
```

### DOCKER COMMANDS
1. Build the image: ```docker build -t url_shortener -f build/package/Dockerfile .```
2. Run the image: ```docker run -p 9091:9091 -e PORT=9091 -e HOST=0.0.0.0 url_shortener```
3. Tag image: ```docker tag 8697e1c1a2e9 hiteshpattanayak/url_shortener:1.0```
4. Push image to docker hub: ```docker push hiteshpattanayak/url_shortener:1.0```


### RUN TESTS
1. All tests: gotestsum ./...
2. Integration test path: internal/app/app_integration_test

### FOLDER STRUCTURE REFERENCE
https://github.com/golang-standards/project-layout

### STANDARDS REFERENCE
https://github.com/uber-go/guide/blob/master/style.md