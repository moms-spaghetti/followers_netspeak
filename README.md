# Followers microservice
Uses Go, gRPC, Protobuf

# MakeFile
Includes various commands to assist rebuilding protobufs, build and running the microservice
- `make proto` regenerates protobuf go code
- `make build` builds executable binary and places in root folder
- `make run` runs the microservice
- `make build_run` builds the executable binary and runs the microservice 

# Protobuf RPCs
- `GetUser` gets a single user
- `GetAllUsers` gets all users
- `CreateUser` creates a single user
- `FollowUser` updates a user to follow another user
- `UnfollowUser` updates a user to unfollow another user

# Development Diary 
- The service is fairly basic, I should have used an existing design architecture to better separate concerns
- Ran out of time writing tests but I've included some basic ones at the server level
- I should have included a business logic layer to move out some of the logic from the RPC functions
- Also should have included a storage layer
- Model converters are missing tests
- Protobuf file is included in the package but I should probably separate this out
- The mocked storage is pretty rudimentary, I could have probably found a package to do this 
- Should have probably separated out the follower data into a separate storage
- The project hasn't been linted
- Project is tightly coupled, separating out would have made it more extensible
- Lack of service interfaces prevents me from mocking
- With more time I'd probably refactor most of the project and follow something like hexagonal architecture
- There's no authentication service
- Getting a user returns users they follow and users who follow them. I probably should have added RPCs to get this data
