# Trust Wallet Assets Info

![Check](https://github.com/trustwallet/assets/workflows/Check/badge.svg)

## Przegląd

Repozytorium tokenów Trust Wallet to kompleksowa, aktualna kolekcja informacji o kilku tysiącach (!) tokenów kryptograficznych.

[Trust Wallet](https://trustwallet.com) używa logotypów tokenów z tego źródła, podobnie jak wiele innych projektów.

Repozytorium zawiera informacje o tokenach z kilku blockchainów, informacje o dApps, walidatorach stakingu itp.
Dla każdego tokena dostępne jest logo oraz opcjonalne dodatkowe informacje (takie dane nie są dostępne on-chain).

Tak duża kolekcja może być utrzymywana tylko dzięki wysiłkowi społeczności, więc _dodaj swój token_.

<center><img src='https://trustwallet.com/assets/images/media/assets/horizontal_blue.png' height="200"></center>

## Jak dodać token

Pamiętaj, że __zupełnie nowe tokeny nie są akceptowane__,
projekty muszą być wiarygodne, posiadać dostępne informacje oraz __niezerowy obieg__
(szczegóły limitów znajdziesz na stronie <https://developer.trustwallet.com/listing-new-assets/requirements>).

### Aplikacja Assets

[Aplikacja webowa Assets](https://assets.trustwallet.com) może być używana do większości nowych zgłoszeń tokenów (wymagane jest konto GitHub).

### Szybki start

Szczegóły struktury repozytorium oraz wytyczne dotyczące wkładu są opisane na
[stronie dla deweloperów](https://developer.trustwallet.com/listing-new-assets/new-asset).
Poniżej znajduje się krótkie podsumowanie dla najczęstszego przypadku użycia.


## Dokumentacja

Szczegóły znajdziesz na [stronie dla deweloperów](https://developer.trustwallet.com):

- [Wytyczne dotyczące wkładu](https://developer.trustwallet.com/listing-new-assets/repository_details)

- [FAQ](https://developer.trustwallet.com/listing-new-assets/faq)

## Skrypty

Dostępnych jest kilka skryptów dla opiekunów repozytorium:

- `make check` -- Wykonaj sprawdzenia walidacyjne; używane również w ciągłej integracji.
- `make fix` -- Wykonaj automatyczne poprawki tam, gdzie to możliwe.
- `make update-auto` -- Uruchom automatyczne aktualizacje z zewnętrznych źródeł, wykonywane regularnie (GitHub action).
- `make add-token asset_id=c60_t0x4Fabb145d64652a948d72533023f6E7A623C7C53` -- Utwórz plik `info.json` jako szablon zasobu.
- `make add-tokenlist asset_id=c60_t0x4Fabb145d64652a948d72533023f6E7A623C7C53` -- Dodaj token do tokenlist.json.
- `make add-tokenlist-extended asset_id=c60_t0x4Fabb145d64652a948d72533023f6E7A623C7C53` -- Dodaj token do tokenlist-extended.json.

## O sprawdzeniach

To repozytorium zawiera zestaw skryptów do weryfikacji wszystkich informacji. Zaimplementowane jako skrypty Golang, dostępne poprzez `make check`, wykonywane podczas budowania CI; sprawdzają całe repozytorium.
Podobna logika sprawdzania jest zaimplementowana:

- w aplikacji assets-management; do sprawdzania zmienionych plików tokenów w PR-ach lub podczas tworzenia PR. Sprawdza różnice, może być uruchamiana ze środowiska przeglądarki.
- w merge-fee-bot, który działa jako aplikacja GitHub i pokazuje wyniki w komentarzu do PR. Wykonywany w środowisku bez przeglądarki.

## Utrzymanie par handlowych

Informacje o obsługiwanych parach handlowych są przechowywane w plikach `tokenlist.json`.
Pary handlowe mogą być aktualizowane --
z Uniswap/Ethereum i PancakeSwap/Smartchain -- za pomocą skryptu aktualizacyjnego (i zatwierdzania zmian).
Minimalne wartości limitów dla włączenia par handlowych są ustawione w [pliku konfiguracyjnym](https://github.com/trustwallet/assets/blob/master/.github/assets.config.yaml).
Dostępne są również opcje wymuszenia włączenia i wykluczenia w konfiguracji.

## Zastrzeżenie

Zespół Trust Wallet zezwala każdemu na przesyłanie nowych zasobów do tego repozytorium. Nie oznacza to jednak, że jesteśmy w bezpośrednim partnerstwie ze wszystkimi projektami.

Zespół Trust Wallet odrzuci projekty uznane za oszustwa lub fraudy po dokładnym przeglądzie.
Zespół Trust Wallet zastrzega sobie prawo do zmiany warunków przesyłania zasobów w dowolnym momencie ze względu na zmieniające się warunki rynkowe, ryzyko oszustwa lub inne czynniki, które uznamy za stosowne.

Ponadto zachowanie przypominające spam, w tym między innymi masowa dystrybucja tokenów na losowe adresy, spowoduje oznaczenie zasobu jako spam i możliwe usunięcie z repozytorium.

## Licencja

Skrypty i dokumentacja w tym projekcie są udostępniane na licencji [MIT](LICENSE)
