package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_Host_Periph struct {
	HCFG     RHCFG
	HFIR     RHFIR
	HFNUM    RHFNUM
	_        uint32
	HPTXSTS  RHPTXSTS
	HAINT    RHAINT
	HAINTMSK RHAINTMSK
}

func (p *USB_OTG_Host_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type HCFG uint32

func (b HCFG) Field(mask HCFG) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCFG) J(v int) HCFG {
	return HCFG(bits.MakeField32(v, uint32(mask)))
}

type RHCFG struct{ mmio.U32 }

func (r *RHCFG) Bits(mask HCFG) HCFG    { return HCFG(r.U32.Bits(uint32(mask))) }
func (r *RHCFG) StoreBits(mask, b HCFG) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCFG) SetBits(mask HCFG)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCFG) ClearBits(mask HCFG)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCFG) Load() HCFG             { return HCFG(r.U32.Load()) }
func (r *RHCFG) Store(b HCFG)           { r.U32.Store(uint32(b)) }

func (r *RHCFG) AtomicStoreBits(mask, b HCFG) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCFG) AtomicSetBits(mask HCFG)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCFG) AtomicClearBits(mask HCFG)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCFG struct{ mmio.UM32 }

func (rm RMHCFG) Load() HCFG   { return HCFG(rm.UM32.Load()) }
func (rm RMHCFG) Store(b HCFG) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FSLSPCS() RMHCFG {
	return RMHCFG{mmio.UM32{&p.HCFG.U32, uint32(FSLSPCS)}}
}

func (p *USB_OTG_Host_Periph) FSLSS() RMHCFG {
	return RMHCFG{mmio.UM32{&p.HCFG.U32, uint32(FSLSS)}}
}

type HFIR uint32

func (b HFIR) Field(mask HFIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HFIR) J(v int) HFIR {
	return HFIR(bits.MakeField32(v, uint32(mask)))
}

type RHFIR struct{ mmio.U32 }

func (r *RHFIR) Bits(mask HFIR) HFIR    { return HFIR(r.U32.Bits(uint32(mask))) }
func (r *RHFIR) StoreBits(mask, b HFIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHFIR) SetBits(mask HFIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHFIR) ClearBits(mask HFIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHFIR) Load() HFIR             { return HFIR(r.U32.Load()) }
func (r *RHFIR) Store(b HFIR)           { r.U32.Store(uint32(b)) }

func (r *RHFIR) AtomicStoreBits(mask, b HFIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHFIR) AtomicSetBits(mask HFIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHFIR) AtomicClearBits(mask HFIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHFIR struct{ mmio.UM32 }

func (rm RMHFIR) Load() HFIR   { return HFIR(rm.UM32.Load()) }
func (rm RMHFIR) Store(b HFIR) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FRIVL() RMHFIR {
	return RMHFIR{mmio.UM32{&p.HFIR.U32, uint32(FRIVL)}}
}

type HFNUM uint32

func (b HFNUM) Field(mask HFNUM) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HFNUM) J(v int) HFNUM {
	return HFNUM(bits.MakeField32(v, uint32(mask)))
}

type RHFNUM struct{ mmio.U32 }

func (r *RHFNUM) Bits(mask HFNUM) HFNUM   { return HFNUM(r.U32.Bits(uint32(mask))) }
func (r *RHFNUM) StoreBits(mask, b HFNUM) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHFNUM) SetBits(mask HFNUM)      { r.U32.SetBits(uint32(mask)) }
func (r *RHFNUM) ClearBits(mask HFNUM)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHFNUM) Load() HFNUM             { return HFNUM(r.U32.Load()) }
func (r *RHFNUM) Store(b HFNUM)           { r.U32.Store(uint32(b)) }

func (r *RHFNUM) AtomicStoreBits(mask, b HFNUM) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHFNUM) AtomicSetBits(mask HFNUM)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHFNUM) AtomicClearBits(mask HFNUM)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHFNUM struct{ mmio.UM32 }

func (rm RMHFNUM) Load() HFNUM   { return HFNUM(rm.UM32.Load()) }
func (rm RMHFNUM) Store(b HFNUM) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FRNUM() RMHFNUM {
	return RMHFNUM{mmio.UM32{&p.HFNUM.U32, uint32(FRNUM)}}
}

func (p *USB_OTG_Host_Periph) FTREM() RMHFNUM {
	return RMHFNUM{mmio.UM32{&p.HFNUM.U32, uint32(FTREM)}}
}

type HPTXSTS uint32

func (b HPTXSTS) Field(mask HPTXSTS) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HPTXSTS) J(v int) HPTXSTS {
	return HPTXSTS(bits.MakeField32(v, uint32(mask)))
}

type RHPTXSTS struct{ mmio.U32 }

func (r *RHPTXSTS) Bits(mask HPTXSTS) HPTXSTS { return HPTXSTS(r.U32.Bits(uint32(mask))) }
func (r *RHPTXSTS) StoreBits(mask, b HPTXSTS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHPTXSTS) SetBits(mask HPTXSTS)      { r.U32.SetBits(uint32(mask)) }
func (r *RHPTXSTS) ClearBits(mask HPTXSTS)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHPTXSTS) Load() HPTXSTS             { return HPTXSTS(r.U32.Load()) }
func (r *RHPTXSTS) Store(b HPTXSTS)           { r.U32.Store(uint32(b)) }

func (r *RHPTXSTS) AtomicStoreBits(mask, b HPTXSTS) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHPTXSTS) AtomicSetBits(mask HPTXSTS)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHPTXSTS) AtomicClearBits(mask HPTXSTS)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHPTXSTS struct{ mmio.UM32 }

func (rm RMHPTXSTS) Load() HPTXSTS   { return HPTXSTS(rm.UM32.Load()) }
func (rm RMHPTXSTS) Store(b HPTXSTS) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) PTXFSAVL() RMHPTXSTS {
	return RMHPTXSTS{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXFSAVL)}}
}

func (p *USB_OTG_Host_Periph) PTXQSAV() RMHPTXSTS {
	return RMHPTXSTS{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXQSAV)}}
}

func (p *USB_OTG_Host_Periph) PTXQTOP() RMHPTXSTS {
	return RMHPTXSTS{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXQTOP)}}
}

type HAINT uint32

func (b HAINT) Field(mask HAINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HAINT) J(v int) HAINT {
	return HAINT(bits.MakeField32(v, uint32(mask)))
}

type RHAINT struct{ mmio.U32 }

func (r *RHAINT) Bits(mask HAINT) HAINT   { return HAINT(r.U32.Bits(uint32(mask))) }
func (r *RHAINT) StoreBits(mask, b HAINT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHAINT) SetBits(mask HAINT)      { r.U32.SetBits(uint32(mask)) }
func (r *RHAINT) ClearBits(mask HAINT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHAINT) Load() HAINT             { return HAINT(r.U32.Load()) }
func (r *RHAINT) Store(b HAINT)           { r.U32.Store(uint32(b)) }

func (r *RHAINT) AtomicStoreBits(mask, b HAINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHAINT) AtomicSetBits(mask HAINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHAINT) AtomicClearBits(mask HAINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHAINT struct{ mmio.UM32 }

func (rm RMHAINT) Load() HAINT   { return HAINT(rm.UM32.Load()) }
func (rm RMHAINT) Store(b HAINT) { rm.UM32.Store(uint32(b)) }

type HAINTMSK uint32

func (b HAINTMSK) Field(mask HAINTMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HAINTMSK) J(v int) HAINTMSK {
	return HAINTMSK(bits.MakeField32(v, uint32(mask)))
}

type RHAINTMSK struct{ mmio.U32 }

func (r *RHAINTMSK) Bits(mask HAINTMSK) HAINTMSK { return HAINTMSK(r.U32.Bits(uint32(mask))) }
func (r *RHAINTMSK) StoreBits(mask, b HAINTMSK)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHAINTMSK) SetBits(mask HAINTMSK)       { r.U32.SetBits(uint32(mask)) }
func (r *RHAINTMSK) ClearBits(mask HAINTMSK)     { r.U32.ClearBits(uint32(mask)) }
func (r *RHAINTMSK) Load() HAINTMSK              { return HAINTMSK(r.U32.Load()) }
func (r *RHAINTMSK) Store(b HAINTMSK)            { r.U32.Store(uint32(b)) }

func (r *RHAINTMSK) AtomicStoreBits(mask, b HAINTMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHAINTMSK) AtomicSetBits(mask HAINTMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHAINTMSK) AtomicClearBits(mask HAINTMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHAINTMSK struct{ mmio.UM32 }

func (rm RMHAINTMSK) Load() HAINTMSK   { return HAINTMSK(rm.UM32.Load()) }
func (rm RMHAINTMSK) Store(b HAINTMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) HAINTM() RMHAINTMSK {
	return RMHAINTMSK{mmio.UM32{&p.HAINTMSK.U32, uint32(HAINTM)}}
}
