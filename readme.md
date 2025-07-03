2 tables
    slips
        id
        stake
        status
    selections - if a selection is won
        id
        slip_id on to many
        outcome id
        status - won | lost

rabbit mq


---
    procedure
    create tables
        then consume the rabbit mq

rabbit -> 167.99.88.27
    port 35672
    user admin
    pass admin

channel name
    BET-SETTLEMENT

--- starting rabbit mq server
    docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management
    
    then create a channel under channels

