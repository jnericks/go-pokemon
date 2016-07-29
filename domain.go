package gopokemon

// Inventory ...
type Inventory struct {
	Success        bool           `json:"success"`
	InventoryDelta InventoryDelta `json:"inventory_delta"`
}

// InventoryDelta ...
type InventoryDelta struct {
	NewTimestampMS uint64          `json:"new_timestamp_ms"`
	InventoryItems []InventoryItem `json:"inventory_items"`
}

// InventoryItem ...
type InventoryItem struct {
	ModifiedTimestampMS uint64            `json:"modified_timestamp_ms"`
	InventoryItemData   InventoryItemData `json:"inventory_item_data"`
}

// InventoryItemData ...
type InventoryItemData struct {
	PlayerStats  PlayerStats  `json:"player_stats"`
	PokemonData  PokemonData  `json:"pokemon_data"`
	Candy        Candy        `json:"candy"`
	PokedexEntry PokedexEntry `json:"pokedex_entry"`
	Item         Item         `json:"item"`
}

// PlayerStats ...
type PlayerStats struct {
	Level                uint    `json:"level"`
	Experience           uint    `json:"experience"`
	PrevLevelXP          uint    `json:"prev_level_xp"`
	NextLevelXP          uint    `json:"next_level_xp"`
	KMWalked             float64 `json:"km_walked"`
	PokemonsEncountered  uint    `json:"pokemons_encountered"`
	UniquePokedexEntries uint    `json:"unique_pokedex_entries"`
	PokemonsCaptured     uint    `json:"pokemons_captured"`
	Evolutions           uint    `json:"evolutions"`
	PokeStopVisits       uint    `json:"poke_stop_visits"`
	PokeballsThrown      uint    `json:"pokeballs_thrown"`
	EggsHatched          uint    `json:"eggs_hatched"`
	BattleAttackWon      uint    `json:"battle_attack_won"`
	BattleAttackTotal    uint    `json:"battle_attack_total"`
	BattleTrainingWon    uint    `json:"battle_training_won"`
	BattleTrainingTotal  uint    `json:"battle_training_total"`
	PrestigeRaisedTotal  uint    `json:"prestige_raised_total"`
	PrestigeDroppedTotal uint    `json:"prestige_dropped_total"`
	PokemonDeployed      uint    `json:"pokemon_deployed"`
	PokemonCaughtByType  string  `json:"pokemon_caught_by_type"`
	SmallRattataCaught   uint    `json:"small_rattata_caught"`
}

// PokemonData ...
type PokemonData struct {
	ID                uint64  `json:"id"`
	PokemonID         uint64  `json:"pokemon_id"`
	CP                uint    `json:"cp"`
	Stamina           uint    `json:"stamina"`
	StaminaMax        uint    `json:"stamina_max"`
	Move1             uint    `json:"move_1"`
	Move2             uint    `json:"move_2"`
	HeightM           float64 `json:"height_m"`
	WeightKG          float64 `json:"weight_kg"`
	IndividualAttack  uint    `json:"individual_attack"`
	IndividualDefense uint    `json:"individual_defense"`
	IndividualStamina uint    `json:"individual_stamina"`
	CPMultiplier      float64 `json:"cp_multiplier"`
	Pokeball          uint    `json:"pokeball"`
	CapturedCellID    uint64  `json:"captured_cell_id"`
	CreationTimeMS    uint64  `json:"creation_time_ms"`
	Nickname          string  `json:"nickname"`
	FromFort          uint    `json:"from_fort"`
	IsEgg             bool    `json:"is_egg"`
	EggKMWalkedTarget uint    `json:"egg_km_walked_target"`
	EggIncubatorID    string  `json:"egg_incubator_id"`
}

// Candy ...
type Candy struct {
	FamilyID uint64 `json:"family_id"`
	Candy    uint   `json:"candy"`
}

// PokedexEntry ...
type PokedexEntry struct {
	PokemonID        uint64 `json:"pokemon_id"`
	TimesEncountered uint   `json:"times_encountered"`
	TimesCaptured    uint   `json:"times_captured"`
}

// Item ...
type Item struct {
	ItemID uint64 `json:"item_id"`
	Count  uint   `json:"count"`
}
