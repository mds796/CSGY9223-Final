# M3: Matheus Vieira Portela, Miguel David Salcedo, Mel Savich

## Architecture
```bash
.
└── CSGY9223-Final
    ├── auth
    │   ├── auth.go
    │   ├── authpb
    │   │   ├── service.pb.go
    │   │   └── service.proto
    │   ├── config.go
    │   ├── errors.go
    │   ├── service.go
    │   └── service_test.go
    ├── cmd
    │   ├── auth.go
    │   ├── feed.go
    │   ├── follow.go
    │   ├── post.go
    │   ├── root.go
    │   ├── server.go
    │   ├── user.go
    │   └── web.go
    ├── feed
    │   ├── config.go
    │   ├── feed.go
    │   ├── feedpb
    │   │   ├── feed.pb.go
    │   │   ├── feed.proto
    │   │   ├── service.pb.go
    │   │   └── service.proto
    │   ├── service.go
    │   └── service_test.go
    ├── follow
    │   ├── config.go
    │   ├── errors.go
    │   ├── follow.go
    │   ├── followpb
    │   │   ├── follow.pb.go
    │   │   ├── follow.proto
    │   │   ├── service.pb.go
    │   │   └── service.proto
    │   ├── service.go
    │   └── service_test.go
    ├── main.go
    ├── post
    │   ├── config.go
    │   ├── errors.go
    │   ├── post.go
    │   ├── postpb
    │   │   ├── post.pb.go
    │   │   ├── post.proto
    │   │   ├── service.pb.go
    │   │   └── service.proto
    │   ├── service.go
    │   └── service_test.go
    ├── storage
    │   ├── errors.go
    │   ├── raft.go
    │   ├── storage.go
    │   └── stub.go
    ├── user
    │   ├── config.go
    │   ├── errors.go
    │   ├── service.go
    │   ├── service_test.go
    │   ├── user.go
    │   └── userpb
    │       ├── service.pb.go
    │       └── service.proto
    └── web
        ├── auth.go
        ├── config.go
        ├── feed.go
        ├── feed_test.go
        ├── follow.go
        ├── follows_test.go
        ├── follow_test.go
        ├── login_test.go
        ├── logout_test.go
        ├── post_test.go
        ├── register_test.go
        ├── static_test.go
        ├── url_parameters.go
        ├── web.go
        └── web_test.go
```
### HashiCorp

## Demo Showing All the Functionality of the UI
* Register(`professor`)
* Login(`professor`)
* Post(`professor`, `There are only two hard problems in distributed systems:  2. Exactly-once delivery 1. Guaranteed order of messages 2. Exactly-once delivery`)
* Logout(`professor`)
* Login(`mks629`)
* Search(`p`)
* Follow(`mks629`, `professor`)
  * Sees post by `professor`
* Unfollow(`mks629`, `professor`)
  * Doesn't see post by `professor`
* Logout(`mks629`)

## Problems
* **RPC Messages from Stage 2 to Stage 3**
  * Our RPC messages contained different types such as `int`, object `Cookie`, and pointers to miscellaneous objects. We ended up using one coherent model; all data was serialized as bytes, and once received, it was deserialized.

## What We Learned
* **Good design early on is *critical*.**
  * We didn't have to go back and redo *too much* work.
* **Unit test often.**
  * Testing made trying new things easier. It helped us decide early on if something was going to work, or it wasn't.
