// +build l1xx_md

package sdio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type SDIO_Periph struct {
	POWER   RPOWER
	CLKCR   RCLKCR
	ARG     RARG
	CMD     RCMD
	RESPCMD RRESPCMD
	RESP    [4]RRESP
	DTIMER  RDTIMER
	DLEN    RDLEN
	DCTRL   RDCTRL
	DCOUNT  RDCOUNT
	STA     RSTA
	ICR     RICR
	MASK    RMASK
	_       [2]uint32
	FIFOCNT RFIFOCNT
	_       [13]uint32
	FIFO    RFIFO
}

func (p *SDIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SDIO = (*SDIO_Periph)(unsafe.Pointer(uintptr(mmap.SDIO_BASE)))

type POWER uint32

func (b POWER) Field(mask POWER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask POWER) J(v int) POWER {
	return POWER(bits.MakeField32(v, uint32(mask)))
}

type RPOWER struct{ mmio.U32 }

func (r *RPOWER) Bits(mask POWER) POWER   { return POWER(r.U32.Bits(uint32(mask))) }
func (r *RPOWER) StoreBits(mask, b POWER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPOWER) SetBits(mask POWER)      { r.U32.SetBits(uint32(mask)) }
func (r *RPOWER) ClearBits(mask POWER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPOWER) Load() POWER             { return POWER(r.U32.Load()) }
func (r *RPOWER) Store(b POWER)           { r.U32.Store(uint32(b)) }

func (r *RPOWER) AtomicStoreBits(mask, b POWER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPOWER) AtomicSetBits(mask POWER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPOWER) AtomicClearBits(mask POWER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPOWER struct{ mmio.UM32 }

func (rm RMPOWER) Load() POWER   { return POWER(rm.UM32.Load()) }
func (rm RMPOWER) Store(b POWER) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) PWRCTRL() RMPOWER {
	return RMPOWER{mmio.UM32{&p.POWER.U32, uint32(PWRCTRL)}}
}

type CLKCR uint32

func (b CLKCR) Field(mask CLKCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLKCR) J(v int) CLKCR {
	return CLKCR(bits.MakeField32(v, uint32(mask)))
}

type RCLKCR struct{ mmio.U32 }

func (r *RCLKCR) Bits(mask CLKCR) CLKCR   { return CLKCR(r.U32.Bits(uint32(mask))) }
func (r *RCLKCR) StoreBits(mask, b CLKCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCLKCR) SetBits(mask CLKCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCLKCR) ClearBits(mask CLKCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCLKCR) Load() CLKCR             { return CLKCR(r.U32.Load()) }
func (r *RCLKCR) Store(b CLKCR)           { r.U32.Store(uint32(b)) }

func (r *RCLKCR) AtomicStoreBits(mask, b CLKCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCLKCR) AtomicSetBits(mask CLKCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCLKCR) AtomicClearBits(mask CLKCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCLKCR struct{ mmio.UM32 }

func (rm RMCLKCR) Load() CLKCR   { return CLKCR(rm.UM32.Load()) }
func (rm RMCLKCR) Store(b CLKCR) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CLKDIV() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(CLKDIV)}}
}

func (p *SDIO_Periph) CLKEN() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(CLKEN)}}
}

func (p *SDIO_Periph) PWRSAV() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(PWRSAV)}}
}

func (p *SDIO_Periph) BYPASS() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(BYPASS)}}
}

func (p *SDIO_Periph) WIDBUS() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(WIDBUS)}}
}

func (p *SDIO_Periph) NEGEDGE() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(NEGEDGE)}}
}

func (p *SDIO_Periph) HWFC_EN() RMCLKCR {
	return RMCLKCR{mmio.UM32{&p.CLKCR.U32, uint32(HWFC_EN)}}
}

type ARG uint32

func (b ARG) Field(mask ARG) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARG) J(v int) ARG {
	return ARG(bits.MakeField32(v, uint32(mask)))
}

type RARG struct{ mmio.U32 }

func (r *RARG) Bits(mask ARG) ARG     { return ARG(r.U32.Bits(uint32(mask))) }
func (r *RARG) StoreBits(mask, b ARG) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RARG) SetBits(mask ARG)      { r.U32.SetBits(uint32(mask)) }
func (r *RARG) ClearBits(mask ARG)    { r.U32.ClearBits(uint32(mask)) }
func (r *RARG) Load() ARG             { return ARG(r.U32.Load()) }
func (r *RARG) Store(b ARG)           { r.U32.Store(uint32(b)) }

func (r *RARG) AtomicStoreBits(mask, b ARG) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RARG) AtomicSetBits(mask ARG)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RARG) AtomicClearBits(mask ARG)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMARG struct{ mmio.UM32 }

func (rm RMARG) Load() ARG   { return ARG(rm.UM32.Load()) }
func (rm RMARG) Store(b ARG) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CMDARG() RMARG {
	return RMARG{mmio.UM32{&p.ARG.U32, uint32(CMDARG)}}
}

type CMD uint32

func (b CMD) Field(mask CMD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMD) J(v int) CMD {
	return CMD(bits.MakeField32(v, uint32(mask)))
}

type RCMD struct{ mmio.U32 }

func (r *RCMD) Bits(mask CMD) CMD     { return CMD(r.U32.Bits(uint32(mask))) }
func (r *RCMD) StoreBits(mask, b CMD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCMD) SetBits(mask CMD)      { r.U32.SetBits(uint32(mask)) }
func (r *RCMD) ClearBits(mask CMD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCMD) Load() CMD             { return CMD(r.U32.Load()) }
func (r *RCMD) Store(b CMD)           { r.U32.Store(uint32(b)) }

func (r *RCMD) AtomicStoreBits(mask, b CMD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCMD) AtomicSetBits(mask CMD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCMD) AtomicClearBits(mask CMD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCMD struct{ mmio.UM32 }

func (rm RMCMD) Load() CMD   { return CMD(rm.UM32.Load()) }
func (rm RMCMD) Store(b CMD) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CMDINDEX() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CMDINDEX)}}
}

func (p *SDIO_Periph) WAITRESP() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITRESP)}}
}

func (p *SDIO_Periph) WAITINT() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITINT)}}
}

func (p *SDIO_Periph) WAITPEND() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(WAITPEND)}}
}

func (p *SDIO_Periph) CPSMEN() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CPSMEN)}}
}

func (p *SDIO_Periph) SDIOSUSPEND() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(SDIOSUSPEND)}}
}

func (p *SDIO_Periph) ENCMDCOMPL() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(ENCMDCOMPL)}}
}

func (p *SDIO_Periph) NIEN() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(NIEN)}}
}

func (p *SDIO_Periph) CEATACMD() RMCMD {
	return RMCMD{mmio.UM32{&p.CMD.U32, uint32(CEATACMD)}}
}

type RESPCMD uint32

func (b RESPCMD) Field(mask RESPCMD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESPCMD) J(v int) RESPCMD {
	return RESPCMD(bits.MakeField32(v, uint32(mask)))
}

type RRESPCMD struct{ mmio.U32 }

func (r *RRESPCMD) Bits(mask RESPCMD) RESPCMD { return RESPCMD(r.U32.Bits(uint32(mask))) }
func (r *RRESPCMD) StoreBits(mask, b RESPCMD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESPCMD) SetBits(mask RESPCMD)      { r.U32.SetBits(uint32(mask)) }
func (r *RRESPCMD) ClearBits(mask RESPCMD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRESPCMD) Load() RESPCMD             { return RESPCMD(r.U32.Load()) }
func (r *RRESPCMD) Store(b RESPCMD)           { r.U32.Store(uint32(b)) }

func (r *RRESPCMD) AtomicStoreBits(mask, b RESPCMD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRESPCMD) AtomicSetBits(mask RESPCMD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESPCMD) AtomicClearBits(mask RESPCMD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESPCMD struct{ mmio.UM32 }

func (rm RMRESPCMD) Load() RESPCMD   { return RESPCMD(rm.UM32.Load()) }
func (rm RMRESPCMD) Store(b RESPCMD) { rm.UM32.Store(uint32(b)) }

type RESP uint32

func (b RESP) Field(mask RESP) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESP) J(v int) RESP {
	return RESP(bits.MakeField32(v, uint32(mask)))
}

type RRESP struct{ mmio.U32 }

func (r *RRESP) Bits(mask RESP) RESP    { return RESP(r.U32.Bits(uint32(mask))) }
func (r *RRESP) StoreBits(mask, b RESP) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESP) SetBits(mask RESP)      { r.U32.SetBits(uint32(mask)) }
func (r *RRESP) ClearBits(mask RESP)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRESP) Load() RESP             { return RESP(r.U32.Load()) }
func (r *RRESP) Store(b RESP)           { r.U32.Store(uint32(b)) }

func (r *RRESP) AtomicStoreBits(mask, b RESP) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRESP) AtomicSetBits(mask RESP)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESP) AtomicClearBits(mask RESP)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESP struct{ mmio.UM32 }

func (rm RMRESP) Load() RESP   { return RESP(rm.UM32.Load()) }
func (rm RMRESP) Store(b RESP) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CARDSTATUS1(n int) RMRESP {
	return RMRESP{mmio.UM32{&p.RESP[n].U32, uint32(CARDSTATUS1)}}
}

type DTIMER uint32

func (b DTIMER) Field(mask DTIMER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTIMER) J(v int) DTIMER {
	return DTIMER(bits.MakeField32(v, uint32(mask)))
}

type RDTIMER struct{ mmio.U32 }

func (r *RDTIMER) Bits(mask DTIMER) DTIMER  { return DTIMER(r.U32.Bits(uint32(mask))) }
func (r *RDTIMER) StoreBits(mask, b DTIMER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTIMER) SetBits(mask DTIMER)      { r.U32.SetBits(uint32(mask)) }
func (r *RDTIMER) ClearBits(mask DTIMER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDTIMER) Load() DTIMER             { return DTIMER(r.U32.Load()) }
func (r *RDTIMER) Store(b DTIMER)           { r.U32.Store(uint32(b)) }

func (r *RDTIMER) AtomicStoreBits(mask, b DTIMER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDTIMER) AtomicSetBits(mask DTIMER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDTIMER) AtomicClearBits(mask DTIMER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDTIMER struct{ mmio.UM32 }

func (rm RMDTIMER) Load() DTIMER   { return DTIMER(rm.UM32.Load()) }
func (rm RMDTIMER) Store(b DTIMER) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DATATIME() RMDTIMER {
	return RMDTIMER{mmio.UM32{&p.DTIMER.U32, uint32(DATATIME)}}
}

type DLEN uint32

func (b DLEN) Field(mask DLEN) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DLEN) J(v int) DLEN {
	return DLEN(bits.MakeField32(v, uint32(mask)))
}

type RDLEN struct{ mmio.U32 }

func (r *RDLEN) Bits(mask DLEN) DLEN    { return DLEN(r.U32.Bits(uint32(mask))) }
func (r *RDLEN) StoreBits(mask, b DLEN) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDLEN) SetBits(mask DLEN)      { r.U32.SetBits(uint32(mask)) }
func (r *RDLEN) ClearBits(mask DLEN)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDLEN) Load() DLEN             { return DLEN(r.U32.Load()) }
func (r *RDLEN) Store(b DLEN)           { r.U32.Store(uint32(b)) }

func (r *RDLEN) AtomicStoreBits(mask, b DLEN) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDLEN) AtomicSetBits(mask DLEN)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDLEN) AtomicClearBits(mask DLEN)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDLEN struct{ mmio.UM32 }

func (rm RMDLEN) Load() DLEN   { return DLEN(rm.UM32.Load()) }
func (rm RMDLEN) Store(b DLEN) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DATALENGTH() RMDLEN {
	return RMDLEN{mmio.UM32{&p.DLEN.U32, uint32(DATALENGTH)}}
}

type DCTRL uint32

func (b DCTRL) Field(mask DCTRL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCTRL) J(v int) DCTRL {
	return DCTRL(bits.MakeField32(v, uint32(mask)))
}

type RDCTRL struct{ mmio.U32 }

func (r *RDCTRL) Bits(mask DCTRL) DCTRL   { return DCTRL(r.U32.Bits(uint32(mask))) }
func (r *RDCTRL) StoreBits(mask, b DCTRL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCTRL) SetBits(mask DCTRL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCTRL) ClearBits(mask DCTRL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCTRL) Load() DCTRL             { return DCTRL(r.U32.Load()) }
func (r *RDCTRL) Store(b DCTRL)           { r.U32.Store(uint32(b)) }

func (r *RDCTRL) AtomicStoreBits(mask, b DCTRL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCTRL) AtomicSetBits(mask DCTRL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCTRL) AtomicClearBits(mask DCTRL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCTRL struct{ mmio.UM32 }

func (rm RMDCTRL) Load() DCTRL   { return DCTRL(rm.UM32.Load()) }
func (rm RMDCTRL) Store(b DCTRL) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DTEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTEN)}}
}

func (p *SDIO_Periph) DTDIR() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTDIR)}}
}

func (p *SDIO_Periph) DTMODE() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DTMODE)}}
}

func (p *SDIO_Periph) DMAEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DMAEN)}}
}

func (p *SDIO_Periph) DBLOCKSIZE() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(DBLOCKSIZE)}}
}

func (p *SDIO_Periph) RWSTART() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWSTART)}}
}

func (p *SDIO_Periph) RWSTOP() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWSTOP)}}
}

func (p *SDIO_Periph) RWMOD() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(RWMOD)}}
}

func (p *SDIO_Periph) SDIOEN() RMDCTRL {
	return RMDCTRL{mmio.UM32{&p.DCTRL.U32, uint32(SDIOEN)}}
}

type DCOUNT uint32

func (b DCOUNT) Field(mask DCOUNT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCOUNT) J(v int) DCOUNT {
	return DCOUNT(bits.MakeField32(v, uint32(mask)))
}

type RDCOUNT struct{ mmio.U32 }

func (r *RDCOUNT) Bits(mask DCOUNT) DCOUNT  { return DCOUNT(r.U32.Bits(uint32(mask))) }
func (r *RDCOUNT) StoreBits(mask, b DCOUNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCOUNT) SetBits(mask DCOUNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCOUNT) ClearBits(mask DCOUNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCOUNT) Load() DCOUNT             { return DCOUNT(r.U32.Load()) }
func (r *RDCOUNT) Store(b DCOUNT)           { r.U32.Store(uint32(b)) }

func (r *RDCOUNT) AtomicStoreBits(mask, b DCOUNT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDCOUNT) AtomicSetBits(mask DCOUNT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDCOUNT) AtomicClearBits(mask DCOUNT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDCOUNT struct{ mmio.UM32 }

func (rm RMDCOUNT) Load() DCOUNT   { return DCOUNT(rm.UM32.Load()) }
func (rm RMDCOUNT) Store(b DCOUNT) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) DATACOUNT() RMDCOUNT {
	return RMDCOUNT{mmio.UM32{&p.DCOUNT.U32, uint32(DATACOUNT)}}
}

type STA uint32

func (b STA) Field(mask STA) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask STA) J(v int) STA {
	return STA(bits.MakeField32(v, uint32(mask)))
}

type RSTA struct{ mmio.U32 }

func (r *RSTA) Bits(mask STA) STA     { return STA(r.U32.Bits(uint32(mask))) }
func (r *RSTA) StoreBits(mask, b STA) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTA) SetBits(mask STA)      { r.U32.SetBits(uint32(mask)) }
func (r *RSTA) ClearBits(mask STA)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSTA) Load() STA             { return STA(r.U32.Load()) }
func (r *RSTA) Store(b STA)           { r.U32.Store(uint32(b)) }

func (r *RSTA) AtomicStoreBits(mask, b STA) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSTA) AtomicSetBits(mask STA)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSTA) AtomicClearBits(mask STA)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSTA struct{ mmio.UM32 }

func (rm RMSTA) Load() STA   { return STA(rm.UM32.Load()) }
func (rm RMSTA) Store(b STA) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CCRCFAIL() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CCRCFAIL)}}
}

func (p *SDIO_Periph) DCRCFAIL() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(DCRCFAIL)}}
}

func (p *SDIO_Periph) CTIMEOUT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CTIMEOUT)}}
}

func (p *SDIO_Periph) DTIMEOUT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(DTIMEOUT)}}
}

func (p *SDIO_Periph) TXUNDERR() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXUNDERR)}}
}

func (p *SDIO_Periph) RXOVERR() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXOVERR)}}
}

func (p *SDIO_Periph) CMDREND() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CMDREND)}}
}

func (p *SDIO_Periph) CMDSENT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CMDSENT)}}
}

func (p *SDIO_Periph) DATAEND() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(DATAEND)}}
}

func (p *SDIO_Periph) STBITERR() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(STBITERR)}}
}

func (p *SDIO_Periph) DBCKEND() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(DBCKEND)}}
}

func (p *SDIO_Periph) CMDACT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CMDACT)}}
}

func (p *SDIO_Periph) TXACT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXACT)}}
}

func (p *SDIO_Periph) RXACT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXACT)}}
}

func (p *SDIO_Periph) TXFIFOHE() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXFIFOHE)}}
}

func (p *SDIO_Periph) RXFIFOHF() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXFIFOHF)}}
}

func (p *SDIO_Periph) TXFIFOF() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXFIFOF)}}
}

func (p *SDIO_Periph) RXFIFOF() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXFIFOF)}}
}

func (p *SDIO_Periph) TXFIFOE() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXFIFOE)}}
}

func (p *SDIO_Periph) RXFIFOE() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXFIFOE)}}
}

func (p *SDIO_Periph) TXDAVL() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(TXDAVL)}}
}

func (p *SDIO_Periph) RXDAVL() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(RXDAVL)}}
}

func (p *SDIO_Periph) SDIOIT() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(SDIOIT)}}
}

func (p *SDIO_Periph) CEATAEND() RMSTA {
	return RMSTA{mmio.UM32{&p.STA.U32, uint32(CEATAEND)}}
}

type ICR uint32

func (b ICR) Field(mask ICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR) J(v int) ICR {
	return ICR(bits.MakeField32(v, uint32(mask)))
}

type RICR struct{ mmio.U32 }

func (r *RICR) Bits(mask ICR) ICR     { return ICR(r.U32.Bits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

func (r *RICR) AtomicStoreBits(mask, b ICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICR) AtomicSetBits(mask ICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICR) AtomicClearBits(mask ICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CCRCFAILC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CCRCFAILC)}}
}

func (p *SDIO_Periph) DCRCFAILC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DCRCFAILC)}}
}

func (p *SDIO_Periph) CTIMEOUTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CTIMEOUTC)}}
}

func (p *SDIO_Periph) DTIMEOUTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DTIMEOUTC)}}
}

func (p *SDIO_Periph) TXUNDERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(TXUNDERRC)}}
}

func (p *SDIO_Periph) RXOVERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(RXOVERRC)}}
}

func (p *SDIO_Periph) CMDRENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CMDRENDC)}}
}

func (p *SDIO_Periph) CMDSENTC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CMDSENTC)}}
}

func (p *SDIO_Periph) DATAENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DATAENDC)}}
}

func (p *SDIO_Periph) STBITERRC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(STBITERRC)}}
}

func (p *SDIO_Periph) DBCKENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(DBCKENDC)}}
}

func (p *SDIO_Periph) SDIOITC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(SDIOITC)}}
}

func (p *SDIO_Periph) CEATAENDC() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(CEATAENDC)}}
}

type MASK uint32

func (b MASK) Field(mask MASK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MASK) J(v int) MASK {
	return MASK(bits.MakeField32(v, uint32(mask)))
}

type RMASK struct{ mmio.U32 }

func (r *RMASK) Bits(mask MASK) MASK    { return MASK(r.U32.Bits(uint32(mask))) }
func (r *RMASK) StoreBits(mask, b MASK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMASK) SetBits(mask MASK)      { r.U32.SetBits(uint32(mask)) }
func (r *RMASK) ClearBits(mask MASK)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMASK) Load() MASK             { return MASK(r.U32.Load()) }
func (r *RMASK) Store(b MASK)           { r.U32.Store(uint32(b)) }

func (r *RMASK) AtomicStoreBits(mask, b MASK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMASK) AtomicSetBits(mask MASK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMASK) AtomicClearBits(mask MASK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMASK struct{ mmio.UM32 }

func (rm RMMASK) Load() MASK   { return MASK(rm.UM32.Load()) }
func (rm RMMASK) Store(b MASK) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) CCRCFAILIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CCRCFAILIE)}}
}

func (p *SDIO_Periph) DCRCFAILIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DCRCFAILIE)}}
}

func (p *SDIO_Periph) CTIMEOUTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CTIMEOUTIE)}}
}

func (p *SDIO_Periph) DTIMEOUTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DTIMEOUTIE)}}
}

func (p *SDIO_Periph) TXUNDERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXUNDERRIE)}}
}

func (p *SDIO_Periph) RXOVERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXOVERRIE)}}
}

func (p *SDIO_Periph) CMDRENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDRENDIE)}}
}

func (p *SDIO_Periph) CMDSENTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDSENTIE)}}
}

func (p *SDIO_Periph) DATAENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DATAENDIE)}}
}

func (p *SDIO_Periph) STBITERRIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(STBITERRIE)}}
}

func (p *SDIO_Periph) DBCKENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(DBCKENDIE)}}
}

func (p *SDIO_Periph) CMDACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CMDACTIE)}}
}

func (p *SDIO_Periph) TXACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXACTIE)}}
}

func (p *SDIO_Periph) RXACTIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXACTIE)}}
}

func (p *SDIO_Periph) TXFIFOHEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOHEIE)}}
}

func (p *SDIO_Periph) RXFIFOHFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOHFIE)}}
}

func (p *SDIO_Periph) TXFIFOFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOFIE)}}
}

func (p *SDIO_Periph) RXFIFOFIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOFIE)}}
}

func (p *SDIO_Periph) TXFIFOEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXFIFOEIE)}}
}

func (p *SDIO_Periph) RXFIFOEIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXFIFOEIE)}}
}

func (p *SDIO_Periph) TXDAVLIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(TXDAVLIE)}}
}

func (p *SDIO_Periph) RXDAVLIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(RXDAVLIE)}}
}

func (p *SDIO_Periph) SDIOITIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(SDIOITIE)}}
}

func (p *SDIO_Periph) CEATAENDIE() RMMASK {
	return RMMASK{mmio.UM32{&p.MASK.U32, uint32(CEATAENDIE)}}
}

type FIFOCNT uint32

func (b FIFOCNT) Field(mask FIFOCNT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FIFOCNT) J(v int) FIFOCNT {
	return FIFOCNT(bits.MakeField32(v, uint32(mask)))
}

type RFIFOCNT struct{ mmio.U32 }

func (r *RFIFOCNT) Bits(mask FIFOCNT) FIFOCNT { return FIFOCNT(r.U32.Bits(uint32(mask))) }
func (r *RFIFOCNT) StoreBits(mask, b FIFOCNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFIFOCNT) SetBits(mask FIFOCNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RFIFOCNT) ClearBits(mask FIFOCNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFIFOCNT) Load() FIFOCNT             { return FIFOCNT(r.U32.Load()) }
func (r *RFIFOCNT) Store(b FIFOCNT)           { r.U32.Store(uint32(b)) }

func (r *RFIFOCNT) AtomicStoreBits(mask, b FIFOCNT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFIFOCNT) AtomicSetBits(mask FIFOCNT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFIFOCNT) AtomicClearBits(mask FIFOCNT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFIFOCNT struct{ mmio.UM32 }

func (rm RMFIFOCNT) Load() FIFOCNT   { return FIFOCNT(rm.UM32.Load()) }
func (rm RMFIFOCNT) Store(b FIFOCNT) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) FIFOCOUNT() RMFIFOCNT {
	return RMFIFOCNT{mmio.UM32{&p.FIFOCNT.U32, uint32(FIFOCOUNT)}}
}

type FIFO uint32

func (b FIFO) Field(mask FIFO) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FIFO) J(v int) FIFO {
	return FIFO(bits.MakeField32(v, uint32(mask)))
}

type RFIFO struct{ mmio.U32 }

func (r *RFIFO) Bits(mask FIFO) FIFO    { return FIFO(r.U32.Bits(uint32(mask))) }
func (r *RFIFO) StoreBits(mask, b FIFO) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFIFO) SetBits(mask FIFO)      { r.U32.SetBits(uint32(mask)) }
func (r *RFIFO) ClearBits(mask FIFO)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFIFO) Load() FIFO             { return FIFO(r.U32.Load()) }
func (r *RFIFO) Store(b FIFO)           { r.U32.Store(uint32(b)) }

func (r *RFIFO) AtomicStoreBits(mask, b FIFO) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFIFO) AtomicSetBits(mask FIFO)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFIFO) AtomicClearBits(mask FIFO)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFIFO struct{ mmio.UM32 }

func (rm RMFIFO) Load() FIFO   { return FIFO(rm.UM32.Load()) }
func (rm RMFIFO) Store(b FIFO) { rm.UM32.Store(uint32(b)) }

func (p *SDIO_Periph) FIFODATA() RMFIFO {
	return RMFIFO{mmio.UM32{&p.FIFO.U32, uint32(FIFODATA)}}
}
