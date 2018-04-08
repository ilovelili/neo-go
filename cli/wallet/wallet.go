package wallet

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var (
	errNoPath          = errors.New("target path where the wallet should be stored is must be passed using (--path, -p) flags")
	errPhraseMissmatch = errors.New("The entered passphrases do not match. Maybe you have misspelled them?")
)

// NewCommand creates a new Wallet command.
func NewCommand() cli.Command {
	return cli.Command{
		Name:  "wallet",
		Usage: "create, open and manage a NEO wallet",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Usage:  "create a new wallet",
				Action: createWallet,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "path, p",
						Usage: "Target location of the wallet file.",
					},
					cli.BoolFlag{
						Name:  "account, a",
						Usage: "Create a new account",
					},
				},
			},
			{
				Name:   "open",
				Usage:  "open a existing NEO wallet",
				Action: openWallet,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "path, p",
						Usage: "Target location of the wallet file.",
					},
				},
			},
		},
	}
}

// openWallet open a wallet
func openWallet(ctx *cli.Context) error {
	return nil
}

// createWallet create a wallet
func createWallet(ctx *cli.Context) error {
	path := ctx.String("path")
	if len(path) == 0 {
		return cli.NewExitError(errNoPath, 1)
	}
	wall, err := wallet.NewWallet(path)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	if err := wall.Save(); err != nil {
		return cli.NewExitError(err, 1)
	}

	if ctx.Bool("account") {
		if err := createAccount(ctx, wall); err != nil {
			return cli.NewExitError(err, 1)
		}
	}

	dumpWallet(wall)
	fmt.Printf("wallet succesfully created, file location is %s\n", wall.Path())
	return nil
}

// createAccount create an account
func createAccount(ctx *cli.Context, wall *wallet.Wallet) error {
	var (
		rawName,
		rawPhrase,
		rawPhraseCheck []byte
	)

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the account > ")
	rawName, _ = buf.ReadBytes('\n')
	fmt.Print("Enter passphrase > ")
	rawPhrase, _ = buf.ReadBytes('\n')
	fmt.Print("Confirm passphrase > ")
	rawPhraseCheck, _ = buf.ReadBytes('\n')

	// Clean data
	var (
		name        = sanitize(rawName)
		phrase      = sanitize(rawPhrase)
		phraseCheck = sanitize(rawPhraseCheck)
	)

	if phrase != phraseCheck {
		return errPhraseMissmatch
	}

	return wall.CreateAccount(name, phrase)
}

// dumpWallet show wallet in JSON
func dumpWallet(wall *wallet.Wallet) {
	b, _ := wall.JSON()
	fmt.Println("")
	fmt.Println(string(b))
	fmt.Println("")
}

func sanitize(input []byte) string {
	return strings.TrimRight(string(input), "\n")
}
