package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"
	return app
}
