# Domblr

DOM assembler web framework using GoLang WASM. Combines the whole web stack into one GoLang codebase.

## Goals

- Build small WASM files with TinyGo and wasm-opt
- Provide a framework for writing pure Go web applications supporting the backend and frontend
- Simplify the integration of frontend and backend with a Java RMI-inspired API system
- Perform as much computation as possible on the backend before serving to the frontend
- Make the developer feel like they are writing one application, rather than many disjointed codesets
- Automate the process of intelligently dividing responsibilities
- Provide a system for declaratively specifying the widget tree, routes, dynamic behaviour and API endpoints

## Considerations

- TinyGo is not ready to produce production WASM code due to memory leaks and partial support of the standard library
- The chance of backend code being leaked into the WASM binary and causing security nightmares is very real and not much of a concern at this stage 
- This project is mainly an experiment

## TODO

1. Fill out the widgets to produce good-looking pages
2. Generate separate CSS files
3. Add a router to dynamically handle links
4. ...?
