package concurrency

type VerificadorWebsite func(string) bool

type resultado struct {
	string
	bool
}

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(targetURL string) {
			canalResultado <- resultado{targetURL, vw(targetURL)}
			//resultados[targetURL] = vw(targetURL)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}

	return resultados
}
