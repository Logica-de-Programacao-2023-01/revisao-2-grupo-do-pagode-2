package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]`
//significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	origem := make(map[string]bool)

	// Percorre a lista de caminhos e adiciona as cidades de origem ao mapa
	for _, caminho := range caminhos {
		origem[caminho[0]] = true
	}

	// Percorre a lista de caminhos novamente e verifica a cidade de destino
	for _, caminho := range caminhos {
		// Se a cidade de destino não está presente no mapa de origem, é a cidade de destino final
		if !origem[caminho[1]] {
			return caminho[1], nil
		}
	}

	// Se não encontrou cidade de destino, retorna um erro
	return "", errors.New("Não existe cidade de destino")
}
