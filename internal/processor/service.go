package processor

import (
	assetsmanager "github.com/trustwallet/assets-go-libs/client/assets-manager"
	"github.com/trustwallet/assets-go-libs/file"
	"github.com/trustwallet/assets/internal/config"
)

type Service struct {
	fileService   *file.Service
	assetsManager assetsmanager.Client
}

func NewService(fileProvider *file.Service) *Service {
	return &Service{
		fileService:   fileProvider,
		assetsManager: assetsmanager.InitClient(config.Default.ClientURLs.AssetsManagerAPI, nil),
	}
}

func (s *Service) GetValidator(f *file.AssetFile) []Validator {
	jsonValidator := Validator{Name: "Walidacja JSON", Run: s.ValidateJSON}

	switch f.Type() {
	case file.TypeAssetFolder:
		return []Validator{
			{Name: "Każdy folder zasobu ma prawidłowy adres i zawiera tylko dozwolone pliki", Run: s.ValidateAssetFolder},
		}
	case file.TypeChainFolder:
		return []Validator{
			{Name: "Foldery łańcucha są pisane małymi literami i zawierają tylko dozwolone pliki", Run: s.ValidateChainFolder},
		}
	case file.TypeChainInfoFolder:
		return []Validator{
			{Name: "Folder info łańcucha (zawiera pliki)", Run: s.ValidateInfoFolder},
		}
	case file.TypeDappsFolder:
		return []Validator{
			{Name: "Folder dapps zawiera tylko dozwolone pliki png pisane małymi literami", Run: s.ValidateDappsFolder},
		}
	case file.TypeRootFolder:
		return []Validator{
			{Name: "Folder główny zawiera tylko dozwolone pliki", Run: s.ValidateRootFolder},
		}
	case file.TypeValidatorsAssetFolder:
		return []Validator{
			{Name: "Folder zasobu walidatorów posiada logo i prawidłowy adres)", Run: s.ValidateValidatorsAssetFolder},
		}

	case file.TypeAssetLogoFile, file.TypeChainLogoFile, file.TypeDappsLogoFile, file.TypeValidatorsLogoFile:
		return []Validator{
			{Name: "Rozmiar i wymiary logo są prawidłowe", Run: s.ValidateImage},
		}
	case file.TypeAssetInfoFile:
		return []Validator{
			jsonValidator,
			{Name: "Plik info zasobu jest prawidłowy", Run: s.ValidateAssetInfoFile},
		}
	case file.TypeChainInfoFile:
		return []Validator{
			{Name: "Plik info łańcucha jest prawidłowy", Run: s.ValidateChainInfoFile},
		}
	case file.TypeTokenListFile:
		return []Validator{
			jsonValidator,
			{Name: "Plik tokenlist jest prawidłowy", Run: s.ValidateTokenListFile},
		}
	case file.TypeTokenListExtendedFile:
		return []Validator{
			jsonValidator,
			{Name: "Rozszerzony plik tokenlist jest prawidłowy", Run: s.ValidateTokenListExtendedFile},
		}
	case file.TypeValidatorsListFile:
		return []Validator{
			jsonValidator,
			{Name: "Plik listy walidatorów jest prawidłowy", Run: s.ValidateValidatorsListFile},
		}
	}

	return nil
}

func (s *Service) GetFixers(f *file.AssetFile) []Fixer {
	jsonFixer := Fixer{
		Name: "Formatowanie wszystkich plików json",
		Run:  s.FixJSON,
	}

	switch f.Type() {
	case file.TypeAssetFolder:
		return []Fixer{
			{Name: "Zmiana nazwy folderu zasobu EVM na prawidłową sumę kontrolną adresu", Run: s.FixETHAddressChecksum},
		}
	case file.TypeAssetInfoFile:
		return []Fixer{
			jsonFixer,
			{Name: "Naprawa plików info.json zasobu", Run: s.FixAssetInfo},
		}
	case file.TypeChainInfoFile:
		return []Fixer{
			jsonFixer,
			{Name: "Naprawa plików info.json łańcucha", Run: s.FixChainInfoJSON},
		}
	case file.TypeChainLogoFile, file.TypeAssetLogoFile, file.TypeValidatorsLogoFile, file.TypeDappsLogoFile:
		return []Fixer{
			{Name: "Zmiana rozmiaru i kompresja obrazów logo", Run: s.FixLogo},
		}
	case file.TypeValidatorsListFile:
		return []Fixer{
			jsonFixer,
		}
	}

	return nil
}

func (s *Service) GetUpdatersAuto() []Updater {
	return []Updater{}
}
