package main

import "math/big"

type complexBigF struct {
	r, i big.Float
}

type complexBigR struct {
	r, i big.Rat
}

func (c1 *complexBigF) add(c2 *complexBigF) *complexBigF {
	c1.r.Add(&c1.r, &c2.r)
	c1.i.Add(&c1.i, &c2.i)
	return c1
}

func (c1 *complexBigF) sub(c2 *complexBigF) *complexBigF {
	c1.r.Sub(&c1.r, &c2.r)
	c1.i.Sub(&c1.i, &c2.i)
	return c1
}

func (c1 *complexBigF) mul(c2 *complexBigF) *complexBigF {
	c3 := complexBigF{}
	c3.r.Add(c3.r.Mul(&c1.r, &c2.r), c1.
	c1.i.Mul(&c1.i, &c2.i)
	return c1
}

func (c1 *complexBigR) add(c2 *complexBigR) *complexBigR {
	c1.r.Add(&c1.r, &c2.r)
	c1.i.Add(&c1.i, &c2.i)
	return c1
}

func (c1 *complexBigR) sub(c2 *complexBigR) *complexBigR {
	c1.r.Sub(&c1.r, &c2.r)
	c1.i.Sub(&c1.i, &c2.i)
	return c1
}

func (c1 *complexBigR) mul(c2 *complexBigR) *complexBigR {
	c1.r.Mul(&c1.r, &c2.r)
	c1.i.Mul(&c1.i, &c2.i)
	return c1
}
