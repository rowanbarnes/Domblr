# Domblr

DOM assembler web framework using GoLang WASM. Combines the whole web stack into one GoLang codebase.

## Goals

### Efficiency

- Build small WASM files with TinyGo and wasm-opt
- Perform as much computation as possible on the backend before serving to the frontend
- Serve the least content possible for first paint, then slowly stream other data (e.g. WASM file)

### Programmability

- Provide a framework for writing pure Go web applications supporting the backend and frontend
- Simplify the integration of frontend and backend with a Java RMI-inspired API system
- Make the developer feel like they are writing one application, rather than many disjointed codesets
- Automate the process of intelligently dividing responsibilities
- Provide a system for declaratively specifying the widget tree, routes, dynamic behaviour and API endpoints

#### Composition Justification

- Widgets are all created by their struct to give the programmer named parameters
- Widgets setup by tree traversal, ingesting the parameters stored in the struct
- Widgets compose the behaviours required to render the CSS & HTML

#### Sensible Widget Constraints (WIP)

- Widgets expand to parent's constraints on its minor and/or major axis (`100%`)
- Widgets can shrink to child (`fit-content`)
- Widgets can give a specific size (`?px`)
- Behind the scenes, flexbox mode is exclusively used for simplicity
- shrink(expand) = expand(expand)
  - Constraint is passed up tree, affecting CSS of parent
- expand(shrink) = expand(shrink + empty space)?

## Considerations

- TinyGo is not ready to produce production WASM code due to memory leaks and partial support of the standard library
- The chance of backend code being leaked into the WASM binary and causing security nightmares is very real and not much of a concern at this stage 
- This project is mainly an experiment


## TODO
- [ ] Fill out the widgets to produce good-looking pages 
- [x] Generate separate CSS files 
- [ ] Add a router to dynamically handle links
- [ ] Refactor into library with reasonable programmability (reflection?)
- [ ] ...everything else
- [ ] Fix security nightmares (code leakage) during WASM compilation
- [ ] Optimize file sizes, minification, compression, etc
