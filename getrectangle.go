package sprint

type Coords struct {
X int
Y int

} 
type Rectangle struct {
Width int
Height int
Area int
Perimeter int

} 

func GetRectangle(min, max Coords) Rectangle {
	var rec Rectangle
	rec.Width = max.X - min.X 
	rec.Height = max.Y - min.Y 
	rec.Area = rec.Width * rec.Height
	rec.Perimeter = (rec.Width + rec.Height) * 2

	return rec
}
