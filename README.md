Welcome to my personal Weather API powered by Go! This has been a fun personal project for really learning Golang,
as well as thinking more about what I want in a backend service. As of this version there are three main components:

1) API - using gorilla/mux you can currently call four endpoints from the free rapidapi weather service.
This will ONLY WORK with the /current/ endpoint right now because of the transform method
2) Transform - each API call returns data that is not useful for me. Rather than store and display things I do not need
I am using a simple transform function. As a POC I only have this implemented for the /current/ endpoint, so you will not
get usable data from the other three endpoints for now
3) Database - this is non-functional at the moment and the database.go file is just a basic template/placeholder. Eventually
this service will call one or all of the endpoints at regular intervals and write the data to a SQL database. This will 
eventually EVENTUALLY be used with some ML to make predicitions for weather and see how it lines up with the actual weather.

Once these pieces are all functional this will live on my personal website (oh yeah, I need to build that, too...) as a widget,
with the goal of implementing this same basic format in multiple languages so you can see how they perform relative to the others.
This will likely be a running project for many years as I continue to add fun bits to it!

I am always open to constructive criticism and suggestions, so please do not hesitate to point out errors or potential optimisations.
