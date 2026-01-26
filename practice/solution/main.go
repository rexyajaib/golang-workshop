package main

import (
	"fmt"
	"golang-workshop/practice/solution/helper"
)

type CoinInterface interface {
	GetName() string
	SetName(string)
}

type SpotCoin struct {
	Name string
}

func (s *SpotCoin) GetName() string {
	return s.Name
}

func (s *SpotCoin) SetName(name string) {
	s.Name = name
}

type FutureCoin struct {
	Name string
}

func (f *FutureCoin) GetName() string {
	return fmt.Sprintf("PERP-%s", f.Name)
}

func (f *FutureCoin) SetName(name string) {
	f.Name = name
}

type Coins struct {
	Spot   []SpotCoin
	Future []FutureCoin
}

func NewCoins() *Coins {
	return &Coins{
		Spot:   make([]SpotCoin, 0),   // Initialize slice
		Future: make([]FutureCoin, 0), // Initialize slice
	}
}

func (c *Coins) AddSpotCoin(coin SpotCoin) {
	c.Spot = helper.AddToSliceIfNotExists(c.Spot, coin)
}

func (c *Coins) AddFutureCoin(coin FutureCoin) {
	c.Future = helper.AddToSliceIfNotExists(c.Future, coin)
}

func (c *Coins) GetSpotCoin(index int) SpotCoin {
	return helper.GetFromSlice(c.Spot, index)
}

func (c *Coins) GetFutureCoin(index int) FutureCoin {
	return helper.GetFromSlice(c.Future, index)
}

func (c *Coins) ReplaceSpotCoin(index int, coin SpotCoin) {
	c.Spot[index] = coin
}

func (c *Coins) ReplaceFutureCoin(index int, coin FutureCoin) {
	c.Future[index] = coin
}

func (c *Coins) RemoveSpotCoin(index int) {
	c.Spot = helper.RemoveFromSlice(c.Spot, index)
}

func (c *Coins) RemoveFutureCoin(index int) {
	c.Future = helper.RemoveFromSlice(c.Future, index)
}

func (c *Coins) LogCoins() {
	fmt.Println("Spot Coins:")
	for _, coin := range c.Spot {
		fmt.Printf("%s ", coin.GetName())
	}
	fmt.Println()
	fmt.Println("Future Coins:")
	for _, coin := range c.Future {
		fmt.Printf("%s ", coin.GetName())
	}
	fmt.Println()

}

func main() {
	coins := NewCoins()

	bitcoin := SpotCoin{Name: "Bitcoin"}
	ethereum := FutureCoin{Name: "Ethereum"}

	coins.AddSpotCoin(bitcoin)
	coins.AddFutureCoin(ethereum)

	retrievedBitcoin := coins.GetSpotCoin(0)
	retrievedEthereum := coins.GetFutureCoin(0)

	retrievedBitcoin.SetName("BTC")
	retrievedEthereum.SetName("ETH")

	coins.ReplaceSpotCoin(0, retrievedBitcoin)
	coins.ReplaceFutureCoin(0, retrievedEthereum)

	solana := SpotCoin{Name: "SOL"}
	coins.AddSpotCoin(solana)

	// Expect coins.Spot[0].Name to be "BTC" and coins.Future[0].Name to be "PERP-ETH"
	coins.LogCoins()

	fmt.Println("After removing the first spot coin:")
	coins.RemoveSpotCoin(0)
	coins.LogCoins()
}
