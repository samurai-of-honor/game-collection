package games

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Catastrophe struct {
	Lore              string
	StayTime          string
	SurvivePopulation float32
	DestructionDegree int
}

type Info struct {
	Square          int
	FoodStocks      string
	MedicalOffice   string
	GreenHouses     string
	Armory          string
	AdditionalItem1 string
	AdditionalItem2 string
}

// TimeFormatter generate time and add an ending
func TimeFormatter() string {
	timeCategory := rand.Intn(3)
	timeStr := strconv.Itoa(Time[timeCategory][rand.Intn(len(Time[timeCategory]))])
	switch timeCategory {
	case 0:
		timeStr += " днів"
	case 1:
		timeStr += " місяць(-ів)"
	case 2:
		timeStr += " років"
	}
	return timeStr
}

func BunkerCatastrophe() string {
	c := new(Catastrophe)
	p := "%"

	c.Lore = Catastrophes[rand.Intn(len(Catastrophes))]
	c.StayTime = TimeFormatter()

	c.SurvivePopulation = SurvivePopulation[rand.Intn(len(SurvivePopulation))]
	c.DestructionDegree = DestructionDegree[rand.Intn(len(DestructionDegree))]

	ctrph := fmt.Sprintf("💥 <b>Катастрофа:</b> <i>%s</i>\n⏱ <b>Час в бункері:</b> <i>%s</i>\n", c.Lore, c.StayTime)
	ctrph += fmt.Sprintf("👨‍👩‍👧‍👦 <b>Кількість виживших:</b> <i>%.2f %s</i>\n\U0001F9F1 <b>Руйнування на поверхні:</b> <i>%d %s</i>\n",
		c.SurvivePopulation, p, c.DestructionDegree, p)
	return ctrph
}

func BunkerInfo() string {
	i := new(Info)
	rand.Seed(time.Now().Unix()) // initialize pseudo random generator

	i.Square = Square[rand.Intn(len(Square))]

	// Formation of additional rooms
	if i.Square < 49 {
		i.MedicalOffice = IsOrIsnt[1]
		i.GreenHouses = IsOrIsnt[1]
		i.Armory = IsOrIsnt[1]
	} else {
		i.MedicalOffice = IsOrIsnt[rand.Intn(2)]
		i.GreenHouses = IsOrIsnt[rand.Intn(2)]
		i.Armory = IsOrIsnt[rand.Intn(2)]
	}

	s := TimeFormatter()
	if s == "40 років" || s == "69 років" || s == "100 років" || s == "999 років" {
		s = TimeFormatter()
	}
	i.FoodStocks = s

	// Formation of additional items
	if rand.Intn(7) == 0 {
		i.AdditionalItem1 = " нічого"
	} else {
		i.AdditionalItem1 = AdditionalItems[rand.Intn(len(AdditionalItems))]
	}
	if rand.Intn(7) == 0 {
		i.AdditionalItem2 = " нічого"
	} else {
		i.AdditionalItem2 = AdditionalItems[rand.Intn(len(AdditionalItems))]
	}

	info := fmt.Sprintf("🔲 <b>Площа:</b> <i>%d м2</i>\n🍲 <b>Запасів води і їжі на:</b> <i>%s</i>\n",
		i.Square, i.FoodStocks)
	info += fmt.Sprintf("🏥 <b>Медичний кабінет:</b> <i>%s</i>\n🌿 <b>Город із теплицями:</b> <i>%s</i>\n🔫 <b>Зброярня:</b> <i>%s</i>\n",
		i.MedicalOffice, i.GreenHouses, i.Armory)
	info += fmt.Sprintf("📦 <b>Додатковий предмет 1:</b>\n<i>%s</i>\n📦 <b>Додатковий предмет 2:</b>\n<i>%s</i>\n",
		i.AdditionalItem1, i.AdditionalItem2)

	return info
}
