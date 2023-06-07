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

	//1층 무대 앞부분
	hall1 := eraser.HorizontalHallway(16)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 4}, pos.Absolute{X: 3,Y: 15})
	rectPosition1 := eraser.Rectangle(pos.Absolute{X: 4,Y: 4}, pos.Absolute{X: 4,Y: 14})
	rectPosition2 := eraser.Rectangle(pos.Absolute{X: 5,Y: 4},pos.Absolute{X: 5,Y: 5})

	rect2 := eraser.Rectangle(pos.Absolute{X: 49,Y: 4}, pos.Absolute{X: 50,Y: 15})

	rect3 := eraser.Rectangle(pos.Absolute{X: 11,Y: 4}, pos.Absolute{X: 11,Y: 15})
	rectPosition3 := eraser.Rectangle(pos.Absolute{X: 12,Y: 4},pos.Absolute{X: 12,Y: 11})
	rectPosition4 := eraser.Rectangle(pos.Absolute{X: 13,Y: 4},pos.Absolute{X: 13,Y: 6})

	rect4 := eraser.Rectangle(pos.Absolute{X: 21,Y: 4}, pos.Absolute{X: 21,Y: 15})
	rectPosition5 := eraser.Rectangle(pos.Absolute{X: 22,Y: 4},pos.Absolute{X: 22,Y: 13})
	rectPosition6 := eraser.Rectangle(pos.Absolute{X: 30,Y: 4},pos.Absolute{X: 30,Y: 7})

	rect5 := eraser.Rectangle(pos.Absolute{X: 31,Y: 4}, pos.Absolute{X: 31,Y: 15})
	rectPosition7 := eraser.Rectangle(pos.Absolute{X: 39, Y: 4}, pos.Absolute{X: 39,Y: 6})
	rectPosition8 := eraser.Rectangle(pos.Absolute{X: 40, Y: 4},pos.Absolute{X: 40,Y: 11})

	rect6 := eraser.Rectangle(pos.Absolute{X: 41,Y: 4}, pos.Absolute{X: 41,Y: 15})
	rectPosition9 := eraser.Rectangle(pos.Absolute{X: 47,Y: 4},pos.Absolute{X: 47,Y: 5})
	rectPosition10 := eraser.Rectangle(pos.Absolute{X: 48,Y: 4},pos.Absolute{X: 48,Y: 14})

	
	//1층 무대 뒷부분
	rectBack1 := eraser.Rectangle(pos.Absolute{X: 10, Y: 17}, pos.Absolute{X: 10, Y: 22})
	rectBackPosition1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 19}, pos.Absolute{X: 5, Y: 22})
	rectBackPosition2 := eraser.Rectangle(pos.Absolute{X: 6,Y: 20}, pos.Absolute{X: 6, Y: 22})
	rectBackPosition3 := eraser.Rectangle(pos.Absolute{X: 7,Y: 22}, pos.Absolute{X: 7, Y: 22})

	rectBack2 := eraser.Rectangle(pos.Absolute{X: 21, Y: 17}, pos.Absolute{X: 21, Y: 22})
	rectBackPosition4 := eraser.Rectangle(pos.Absolute{X: 22,Y: 22}, pos.Absolute{X: 47, Y: 22})
	
	rectBack3 := eraser.Rectangle(pos.Absolute{X: 32, Y: 17}, pos.Absolute{X: 32, Y: 22})
	rectBack4 := eraser.Rectangle(pos.Absolute{X: 43, Y: 17}, pos.Absolute{X: 43, Y: 22})
	rectBackPosition5 := eraser.Rectangle(pos.Absolute{X: 48, Y: 19}, pos.Absolute{X: 50, Y: 22})
	rectBackPosition6 := eraser.Rectangle(pos.Absolute{X: 47, Y: 20}, pos.Absolute{X: 47, Y: 22})

	integrated := eraser.Integrated(hall1, rect1, rect2, rect3, rect4, rect5, rect6,
									rectPosition1,rectPosition2,rectPosition3,rectPosition4,rectPosition5,
									rectPosition6,rectPosition7,rectPosition8,rectPosition9,rectPosition10,
								
									rectBack1,rectBack2,rectBack3,rectBack4,
									rectBackPosition1,rectBackPosition2,rectBackPosition3, rectBackPosition4,
									rectBackPosition5,rectBackPosition6)

									
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         48,
		YSize:         19,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific := eraser.Position(pos.Absolute{X: 43, Y: 22})
	integrated2 := eraser.Integrated(specific)
	base = model.NewSeatBase(33, 22, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 23, 18),
		XSize:         13,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	rectWheel := eraser.Rectangle(pos.Absolute{X: 17, Y: 23}, pos.Absolute{X: 36,Y: 23})
	integrated3 := eraser.Integrated(rectWheel)
	base = model.NewSeatBase(11, 23, "휠체어석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 23, 18),
		XSize:         32,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
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
		Key:             "1F",
		Seats:           seats,
		XSize:           48,
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