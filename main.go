package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()	

	
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	token := os.Getenv("TOKEN")
	
	b, err := bot.New(token, opts...)

	if err != nil {
		log.Fatalln("Error al iniciar el bot")
	}
	log.Println("bot running")
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	value := update.Message.Text
	
	costo, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Print("Error al convertir")
		
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text: "Error al ingresar el precio del producto, Intente nuevamente!...",
		})

		return
	}

	margenPorcentaje := 40.0
	iva := 21.0
	impuestoProvincial := 9.0

	precioConMargen := costo * (1 + margenPorcentaje/100)
	
	precioFinal := precioConMargen * (1 + (iva + impuestoProvincial)/100)
	
	precioFinalString := fmt.Sprintf("%.2f", precioFinal)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "Precio Final: $ "+precioFinalString,
	})
}