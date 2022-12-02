# db-controller-example
Example of db connection to app using controllers.

Thanks to: https://gist.githubusercontent.com/praveen001, with blog post: https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

Also succinct blog post here: https://medium.com/@kramankishore/data-access-layer-dao-why-is-it-needed-how-to-structure-it-47d00d84f00c

# Proof-of-Concept
This repo is simply a proof-of-concept based on praveeen001's database organization. I didn't particularly like the organization and naming conventions he used, but the interface-model to controller project layout was interesting.

I think the naming conventions are better suited as controller (func main() where the router and server lives) models (where the structs live) repo (where the persistent data lives) and DAL (which provides and interface between controller <--> repo and models <--> repo.

I've since moved these ideas into a private working repository, but left the basic framework here for somebody who is doing their own research and curious about ways to separate the model/db dependencies.

# Base on Redis Json
The project this was ultimately moved to was a RedisJson supported data layer, so the function inputs and signature returns may not be ideal if this was moving to a tradiational Relational db. It likely would be similar for other key/value stores or document databases.

# Testing
Testing with this layout can be done without connecting to a live db as it separates the db connection (and configuration) from the models. This makes TDD much more enjoyable during development.

# Did not Complete
I didn't complete this repo, so I'm sure there are a bunch of bugs and code, such as things that aren't needed or warranted. It's simpy a quick-and-dirty proof of concept. Use it as you like.
