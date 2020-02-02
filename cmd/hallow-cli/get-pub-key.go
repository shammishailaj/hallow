package main

import (
	"encoding/base64"
	"fmt"

	"github.com/alex/hallow/kmssigner"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh"
)

var (
	GetPubKeyCommand = &cli.Command{
		Name:   "get-pub-key",
		Usage:  "Gets the SSH public key for the CA KMS key",
		Action: GetPubKey,
	}
)

func GetPubKey(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("Wrong number of arguments")
	}

	sess := session.New()
	signer, err := kmssigner.New(kms.New(sess), c.Args().Get(0))
	if err != nil {
		return err
	}
	sshPubKey, err := ssh.NewPublicKey(signer.Public())
	if err != nil {
		return err
	}
	fmt.Printf(
		"%s %s\n",
		sshPubKey.Type(),
		base64.StdEncoding.EncodeToString(sshPubKey.Marshal()),
	)

	return nil
}
