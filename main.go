package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inconshreveable/log15"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	version   string
	gitCommit string
	gitTag    string
	log       = log15.New()
)

func fullVersion() string {
	versionMeta := "release"
	if gitTag == "" {
		versionMeta = "dev"
	}
	return fmt.Sprintf("%s-%s-%s", version, gitCommit, versionMeta)
}

func main() {
	app := cli.App{
		Version:   fullVersion(),
		Name:      "Floyd",
		Usage:     "Node of Floyd Economy Network",
		Copyright: "2020 Floyd Economy <https://www.github.com/floydeconomy/>",
		Flags:     []cli.Flag{
			// networkFlag,
			// configDirFlag,
			// dataDirFlag,
			// cacheFlag,
			// beneficiaryFlag,
			// targetGasLimitFlag,
			// apiAddrFlag,
			// apiCorsFlag,
			// apiTimeoutFlag,
			// apiCallGasLimitFlag,
			// apiBacktraceLimitFlag,
			// verbosityFlag,
			// maxPeersFlag,
			// p2pPortFlag,
			// natFlag,
			// bootNodeFlag,
			// skipLogsFlag,
			// pprofFlag,
			// verifyLogsFlag,
		},
		Action: defaultAction,
		// Commands: []cli.Command{
		// 	{
		// 		Name:  "solo",
		// 		Usage: "client runs in solo mode for test & dev",
		// 		Flags: []cli.Flag{
		// 			dataDirFlag,
		// 			apiAddrFlag,
		// 			apiCorsFlag,
		// 			apiTimeoutFlag,
		// 			apiCallGasLimitFlag,
		// 			apiBacktraceLimitFlag,
		// 			onDemandFlag,
		// 			persistFlag,
		// 			gasLimitFlag,
		// 			verbosityFlag,
		// 			pprofFlag,
		// 			verifyLogsFlag,
		// 		},
		// 		Action: soloAction,
		// 	},
		// 	{
		// 		Name:  "master-key",
		// 		Usage: "master key management",
		// 		Flags: []cli.Flag{
		// 			configDirFlag,
		// 			importMasterKeyFlag,
		// 			exportMasterKeyFlag,
		// 		},
		// 		Action: masterKeyAction,
		// 	},
		// },
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func defaultAction(ctx *cli.Context) error {
	// exitSignal := handleExitSignal()
	//
	// defer func() { log.Info("exited") }()

	// initLogger(ctx)
	// gene, forkConfig := selectGenesis(ctx)
	// instanceDir := makeInstanceDir(ctx, gene)
	//
	// mainDB := openMainDB(ctx, instanceDir)
	// defer func() { log.Info("closing main database..."); mainDB.Close() }()
	//
	// skipLogs := ctx.Bool(skipLogsFlag.Name)
	//
	// logDB := openLogDB(ctx, instanceDir)
	// defer func() { log.Info("closing log database..."); logDB.Close() }()
	//
	// chain := initChain(gene, mainDB, logDB)
	// master := loadNodeMaster(ctx)
	//
	// printStartupMessage1(gene, chain, master, instanceDir, forkConfig)
	//
	// if !skipLogs {
	// 	if err := syncLogDB(exitSignal, chain, logDB, ctx.Bool(verifyLogsFlag.Name)); err != nil {
	// 		return err
	// 	}
	// }
	//
	// txPool := txpool.New(chain, state.NewCreator(mainDB), defaultTxPoolOptions)
	// defer func() { log.Info("closing tx pool..."); txPool.Close() }()
	//
	// p2pcom := newP2PComm(ctx, chain, txPool, instanceDir)
	// apiHandler, apiCloser := api.New(
	// 	chain,
	// 	state.NewCreator(mainDB),
	// 	txPool,
	// 	logDB,
	// 	p2pcom.comm,
	// 	ctx.String(apiCorsFlag.Name),
	// 	uint32(ctx.Int(apiBacktraceLimitFlag.Name)),
	// 	uint64(ctx.Int(apiCallGasLimitFlag.Name)),
	// 	ctx.Bool(pprofFlag.Name),
	// 	skipLogs,
	// 	forkConfig)
	// defer func() { log.Info("closing API..."); apiCloser() }()
	//
	// apiURL, srvCloser := startAPIServer(ctx, apiHandler, chain.GenesisBlock().Header().ID())
	// defer func() { log.Info("stopping API server..."); srvCloser() }()
	//
	// printStartupMessage2(apiURL, getNodeID(ctx))
	//
	// p2pcom.Start()
	// defer p2pcom.Stop()
	//
	// return node.New(
	// 	master,
	// 	chain,
	// 	state.NewCreator(mainDB),
	// 	logDB,
	// 	txPool,
	// 	filepath.Join(instanceDir, "tx.stash"),
	// 	p2pcom.comm,
	// 	uint64(ctx.Int(targetGasLimitFlag.Name)),
	// 	skipLogs,
	// 	forkConfig).
	// 	Run(exitSignal)

	return nil
}

func handleExitSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exitSignalCh := make(chan os.Signal)
		signal.Notify(exitSignalCh, os.Interrupt, os.Kill, syscall.SIGTERM)

		select {
		case sig := <-exitSignalCh:
			log.Info("exit signal received", "signal", sig)
			cancel()
		}
	}()
	return ctx
}
