# Unleash Demo

This repository is a demo app for showcasing Unleash together with the Go
client. To run the demo, simply clone the repo and run:

```shell script
docker-compose up
```

This command will download the relevant Docker images and build the demo app
locally. It will then run 3 containers:

 * PostgreSQL 11.2 (http://localhost:5432)
 * Unleash 3.2 (http://localhost:4242)
 * Demo App (http://localhost:3000)
 
 First you should verify that everything is working by navigating to the
 [app](http://localhost:3000). You should see a fairly boring webapp with
 some explanatory text.
 
 The next step is to enable some feature toggles. This is done by navigating
 to the [Unleash](http://localhost:4242) service. The demo app currently
 supports these features:
 
  * NyanCat
  * RickRoll
  * PinkNavBar
  * RedirectToNewPage
  
These features need to be added to Unleash before they have any effect on the
demo app. The simplest way to enable them is to use the **default** activation
strategy and then enable the feature.

NOTE: When entering the feature toggles in Unleash, the name must correspond
*exactly* (case sensitive) to the code.

The Demo App uses [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
so changes to Go code will automatically rebuild the app.

For more information, please see https://unleash.github.io/  
 
 