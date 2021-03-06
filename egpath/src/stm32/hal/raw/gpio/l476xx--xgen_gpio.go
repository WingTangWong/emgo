// +build l476xx

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type GPIO_Periph struct {
	MODER   RMODER
	OTYPER  ROTYPER
	OSPEEDR ROSPEEDR
	PUPDR   RPUPDR
	IDR     RIDR
	ODR     RODR
	BSRR    RBSRR
	LCKR    RLCKR
	AFR     [2]RAFR
	BRR     RBRR
	ASCR    RASCR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

//emgo:const
var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

//emgo:const
var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

//emgo:const
var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

//emgo:const
var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

//emgo:const
var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

//emgo:const
var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

//emgo:const
var GPIOH = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE)))

type MODER uint32

func (b MODER) Field(mask MODER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MODER) J(v int) MODER {
	return MODER(bits.MakeField32(v, uint32(mask)))
}

type RMODER struct{ mmio.U32 }

func (r *RMODER) Bits(mask MODER) MODER   { return MODER(r.U32.Bits(uint32(mask))) }
func (r *RMODER) StoreBits(mask, b MODER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) SetBits(mask MODER)      { r.U32.SetBits(uint32(mask)) }
func (r *RMODER) ClearBits(mask MODER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMODER) Load() MODER             { return MODER(r.U32.Load()) }
func (r *RMODER) Store(b MODER)           { r.U32.Store(uint32(b)) }

func (r *RMODER) AtomicStoreBits(mask, b MODER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) AtomicSetBits(mask MODER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMODER) AtomicClearBits(mask MODER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMODER struct{ mmio.UM32 }

func (rm RMMODER) Load() MODER   { return MODER(rm.UM32.Load()) }
func (rm RMMODER) Store(b MODER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODE0() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE0)}}
}

func (p *GPIO_Periph) MODE1() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE1)}}
}

func (p *GPIO_Periph) MODE2() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE2)}}
}

func (p *GPIO_Periph) MODE3() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE3)}}
}

func (p *GPIO_Periph) MODE4() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE4)}}
}

func (p *GPIO_Periph) MODE5() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE5)}}
}

func (p *GPIO_Periph) MODE6() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE6)}}
}

func (p *GPIO_Periph) MODE7() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE7)}}
}

func (p *GPIO_Periph) MODE8() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE8)}}
}

func (p *GPIO_Periph) MODE9() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE9)}}
}

func (p *GPIO_Periph) MODE10() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE10)}}
}

func (p *GPIO_Periph) MODE11() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE11)}}
}

func (p *GPIO_Periph) MODE12() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE12)}}
}

func (p *GPIO_Periph) MODE13() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE13)}}
}

func (p *GPIO_Periph) MODE14() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE14)}}
}

func (p *GPIO_Periph) MODE15() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODE15)}}
}

type OTYPER uint32

func (b OTYPER) Field(mask OTYPER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTYPER) J(v int) OTYPER {
	return OTYPER(bits.MakeField32(v, uint32(mask)))
}

type ROTYPER struct{ mmio.U32 }

func (r *ROTYPER) Bits(mask OTYPER) OTYPER  { return OTYPER(r.U32.Bits(uint32(mask))) }
func (r *ROTYPER) StoreBits(mask, b OTYPER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) SetBits(mask OTYPER)      { r.U32.SetBits(uint32(mask)) }
func (r *ROTYPER) ClearBits(mask OTYPER)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROTYPER) Load() OTYPER             { return OTYPER(r.U32.Load()) }
func (r *ROTYPER) Store(b OTYPER)           { r.U32.Store(uint32(b)) }

func (r *ROTYPER) AtomicStoreBits(mask, b OTYPER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) AtomicSetBits(mask OTYPER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROTYPER) AtomicClearBits(mask OTYPER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOTYPER struct{ mmio.UM32 }

func (rm RMOTYPER) Load() OTYPER   { return OTYPER(rm.UM32.Load()) }
func (rm RMOTYPER) Store(b OTYPER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT0() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT0)}}
}

func (p *GPIO_Periph) OT1() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT1)}}
}

func (p *GPIO_Periph) OT2() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT2)}}
}

func (p *GPIO_Periph) OT3() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT3)}}
}

func (p *GPIO_Periph) OT4() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT4)}}
}

func (p *GPIO_Periph) OT5() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT5)}}
}

func (p *GPIO_Periph) OT6() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT6)}}
}

func (p *GPIO_Periph) OT7() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT7)}}
}

func (p *GPIO_Periph) OT8() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT8)}}
}

func (p *GPIO_Periph) OT9() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT9)}}
}

func (p *GPIO_Periph) OT10() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT10)}}
}

func (p *GPIO_Periph) OT11() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT11)}}
}

func (p *GPIO_Periph) OT12() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT12)}}
}

func (p *GPIO_Periph) OT13() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT13)}}
}

func (p *GPIO_Periph) OT14() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT14)}}
}

func (p *GPIO_Periph) OT15() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT15)}}
}

type OSPEEDR uint32

func (b OSPEEDR) Field(mask OSPEEDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OSPEEDR) J(v int) OSPEEDR {
	return OSPEEDR(bits.MakeField32(v, uint32(mask)))
}

type ROSPEEDR struct{ mmio.U32 }

func (r *ROSPEEDR) Bits(mask OSPEEDR) OSPEEDR { return OSPEEDR(r.U32.Bits(uint32(mask))) }
func (r *ROSPEEDR) StoreBits(mask, b OSPEEDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) SetBits(mask OSPEEDR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROSPEEDR) ClearBits(mask OSPEEDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROSPEEDR) Load() OSPEEDR             { return OSPEEDR(r.U32.Load()) }
func (r *ROSPEEDR) Store(b OSPEEDR)           { r.U32.Store(uint32(b)) }

func (r *ROSPEEDR) AtomicStoreBits(mask, b OSPEEDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) AtomicSetBits(mask OSPEEDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROSPEEDR) AtomicClearBits(mask OSPEEDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOSPEEDR struct{ mmio.UM32 }

func (rm RMOSPEEDR) Load() OSPEEDR   { return OSPEEDR(rm.UM32.Load()) }
func (rm RMOSPEEDR) Store(b OSPEEDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OSPEED0() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED0)}}
}

func (p *GPIO_Periph) OSPEED1() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED1)}}
}

func (p *GPIO_Periph) OSPEED2() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED2)}}
}

func (p *GPIO_Periph) OSPEED3() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED3)}}
}

func (p *GPIO_Periph) OSPEED4() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED4)}}
}

func (p *GPIO_Periph) OSPEED5() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED5)}}
}

func (p *GPIO_Periph) OSPEED6() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED6)}}
}

func (p *GPIO_Periph) OSPEED7() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED7)}}
}

func (p *GPIO_Periph) OSPEED8() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED8)}}
}

func (p *GPIO_Periph) OSPEED9() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED9)}}
}

func (p *GPIO_Periph) OSPEED10() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED10)}}
}

func (p *GPIO_Periph) OSPEED11() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED11)}}
}

func (p *GPIO_Periph) OSPEED12() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED12)}}
}

func (p *GPIO_Periph) OSPEED13() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED13)}}
}

func (p *GPIO_Periph) OSPEED14() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED14)}}
}

func (p *GPIO_Periph) OSPEED15() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED15)}}
}

type PUPDR uint32

func (b PUPDR) Field(mask PUPDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PUPDR) J(v int) PUPDR {
	return PUPDR(bits.MakeField32(v, uint32(mask)))
}

type RPUPDR struct{ mmio.U32 }

func (r *RPUPDR) Bits(mask PUPDR) PUPDR   { return PUPDR(r.U32.Bits(uint32(mask))) }
func (r *RPUPDR) StoreBits(mask, b PUPDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) SetBits(mask PUPDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPUPDR) ClearBits(mask PUPDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPUPDR) Load() PUPDR             { return PUPDR(r.U32.Load()) }
func (r *RPUPDR) Store(b PUPDR)           { r.U32.Store(uint32(b)) }

func (r *RPUPDR) AtomicStoreBits(mask, b PUPDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) AtomicSetBits(mask PUPDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPUPDR) AtomicClearBits(mask PUPDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPUPDR struct{ mmio.UM32 }

func (rm RMPUPDR) Load() PUPDR   { return PUPDR(rm.UM32.Load()) }
func (rm RMPUPDR) Store(b PUPDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPD0() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD0)}}
}

func (p *GPIO_Periph) PUPD1() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD1)}}
}

func (p *GPIO_Periph) PUPD2() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD2)}}
}

func (p *GPIO_Periph) PUPD3() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD3)}}
}

func (p *GPIO_Periph) PUPD4() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD4)}}
}

func (p *GPIO_Periph) PUPD5() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD5)}}
}

func (p *GPIO_Periph) PUPD6() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD6)}}
}

func (p *GPIO_Periph) PUPD7() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD7)}}
}

func (p *GPIO_Periph) PUPD8() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD8)}}
}

func (p *GPIO_Periph) PUPD9() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD9)}}
}

func (p *GPIO_Periph) PUPD10() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD10)}}
}

func (p *GPIO_Periph) PUPD11() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD11)}}
}

func (p *GPIO_Periph) PUPD12() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD12)}}
}

func (p *GPIO_Periph) PUPD13() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD13)}}
}

func (p *GPIO_Periph) PUPD14() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD14)}}
}

func (p *GPIO_Periph) PUPD15() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPD15)}}
}

type IDR uint32

func (b IDR) Field(mask IDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR) J(v int) IDR {
	return IDR(bits.MakeField32(v, uint32(mask)))
}

type RIDR struct{ mmio.U32 }

func (r *RIDR) Bits(mask IDR) IDR     { return IDR(r.U32.Bits(uint32(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U32.Load()) }
func (r *RIDR) Store(b IDR)           { r.U32.Store(uint32(b)) }

func (r *RIDR) AtomicStoreBits(mask, b IDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) AtomicSetBits(mask IDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIDR) AtomicClearBits(mask IDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIDR struct{ mmio.UM32 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM32.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ID0() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID0)}}
}

func (p *GPIO_Periph) ID1() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID1)}}
}

func (p *GPIO_Periph) ID2() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID2)}}
}

func (p *GPIO_Periph) ID3() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID3)}}
}

func (p *GPIO_Periph) ID4() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID4)}}
}

func (p *GPIO_Periph) ID5() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID5)}}
}

func (p *GPIO_Periph) ID6() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID6)}}
}

func (p *GPIO_Periph) ID7() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID7)}}
}

func (p *GPIO_Periph) ID8() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID8)}}
}

func (p *GPIO_Periph) ID9() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID9)}}
}

func (p *GPIO_Periph) ID10() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID10)}}
}

func (p *GPIO_Periph) ID11() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID11)}}
}

func (p *GPIO_Periph) ID12() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID12)}}
}

func (p *GPIO_Periph) ID13() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID13)}}
}

func (p *GPIO_Periph) ID14() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID14)}}
}

func (p *GPIO_Periph) ID15() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(ID15)}}
}

type ODR uint32

func (b ODR) Field(mask ODR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR) J(v int) ODR {
	return ODR(bits.MakeField32(v, uint32(mask)))
}

type RODR struct{ mmio.U32 }

func (r *RODR) Bits(mask ODR) ODR     { return ODR(r.U32.Bits(uint32(mask))) }
func (r *RODR) StoreBits(mask, b ODR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RODR) SetBits(mask ODR)      { r.U32.SetBits(uint32(mask)) }
func (r *RODR) ClearBits(mask ODR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RODR) Load() ODR             { return ODR(r.U32.Load()) }
func (r *RODR) Store(b ODR)           { r.U32.Store(uint32(b)) }

func (r *RODR) AtomicStoreBits(mask, b ODR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RODR) AtomicSetBits(mask ODR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RODR) AtomicClearBits(mask ODR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMODR struct{ mmio.UM32 }

func (rm RMODR) Load() ODR   { return ODR(rm.UM32.Load()) }
func (rm RMODR) Store(b ODR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OD0() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD0)}}
}

func (p *GPIO_Periph) OD1() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD1)}}
}

func (p *GPIO_Periph) OD2() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD2)}}
}

func (p *GPIO_Periph) OD3() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD3)}}
}

func (p *GPIO_Periph) OD4() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD4)}}
}

func (p *GPIO_Periph) OD5() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD5)}}
}

func (p *GPIO_Periph) OD6() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD6)}}
}

func (p *GPIO_Periph) OD7() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD7)}}
}

func (p *GPIO_Periph) OD8() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD8)}}
}

func (p *GPIO_Periph) OD9() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD9)}}
}

func (p *GPIO_Periph) OD10() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD10)}}
}

func (p *GPIO_Periph) OD11() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD11)}}
}

func (p *GPIO_Periph) OD12() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD12)}}
}

func (p *GPIO_Periph) OD13() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD13)}}
}

func (p *GPIO_Periph) OD14() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD14)}}
}

func (p *GPIO_Periph) OD15() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(OD15)}}
}

type BSRR uint32

func (b BSRR) Field(mask BSRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRR) J(v int) BSRR {
	return BSRR(bits.MakeField32(v, uint32(mask)))
}

type RBSRR struct{ mmio.U32 }

func (r *RBSRR) Bits(mask BSRR) BSRR    { return BSRR(r.U32.Bits(uint32(mask))) }
func (r *RBSRR) StoreBits(mask, b BSRR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBSRR) SetBits(mask BSRR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBSRR) ClearBits(mask BSRR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBSRR) Load() BSRR             { return BSRR(r.U32.Load()) }
func (r *RBSRR) Store(b BSRR)           { r.U32.Store(uint32(b)) }

func (r *RBSRR) AtomicStoreBits(mask, b BSRR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBSRR) AtomicSetBits(mask BSRR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBSRR) AtomicClearBits(mask BSRR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBSRR struct{ mmio.UM32 }

func (rm RMBSRR) Load() BSRR   { return BSRR(rm.UM32.Load()) }
func (rm RMBSRR) Store(b BSRR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BS0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS0)}}
}

func (p *GPIO_Periph) BS1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS1)}}
}

func (p *GPIO_Periph) BS2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS2)}}
}

func (p *GPIO_Periph) BS3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS3)}}
}

func (p *GPIO_Periph) BS4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS4)}}
}

func (p *GPIO_Periph) BS5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS5)}}
}

func (p *GPIO_Periph) BS6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS6)}}
}

func (p *GPIO_Periph) BS7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS7)}}
}

func (p *GPIO_Periph) BS8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS8)}}
}

func (p *GPIO_Periph) BS9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS9)}}
}

func (p *GPIO_Periph) BS10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS10)}}
}

func (p *GPIO_Periph) BS11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS11)}}
}

func (p *GPIO_Periph) BS12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS12)}}
}

func (p *GPIO_Periph) BS13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS13)}}
}

func (p *GPIO_Periph) BS14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS14)}}
}

func (p *GPIO_Periph) BS15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS15)}}
}

func (p *GPIO_Periph) BR0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR15)}}
}

type LCKR uint32

func (b LCKR) Field(mask LCKR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR) J(v int) LCKR {
	return LCKR(bits.MakeField32(v, uint32(mask)))
}

type RLCKR struct{ mmio.U32 }

func (r *RLCKR) Bits(mask LCKR) LCKR    { return LCKR(r.U32.Bits(uint32(mask))) }
func (r *RLCKR) StoreBits(mask, b LCKR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) SetBits(mask LCKR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLCKR) ClearBits(mask LCKR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLCKR) Load() LCKR             { return LCKR(r.U32.Load()) }
func (r *RLCKR) Store(b LCKR)           { r.U32.Store(uint32(b)) }

func (r *RLCKR) AtomicStoreBits(mask, b LCKR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) AtomicSetBits(mask LCKR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLCKR) AtomicClearBits(mask LCKR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLCKR struct{ mmio.UM32 }

func (rm RMLCKR) Load() LCKR   { return LCKR(rm.UM32.Load()) }
func (rm RMLCKR) Store(b LCKR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) LCK0() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *GPIO_Periph) LCK1() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *GPIO_Periph) LCK2() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *GPIO_Periph) LCK3() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *GPIO_Periph) LCK4() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *GPIO_Periph) LCK5() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *GPIO_Periph) LCK6() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *GPIO_Periph) LCK7() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *GPIO_Periph) LCK8() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *GPIO_Periph) LCK9() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *GPIO_Periph) LCK10() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *GPIO_Periph) LCK11() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *GPIO_Periph) LCK12() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *GPIO_Periph) LCK13() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *GPIO_Periph) LCK14() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *GPIO_Periph) LCK15() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *GPIO_Periph) LCKK() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}

type AFR uint32

func (b AFR) Field(mask AFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFR) J(v int) AFR {
	return AFR(bits.MakeField32(v, uint32(mask)))
}

type RAFR struct{ mmio.U32 }

func (r *RAFR) Bits(mask AFR) AFR     { return AFR(r.U32.Bits(uint32(mask))) }
func (r *RAFR) StoreBits(mask, b AFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) SetBits(mask AFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAFR) ClearBits(mask AFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAFR) Load() AFR             { return AFR(r.U32.Load()) }
func (r *RAFR) Store(b AFR)           { r.U32.Store(uint32(b)) }

func (r *RAFR) AtomicStoreBits(mask, b AFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) AtomicSetBits(mask AFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAFR) AtomicClearBits(mask AFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAFR struct{ mmio.UM32 }

func (rm RMAFR) Load() AFR   { return AFR(rm.UM32.Load()) }
func (rm RMAFR) Store(b AFR) { rm.UM32.Store(uint32(b)) }

type BRR uint32

func (b BRR) Field(mask BRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR) J(v int) BRR {
	return BRR(bits.MakeField32(v, uint32(mask)))
}

type RBRR struct{ mmio.U32 }

func (r *RBRR) Bits(mask BRR) BRR     { return BRR(r.U32.Bits(uint32(mask))) }
func (r *RBRR) StoreBits(mask, b BRR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) SetBits(mask BRR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBRR) ClearBits(mask BRR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBRR) Load() BRR             { return BRR(r.U32.Load()) }
func (r *RBRR) Store(b BRR)           { r.U32.Store(uint32(b)) }

func (r *RBRR) AtomicStoreBits(mask, b BRR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) AtomicSetBits(mask BRR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBRR) AtomicClearBits(mask BRR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBRR struct{ mmio.UM32 }

func (rm RMBRR) Load() BRR   { return BRR(rm.UM32.Load()) }
func (rm RMBRR) Store(b BRR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BR0() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR15)}}
}

type ASCR uint32

func (b ASCR) Field(mask ASCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASCR) J(v int) ASCR {
	return ASCR(bits.MakeField32(v, uint32(mask)))
}

type RASCR struct{ mmio.U32 }

func (r *RASCR) Bits(mask ASCR) ASCR    { return ASCR(r.U32.Bits(uint32(mask))) }
func (r *RASCR) StoreBits(mask, b ASCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RASCR) SetBits(mask ASCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RASCR) ClearBits(mask ASCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RASCR) Load() ASCR             { return ASCR(r.U32.Load()) }
func (r *RASCR) Store(b ASCR)           { r.U32.Store(uint32(b)) }

func (r *RASCR) AtomicStoreBits(mask, b ASCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RASCR) AtomicSetBits(mask ASCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RASCR) AtomicClearBits(mask ASCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMASCR struct{ mmio.UM32 }

func (rm RMASCR) Load() ASCR   { return ASCR(rm.UM32.Load()) }
func (rm RMASCR) Store(b ASCR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ASC0() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC0)}}
}

func (p *GPIO_Periph) ASC1() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC1)}}
}

func (p *GPIO_Periph) ASC2() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC2)}}
}

func (p *GPIO_Periph) ASC3() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC3)}}
}

func (p *GPIO_Periph) ASC4() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC4)}}
}

func (p *GPIO_Periph) ASC5() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC5)}}
}

func (p *GPIO_Periph) ASC6() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC6)}}
}

func (p *GPIO_Periph) ASC7() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC7)}}
}

func (p *GPIO_Periph) ASC8() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC8)}}
}

func (p *GPIO_Periph) ASC9() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC9)}}
}

func (p *GPIO_Periph) ASC10() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC10)}}
}

func (p *GPIO_Periph) ASC11() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC11)}}
}

func (p *GPIO_Periph) ASC12() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC12)}}
}

func (p *GPIO_Periph) ASC13() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC13)}}
}

func (p *GPIO_Periph) ASC14() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC14)}}
}

func (p *GPIO_Periph) ASC15() RMASCR {
	return RMASCR{mmio.UM32{&p.ASCR.U32, uint32(ASC15)}}
}
