# gRPC Tickets

This project is an implementation example of microservices interacting using gRPC protocol via protobuf.
Basically, there are two services

- calculator: responsible for dealing with calc operations 
- factors: responsible for calculating the sum factor based on the criteria

Services:

- tickets --> golang 
    rpc ListTickets() ListTicketResponse
    rpc CalcTickets(TicketID, Age) (int32, error)

- discounts --> TypeScript
    rpc ApplyDiscounts(Age) (int32, error)