package Person

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var Students []Person

func init() {
	Students = []Person{
		{
			Nama:      "Ivan Abdurrafie",
			Alamat:    "Ponorogo",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Golang asyik",
		}, {
			Nama:      "Budi",
			Alamat:    "malang",
			Pekerjaan: "progamer",
			Alasan:    "Golang top",
		}, {
			Nama:      "Andi",
			Alamat:    "swiss",
			Pekerjaan: "satpam",
			Alasan:    "Golang uye",
		}, {
			Nama:      "Siska",
			Alamat:    "Jakarta",
			Pekerjaan: "idol",
			Alasan:    "Golang uhui",
		}, {
			Nama:      "Muthe",
			Alamat:    "jogja",
			Pekerjaan: "polisi",
			Alasan:    "Golang mantul",
		}, {
			Nama:      "Alex",
			Alamat:    "surabaya",
			Pekerjaan: "dokter",
			Alasan:    "Golang enak",
		}, {
			Nama:      "Toni",
			Alamat:    "solo",
			Pekerjaan: "atlet",
			Alasan:    "Golang top",
		}, {
			Nama:      "Ringo",
			Alamat:    "Bandung",
			Pekerjaan: "Guru",
			Alasan:    "Golang epic",
		}, {
			Nama:      "Aiziz",
			Alamat:    "Semarang",
			Pekerjaan: "aktor",
			Alasan:    "Golang mudah",
		}, {
			Nama:      "marsha",
			Alamat:    "batam",
			Pekerjaan: "suster",
			Alasan:    "Golang menyenangkan",
		},
	}
}
