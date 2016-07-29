package main

type inventoryDelta struct {
	newTimestampMS int64
	inventoryItems []inventoryItem
}

type inventoryItem struct {
	modifiedTimestampMS int64
	inventoryIdemData   interface{}
}

type playerStats struct {
	level                int
	experience           int
	prevLevelXP          int
	nextLevelXP          int
	kmWalked             float32
	pokemonsEncountered  int
	uniquePokedexEntries int
	pokemonsCaptured     int
	evolutions           int
	pokeStopVisits       int
	pokeballsThrown      int
	eggsHatched          int
	battleAttackWon      int
	battleAttackTotal    int
	battleTrainingWon    int
	battleTrainingTotal  int
	prestigeRaisedTotal  int
	prestigeDroppedTotal int
	pokemonDeployed      int
	pokemonCaughtByType  string
	smallRattataCaught   int
}

type pokemonData struct {
	id                int
	pokemonID         int
	cp                int
	stamina           int
	staminaMax        int
	move1             int
	move2             int
	heightM           float64
	weightKG          float64
	individualAttack  int
	individualDefense int
	individualStamina int
	cpMultiplier      float64
	pokeball          int
	capturedCellID    int64
	creationTimeMS    int64
	nickname          string
	fromFort          int
}

type candy struct {
	familyID int
	candy    int
}

type pokedexEntry struct {
	pokemonID        int
	timesEncountered int
	timesCaptured    int
}

type item struct {
}
