package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	if len(caminhos) == 0 {
		return "", errors.New("Não foi possível determinar a cidade de destino")
	}

	destinos := make(map[string]int)

	for _, caminho := range caminhos {
		destinos[caminho[1]]++
		destinos[caminho[0]]--
	}

	for cidade, ocorrencias := range destinos {
		if ocorrencias > 0 {
			return cidade, nil
		}
	}

	return "", errors.New("Não foi possível determinar a cidade de destino")
}
