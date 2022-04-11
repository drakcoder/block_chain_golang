package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var bc BlockChain
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to db")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("block_chain").Collection("block_chain")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/getData", func(c *fiber.Ctx) error {
		blocks_ref, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var blocks []Block
		err = blocks_ref.All(ctx, &blocks)
		if err != nil {
			log.Fatal(err)
		}
		TOP = 0
		for i := range blocks {
			if !blocks[i].Mined {
				break
			}
			TOP++
		}
		LENGTH = len(blocks)
		fmt.Println(TOP, LENGTH)
		bc = BlockChain{Blocks: blocks}
		return c.JSON(bc)
	})
	app.Get("/mine", func(c *fiber.Ctx) error {
		id := c.Query("id")
		idi, error := strconv.Atoi(id)
		if error != nil {
			return c.SendStatus(502)
		}
		status, b1, b2 := bc.mine(idi)
		collection.UpdateOne(ctx, bson.M{"StringData": b1.StringData}, bson.D{{"$set", bson.D{{"Hash", b1.Hash}, {"Mined", b1.Mined}}}})
		collection.UpdateOne(ctx, bson.M{"StringData": b2.StringData}, bson.D{{"$set", bson.D{{"PrevHash", b2.PrevHash}}}})
		if !status {
			return c.SendStatus(502)
		}
		return c.SendStatus(200)
	})
	app.Get("/add_block", func(c *fiber.Ctx) error {
		return c.Render("add_block", fiber.Map{})
	})
	app.Post("/add_block", func(c *fiber.Ctx) error {
		payload := struct {
			Data string `json:"Data"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			fmt.Println(err)
			return c.SendStatus(502)
		}
		b := bc.add_block(payload.Data)
		d := []interface{}{bson.D{{"StringData", b.StringData}, {"Data", b.Data}, {"PrevHash", b.PrevHash}, {"Hash", b.Hash}, {"nonce", b.nonce}, {"Mined", b.Mined}}}
		res, err := collection.InsertOne(ctx, d[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		return c.SendStatus(201)
	})

	log.Fatal(app.Listen(":8080"))
}
