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

	Arect1 := eraser. Rectangle(pos.Absolute{X: 3, Y: 22}, pos.Absolute{X: 3, Y: 26})
	Arect2 := eraser. Rectangle(pos.Absolute{X: 4, Y: 24}, pos.Absolute{X: 4, Y: 26})
	Arect3 := eraser. Rectangle(pos.Absolute{X: 5, Y: 25}, pos.Absolute{X: 5, Y: 26})
	integrated1 := eraser.Integrated(Arect1,Arect2,Arect3)
	base := model.NewSeatBase(3, 21, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         6,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	Brect1 := eraser.Rectangle(pos.Absolute{X: 18, Y: 21}, pos.Absolute{X: 18, Y: 25})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 41, Y: 21}, pos.Absolute{X: 41, Y: 24})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 19, Y: 21}, pos.Absolute{X: 19, Y: 23})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 40, Y: 21}, pos.Absolute{X: 40, Y: 22})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 20, Y: 21}, pos.Absolute{X: 20, Y: 21})
	integrated2 := eraser.Integrated(Brect1,Brect2,Brect3,Brect4,Brect5)
	base = model.NewSeatBase(18, 21, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         6,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 56, Y: 22}, pos.Absolute{X: 56, Y: 26})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 55, Y: 24}, pos.Absolute{X: 55, Y: 26})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 54, Y: 25}, pos.Absolute{X: 54, Y: 26})
	integrated3 := eraser.Integrated(Crect1,Crect2,Crect3)
	base = model.NewSeatBase(43, 21, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         6,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 3, "L10석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 9, "L11석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 14, "L12석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 16, "L12석")
	blockInput11 := group.BlockInput{
		Criteria:      model.NewSeat(base, 7, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(55, 3, "R10석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(55, 9, "R11석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(55, 14, "R12석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(56, 16, "R12석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 7, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.VerticalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	block7 := group.VerticalBlock(blockInput7)
	block8 := group.VerticalBlock(blockInput8)
	block9 := group.VerticalBlock(blockInput9)
	block10 := group.VerticalBlock(blockInput10)
	block11 := group.VerticalBlock(blockInput11)

	return group.Mixed(block1, block2, block3, block4, block5,block6,block7,block8,block9,block10, block11)
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
		Key:             "4F",
		Seats:           seats,
		XSize:           58,
		YSize:           24,
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
