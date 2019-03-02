// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "jwt": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/hirokisan/go-sandbox/go-api/design/jwt
// --out=$(GOPATH)/src/github.com/hirokisan/go-sandbox/go-api/jwt
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/hirokisan/go-sandbox/go-api/jwt/client"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// SecureBasicCommand is the command line data structure for the secure action of basic
	SecureBasicCommand struct {
		PrettyPrint bool
	}

	// UnsecureBasicCommand is the command line data structure for the unsecure action of basic
	UnsecureBasicCommand struct {
		PrettyPrint bool
	}

	// SecureJWTCommand is the command line data structure for the secure action of jwt
	SecureJWTCommand struct {
		// Force auth failure via JWT validation middleware
		Fail        string
		PrettyPrint bool
	}

	// SigninJWTCommand is the command line data structure for the signin action of jwt
	SigninJWTCommand struct {
		PrettyPrint bool
	}

	// UnsecureJWTCommand is the command line data structure for the unsecure action of jwt
	UnsecureJWTCommand struct {
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "secure",
		Short: `secure action`,
	}
	tmp1 := new(SecureBasicCommand)
	sub = &cobra.Command{
		Use:   `basic ["/basic"]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(SecureJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/jwt"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `Creates a valid JWT`,
	}
	tmp3 := new(SigninJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/jwt/signin"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "unsecure",
		Short: `unsecure action`,
	}
	tmp4 := new(UnsecureBasicCommand)
	sub = &cobra.Command{
		Use:   `basic ["/basic/unsecure"]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(UnsecureJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/jwt/unsecure"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the SecureBasicCommand command.
func (cmd *SecureBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecureBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecureBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecureBasicCommand command.
func (cmd *UnsecureBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic/unsecure"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecureBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecureBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecureJWTCommand command.
func (cmd *SecureJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp6 *bool
	if cmd.Fail != "" {
		var err error
		tmp6, err = boolVal(cmd.Fail)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--fail", "err", err)
			return err
		}
	}
	resp, err := c.SecureJWT(ctx, path, tmp6)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecureJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var fail string
	cc.Flags().StringVar(&cmd.Fail, "fail", fail, `Force auth failure via JWT validation middleware`)
}

// Run makes the HTTP request corresponding to the SigninJWTCommand command.
func (cmd *SigninJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/signin"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninJWT(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecureJWTCommand command.
func (cmd *UnsecureJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/unsecure"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecureJWT(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecureJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
