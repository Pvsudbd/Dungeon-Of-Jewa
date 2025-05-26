package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type player struct {
	NamaP     string
	HP        int
	Mana      int
	PAttack   int
	Armor     int
	Weapon    equipment
	ranged    equipment
	Plate     equipment
	Potion    []Consumamble
	Skill     skill
	inventory []equipment
	class     string
	Gold      int
}

type equipment struct {
	NamaE                          string
	TypeE                          string
	hp, armour, damage, Cost, mana int
}

type skill struct {
	NamaS     string
	damage    int
	ManaCost  int
	SkillType string
}

type class struct {
	NamaC string
	Power int
	Skill skill
}

type Consumamble struct {
	nama                           string
	hp, mana, DamageC, armor, cost int
}

type monster struct {
	nama                      string
	hp, mana, attackM, armour int
	skill                     skill
}

type ShopAjiAW struct {
	NamaB  string
	Type   string
	damage int
	def    int
	cost   int
}

type NPCTOKO struct {
	NamaNPC string
}

var npc = []NPCTOKO{
	{NamaNPC: "Bro Agus"},
	{NamaNPC: "Cintya"},
	{NamaNPC: "Hannah Bule"},
	{NamaNPC: "Cak Ji"},
	{NamaNPC: "Sugeng Tirta"},
	{NamaNPC: "Roro Karomah"},
	{NamaNPC: "Nyonya Patricia"},
	{NamaNPC: "Tirta Adisatya"},
	{NamaNPC: "Hasan Malioboro"},
}

func bubbleShopA(arr []ShopAjiAW) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		tukar := false
		for j := 0; j < n-i-1; j++ {
			if arr[j].cost > arr[j+1].cost {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
}

func bubbleShopC(arr []Consumamble) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		tukar := false
		for j := 0; j < n-i-1; j++ {
			if arr[j].cost > arr[j+1].cost {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
}

func SequentialAW(cihuy []ShopAjiAW, name string) *ShopAjiAW {
	for i := 0; i < len(cihuy); i++ {
		if cihuy[i].NamaB == name {
			return &cihuy[i]
		}
	}
	return nil
}

func SequentialC(cihuy2 []Consumamble, name string) *Consumamble {
	for i := 0; i < len(cihuy2); i++ {
		if cihuy2[i].nama == name {
			return &cihuy2[i]
		}
	}
	return nil
}

func beliArmor(p *player) {
	var Armour = []ShopAjiAW{
		{NamaB: "Majapahit Zirah", Type: "Armor", def: 12, cost: 15},
		{NamaB: "Iron Chestplate", Type: "Armor", def: 8, cost: 10},
		{NamaB: "Good Iron Chestplate", Type: "Armor", def: 10, cost: 12},
		{NamaB: "Totem Of Saranjana", Type: "Armor", def: 60, cost: 82},
		{NamaB: "Mongolian Laminar Armor", Type: "Armor", def: 16, cost: 20},
		{NamaB: "Kriptonian Steel Armour", Type: "Armor", def: 35, cost: 29},
		{NamaB: "Kostum pilo", Type: "Armor", def: 34, cost: 25},
		{NamaB: "Agartha Full Composite", Type: "Armor", def: 45, cost: 40},
		{NamaB: "Nyi Sandi Chestplate", Type: "Armor", def: 22, cost: 27},
		{NamaB: "Northen Armor", Type: "Armor", def: 42, cost: 62},
		{NamaB: "Steel Plate with Jimat", Type: "Armor", def: 100, cost: 162},
		{NamaB: "Low Quality SP w Jimat", Type: "Armor", def: 62, cost: 72},
		{NamaB: "Agartha Half Composite", Type: "Armor", def: 38, cost: 31},
		{NamaB: "Dog Lether Shirt", Type: "Armor", def: 3, cost: 2},
		{NamaB: "Cow Leather Shirt", Type: "Armor", def: 4, cost: 4},
		{NamaB: "White People Robe", Type: "Armor", def: 7, cost: 10},
		{NamaB: "Cursed Saranjana Armor", Type: "Armor", def: 150, cost: 210},
	}
	var Armor []ShopAjiAW

	for i := 0; i < 5; i++ {
		randomIndexE := rand.Intn(len(Armour))
		Armor = append(Armor, Armour[randomIndexE])

	}

	bubbleShopA(Armor)

	for i, s := range Armor {
		fmt.Printf("%d. %s (Armor: %d) - %d gold\n", i+1, s.NamaB, s.def, s.cost)
	}
	fmt.Printf("Duit:%d\n", p.Gold)
	var pilih string
	fmt.Println("=========Your Choice=========")
	fmt.Print("Masukkan nomor atau nama armor yang ingin dibeli: ")
	fmt.Scanln(&pilih)
	fmt.Printf("=============================\n")

	var itemA *ShopAjiAW
	if n, err := strconv.Atoi(pilih); err == nil {
		if n >= 1 && n <= len(Armor) {
			itemA = &Armor[n-1]
		}
	} else {
		itemA = SequentialAW(Armor, pilih)
	}

	if itemA == nil {
		fmt.Println("Cak Ji : Mana ada")
		return
	}
	if p.Gold < itemA.cost {
		fmt.Println("Cak Ji : Duitmu kurang")
		return
	}

	p.Armor -= p.Plate.armour
	p.Plate = equipment{
		NamaE:  itemA.NamaB,
		TypeE:  itemA.Type,
		armour: itemA.def,
		Cost:   itemA.cost,
	}
	p.Armor += itemA.def
	p.Gold -= itemA.cost
	fmt.Printf("Cak Ji : Kamu membeli %s seharga %d g-coin\n", itemA.NamaB, itemA.cost)
}

func beliwepong(p *player) {
	var bedil = []ShopAjiAW{
		{NamaB: "Keris Ambasung", Type: "weapon", damage: 15, cost: 6},
		{NamaB: "Golok", Type: "weapon", damage: 20, cost: 4},
		{NamaB: "Arit Mistis", Type: "weapon", damage: 25, cost: 8},
		{NamaB: "Jawirian Hell Sword", Type: "weapon", damage: 30, cost: 17},
		{NamaB: "Mongolian Dagger", Type: "weapon", damage: 40, cost: 25},
		{NamaB: "Persian Sword", Type: "weapon", damage: 50, cost: 40},
		{NamaB: "Jati Bow", Type: "ranged", damage: 20, cost: 6},
		{NamaB: "Guard Bow", Type: "ranged", damage: 30, cost: 15},
		{NamaB: "Bow With Jimat I", Type: "ranged", damage: 40, cost: 30},
		{NamaB: "Mongolian Sacred Bow", Type: "ranged", damage: 46, cost: 35},
		{NamaB: "Sultanate Of Aceh Bow", Type: "ranged", damage: 56, cost: 52},
		{NamaB: "Bow With Jimat II", Type: "ranged", damage: 70, cost: 82},
		{NamaB: "Pedang Raja Ngawi", Type: "weapon", damage: 40, cost: 42},
		{NamaB: "Busur Derajat Raja Rusdi", Type: "ranged", damage: 60, cost: 66},
		{NamaB: "Gada Gajah Mada", Type: "weapon", damage: 85, cost: 100},
		{NamaB: "Tombak Pecel", Type: "weapon", damage: 20, cost: 32},
		{NamaB: "Enchanted Manuk", Type: "ranged", damage: 90, cost: 90},
		{NamaB: "Pedang Ireng Lagendaris", Type: "weapon", damage: 60, cost: 77},
		{NamaB: "Busur Bujur", Type: "ranged", damage: 30, cost: 42},
		{NamaB: "Sarung Tangan Jawirian", Type: "weapon", damage: 25, cost: 33},
		{NamaB: "Rudal Balistik", Type: "ranged", damage: 100, cost: 92},
	}

	var equipments []ShopAjiAW
	for i := 0; i < 5; i++ {
		randomIndexE := rand.Intn(len(bedil))
		equipments = append(equipments, bedil[randomIndexE])

	}

	bubbleShopA(equipments)
	for i, s := range equipments {
		fmt.Printf("%d. %s (Damage: %d) - %d gold\n", i+1, s.NamaB, s.damage, s.cost)
	}
	fmt.Printf("Duit: %d\n", p.Gold)
	var pilih string
	fmt.Print("Masukkan nomor atau nama senjata yang ingin dibeli: ")
	fmt.Scanln(&pilih)
	fmt.Printf("=============================\n")

	var item *ShopAjiAW
	if n, err := strconv.Atoi(pilih); err == nil {
		if n >= 1 && n <= len(equipments) {
			item = &equipments[n-1]
		}
	} else {
		item = SequentialAW(equipments, pilih)
	}

	if item == nil {
		fmt.Println("Cak Ji : Senjata tidak ditemukan")
		return
	}
	if p.Gold < item.cost {
		fmt.Println("Cak Ji : Duitmu kurang")
		return
	}

	p.Weapon = equipment{
		NamaE:  item.NamaB,
		TypeE:  item.Type,
		damage: item.damage,
		Cost:   item.cost,
	}
	p.inventory = append(p.inventory, p.Weapon)
	p.Gold -= item.cost

	fmt.Printf("Cak Ji : Kamu membeli %s seharga %d g-coin\n", item.NamaB, item.cost)
}

func beliConsumable(p *player) {
	var Consums = []Consumamble{
		{nama: "Healer I", hp: 20, mana: 0, DamageC: 0, armor: 0, cost: 5},
		{nama: "Healer II", hp: 30, mana: 0, DamageC: 0, armor: 0, cost: 10},
		{nama: "Jeruk Bali", hp: 15, mana: 10, DamageC: 0, armor: 0, cost: 15},
		{nama: "Pisang Keju Majapahit", hp: 30, mana: 50, DamageC: 0, armor: 0, cost: 45},
		{nama: "Salmon Melayu", hp: 15, mana: 50, DamageC: 4, armor: 0, cost: 50},
		{nama: "Bika Batak", hp: 20, mana: 10, DamageC: 30, armor: 0, cost: 60},
		{nama: "Ketupat Mataram", hp: 5, mana: 5, DamageC: 23, armor: 0, cost: 27},
		{nama: "Teh Malang", hp: 10, mana: 10, DamageC: 10, armor: 10, cost: 20},
		{nama: "Aceh Cigar II", hp: 10, mana: 5, DamageC: 60, armor: 0, cost: 120},
		{nama: "Aceh Gayo Coffee", hp: 60, mana: 0, DamageC: 0, armor: 0, cost: 62},
		{nama: "Pecel Madiun", hp: 35, mana: 0, DamageC: 0, armor: 0, cost: 23},
		{nama: "Lele Bakar", hp: 5, mana: 0, DamageC: 0, armor: 0, cost: 2},
		{nama: "Kaki Kalkun", hp: 15, mana: 5, DamageC: 0, armor: 0, cost: 5},
		{nama: "Penyet Trex", hp: 50, mana: 0, DamageC: 20, armor: 0, cost: 30},
		{nama: "Batak Soup", hp: 10, mana: 10, DamageC: 10, armor: 10, cost: 10},
		{nama: "Kue lapis Argo Ngawi", hp: 100, mana: 50, DamageC: 10, armor: 20, cost: 51},
		{nama: "Jagung Ambasung", hp: 20, mana: 20, DamageC: 10, armor: 10, cost: 26},
		{nama: "Minyak Lintah Wonogiri", hp: 0, mana: 50, DamageC: 10, armor: 10, cost: 24},
		{nama: "Terang Bulan Megalodon", hp: 70, mana: 10, DamageC: 0, armor: 10, cost: 35},
		{nama: "Kunyit Papua", hp: 100, mana: 0, DamageC: 0, armor: 0, cost: 45},
		{nama: "Sumatra Honey", hp: 0, mana: 20, DamageC: 0, armor: 0, cost: 7},
		{nama: "Javanese Delight", hp: 0, mana: 35, DamageC: 0, armor: 0, cost: 12},
		{nama: "Mongolian Phantom", hp: 4, mana: 55, DamageC: 0, armor: 0, cost: 24},
		{nama: "Aceh Cigar", hp: 10, mana: 100, DamageC: 0, armor: 0, cost: 47},
	}
	var Consum []Consumamble
	for i := 0; i < 5; i++ {
		randomIndexE := rand.Intn(len(Consums))
		Consum = append(Consum, Consums[randomIndexE])

	}

	bubbleShopC(Consum)
	for i, s := range Consum {
		fmt.Printf("%d. %s, Heal: %d, Mana: %d, Cost: %d\n", i+1, s.nama, s.hp, s.mana, s.cost)
	}
	fmt.Printf("Duit:%d\n", p.Gold)
	var pilih string
	fmt.Print("Masukkan nomor atau nama potion yang ingin dibeli: ")
	fmt.Scanln(&pilih)
	fmt.Printf("=============================\n")

	var itemB *Consumamble
	if n, err := strconv.Atoi(pilih); err == nil {
		if n >= 1 && n <= len(Consum) {
			itemB = &Consum[n-1]
		}
	} else {
		itemB = SequentialC(Consum, pilih)
	}

	if itemB == nil {
		fmt.Println("Cak Ji : Mana ada")
		return
	}
	if p.Gold < itemB.cost {
		fmt.Println("Cak Ji : Duitmu kurang")
		return
	}
	p.Potion = append(p.Potion, *itemB)
	p.Gold -= itemB.cost
	fmt.Printf("Cak Ji : Kamu membeli %s seharga %d g-coin\n", itemB.nama, itemB.cost)
}

func shop(p *player) {
	rand.Seed(time.Now().UnixNano())
	RandomNpc := npc[rand.Intn(len(npc))]

	for {
		fmt.Printf("Selamat datang di Toko %s \nA. Armor\nB. Weapon\nC. Jamu\nE. Keluar\n", RandomNpc.NamaNPC)
		fmt.Printf("Duit:%d\n", p.Gold)
		var pilihan string
		fmt.Printf("Apa pilihanmu nak?: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		switch pilihan {
		case "A":
			beliArmor(p)
		case "B":
			beliwepong(p)
		case "C":
			beliConsumable(p)
		case "E":
			fmt.Printf("\n%s: Semoga kau selalu dalam lindungan tuhan!", RandomNpc.NamaNPC)
			return
		}
	}
}

func enemyAttack(e *monster, p *player) int {
	damage := e.attackM - p.Armor
	if e.attackM <= p.Armor {
		damage = 1
	}
	p.HP -= damage
	return damage
}

func enemySkill(e *monster, p *player) int {
	e.mana -= e.skill.ManaCost
	damage := e.skill.damage - p.Armor
	if e.skill.damage <= p.Armor {
		damage = 1
	}
	p.HP -= damage
	return damage
}

func enemyTurn(p *player, e *monster) {
	if e.hp <= 0 {
		return
	}
	randomIndex := rand.Intn(2)
	var enemydamage int
	if e.mana == 0 {
		randomIndex = 0
	}
	switch randomIndex {
	case 0:
		enemydamage = enemyAttack(e, p)
		fmt.Println("Musuh menggunakan Attack biasa dengan damage: ", enemydamage)

	case 1:
		enemydamage = enemySkill(e, p)
		fmt.Println("Musuh menggunakan SKill dengan damage: ", enemydamage)

	}

}

func attack(p *player, e *monster) {
	damage := p.PAttack - e.armour
	e.hp -= damage
	fmt.Println("Kamu menyerang dengan damage ", damage)
	enemyTurn(p, e)

}

func Skill(p *player, e *monster) {
	if p.Mana < p.Skill.ManaCost {
		return
	}
	p.Mana -= p.Skill.ManaCost
	damage := p.Skill.damage - e.armour
	e.hp -= damage
	fmt.Println("Kamu menyerang dengan damage ", damage)
	enemyTurn(p, e)
}

func enemyDrop(c []Consumamble, e []equipment, p *player) {
	var dropC []Consumamble
	var dropE []equipment
	min := 10
	max := 30
	rangeSize := max - min + 1

	randomGC := rand.Intn(rangeSize) + min
	p.Gold += randomGC

	for i := 0; i < 5; i++ {
		randomIndexE := rand.Intn(len(e))
		dropE = append(dropE, e[randomIndexE])
		randomIndexC := rand.Intn(len(c))
		dropC = append(dropC, c[randomIndexC])

	}

	fmt.Println("Kamu Berhasi Mengalahkan Musuh! dan mendapatkan ", randomGC, " G-Coin")

	for i := 0; i < 2; i++ {
		var cek int
		var tipe string
		var namaitem int
		fmt.Println("======== Potion: ========")
		for i, c := range dropC {
			fmt.Printf("%d %s Hp: %d damage: %d Mana: %d Armour: %d Harga %d\n", i+1,c.nama, c.hp, c.DamageC, c.mana, c.armor, c.cost)
		}
		fmt.Println("======== Equipment ========")
		for i, e := range dropE {
			fmt.Printf("%d %s Hp: %d damage: %d Mana: %d Armour: %d Harga %d\n", i+1, e.NamaE,e.hp,e.damage,e.mana,e.armour,e.Cost)
		}
		fmt.Print("Pilih Item yang ingin di ambil maks 2, equipment atau potion?(E equipment P potion): ")
		fmt.Scan(&tipe)

		switch tipe {
		case "P":
			fmt.Print("Masukan Nomor Item: ")
			fmt.Scan(&namaitem)
			for i := 0; i <= len(c); i++ {
				if i == namaitem-1 {
					p.Potion = append(p.Potion, dropC[i])
					fmt.Println("Kamu Mengambil item :", dropC[i].nama)
					dropC = append(dropC[:i], dropC[i+1:]...)
					fmt.Println("OK dapet")
					continue
				} else {
					cek++
				}
			}
		case "E":
			fmt.Print("Masukan Nomor Item: ")
			fmt.Scan(&namaitem)

			for i := 0; i <= len(e); i++ {
				if i == namaitem-1 {
					p.inventory = append(p.inventory, dropE[i])
					fmt.Println("Kamu Mengambil item :", dropE[i].NamaE)
					dropE = append(dropE[:i], dropE[i+1:]...)
					fmt.Println("OK dapet")
					continue
				} else {
					cek++
				}

			}
		default:
			fmt.Println("Ketik ulang kamu salah ketik")
			i--
		}

		if cek == len(dropC) && len(dropE) == cek {
			i--
			fmt.Println("Barang yang ingin anda ambil tidak ada, coba lagi!")
		}

	}

}


func battle(p *player, M []monster, level int, consum []Consumamble, equip []equipment) int {
	var aksi int
	randomIndex := rand.Intn(len(M))
	e := M[randomIndex]
	e.hp += level
	e.attackM += level
	e.armour += level

	fmt.Printf("kamu bertemu dengan monster: %s\n", e.nama)

	

	for p.HP > 0 && e.hp > 0 {
		fmt.Printf("%s\n", e.nama)
		fmt.Printf("HP\t\t: %d\n", e.hp)
		fmt.Printf("Mana\t\t: %d\n", e.mana)
		fmt.Printf("Attack\t\t: %d\n", e.attackM)
		fmt.Printf("Armour\t\t: %d\n", e.armour)
		fmt.Printf("Nama Skill\t: %s\n", e.skill.NamaS)
		fmt.Printf("Damage Skill\t: %d\n", e.skill.damage)
		fmt.Print("\n\n=================================\n\n")
		fmt.Printf("%s (Kamu)\n", p.NamaP)
		fmt.Printf("HP\t\t: %d\n", p.HP)
		fmt.Printf("Mana\t\t: %d\n", p.Mana)
		fmt.Printf("Attack\t\t: %d\n", p.PAttack)
		fmt.Printf("Armour\t\t: %d\n", p.Armor)
		fmt.Println("======== Menu ========")
		fmt.Println("1. Attack")
		fmt.Println("2. Skill")
		fmt.Println("3. Inventory")
		fmt.Scan(&aksi)

		switch aksi {
		case 1:
			attack(p, &e)
		case 2:
			if p.Mana == 0 {
				fmt.Println("Mana tidak cukup")
			} else {
				Skill(p, &e)
			}
		case 3:
			fmt.Scanln()
			invetokry(p)
		}

		if p.HP <= 0 {
			fmt.Println("Kamu Kalah")
			return -1
		}

	}

	enemyDrop(consum, equip, p)
	return 1

}

func (p *player) menu() {
	for {
		fmt.Println("\n==== PLAYER INFO ====")
		fmt.Println("Name\t\t:", p.NamaP)
		fmt.Println("Attack\t\t:", p.PAttack)
		fmt.Println("Health\t\t:", p.HP)
		fmt.Println("Mana\t\t:", p.Mana)
		fmt.Println("Senjata\t\t:", p.Weapon.NamaE)
		fmt.Println("Armor\t\t:", p.Plate.NamaE)
		fmt.Println("Duit\t\t:", p.Gold)
		fmt.Println("====== Menu ======")
		fmt.Println("[1] Buka Inventory")
		fmt.Println("[2] Lanjutkan Petualangan")
		var pilihan int
		fmt.Println("===================")
		fmt.Printf("Masukan Input: ")
		fmt.Scan(&pilihan)
		fmt.Println("===================")

		switch pilihan {
		case 1:
			fmt.Scanln()
			invetokry(p)
		case 2:
			return

		}

	}
}
func SequentialInventory(inv []equipment, name string) *equipment {
	for i := 0; i < len(inv); i++ {
		if strings.EqualFold(inv[i].NamaE, name) {
			return &inv[i]
		}
	}
	return nil
}

func binaryPotion(potions []Consumamble, name string) *Consumamble {
    low, high := 0, len(potions)-1
    name = strings.ToLower(name)
    for low <= high {
        mid := (low + high) / 2
        midName := strings.ToLower(potions[mid].nama)
        if midName == name {
            return &potions[mid]
        } else if midName < name {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return nil
}

func insertionequipment(inv []equipment) {
	for i := 1; i < len(inv); i++ {
		key := inv[i]
		j := i - 1
		for j >= 0 && inv[j].Cost > key.Cost {
			inv[j+1] = inv[j]
			j--
		}
		inv[j+1] = key
	}
}

func insertionSortPotion(potions []Consumamble) {
	for i := 1; i < len(potions); i++ {
		key := potions[i]
		j := i - 1
		for j >= 0 && potions[j].cost > key.cost {
			potions[j+1] = potions[j]
			j--
		}
		potions[j+1] = key
	}
}


func (p *player) Equip(index int) {
    if index < 0 || index >= len(p.inventory) {
        fmt.Println("Index item tidak valid.")
        return
    }
    item := p.inventory[index]

    if p.Weapon.NamaE == item.NamaE  {
        fmt.Println("Kamu sudah pakai item yang sama")
        return
    }
    if p.Plate.NamaE == item.NamaE {
        fmt.Println("Kamu sudah pakai item yang sama")
        return
    }

    if item.TypeE == "ranged" && (p.class == "Fighter" || p.class == "Tanker" || p.class == "Assassino") {
        fmt.Println("Class ini tidak boleh pakai senjata jarak jauh!")
        return
    }

    fmt.Println("Kamu memakai:", item.NamaE)

    switch strings.ToLower(item.TypeE) {
    case "weapon", "ranged":
        p.PAttack -= p.Weapon.damage
        p.Weapon = item
        p.PAttack += p.Weapon.damage
    case "armor":
        p.Armor -= p.Plate.armour
        p.HP -= p.Plate.hp
        p.Plate = item
        p.Armor += p.Plate.armour
        p.HP += item.hp
    default:
        fmt.Println("Tipe item tidak dikenali")
    }
}


func (p *player) usepotion(index int) {
	if index < 0 || index >= len(p.Potion) {
		fmt.Println("Index potion tidak valid.")
		return
	}
	item := p.Potion[index]
	p.HP += item.hp
	p.Mana += item.mana
	fmt.Printf("Kamu memakai %s! HP +%d, Mana +%d\n", item.nama, item.hp, item.mana)
	p.Potion = append(p.Potion[:index], p.Potion[index+1:]...)
}

func invetokry(p *player) {
	insertionequipment(p.inventory)
	insertionSortPotion(p.Potion)
	fmt.Println("\n ==== Inventory ====")
	fmt.Println("-Equipment:")
	for i, eq := range p.inventory {
		fmt.Printf("%d. %s (Tipe: %s, Damage: %d, Armor: %d, Cost: %d)\n", i+1, eq.NamaE, eq.TypeE, eq.damage, eq.armour, eq.Cost)
	}

	fmt.Println("-Potion:")
	for i, pot := range p.Potion {
		fmt.Printf("%d. %s (Heal: %d, Mana: %d, Cost: %d)\n", len(p.inventory)+i+1, pot.nama, pot.hp, pot.mana, pot.cost)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Cari equipment berdasarkan nama (atau tekan enter untuk skip): \n")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		item := SequentialInventory(p.inventory, input)
		if item != nil {
			fmt.Println("Item ditemukan:", item.NamaE, "(Tipe:", item.TypeE, ", Damage:", item.damage, ", Armor:", item.armour, "Cost", item.Cost, "). Pakai?")
		} else {
			fmt.Println("Item tidak ditemukan.")
		}
	}

	fmt.Print("Ketik nomer item untuk equip/pakai, atau nama equipment (atau '-1' untuk keluar): ")
	pilihanRaw, _ := reader.ReadString('\n')
	pilihan := strings.TrimSpace(pilihanRaw)

	if pilihan == "-1" {
		return
	}

	if idx, err := strconv.Atoi(pilihan); err == nil {
    if idx >= 1 && idx <= len(p.inventory) {
        before := p.Weapon.NamaE
        p.Equip(idx - 1)
        if before != p.Weapon.NamaE {
            fmt.Println("======== Equipment berhasil dipakai! ========")
        }
    } else if idx > len(p.inventory) && idx <= len(p.inventory)+len(p.Potion) {
        p.usepotion(idx - len(p.inventory) - 1)
        fmt.Println("======== Potion berhasil dikonsumsi! ========")
    } else {
        fmt.Println("Item tidak ditemukan.")
    }
	} else {
		item := SequentialInventory(p.inventory, pilihan)
		if item != nil {
			for i, eq := range p.inventory {
				if strings.Contains(strings.ToLower(eq.NamaE), strings.ToLower(pilihan)) {
					p.Equip(i)
					fmt.Println("======== Equipment berhasil dipakai! ========")
					return
				}
			}
		} else {
			pot := binaryPotion(p.Potion, pilihan)
			if pot != nil {
				for i, ptn := range p.Potion {
					if strings.Contains(strings.ToLower(ptn.nama), strings.ToLower(pilihan)) {
						p.usepotion(i)
						fmt.Println("======== Potion berhasil dikonsumsi! ========")
						return
					}
				}
			} else {
				fmt.Println("Item tidak ditemukan.")
			}
		}
	}
}

func main() {
	var next int
	var class int
	var player = player{
		NamaP:   "Hero",
		HP:      100,
		PAttack: 10,
		Gold:    5,
		Plate: equipment{
			NamaE:  "Kosong",
			hp:     0,
			damage: 0,
			armour: 0,
		},
		Potion: []Consumamble{},
		Weapon: equipment{
			NamaE:  "Kosong",
			hp:     0,
			damage: 0,
			armour: 0,
		},
	}
	var monsters = []monster{
		{nama: "Looters", hp: 20, mana: 40, attackM: 10,armour: 12, skill: skill{NamaS: "Looters Attack", ManaCost: 20, SkillType: "Physical", damage: 30}},
		{nama: "Trained Looters", hp: 35, mana: 40, attackM: 20,armour: 5, skill: skill{NamaS: "Taekwondo", ManaCost: 25, SkillType: "Physical", damage: 35}},
		{nama: "Dungeon Bandit", hp: 40, mana: 40, attackM: 25, skill: skill{NamaS: "Knive Charge", ManaCost: 20, SkillType: "Physical", damage: 45}},
		{nama: "Wong Mati", hp: 55, mana: 70, attackM: 32, skill: skill{NamaS: "Bite", ManaCost: 20, SkillType: "Physical", damage: 57}},
		{nama: "Wong Mati Armored", hp: 120, mana: 40, attackM: 32,armour: 20, skill: skill{NamaS: "Bite", ManaCost: 20, SkillType: "Physical", damage: 57}},
		{nama: "Orang Marah", hp: 120, mana: 40, attackM: 10, skill: skill{NamaS: "Drunk Punch", ManaCost: 20, SkillType: "Physical", damage: 50}},
		{nama: "Weak Jawirian", hp: 50, mana: 40, attackM: 50, skill: skill{NamaS: "Sword Dash", ManaCost: 20, SkillType: "Physical", damage: 42}},
		{nama: "Jawirian Knight", hp: 65, mana: 40, attackM: 60, skill: skill{NamaS: "Sword stab", ManaCost: 20, SkillType: "Physical", damage: 70}},
		{nama: "Deep Dungeon Bandit", hp: 80, mana: 40, attackM: 10,armour: 5, skill: skill{NamaS: "Ambush", ManaCost: 20, SkillType: "Physical", damage: 42}},
		{nama: "Madura Scavenggers", hp: 90, mana: 20, attackM: 10,armour: 5, skill: skill{NamaS: "Te sateeee!", ManaCost: 20, SkillType: "Physical", damage: 30}},
		{nama: "Madman", hp: 140, mana: 40, attackM: 20,armour: 10, skill: skill{NamaS: "Weak punch", ManaCost: 20, SkillType: "Physical", damage: 25}},
		{nama: "Bandit Medan", hp: 120, mana: 10, attackM: 40,armour: 5, skill: skill{NamaS: "BESI!!!!", ManaCost: 10, SkillType: "Physical", damage: 80}},
		{nama: "Antek Dewa Junggo", hp: 120, mana: 43, attackM: 10,armour: 10, skill: skill{NamaS: "Ku Buru Kau Komdis", ManaCost: 210, SkillType: "Physical", damage:30}},
    }
    
	var boss = []monster{
		{nama: "Dungeon Bandit Lord", hp: 150, mana: 40, attackM: 15,armour: 20, skill: skill{NamaS: "Rage of the poor guy", ManaCost: 20, SkillType: "Physical", damage: 25}},
		{nama: "Cursed King", hp: 200, mana: 20, attackM: 15,armour: 15, skill: skill{NamaS: "Jawa Rage", ManaCost: 10, SkillType: "Magic", damage: 20}},
		{nama: "Jawirian Lord", hp: 180, mana: 0, attackM: 30,armour: 21, skill: skill{NamaS: "ludah", ManaCost: 0, SkillType: "Physical", damage: 35}},
		{nama: "Jawirian Demon King", hp: 250, mana: 50, attackM: 45,armour: 20, skill: skill{NamaS: "Yapping", ManaCost: 10, SkillType: "Physical", damage: 60}},
		{nama: "Dungeon Lord", hp: 305, mana: 100, attackM: 35,armour: 30, skill: skill{NamaS: "Branch Of Luminate", ManaCost: 20, SkillType: "Physhical", damage: 50}},
	}

	var Consum = []Consumamble{
		{nama: "Healer I", hp: 20, mana: 0, cost: 5},
		{nama: "Healer II", hp: 30, mana: 0, cost: 15},
		{nama: "Jamu Blitar", hp: 50, mana: 0, cost: 25},
		{nama: "Kunyit", hp: 100, mana: 0, cost: 45},
		{nama: "Honey", hp: 0, mana: 20, cost: 7},
		{nama: "Delight", hp: 0, mana: 35, cost: 17},
		{nama: "Mongolian Phantom", hp: 4, mana: 55, cost: 34},
		{nama: "Aceh Cigar", hp: 10, mana: 100, cost: 50},
	}

	var equipments = []equipment{
		{NamaE: "Keris Ambasung", damage: 15, Cost: 6, TypeE: "weapon"},
		{NamaE: "Golok Bambu", damage: 10, Cost: 4, TypeE: "weapon"},
		{NamaE: "Kujang Mistis", damage: 20, Cost: 8, TypeE: "weapon"},
		{NamaE: "Tombak Jawara", damage: 25, Cost: 12, TypeE: "weapon"},
		{NamaE: "Wooden Bow", damage: 30, Cost: 20, TypeE: "ranged"},
		{NamaE: "Hunting Bow", damage: 42, Cost: 45, TypeE: "ranged"},
		{NamaE: "Bandit Bow", damage: 50, Cost: 50, TypeE: "ranged"},
		{NamaE: "Red Bow", damage: 60, Cost: 55, TypeE: "ranged"},
	}
    
	fmt.Print("========= DUNGEON OF JEWA =========\n\n")
	fmt.Println("Selamat datang di Dungeon Of Jewa, Ini adalah sebuah zaman dimana jawirian menguasai bumi.\nDungeon-Dungeon tercipta dan dipenuhi mosnter mengerikan\nDi saat itulah seorang pahlwan di butuhkan, dengan demikian kamu di tunjuk sebagai yang Agung terpilih oleh ilahi untuk menyelawatkan kita semua")
	fmt.Print("\nMasukan nama yang maha agung terpilih : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	player.NamaP = strings.TrimSpace(input)
	fmt.Println("Baiklah tuan: ", player.NamaP)

	fmt.Println("Class apa yanng anda ingin kan untuk membasmi jawirian?")

	for i := 0; i == 0; {
		i++
		fmt.Println("1.Archer\n2.Figter\n3.Tanker\n4.Assassino")
		fmt.Scan(&class)
		switch class {
		case 1:
			player.class = "Archer"
			player.HP = 150
			player.PAttack = 45
			player.Mana = 120
			player.Skill.NamaS = "Panah Pringgodani"
			player.Skill.ManaCost = 20
			player.Skill.damage = 70
			fmt.Println("Oke kamu sekarang adalah pemanah dan ini adalah stat mu")
			fmt.Print("\n")
			fmt.Println("Data Player:")
			fmt.Printf("Nama\t\t: %s\n", player.NamaP)
			fmt.Printf("HP\t\t: %d\n", player.HP)
			fmt.Printf("Armour\t\t: %d\n", player.Armor)
			fmt.Printf("Attack\t\t: %d\n", player.PAttack)
			fmt.Printf("Mana\t\t: %d\n", player.Mana)
			fmt.Printf("Nama Skill\t: %s\n", player.Skill.NamaS)
			fmt.Printf("Mana Cost\t: %d\n", player.Skill.ManaCost)
			fmt.Printf("Damage\t\t: %d\n", player.Skill.damage)

		case 2:
			player.class = "Fighter"
			player.HP = 275
			player.PAttack = 25
			player.Mana = 125
			player.Skill.NamaS = "Mandau Terbang"
			player.Skill.ManaCost = 25
			player.Skill.damage = 55
			fmt.Println("Oke kamu sekarang adalah Fighter dan ini adalah stat mu")
			fmt.Print("\n")
			fmt.Println("Data Player:")
			fmt.Printf("Nama\t\t: %s\n", player.NamaP)
			fmt.Printf("HP\t\t: %d\n", player.HP)
			fmt.Printf("Armour\t\t: %d\n", player.Armor)
			fmt.Printf("Attack\t\t: %d\n", player.PAttack)
			fmt.Printf("Mana\t\t: %d\n", player.Mana)
			fmt.Printf("Nama Skill\t: %s\n", player.Skill.NamaS)
			fmt.Printf("Mana Cost\t: %d\n", player.Skill.ManaCost)
			fmt.Printf("Damage\t\t: %d\n", player.Skill.damage)

		case 3:
			player.class = "Tanker"
			player.HP = 325
			player.PAttack = 30
			player.Mana = 100
			player.Skill.NamaS = "Javanese Punch"
			player.Skill.ManaCost = 20
			player.Skill.damage = 45
			fmt.Println("Oke kamu sekarang adalah Thanker dan ini adalah stat mu")
			fmt.Print("\n")
			fmt.Println("Data Player:")
			fmt.Printf("Nama\t\t: %s\n", player.NamaP)
			fmt.Printf("HP\t\t: %d\n", player.HP)
			fmt.Printf("Armour\t\t: %d\n", player.Armor)
			fmt.Printf("Attack\t\t: %d\n", player.PAttack)
			fmt.Printf("Mana\t\t: %d\n", player.Mana)
			fmt.Printf("Nama Skill\t: %s\n", player.Skill.NamaS)
			fmt.Printf("Mana Cost\t: %d\n", player.Skill.ManaCost)
			fmt.Printf("Damage\t\t: %d\n", player.Skill.damage)

		case 4:
			player.class = "Assassino"
			player.HP = 175
			player.PAttack = 40
			player.Mana = 150
			player.Skill.NamaS = "Stab of Injustice"
			player.Skill.ManaCost = 25
			player.Skill.damage = 60
			fmt.Println("Oke kamu sekarang adalah Assassino dan ini adalah stat mu")
			fmt.Print("\n")
			fmt.Println("Data Player:")
			fmt.Printf("Nama\t\t: %s\n", player.NamaP)
			fmt.Printf("HP\t\t: %d\n", player.HP)
			fmt.Printf("Armour\t\t: %d\n", player.Armor)
			fmt.Printf("Attack\t\t: %d\n", player.PAttack)
			fmt.Printf("Mana\t\t: %d\n", player.Mana)
			fmt.Printf("Nama Skill\t: %s\n", player.Skill.NamaS)
			fmt.Printf("Mana Cost\t: %d\n", player.Skill.ManaCost)
			fmt.Printf("Damage\t\t: %d\n", player.Skill.damage)

		default:
			fmt.Println("Class tidak ada, pilih kembali !")
			i = 0
		}
	}
	fmt.Print("\n\n")
	fmt.Println("Oke kamu akan langsung ke dungeon level satu")

	level := 0
	for {
		var f int
		fmt.Println("======== Apakah kamu siap? ========")
		fmt.Println("[1] Gas!")
		fmt.Println("[0] Keluar!")
		fmt.Println("===================")
		fmt.Printf("Pilih menu [1/2]: ")
		fmt.Scan(&f)
		fmt.Println("===================")
		if f == 0 {
			fmt.Println("Keluar dari dungeon...")
			break
		}
		if f == 1 {
            if level == 15 {
                fmt.Println("Selamat kamu berhasil memenenagkan game ini")
                return
            }
			fmt.Print("\n\n\n\n\n")
			fmt.Println("Kamu sedang sedang ada di level : ", level+1)
			level++

			if level%5 == 0 {
				fmt.Println("======== Kamu akan melawan Boss, Ini lantai Boss!! ========")
				next = battle(&player, boss, level, Consum, equipments)
			} else {
				next = battle(&player, monsters, level, Consum, equipments)
			}

			if next == -1 {
				fmt.Println("Game Over!")
				break
			}
			if level%3 == 0 {
				fmt.Println("======== Shop terbuka! ========")
				shop(&player)
			}

		}
		player.menu()

	}
}
