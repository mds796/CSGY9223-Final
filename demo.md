# M3: Matheus Vieira Portela, Miguel David Salcedo, Mel Savich

## Architecture
![M3 Architecture](https://github.com/mds796/CSGY9223-Final/blob/master/m3.png)

### HashiCorp
* CoreOS is the most widely used Raft library in production, we wanted to try something different.
* There were many well written third-party examples on Hashicorp.

## Demo All the Functionality of the UI
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

## Demo User Scenario That Would Use Replication
Any action that interacts with `storage` uses replication.
* Login(`mks629`)
* Follow(`mks629`, `professor`)
* Logout(`mks629`)
* Login(`professor`)
* \**kill one of the backend servers*\*
* Post(`professor`, `I was gonna tell you guys a joke about UDP, but you might not get it.`)
* Logout(`professor`)
* Login(`mks629`)
  * Sees most recent post by `professor`

## Problems
* **RPC Messages from Stage 2 to Stage 3**
  * Our RPC messages contained different types such as `int`, object `Cookie`, and pointers to miscellaneous objects. We ended up using one coherent model; all data was serialized as bytes, and once received, it was deserialized.

## What We Learned
* **Good design early on is *critical*.**
  * We didn't have to go back and redo *too much* work.
* **Unit test often.**
  * Testing made trying new things easier. It helped us decide early on if something was going to work, or it wasn't.
* **How to build a distributed system with fault tolerance, replication, and consistency.**
