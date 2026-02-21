package manager

import (
	"os"
	"strings"

	"github.com/trustwallet/assets-go-libs/file"
	"github.com/trustwallet/assets-go-libs/path"
	"github.com/trustwallet/assets/internal/config"
	"github.com/trustwallet/assets/internal/processor"
	"github.com/trustwallet/assets/internal/report"
	"github.com/trustwallet/assets/internal/service"
	"github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configPath, root string

func InitCommands() {
	rootCmd.Flags().StringVar(&configPath, "config", ".github/assets.config.yaml",
		"config file (default is $HOME/.github/assets.config.yaml)")
	rootCmd.Flags().StringVar(&root, "root", ".", "root path to files")

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(fixCmd)
	rootCmd.AddCommand(updateAutoCmd)
	rootCmd.AddCommand(addTokenCmd)
	rootCmd.AddCommand(addTokenlistCmd)
	rootCmd.AddCommand(addTokenlistExtendedCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:   "assets",
		Short: "",
		Long:  "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Wykonaj sprawdzenia walidacyjne",
		Run: func(cmd *cobra.Command, args []string) {
			assetsService := InitAssetsService()
			assetsService.RunJob(assetsService.Check)
		},
	}
	fixCmd = &cobra.Command{
		Use:   "fix",
		Short: "Wykonaj automatyczne poprawki tam, gdzie to możliwe",
		Run: func(cmd *cobra.Command, args []string) {
			assetsService := InitAssetsService()
			assetsService.RunJob(assetsService.Fix)
		},
	}
	updateAutoCmd = &cobra.Command{
		Use:   "update-auto",
		Short: "Uruchom automatyczne aktualizacje z zewnętrznych źródeł",
		Run: func(cmd *cobra.Command, args []string) {
			assetsService := InitAssetsService()
			assetsService.RunUpdateAuto()
		},
	}

	addTokenCmd = &cobra.Command{
		Use:   "add-token",
		Short: "Tworzy szablon info.json dla zasobu",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("Oczekiwano 1 argumentu")
			}

			err := CreateAssetInfoJSONTemplate(args[0])
			if err != nil {
				log.Fatalf("Nie można utworzyć szablonu info.json zasobu: %v", err)
			}
		},
	}

	addTokenlistCmd = &cobra.Command{
		Use:   "add-tokenlist",
		Short: "Dodaje token do tokenlist.json",
		Run: func(cmd *cobra.Command, args []string) {
			handleAddTokenList(args, path.TokenlistDefault)
		},
	}

	addTokenlistExtendedCmd = &cobra.Command{
		Use:   "add-tokenlist-extended",
		Short: "Dodaje token do tokenlist-extended.json",
		Run: func(cmd *cobra.Command, args []string) {
			handleAddTokenList(args, path.TokenlistExtended)
		},
	}
)

func handleAddTokenList(args []string, tokenlistType path.TokenListType) {
	if len(args) != 1 {
		log.Fatal("Oczekiwano 1 argumentu")
	}

	c, tokenID, err := asset.ParseID(args[0])
	if err != nil {
		log.Fatalf("Nie można przetworzyć tokena: %v", err)
	}

	chain, ok := coin.Coins[c]
	if !ok {
		log.Fatal("Nieprawidłowy token")
	}

	err = AddTokenToTokenListJSON(chain, args[0], tokenID, tokenlistType)
	if err != nil {
		log.Fatalf("Nie można dodać tokena: %v", err)
	}
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func InitAssetsService() *service.Service {
	setup()

	paths, err := file.ReadLocalFileStructure(root, config.Default.ValidatorsSettings.RootFolder.SkipFiles)
	if err != nil {
		log.WithError(err).Fatal("Nie udało się załadować struktury plików.")
	}

	paths = filter(paths, func(path string) bool {
		for _, dir := range config.Default.ValidatorsSettings.RootFolder.SkipDirs {
			if strings.Contains(path, dir) {
				return false
			}
		}
		return true
	})

	fileService := file.NewService(paths...)
	validatorsService := processor.NewService(fileService)
	reportService := report.NewService()

	return service.NewService(fileService, validatorsService, reportService, paths)
}

func setup() {
	if err := config.SetConfig(configPath); err != nil {
		log.WithError(err).Fatal("Nie udało się ustawić konfiguracji.")
	}

	logLevel, err := log.ParseLevel(config.Default.App.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Nie udało się przetworzyć poziomu logowania.")
	}

	log.SetLevel(logLevel)
}

// Execute dodaje wszystkie podpolecenia do polecenia głównego i odpowiednio ustawia flagi.
// Jest wywoływane przez main.main(). Musi zostać wykonane tylko raz dla rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
