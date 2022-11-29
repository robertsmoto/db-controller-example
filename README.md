# db-controller-example
Example of db connection to app using controllers.

Thanks to: https://gist.githubusercontent.com/praveen001, with blog post: https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

# Proof-of-Concept
This repo is simply a proof-of-concept based on praveeen001's database organization. I didn't particularly like the organization and naming conventions he used, but the interface-model to controller project layout was interesting.

I've since moved these ideas into a private working repository, but left the basic framework here for somebody who is doing their own research and curious about some db organization ideas.

# Base on Redis Json
The project this was ultimately moved to was a RedisJson supported data layer, so the function inputs and signature returns may not be ideal if this was moving to a tradiational Relational db.

# Testing
Testing with this layout can be done without connecting to a live db as it separates the db connection (and configuration) from the models. This makes TDD much more enjoyable during development.

# Did not Complete
I didn't complete this repo, so I'm sure there are a bunch of bugs and code that isn't needed or warranted. It's simpy a quick-and-dirty proof of concept. Use it as you like.
