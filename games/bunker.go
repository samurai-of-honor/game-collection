package games

import (
	"fmt"
	s "game-collection/settings"
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

func BunkerCharacter() string {
	c := new(Character)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	c.Gender = s.Gender[rand.Intn(len(s.Gender))]
	c.Age = s.Age[rand.Intn(len(s.Age))]
	c.Profession = s.Profession[rand.Intn(len(s.Profession))]

	// Formation of experience
	exp := s.ProfessionExp[rand.Intn(len(s.ProfessionExp))]
	expStr := strconv.Itoa(exp)
	if exp >= c.Age-18 && exp > 10 {
		exp = s.ProfessionExp[rand.Intn(len(s.ProfessionExp)-7)]
		expStr = strconv.Itoa(exp)
		c.ProfessionExp = expStr + s.ProfessionExpDate[0]
	} else if exp <= 10 && rand.Intn(2) == 1 {
		c.ProfessionExp = expStr + s.ProfessionExpDate[0]
	} else {
		c.ProfessionExp = expStr + s.ProfessionExpDate[1]
	}

	// Health
	if rand.Intn(6) == 0 {
		c.HealthStatus = "здоровий"
	} else {
		c.HealthStatus = s.HealthStatus[rand.Intn(len(s.HealthStatus))]
	}

	c.Fertility = s.Fertility[rand.Intn(len(s.Fertility))]
	c.BodyType = s.BodyType[rand.Intn(len(s.BodyType))]

	// Phobias
	if rand.Intn(12) == 0 {
		c.Phobia = "немає фобій"
	} else {
		c.Phobia = s.Phobia[rand.Intn(len(s.Phobia))]
	}

	c.Hobby = s.Hobby[rand.Intn(len(s.Hobby))]
	c.AdditionalInfo = s.AdditionalInfo[rand.Intn(len(s.AdditionalInfo))]
	c.Baggage = s.Baggage[rand.Intn(len(s.Baggage))]
	c.Card1 = s.Cards[rand.Intn(len(s.Cards))]
	c.Card2 = s.Cards[rand.Intn(len(s.Cards))]

	return fmt.Sprintf("🚻 <b>Стать:</b> <i>%s</i>\n🔞 <b>Вік:</b> %d\n🥋 <b>Професія:</b> <i>%s</i>\n  \t \t<b>Стаж:</b> <i>%s</i>\n🤰 <b>Здатність до запліднення:</b> <i>%s</i>\n💊 <b>Стан здоров'я:</b> <i>%s</i>\n💪 <b>Тип тіла:</b> <i>%s</i>\n😱 <b>Фобія:</b> <i>%s</i>\n🏆 <b>Хобі:</b> <i>%s</i>\n🎟 <b>Дод. інформація:</b> <i>%s</i>\n💼 <b>Багаж:</b> <i>%s</i>\n1️⃣ <b>Карта 1:</b> <i>%s</i>\n2️⃣ <b>Карта 2:</b> <i>%s</i>\n",
		c.Gender, c.Age, c.Profession, c.ProfessionExp, c.Fertility, c.HealthStatus, c.BodyType, c.Phobia, c.Hobby, c.AdditionalInfo, c.Baggage, c.Card1, c.Card2)
}
