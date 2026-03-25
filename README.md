# GoaTTStack
A GO + Alpine.js + Templ + Tailwind Css stack implementation

## Add Dependencies
1. Install templ cli locally in the project as a go tool:
```sh
go get -tool github.com/a-h/templ/cmd/templ@latest
```

    To run templ once installed, use go tool templ instead of templ.
    
2. Add templ package to our project

```sh
go get github.com/a-h/templ
```

## Live Reload

```sh
go tool templ generate -watch --cmd="go run microservices/frontend/main.go" --proxy="http://172.28.77.167:8080"
```
