//go:build gogensed

//go:generate gogensed gen/tint TYPE=int
package template

type Quad struct {
	D    [4]TYPE
	X    *TYPE
	Y    *TYPE
	U    *TYPE
	V    *TYPE
	XY   []TYPE
	XYZ  []TYPE
	XYUV []TYPE
}

func (q *Quad) Set(data ...TYPE) *Quad {
	for i, d := range data {
		q.D[i] = d
	}
	q.X = &q.D[0]
	q.Y = &q.D[1]
	q.U = &q.D[2]
	q.V = &q.D[3]
	q.XY = q.D[0:2]
	q.XYZ = q.D[0:3]
	q.XYUV = q.D[0:4]
	return q
}

func (q *Quad) QUAD() (TYPE, TYPE, TYPE, TYPE) {
	return q.D[0], q.D[1], q.D[2], q.D[3]
}

func (q *Quad) RGBA() (TYPE, TYPE, TYPE, TYPE) {
	return q.QUAD()
}
