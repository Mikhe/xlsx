package xlsx

type Row struct {
	Cells  []*Cell
	Hidden bool
	Sheet  *Sheet
}

func (r *Row) AddCell(width float64) *Cell {
	cell := NewCell(r)
	r.Cells = append(r.Cells, cell)
	r.Sheet.maybeAddCol(len(r.Cells), width)
	return cell
}
