# golang-web-development

The repo shows ways to create a basic webserver and the common things that should be known before starting web development in go. The standard go package net/http is predominently used. Also for extending some functionalities, (mathcing request with -- (specified method,query parameters) I have use gorilla/mux.

When Creating a web App - These are all you need

- net/http package as fundamental programming block for web development
- negroni for implementing middleware
- Gorilla mux as router
- Gorilla context for sharing values during request lifetime
