# M3
> Matheus Vieira Portela, Mel Savich, Miguel David Salcedo

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
* Register user `professor`
* Login user `professor`
* Post under user `professor`
* Logout user `professor`
* Login user `mks629`
* Search for user `professor`
* Follow `professor`
  * Sees post by `professor`
* Unfollow `professor`
  * Doesn't see post by `professor`
* Logout user `mks629`

## Problems
* **RPC messages from Stage 2 to Stage 3**
  * Our RPC messages contained different types such as `int`, object `Cookie`, and pointers to miscellaneous objects. We ended up using one coherent model; all data was serialized as bytes, and once received, it was deserialized.

## What We Learned
* Good design early on is *critical*.
