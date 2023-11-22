# go-todo
A todo webapp written in go using [htmx](https://htmx.org)
~~and [tailwindcss](https://tailwindcss.com/),~~
and [Bootstrap 5](https://getbootstrap.com/)
with [sqlite](https://github.com/mattn/go-sqlite3)
on the backend.
(I decided to go with Bootstrap instead of Tailwind right now
because I'm not that familiar with CSS and my focus is just a 
UI that's _good enough_.)

## Roadmap/todos/ideas
- [x] make sure app works by checking it with NerdCademy's template
- [ ] design web app layout on paper
- [ ] implement design with bootstrap
- [ ] make sure app looks good on desktop
- [ ] make sure app looks good on mobile
- [ ] consider dockerizing backend part of app
- [ ] consider trying to add new api route to allow for mobile or cli apps
      that connect to backend
- [ ] see how app plays under a proxy like nginx

## Deployment
### Simple
This method assumes you already have a compiled binary file on your 
machine and that you intend to deploy on some other machine.

Copy the binary file, the templates directoy, and the `run.sh` script
into the app's directory on the remote machine.
A potential workflow could be:

<!-- TODO check if this is right syntax and works -->
``` sh
# susbtitute $REMOTE_HOST for your ssh host and $APP_DIRECTORY for the folder 
# you'll have this app run in on the server.
scp -r ./go-todo ./views ./run.sh $REMOTE_HOST:$APP_DIRECTORY
```

If you'd like to change the port that the app runs on,
you can modify it in `run.sh`'s environment variable.

### Building from source
<!-- TODO check if this is actually correct -->
You can just git clone this repo and run 
`CGO_ENGABLED=1 go build`. 
You have to run it with the `cgo` flag because the dependency
to allow for database interaction needs it to compile
(see the
[installation section](https://github.com/mattn/go-sqlite3#installation)
of the `go-sqlite3` for more details).
A potential workflow could be the below:

``` sh
git clone https://github.com/vim-is-life/go-todo/
CGO_ENABLED=1 go build .
```

You should then be able to run the app using `run.sh`, modifying 
the port it runs on as you desire.

## References
Below are some of the resources I used in creating this app:
- [Code from NerdCademy's webapp in go tutorial](https://github.com/NerdCademyDev/gophat)
  - I kind of followed the structure of his code and used it 
    as a loose reference for the code I wrote for my project
- [HTML Templating With Go](https://www.makeuseof.com/go-html-templating/)
- [How to Use Templates in Go](https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go)
  - I used these as references for using templates in go,
    as I was not familiar with using templates.
- [Using MVC to Structure Go Web Applications](https://www.calhoun.io/using-mvc-to-structure-go-web-applications/)
  - I used this article as a loose reference for ways to
    implement the MVC architecture in Go.
- <http://go-database-sql.org/retrieving.html> 
- <http://go-database-sql.org/modifying.html>
  - I used this reference to understand a bit more how to
    work with Go's database/sql package.
- [Which Go router should I use? (with flowchart)](https://www.alexedwards.net/blog/which-go-router-should-i-use)
  - I used this to understand why NerdCademy used 
    [`gorilla/mux`](https://pkg.go.dev/github.com/gorilla/mux)
    over the `http.ServeMux` in the standard library, and to
    just learn more about the different available options for
    muxes in Go.
- [Bootstrap5 Examples](https://getbootstrap.com/docs/5.3/examples/)
  - I used these examples to help with styling my app's HTML.
