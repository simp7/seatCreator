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

func ArtriumSmall() model.Group {

	nameFormatter := nameformatter.Standard()
	integrated := eraser.Integrated()

	base := model.NewSeatBase(3, 6, "가석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 3),
		XSize:         10,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(5, 4, "가석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 1),
		XSize:         8,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 5, "가석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 2),
		XSize:         9,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(14, 4, "나석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	position1 := eraser.Position(pos.Absolute{X: 37,Y: 4})
	rect1 := eraser.Rectangle(pos.Absolute{X: 38, Y: 4}, pos.Absolute{X: 38, Y: 5})
	integrated2 := eraser.Integrated(position1, rect1)
	base = model.NewSeatBase(29, 4, "다석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         10,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 7, Y: 21}, pos.Absolute{X: 12, Y: 21})
	integrated3 := eraser.Integrated(rect2)
	base = model.NewSeatBase(3, 15, "라석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         7,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(14, 15, "마석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(29, 15, "바석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(35, 21, "바석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 7, 7),
		XSize:         4,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 10, Y: 21}, pos.Absolute{X: 31, Y: 21})
	integrated4 := eraser.Integrated(rect3)
	base = model.NewSeatBase(7, 21, "휠체어석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         28,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.HorizontalBlock(blockInput8)
	block9 := group.HorizontalBlock(blockInput9)
	block10 := group.HorizontalBlock(blockInput10)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8,block9, block10)
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
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           36,
		YSize:           18,
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