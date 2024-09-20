package main 

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)




func main()  {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    FailOnErr(err, errs.ConnErr)
    ch, err := conn.Channel()
    FailOnErr(err, errs.ChannelErr)
    q, err := ch.QueueDeclare(

        "orders",
        false,
        false,
        false,
        false,
        nil,


        )
    FailOnErr(err, errs.QueueErr)
    for {
        order := Order{
            ID: "hx0012",
            CutomerID: "hx0012e",
            ItemID: "Some gebberish",
            Quantity: 123432,
        }
        jsonData, _ := json.Marshal(order)
        err := ch.Publish(
            "",
            q.Name,
            false, 
            false, 
            amqp.Publishing{
                ContentType: "application/json",
                Body: jsonData,

            },

            )
        if err != nil {
            log.Fatal("Error publishing to the queue")
        }

        log.Print("Success!")
        time.Sleep(time.Second * 12)
        
    }


}

func FailOnErr(err error, msg string)  {
    if err != nil{
        log.Fatal(err, msg)
    }
}

