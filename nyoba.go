package main
import "fmt"

// =========================
//        MODEL
// =========================

type Pegawai struct {
    Jabatan   string
    MasaKerja int
    JumlahAnak int
}

type Request struct {
    Data Pegawai
}

type Response struct {
    Body string
}

type Handler interface {
    Execute(Request) Response
}

// =========================
//   HITUNG GAJI (LOGIC)
// =========================

func hitungGaji(p Pegawai) (int, int, int) {

    gajiPokok := 0
    tunjanganPerAnak := 0

    switch p.Jabatan {
    case "Staf":
        if p.MasaKerja < 5 {
            gajiPokok = 4000
        } else if p.MasaKerja > 10 {
            gajiPokok = 5000
            tunjanganPerAnak = 100
        } else {
            gajiPokok = 4000
        }

    case "Manager":
        if p.MasaKerja > 10 {
            gajiPokok = 10000
            tunjanganPerAnak = 300
        } else {
            gajiPokok = 8500
            tunjanganPerAnak = 300
        }

    case "Direktur":
        gajiPokok = 20000
        tunjanganPerAnak = 500
    }

    // Maksimal 3 anak
    anak := p.JumlahAnak
    if anak > 3 {
        anak = 3
    }

    totalTunjangan := tunjanganPerAnak * anak
    total := gajiPokok + totalTunjangan

    return gajiPokok, totalTunjangan, total
}

// =========================
//       HANDLER
// =========================

type GajiHandler struct{}

func (h GajiHandler) Execute(req Request) Response {

    pegawai := req.Data

    gajiPokok, tunjangan, total := hitungGaji(pegawai)

    out :=
        "=== HASIL PERHITUNGAN GAJI ===\n" +
            "Jabatan       : " + pegawai.Jabatan + "\n" +
            "Masa Kerja    : " + fmt.Sprint(pegawai.MasaKerja) + " tahun\n" +
            "Jumlah Anak   : " + fmt.Sprint(pegawai.JumlahAnak) + "\n" +
            "--------------------------------\n" +
            "Gaji Pokok    : " + fmt.Sprint(gajiPokok) + "\n" +
            "Tunjangan     : " + fmt.Sprint(tunjangan) + "\n" +
            "Total Gaji    : " + fmt.Sprint(total) + "\n"

    return Response{Body: out}
}

// =========================
//         ROUTER
// =========================

type Router struct {
    routes map[string]Handler
}

func NewRouter() *Router {
    return &Router{
        routes: make(map[string]Handler),
    }
}

func (r *Router) AddRoute(name string, h Handler) {
    r.routes[name] = h
}

func (r *Router) Run(name string, req Request) Response {
    return r.routes[name].Execute(req)
}

// =========================
//           MAIN
// =========================

func main() {

    var jabatan string
    var masaKerja int
    var anak int

    fmt.Print("Masukkan Jabatan (Staf/Manager/Direktur): ")
    fmt.Scan(&jabatan)

    fmt.Print("Masukkan Masa Kerja (tahun): ")
    fmt.Scan(&masaKerja)

    fmt.Print("Masukkan Jumlah Anak: ")
    fmt.Scan(&anak)

    pegawai := Pegawai{
        Jabatan: jabatan,
        MasaKerja: masaKerja,
        JumlahAnak: anak,
    }

    router := NewRouter()
    router.AddRoute("GAJI", GajiHandler{})

    res := router.Run("GAJI", Request{Data: pegawai})
    fmt.Println(res.Body)
}
