package main

// Inventory ...
type inventory struct {
	Success        bool           `json:"success"`
	InventoryDelta inventoryDelta `json:"inventory_delta"`
}

type inventoryDelta struct {
	NewTimestampMS int64           `json:"new_timestamp_ms"`
	InventoryItems []inventoryItem `json:"inventory_items"`
}

type inventoryItem struct {
	ModifiedTimestampMS int64 `json:"modified_timestamp_ms"`
	InventoryItemData   interface{}
}

type inventoryItemData struct {
	PlayerStats  playerStats  `json:"player_stats"`
	PokemonData  pokemonData  `json:"pokemon_data"`
	Candy        candy        `json:"candy"`
	PokedexEntry pokedexEntry `json:"pokedex_entry"`
	Item         item         `json:"item"`
}

type playerStats struct {
	Level                int     `json:"level"`
	Experience           int     `json:"experience"`
	PrevLevelXP          int     `json:"prev_level_xp"`
	NextLevelXP          int     `json:"next_level_xp"`
	KMWalked             float32 `json:"km_walked"`
	PokemonsEncountered  int     `json:"pokemons_encountered"`
	UniquePokedexEntries int     `json:"unique_pokedex_entries"`
	PokemonsCaptured     int     `json:"pokemons_captured"`
	Evolutions           int     `json:"evolutions"`
	PokeStopVisits       int     `json:"poke_stop_visits"`
	PokeballsThrown      int     `json:"pokeballs_thrown"`
	EggsHatched          int     `json:"eggs_hatched"`
	BattleAttackWon      int     `json:"battle_attack_won"`
	BattleAttackTotal    int     `json:"battle_attack_total"`
	BattleTrainingWon    int     `json:"battle_training_won"`
	BattleTrainingTotal  int     `json:"battle_training_total"`
	PrestigeRaisedTotal  int     `json:"prestige_raised_total"`
	PrestigeDroppedTotal int     `json:"prestige_dropped_total"`
	PokemonDeployed      int     `json:"pokemon_deployed"`
	PokemonCaughtByType  string  `json:"pokemon_caught_by_type"`
	SmallRattataCaught   int     `json:"small_rattata_caught"`
}

type pokemonData struct {
	ID                int     `json:"id"`
	PokemonID         int     `json:"pokemon_id"`
	CP                int     `json:"cp"`
	Stamina           int     `json:"stamina"`
	StaminaMax        int     `json:"stamina_max"`
	Move1             int     `json:"move_1"`
	Move2             int     `json:"move_2"`
	HeightM           float64 `json:"height_m"`
	WeightKG          float64 `json:"weight_kg"`
	IndividualAttack  int     `json:"individual_attack"`
	IndividualDefense int     `json:"individual_defense"`
	IndividualStamina int     `json:"individual_stamina"`
	CPMultiplier      float64 `json:"cp_multiplier"`
	Pokeball          int     `json:"pokeball"`
	CapturedCellID    int64   `json:"captured_cell_id"`
	CreationTimeMS    int64   `json:"creation_time_ms"`
	Nickname          string  `json:"nickname"`
	FromFort          int     `json:"from_fort"`
	IsEgg             bool    `json:"is_egg"`
	EggKMWalkedTarget int     `json:"egg_km_walked_target"`
	EggIncubatorID    string  `json:"egg_incubator_id"`
}

type candy struct {
	FamilyID int `json:"family_id"`
	Candy    int `json:"candy"`
}

type pokedexEntry struct {
	PokemonID        int `json:"pokemon_id"`
	TimesEncountered int `json:"times_encountered"`
	TimesCaptured    int `json:"times_captured"`
}

type item struct {
	ItemID int `json:"item_id"`
	Count  int `json:"count"`
}
