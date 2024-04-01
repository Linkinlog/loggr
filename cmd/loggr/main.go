package main

import (
	"log/slog"
	"os"

	"github.com/Linkinlog/loggr/internal/handlers"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

const addr = ":8080"

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	field1 := models.NewField("field 1", "field 1 description")
	field2 := models.NewField("field 2", "field 2 description")
	item1 := models.NewItem("item 1", models.NewImage("id", "https://i.ibb.co/N9kj08h/bandit.jpg", "https://i.ibb.co/qR0Qz05/image.jpg", ""), models.Plant, [5]*models.Field{field1, field2})
	b := models.NewGarden("Bandit Acres", "Somewhere", "Worlds sweetest boy", models.NewImage("qR0Qz05", "https://i.ibb.co/N9kj08h/bandit.jpg", "https://i.ibb.co/qR0Qz05/image.jpg", ""), []*models.Item{item1})

	ma := models.NewGarden("Maggie Falls", "All around", "The bestest girl", models.NewImage("HCGb3p4", "https://i.ibb.co/ZScF71X/IMG-5361.jpg", "https://i.ibb.co/HCGb3p4/IMG-5361.jpg", ""), []*models.Item{})

	mi := models.NewGarden("Mimarigolds", "Blooms everywhere", "Needs direct sunlight And plenty of attention", models.NewImage("G7dpvmZ", "https://i.ibb.co/y0NPn2z/IMG-2115.jpg", "https://i.ibb.co/G7dpvmZ/IMG-2115.jpg", ""), []*models.Item{})

	long := models.NewGarden("Long Garden is quite Long", "Long Island, far far far away from Long Garden, which is not where this location is at", "Bacon ipsum dolor amet venison drumstick jerky, cupim andouille kevin doner salami pastrami beef. Turducken pork loin doner chislic frankfurter. Turkey bresaola ham, picanha landjaeger short loin rump fatback meatloaf bacon pork belly jerky cupim pork chop. Prosciutto short loin doner shoulder. Landjaeger meatball porchetta, beef swine pork filet mignon bresaola pig beef ribs kevin hamburger.  Cow cupim short ribs salami ground round chicken tongue. Porchetta cupim swine sirloin corned beef alcatra picanha prosciutto pastrami chuck. Chicken hamburger tri-tip fatback. Landjaeger venison kevin capicola tenderloin. Hamburger buffalo chuck bacon leberkas pork loin venison turkey.  Meatball cupim jowl meatloaf andouille, burgdoggen salami. Tenderloin strip steak pork pork loin picanha, burgdoggen kevin drumstick ball tip landjaeger spare ribs filet mignon chicken alcatra rump. Bresaola capicola leberkas, fatback drumstick meatloaf pork brisket cupim. Pastrami pork chop boudin shankle bresaola leberkas picanha tenderloin. Meatloaf flank filet mignon, brisket short ribs bresaola alcatra cow beef turducken.  Leberkas doner salami, tri-tip boudin hamburger landjaeger cow meatloaf pork loin burgdoggen pancetta bacon. Kielbasa ham flank tenderloin pastrami ball tip shankle. Beef ribs picanha short ribs boudin, landjaeger meatloaf turkey. Tri-tip tenderloin ham hock hamburger shoulder beef ribs. Jowl burgdoggen pork tail. Frankfurter short ribs bacon sirloin, ham kielbasa meatloaf leberkas andouille cupim pig swine biltong. Chicken short loin sirloin, venison porchetta salami frankfurter flank ball tip ribeye hamburger.  Chuck pig andouille strip steak filet mignon ball tip cupim chislic. Pork belly flank burgdoggen filet mignon pork chop beef cupim capicola ball tip alcatra chislic sirloin. Pastrami short ribs porchetta picanha short loin chicken sausage ribeye. Landjaeger tri-tip biltong rump, chislic pig bresaola tongue jerky ham fatback meatloaf leberkas. Porchetta strip steak shank tongue andouille hamburger shoulder jowl.  Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!", models.NewImage("G7dpvmZ", "https://i.ibb.co/kKcfngX/craiyon-120715-pixel-art-of-smart-doctor-corgi.png", "https://i.ibb.co/hs8j5RV/craiyon-120715-pixel-art-of-smart-doctor-corgi.png", ""), []*models.Item{})

	m := stores.NewInMemory([]*models.Garden{b, ma, mi, long})
	ssr := handlers.NewSSR(l, addr, m)

	l.Info("its_alive!", slog.String("addr", addr))

	if err := ssr.ServeHTTP(); err != nil {
		l.Error("error serving http", err)
	}
}
