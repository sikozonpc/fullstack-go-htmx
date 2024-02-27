# An example of a Full Stack Go web app with HTMX + Templ + Tailwindcss

This projects shows how to build a minimal full stack web app using Go, HTMX, Templ and Tailwindcss.
Everything from the backend and no JavaScript shenanigans.

## Structure

All the HTML are stored in `*.templ` files in the `/views` and `/components` directories.
The `/handlers` directory contains the Go handlers for the different routes that serve those Templ components.

## Installation

There are a few tools that you need to install to run the project.
So make sure you have the following tools installed on your machine.

- [Templ (for the UI layer)](https://templ.guide/quick-start/installation)
- [Tailwindcss CLI (CSS styling library)](https://tailwindcss.com/docs/installation)
- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)
- [Air (for live reloading)](https://github.com/cosmtrek/air)

## Running the project

Firstly make sure you have a MySQL database running on your machine or just swap for any storage you like under `/store`.

Then, for the best development experience, run the project in 3 different processes by running:
```bash
make air # for the go server live reloading
make tailwind # for the tailwindcss live reloading
make templ # for the templ files live generation and reloading
```

This will create a live reloading experience in the browser with real-time changes to the UI and the Go server.