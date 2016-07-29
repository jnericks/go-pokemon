package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	protos "github.com/pkmngo-odi/pogo-protos"
	"github.com/pkmngo-odi/pogo/api"
	"github.com/pkmngo-odi/pogo/auth"
)

var (
	p                 = flag.String("provider", "google", "auth provider ('google' or 'ptc')")
	username          = flag.String("u", "", "username")
	password          = flag.String("p", "", "password")
	refresh           = flag.Bool("r", false, "refresh pokemon inventory from server")
	inventoryDataFile = "inventory.json"
)

func main() {
	flag.Parse()

	if *refresh {
		if *username == "" || *password == "" {
			log.Fatalf("Need to supply username & password in order to refresh pokemon inventory.\n")
			return
		}

		refreshInventory(inventoryDataFile)
	}

	inventory := getInventoryFromFile(inventoryDataFile)
	var pokemonData []interface{}
	for _, i := range inventory.InventoryDelta.InventoryItems {
		data := i.InventoryItemData
		if &data.PokemonData != nil {
			pokemonData = append(pokemonData, data.PokemonData)
		}
	}

	StructsToCsv("pokemondata.csv", pokemonData...)
}

func getInventoryFromFile(file string) (inventory Inventory) {
	bytes, err := ReadFile(inventoryDataFile)
	checkError(err, true)
	json.Unmarshal(bytes, &inventory)
	return
}

func refreshInventory(file string) {
	// Initialize a new authentication provider to retrieve an access token
	provider, err := auth.NewProvider(*p, *username, *password)
	checkError(err, true)

	// Set the coordinates from where you're connecting
	location := &api.Location{
		Lon: 0.0,
		Lat: 0.0,
		Alt: 0.0,
	}

	// Start new session and connect
	debug := false
	session := api.NewSession(provider, location, debug)
	session.Init()

	// Start querying the API
	var inv []byte
	inv, err = getInventory(session)
	checkError(err, true)

	err = WriteFile(file, inv)
	checkError(err, true)
}

func checkError(err error, exit bool) {
	if err != nil {
		log.Panic(err)
		if exit {
			os.Exit(1)
		}
	}
}

func generateRequests() []*protos.Request {
	return make([]*protos.Request, 0)
}

func getPlayer(s *api.Session) ([]byte, error) {
	player, err1 := s.GetPlayer()
	if err1 != nil {
		log.Panic(err1)
		return nil, err1
	}

	json, err2 := json.MarshalIndent(player, "", "  ")
	if err2 != nil {
		log.Panic(err2)
		return nil, err2
	}

	return json, nil
}

func getInventory(s *api.Session) ([]byte, error) {
	req := generateRequests()
	req = append(req, &protos.Request{RequestType: protos.RequestType_GET_INVENTORY})

	res, err1 := s.Call(req)
	if err1 != nil {
		log.Panic(err1)
		return nil, err1
	}

	inv := &protos.GetInventoryResponse{}
	proto.Unmarshal(res.Returns[0], inv)

	json, err2 := json.MarshalIndent(inv, "", "  ")
	if err2 != nil {
		log.Panic(err2)
		return nil, err2
	}

	return json, nil
}
