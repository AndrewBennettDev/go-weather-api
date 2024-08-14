Welcome to my personal Weather API powered by Go! This has been a fun personal project for really learning Golang,
as well as thinking more about what I want in a backend service. As of this version there are three main components:

1) API - using gorilla/mux you can currently call a single endpoint that calls both the Current Weather and Astronomy endpoints from Weather API.
The Transform class then combines these and returns one set of data. There are placeholders for healthchecks now and a "list" endpoint
in case I implement more endpoints later.
2) Transform - each API call returns data that is not useful for me. Rather than store and display things I do not need
I am using a simple transform function. 
3) Database - as of my most recent work this layer now interacts with a locally hosted instance of MariaDB/MySQL. I only have
a basic implementation that allows me to create the table if it does not exist, and write data to that table whenever the primary
endpoint is hit. Other queries to be written soon!

Once these pieces are all functional this will live on my personal website (oh yeah, I need to build that, too...) as a widget,
with the goal of implementing this same basic format in multiple languages so you can see how they perform relative to the others.
This will likely be a running project for many years as I continue to add fun bits to it!

Note: this repo lives on GitHub and GitLab; I love both so there may eventually be deploy files for either/both!

I am always open to constructive criticism and suggestions, so please do not hesitate to point out errors or potential optimisations.
