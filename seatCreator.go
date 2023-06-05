package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/area"
	"github.com/simp7/seatCreator/model/eraser"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
)

func ConcertHall1F() model.Group {
	nameFormatter := nameformatter.Standard()
	
	integrated := eraser.Integrated()
	base := model.NewSeatBase(4, 4, "가석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(3, 4, "가석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 9, 1),
		XSize:         1,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 14, "나석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(3, 14, "나석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 9, 1),
		XSize:         1,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific := eraser.Position(pos.Absolute{X: 35, Y: 12}, pos.Absolute{X: 35,Y: 22})
	integrated2 := eraser.Integrated(specific)
	base = model.NewSeatBase(35, 4, "다석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         2,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('A'),
	}

	base = model.NewSeatBase(35, 14, "라석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         2,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('A'),
	}
	

	block1 := group.VerticalBlock(blockInput)
	block2 := group.VerticalBlock(blockInput2)
	block3 := group.VerticalBlock(blockInput3)
	block4 := group.VerticalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	return group.Mixed(block1, block2, block3, block4, block5, block6)
}


func copy(target model.Group) {
	err := clipboard.WriteAll(target.String())
	if err != nil {
		fmt.Println("Error occured when copying")
		return
	}
	fmt.Println("Copied!")
}

func handler(c *gin.Context) {
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           36,
		YSize:           35,
		BackgroundImage: "",
		Color:           "#ff9f00",
	}

	fmt.Fprintf(c.Writer, "<h1>Seats Map</h1><body>\n%s\n</body>", target.Html())
	copy(target)
}
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", handler)
	r.Run()
}