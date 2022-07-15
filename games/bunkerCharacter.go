package games

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Character struct {
	Gender         string
	Age            int
	Profession     string
	ProfessionExp  string
	Fertility      string
	HealthStatus   string
	BodyType       string
	Phobia         string
	Hobby          string
	AdditionalInfo string
	Baggage        string
	Card1          string
	Card2          string
}

func (c *Character) CharacterFormatting() string {
	character := fmt.Sprintf("🚻 <b>Стать:</b> <i>%s</i>\n🔞 <b>Вік:</b> %d\n🥋 <b>Професія:</b> <i>%s</i>\n  \t \t<b>Стаж:</b> <i>%s</i>\n",
		c.Gender, c.Age, c.Profession, c.ProfessionExp)
	character += fmt.Sprintf("🤰 <b>Здатність до запліднення:</b> <i>%s</i>\n💊 <b>Стан здоров'я:</b> <i>%s</i>\n💪 <b>Тип тіла:</b> <i>%s</i>\n",
		c.Fertility, c.HealthStatus, c.BodyType)
	character += fmt.Sprintf("😱 <b>Фобія:</b> <i>%s</i>\n🏆 <b>Хобі:</b> <i>%s</i>\n🎟 <b>Дод. інформація:</b> <i>%s</i>\n",
		c.Phobia, c.Hobby, c.AdditionalInfo)
	character += fmt.Sprintf("💼 <b>Багаж:</b> <i>%s</i>\n1️⃣ <b>Карта 1:</b> <i>%s</i>\n2️⃣ <b>Карта 2:</b> <i>%s</i>",
		c.Baggage, c.Card1, c.Card2)
	return character
}

func BunkerCharacter() string {
	c := new(Character)
	rand.Seed(time.Now().Unix()) // initialize pseudo random generator

	c.Gender = Gender[rand.Intn(len(Gender))]
	c.Age = Age[rand.Intn(len(Age))]
	c.Profession = Profession[rand.Intn(len(Profession))]

	// Formation of experience
	exp := ProfessionExp[rand.Intn(len(ProfessionExp))]
	expStr := strconv.Itoa(exp)
	if exp >= c.Age-18 && exp > 10 {
		exp = ProfessionExp[rand.Intn(len(ProfessionExp)-7)]
		expStr = strconv.Itoa(exp)
		c.ProfessionExp = expStr + ProfessionExpDate[0]
	} else if exp <= 10 && rand.Intn(2) == 1 {
		c.ProfessionExp = expStr + ProfessionExpDate[0]
	} else {
		c.ProfessionExp = expStr + ProfessionExpDate[1]
	}

	// Formation of fertility
	if c.Gender == Gender[0] && c.Age >= 55 {
		c.Fertility = Fertility[0]
	} else if c.Gender == Gender[1] && c.Age >= 65 {
		c.Fertility = Fertility[0]
	} else {
		c.Fertility = Fertility[rand.Intn(len(Fertility))]
	}

	// Formation of health
	if rand.Intn(8) == 0 {
		c.HealthStatus = "здоровий"
	} else {
		health := rand.Intn(len(HealthStatus))
		if health < len(HealthStatus)-4 {
			stage := strconv.Itoa(IllnessStage[rand.Intn(len(IllnessStage))])
			c.HealthStatus = HealthStatus[health] + " " + stage + "%"
		} else {
			c.HealthStatus = HealthStatus[health]
		}
	}

	c.BodyType = BodyType[rand.Intn(len(BodyType))]

	// Phobias
	if rand.Intn(12) == 0 {
		c.Phobia = "немає фобій"
	} else {
		c.Phobia = Phobia[rand.Intn(len(Phobia))]
	}

	c.Hobby = Hobby[rand.Intn(len(Hobby))]
	c.AdditionalInfo = AdditionalInfo[rand.Intn(len(AdditionalInfo))]
	c.Baggage = Baggage[rand.Intn(len(Baggage))]
	c.Card1 = Cards[rand.Intn(len(Cards))]
	c.Card2 = Cards[rand.Intn(len(Cards))]

	return c.CharacterFormatting()
}
