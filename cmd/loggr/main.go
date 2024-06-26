package main

import (
	"log/slog"
	"os"

	"github.com/Linkinlog/loggr/internal/env"
	"github.com/Linkinlog/loggr/internal/handlers"
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/services"
	"github.com/Linkinlog/loggr/internal/stores"
)

func main() {
	e := env.NewEnv()
	programLevel := slog.LevelError
	debug := e.GetOrDefault("DEBUG", "false")
	if debug == "true" {
		programLevel = slog.LevelDebug
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))

	sqliteAddr := e.GetOrDefault("SQLITE_ADDR", "file:./loggr.db")
	sendGridKey := e.GetOrDefault("SENDGRID_API_KEY", "")

	s := stores.NewSqliteStore(sqliteAddr)
	ir := repositories.NewItemRepository(s)
	ur := repositories.NewUserRepository(s)
	gr := repositories.NewGardenRepository(s)
	sr := repositories.NewSessionRepository(s)
	ms := services.NewMailService(sendGridKey)
	addr := e.GetOrDefault("ADDR", ":8080")
	ssr := handlers.NewSSR(l, addr, createExamples(), ur, gr, ir, sr, ms)

	l.Info("its_alive!", slog.String("addr", addr))

	if err := ssr.ServeHTTP(); err != nil {
		l.Error("error serving http", "err", err.Error())
	}
}

func createExamples() []*models.Garden {
	field1 := "Palm Tree"
	field2 := "Soil mix"
	field3 := "Direct sunlight"
	item := models.NewItem("Tall boy", "https://i.ibb.co/r5Ykp62/photo-1506953823976-52e1fdc0149a.jpg", models.Plant, [5]string{field1, field2, field3})
	b := models.NewGarden("Acorn Acres (Demo)", "Out back", "Demo garden description", "https://i.ibb.co/w6LcG0X/acorn-Acres.jpg", []*models.Item{item})

	longField1 := "Bacon ipsum dolor amet venison drumstick jerky, cupim andouille kevin doner salami pastrami beef. Turducken pork loin doner chislic frankfurter. Turkey bresaola ham, picanha landjaeger short loin rump fatback meatloaf bacon pork belly jerky cupim pork chop. Prosciutto short loin doner shoulder. Landjaeger meatball porchetta, beef swine pork filet mignon bresaola pig beef ribs kevin hamburger.  Cow cupim short ribs salami ground round chicken tongue. Porchetta cupim swine sirloin corned beef alcatra picanha prosciutto pastrami chuck. Chicken hamburger tri-tip fatback. Landjaeger venison kevin capicola tenderloin. Hamburger buffalo chuck bacon leberkas pork loin venison turkey.  Meatball cupim jowl meatloaf andouille, burgdoggen salami. Tenderloin strip steak pork pork loin picanha, burgdoggen kevin drumstick ball tip landjaeger spare ribs filet mignon chicken alcatra rump. Bresaola capicola leberkas, fatback drumstick meatloaf pork brisket cupim. Pastrami pork chop boudin shankle bresaola leberkas picanha tenderloin. Meatloaf flank filet mignon, brisket short ribs bresaola alcatra cow beef turducken.  Leberkas doner salami, tri-tip boudin hamburger landjaeger cow meatloaf pork loin burgdoggen pancetta bacon. Kielbasa ham flank tenderloin pastrami ball tip shankle. Beef ribs picanha short ribs boudin, landjaeger meatloaf turkey. Tri-tip tenderloin ham hock hamburger shoulder beef ribs. Jowl burgdoggen pork tail. Frankfurter short ribs bacon sirloin, ham kielbasa meatloaf leberkas andouille cupim pig swine biltong. Chicken short loin sirloin, venison porchetta salami frankfurter flank ball tip ribeye hamburger.  Chuck pig andouille strip steak filet mignon ball tip cupim chislic. Pork belly flank burgdoggen filet mignon pork chop beef cupim capicola ball tip alcatra chislic sirloin. Pastrami short ribs porchetta picanha short loin chicken sausage ribeye. Landjaeger tri-tip biltong rump, chislic pig bresaola tongue jerky ham fatback meatloaf leberkas. Porchetta strip steak shank tongue andouille hamburger shoulder jowl.  Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!"

	longField2 := "Bacon ipsum dolor amet venison drumstick jerky, cupim andouille kevin doner salami pastrami beef. Turducken pork loin doner chislic frankfurter. Turkey bresaola ham, picanha landjaeger short loin rump fatback meatloaf bacon pork belly jerky cupim pork chop. Prosciutto short loin doner shoulder. Landjaeger meatball porchetta, beef swine pork filet mignon bresaola pig beef ribs kevin hamburger.  Cow cupim short ribs salami ground round chicken tongue. Porchetta cupim swine sirloin corned beef alcatra picanha prosciutto pastrami chuck. Chicken hamburger tri-tip fatback. Landjaeger venison kevin capicola tenderloin. Hamburger buffalo chuck bacon leberkas pork loin venison turkey.  Meatball cupim jowl meatloaf andouille, burgdoggen salami. Tenderloin strip steak pork pork loin picanha, burgdoggen kevin drumstick ball tip landjaeger spare ribs filet mignon chicken alcatra rump. Bresaola capicola leberkas, fatback drumstick meatloaf pork brisket cupim. Pastrami pork chop boudin shankle bresaola leberkas picanha tenderloin. Meatloaf flank filet mignon, brisket short ribs bresaola alcatra cow beef turducken.  Leberkas doner salami, tri-tip boudin hamburger landjaeger cow meatloaf pork loin burgdoggen pancetta bacon. Kielbasa ham flank tenderloin pastrami ball tip shankle. Beef ribs picanha short ribs boudin, landjaeger meatloaf turkey. Tri-tip tenderloin ham hock hamburger shoulder beef ribs. Jowl burgdoggen pork tail. Frankfurter short ribs bacon sirloin, ham kielbasa meatloaf leberkas andouille cupim pig swine biltong. Chicken short loin sirloin, venison porchetta salami frankfurter flank ball tip ribeye hamburger.  Chuck pig andouille strip steak filet mignon ball tip cupim chislic. Pork belly flank burgdoggen filet mignon pork chop beef cupim capicola ball tip alcatra chislic sirloin. Pastrami short ribs porchetta picanha short loin chicken sausage ribeye. Landjaeger tri-tip biltong rump, chislic pig bresaola tongue jerky ham fatback meatloaf leberkas. Porchetta strip steak shank tongue andouille hamburger shoulder jowl.  Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!"

	longField3 := "Bacon ipsum dolor amet venison drumstick jerky, cupim andouille kevin doner salami pastrami beef. Turducken pork loin doner chislic frankfurter. Turkey bresaola ham, picanha landjaeger short loin rump fatback meatloaf bacon pork belly jerky cupim pork chop. Prosciutto short loin doner shoulder. Landjaeger meatball porchetta, beef swine pork filet mignon bresaola pig beef ribs kevin hamburger.  Cow cupim short ribs salami ground round chicken tongue. Porchetta cupim swine sirloin corned beef alcatra picanha prosciutto pastrami chuck. Chicken hamburger tri-tip fatback. Landjaeger venison kevin capicola tenderloin. Hamburger buffalo chuck bacon leberkas pork loin venison turkey.  Meatball cupim jowl meatloaf andouille, burgdoggen salami. Tenderloin strip steak pork pork loin picanha, burgdoggen kevin drumstick ball tip landjaeger spare ribs filet mignon chicken alcatra rump. Bresaola capicola leberkas, fatback drumstick meatloaf pork brisket cupim. Pastrami pork chop boudin shankle bresaola leberkas picanha tenderloin. Meatloaf flank filet mignon, brisket short ribs bresaola alcatra cow beef turducken.  Leberkas doner salami, tri-tip boudin hamburger landjaeger cow meatloaf pork loin burgdoggen pancetta bacon. Kielbasa ham flank tenderloin pastrami ball tip shankle. Beef ribs picanha short ribs boudin, landjaeger meatloaf turkey. Tri-tip tenderloin ham hock hamburger shoulder beef ribs. Jowl burgdoggen pork tail. Frankfurter short ribs bacon sirloin, ham kielbasa meatloaf leberkas andouille cupim pig swine biltong. Chicken short loin sirloin, venison porchetta salami frankfurter flank ball tip ribeye hamburger.  Chuck pig andouille strip steak filet mignon ball tip cupim chislic. Pork belly flank burgdoggen filet mignon pork chop beef cupim capicola ball tip alcatra chislic sirloin. Pastrami short ribs porchetta picanha short loin chicken sausage ribeye. Landjaeger tri-tip biltong rump, chislic pig bresaola tongue jerky ham fatback meatloaf leberkas. Porchetta strip steak shank tongue andouille hamburger shoulder jowl.  Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!"

	longItem := models.NewItem("Pneumonoultramicroscopicsilicovolcanoconiosis Item (Demo)", "https://i.ibb.co/r5Ykp62/photo-1506953823976-52e1fdc0149a.jpg", models.Plant, [5]string{longField1, longField2, longField3})

	long := models.NewGarden("Pneumonoultramicroscopicsilicovolcanoconiosis Garden (Demo)", "This is an example with a lot of text in each field", "Bacon ipsum dolor amet venison drumstick jerky, cupim andouille kevin doner salami pastrami beef. Turducken pork loin doner chislic frankfurter. Turkey bresaola ham, picanha landjaeger short loin rump fatback meatloaf bacon pork belly jerky cupim pork chop. Prosciutto short loin doner shoulder. Landjaeger meatball porchetta, beef swine pork filet mignon bresaola pig beef ribs kevin hamburger.  Cow cupim short ribs salami ground round chicken tongue. Porchetta cupim swine sirloin corned beef alcatra picanha prosciutto pastrami chuck. Chicken hamburger tri-tip fatback. Landjaeger venison kevin capicola tenderloin. Hamburger buffalo chuck bacon leberkas pork loin venison turkey.  Meatball cupim jowl meatloaf andouille, burgdoggen salami. Tenderloin strip steak pork pork loin picanha, burgdoggen kevin drumstick ball tip landjaeger spare ribs filet mignon chicken alcatra rump. Bresaola capicola leberkas, fatback drumstick meatloaf pork brisket cupim. Pastrami pork chop boudin shankle bresaola leberkas picanha tenderloin. Meatloaf flank filet mignon, brisket short ribs bresaola alcatra cow beef turducken.  Leberkas doner salami, tri-tip boudin hamburger landjaeger cow meatloaf pork loin burgdoggen pancetta bacon. Kielbasa ham flank tenderloin pastrami ball tip shankle. Beef ribs picanha short ribs boudin, landjaeger meatloaf turkey. Tri-tip tenderloin ham hock hamburger shoulder beef ribs. Jowl burgdoggen pork tail. Frankfurter short ribs bacon sirloin, ham kielbasa meatloaf leberkas andouille cupim pig swine biltong. Chicken short loin sirloin, venison porchetta salami frankfurter flank ball tip ribeye hamburger.  Chuck pig andouille strip steak filet mignon ball tip cupim chislic. Pork belly flank burgdoggen filet mignon pork chop beef cupim capicola ball tip alcatra chislic sirloin. Pastrami short ribs porchetta picanha short loin chicken sausage ribeye. Landjaeger tri-tip biltong rump, chislic pig bresaola tongue jerky ham fatback meatloaf leberkas. Porchetta strip steak shank tongue andouille hamburger shoulder jowl.  Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!", "https://i.ibb.co/w6LcG0X/acorn-Acres.jpg", []*models.Item{longItem})

	return []*models.Garden{b, long}
}
