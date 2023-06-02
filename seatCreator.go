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

	rectDsection1 := eraser.Rectangle(pos.Absolute{X: 5,Y: 18}, pos.Absolute{X: 5,Y: 20})
	rectDsection2 := eraser.Rectangle(pos.Absolute{X: 6,Y: 19}, pos.Absolute{X: 6,Y: 20})

	rectEsection1 := eraser.Rectangle(pos.Absolute{X: 34,Y: 21}, pos.Absolute{X: 36,Y: 24})
	rectEsection2 := eraser.Rectangle(pos.Absolute{X: 33,Y: 22}, pos.Absolute{X: 33,Y: 24})
	rectEsection3 := eraser.Rectangle(pos.Absolute{X: 32,Y: 23}, pos.Absolute{X: 32,Y: 24})
	specificEsection1 := eraser.Position(pos.Absolute{X: 31, Y: 24})

	rectFsection1 := eraser.Rectangle(pos.Absolute{X:50, Y: 18 }, pos.Absolute{X: 50, Y: 20})
	rectFsection2 := eraser.Rectangle(pos.Absolute{X:49, Y: 19 }, pos.Absolute{X: 49, Y: 20})
	
	integrated := eraser.Integrated(rectDsection1, rectDsection2 ,
									rectEsection1, rectEsection2,rectEsection3, specificEsection1,
									rectFsection1, rectFsection2)
	base := model.NewSeatBase(5, 17, "D석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         4,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(18, 17, "E석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         2,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(22, 19, "E석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 5, 3),
		XSize:         17,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(21, 20, "E석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 4),
		XSize:         16,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(40, 17, "F석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         4,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(7, 1, "L1석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(8, 9, "L2석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         7,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(46, 1, "R1석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(45, 9, "R2석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         7,
		EmptyChecker:  integrated,
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
		Key:             "2F",
		Seats:           seats,
		XSize:           45,
		YSize:           25,
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