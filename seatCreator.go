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

	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 4},pos.Absolute{X: 5,Y: 4})
	rect2 := eraser.Rectangle(pos.Absolute{X: 3,Y: 5},pos.Absolute{X: 4,Y: 6})
	rect3 := eraser.Rectangle(pos.Absolute{X: 3,Y: 7},pos.Absolute{X: 3,Y: 8})
	integrated := eraser.Integrated(rect1, rect2, rect3)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('A'), 1),
	}

	rect4 := eraser.Rectangle(pos.Absolute{X: 22,Y: 4}, pos.Absolute{X: 22,Y: 6})
	integrated2 := eraser.Integrated(rect4)
	base = model.NewSeatBase(9, 4, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('B'), 1),
	}

	rect5 := eraser.Rectangle(pos.Absolute{X: 26,Y: 4},pos.Absolute{X: 28,Y: 4})
	rect6 := eraser.Rectangle(pos.Absolute{X: 27,Y: 5},pos.Absolute{X: 28,Y: 6})
	rect7 := eraser.Rectangle(pos.Absolute{X: 28,Y: 7},pos.Absolute{X: 28,Y: 8})
	integrated3 := eraser.Integrated(rect5,rect6,rect7)
	base = model.NewSeatBase(24, 4, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('C'), 1),
	}

	rect8 := eraser.Rectangle(pos.Absolute{X: 3,Y: 22}, pos.Absolute{X: 3,Y: 23})
	integrated4 := eraser.Integrated(rect8)
	base = model.NewSeatBase(3, 14, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         10,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('D'), 1),
	}
	base = model.NewSeatBase(9, 14, "A석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('E'), 1),
	}

	rect9 := eraser.Rectangle(pos.Absolute{X: 28, Y: 22}, pos.Absolute{X: 28, Y: 23})
	integrated5 := eraser.Integrated(rect9)
	base = model.NewSeatBase(24, 14, "A석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         10,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('F'), 1),
	}

	specific1 := eraser.Position(pos.Absolute{X: 8,Y: 2}, pos.Absolute{X: 22,Y: 2}, pos.Absolute{X: 23,Y: 2})
	integrated6 := eraser.Integrated(specific1)
	base = model.NewSeatBase(6, 2, "A석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         1,
		EmptyChecker:  integrated6,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('O'), 1),
	}

	rect10 := eraser.Rectangle(pos.Absolute{X: 8,Y: 24}, pos.Absolute{X: 23,Y: 24})
	integrated7 := eraser.Integrated(rect10)
	base = model.NewSeatBase(4, 24, "휠체어석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         1,
		EmptyChecker:  integrated7,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('W'), 1),
	}

	rect11 := eraser.Rectangle(pos.Absolute{X: 11,Y: 23},pos.Absolute{X: 20,Y: 23})
	integrated8 := eraser.Integrated(rect11)
	base = model.NewSeatBase(9, 23, "휠체어석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 9, 23),
		XSize:         14,
		YSize:         1,
		EmptyChecker:  integrated8,
		NameFormatter: nameFormatter,
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
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9)
}

func ConcertHall1F() model.Group {
	// hall1 := eraser.VerticalHallway(12,13)
	// rect1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 5}, pos.Absolute{X: 2, Y: 6})
	// rect2 := eraser.Rectangle(pos.Absolute{X: 24, Y: 4}, pos.Absolute{X: 24, Y: 10})
	// rect3 := eraser.Rectangle(pos.Absolute{X: 23, Y: 5}, pos.Absolute{X: 23, Y: 6})
	// rect4 := eraser.Rectangle(pos.Absolute{X: 2, Y: 11}, pos.Absolute{X: 2, Y: 13})
	// rect5 := eraser.Rectangle(pos.Absolute{X: 23, Y: 12}, pos.Absolute{X: 24, Y: 13})

	// specific := eraser.Position(pos.Absolute{X: 3,Y: 12}, pos.Absolute{X: 22,Y: 12})

	// integrated := eraser.Integrated(hall1,rect1, rect2,rect3,rect4,rect5, specific)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 4, "A석")
	// blockInput := group.BlockInput{
	// 	Criteria:      model.NewSeat(base, 1, 1),
	// 	XSize:         23,
	// 	YSize:         10,
	// 	EmptyChecker:  integrated,
	// 	NameFormatter: nameFormatter,
	// }


	// integrated2 := eraser.Integrated()
	// base = model.NewSeatBase(4, 14, "A석")
	// blockInput2 := group.BlockInput{
	// 	Criteria:      model.NewSeat(base, 1, 11),
	// 	XSize:         13,
	// 	YSize:         1,
	// 	EmptyChecker:  integrated2,
	// 	NameFormatter: nameFormatter,
	// }

	rect6 := eraser.Rectangle(pos.Absolute{X: 16, Y: 16}, pos.Absolute{X: 17, Y: 17})
	specific2 := eraser.Position(pos.Absolute{X: 20,Y: 16})
	integrated3 := eraser.Integrated(rect6,specific2)

	base = model.NewSeatBase(6, 16, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 12),
		XSize:         15,
		YSize:         2,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	// block1 := group.HorizontalBlock(blockInput)
	// block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block3)
}

func ConcertHall2F() model.Group {
		nameFormatter := nameformatter.Standard()
		top1 := func() model.Group {
			base := model.NewSeatBase(6, 2, "B석 C1열")
			seat := model.NewSeat(base, 1, 3)
			integrated := eraser.Integrated()
			return group.HorizontalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         32,
				YSize:         1,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}
		top2 := func() model.Group {
			base := model.NewSeatBase(7, 3, "B석 C2열")
			seat := model.NewSeat(base, 2, 3)
			integrated := eraser.Integrated()
			return group.HorizontalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         30,
				YSize:         1,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		top3 := func() model.Group {
			base := model.NewSeatBase(8, 4, "B석 C3열")
			seat := model.NewSeat(base, 3, 3)
			integrated := eraser.Integrated()
			return group.HorizontalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         28,
				YSize:         1,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		left1 := func() model.Group {
			base := model.NewSeatBase(3, 5, "A석 L3열")
			seat := model.NewSeat(base, 1, 12)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         21,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		left2 := func() model.Group {
			base := model.NewSeatBase(2, 6, "A석 L2열")
			seat := model.NewSeat(base, 2, 12)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         20,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		left3 := func() model.Group {
			base := model.NewSeatBase(1, 7, "A석 L1열")
			seat := model.NewSeat(base, 3, 12)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         19,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		right1 := func() model.Group {
			base := model.NewSeatBase(40, 5, "A석 R1열")
			seat := model.NewSeat(base, 1, 18)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         21,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		right2 := func() model.Group {
			base := model.NewSeatBase(41, 6, "A석 R2열")
			seat := model.NewSeat(base, 2, 18)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         20,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		right3 := func() model.Group {
			base := model.NewSeatBase(42, 7, "A석 R3열")
			seat := model.NewSeat(base, 3, 18)
			integrated := eraser.Integrated()
			return group.VerticalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         1,
				YSize:         19,
				NameFormatter: nameFormatter,
				EmptyChecker:  integrated,
			})
		}

		bottom := func() model.Group {
			base := model.NewSeatBase(5, 24, "S석")
			seat := model.NewSeat(base, 1, 1)
			empty := eraser.VerticalHallway(14, 29)
			return group.HorizontalBlock(group.BlockInput{
				Criteria:      seat,
				XSize:         34,
				YSize:         7,
				NameFormatter: nameformatter.Floor(nameformatter.Standard(), 2),
				EmptyChecker:  empty,
			})
		}
		hall := group.Mixed(top1(), top2(),top3(), left1(),left2(),left3(), right1(),right2(),right3(), bottom())
		return hall
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
		XSize:           26,
		YSize:           20,
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
