package app

import (
	"context"
	"golang.org/x/sync/errgroup"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewApp(config *Config) (*App, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.Telegram.Token,
		Poller: &telebot.LongPoller{Timeout: config.Telegram.LongPollerTimeout * time.Second},
	})
	if err != nil {
		return nil, err
	}
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: &serverMux,
	}
	app := &App{
		config: config,
		bot:    bot,
		server: server,
	}
	return app, nil
}

type App struct {
	config *Config
	bot    *telebot.Bot
	server *http.Server
}

func (a *App) Run(ctx context.Context) {
	log.Println("[app] starting")
	signalCtx, signalStop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	defer signalStop()

	wg, gCtx := errgroup.WithContext(signalCtx)

	wg.Go(func() error {
		log.Printf("[app] starting telegram bot in long poller mode")
		a.bot.Start()
		return nil
	})
	wg.Go(func() error {
		<-gCtx.Done()
		log.Println("[app] stopping long polling for telegram bot")
		a.bot.Stop()
		return nil
	})
	wg.Go(func() error {
		log.Printf("[app] starting http server on port %s", a.config.Port)
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println("[app] http server stopped unexpectedly")
			return err
		}
		return nil
	})
	wg.Go(func() error {
		<-gCtx.Done()
		log.Println("[app] shutting down http server")
		serverCtx, serverCancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
		defer serverCancel()
		return a.server.Shutdown(serverCtx)
	})
	wg.Go(func() error {
		<-gCtx.Done()
		signalStop()
		return nil
	})

	if err := wg.Wait(); err != nil {
		log.Println("[app] unexpected exit reason: ", err)
	}
}
