# practice-grpc-server-to-server

test server-to-server transitions using only unary RPCs between 2 Connect servers.

![sequence dialog](http://www.plantuml.com/plantuml/proxy?src=https://gist.githubusercontent.com/kusa-mochi/8b0901ba27fe2097a00b764e2ec391a7/raw/3158096001d2841335dabb06d75f688a0a452fb4/server_to_server_connect.puml)
<!-- ```plantuml
@startuml

actor You as you
box ServerA
    participant "main-A Go routine" as main_a
    participant "server A Go routine" as server_a
    participant "RPC handler A Go routine" as handler_a
end box
box ServerB
    participant "main-B Go routine" as main_b
    participant "server B Go routine" as server_b
    participant "RPC handler B Go routine" as handler_b
end box

== Start servers ==
you ->> main_a ** : start
activate main_a
main_a -> server_a **
activate server_a
server_a -> server_a : start server
main_a -> main_a : listen keyboard input

you ->> main_b ** : start
activate main_b
main_b -> server_b **
activate server_b
server_b -> server_b : start server
main_b -> main_b : listen keyboard input

== Server A to Server B ==
you ->> main_a : keyboard input
main_a -> server_b : call unary RPC with data.
server_b -> handler_b **
activate handler_b
handler_b -> handler_b : print "Unary RPC B is called."
handler_b -> handler_b : print data.
return
destroy handler_b

== Server B to Server A ==
you ->> main_b : keyboard input
main_b -> server_a : call unary RPC with data.
server_a -> handler_a **
activate handler_a
handler_a -> handler_a : print "Unary RPC A is called."
handler_a -> handler_a : print data.
return
destroy handler_a

@enduml
``` -->
