# CrudApi using GoLang
#### this api calls different routes and then posts or gets data from mongodb accordingly.

### How to go about the codebase
1. ``contracts`` contain the structure for the posts and users
2. ``handler`` contains the functions for each separate task respectively for user and post
3. ``server`` contains the routes and the server config
4. ``database`` contains the functions for MongoClient for each task
5. ``response`` contains the function which returns the JSON response

### Steps to run
1. clone the repo
2. ```$ make compile && make start ```
3. to exit simply press ``ctrl + C``

#### IMP : in  ```database/mongo.go``` please give a uri for your mongodb database