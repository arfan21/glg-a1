package data

import "fmt"

type ListBiodata struct {
	Data []Biodata
}

func NewListBiodata() *ListBiodata {
	return &ListBiodata{}
}

func (lb *ListBiodata) Add(data Biodata) {
	no := len(lb.Data)
	data.No = no + 1
	lb.Data = append(lb.Data, data)
}

func (lb *ListBiodata) GetDataByNo(no int) *Biodata {
	for _, val := range lb.Data {
		if val.No == no {
			{
				return &val
			}
		}
	}

	return nil
}

type Biodata struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
}

func NewBiodata(nama string, alamat string, pekerjaan string) Biodata {
	return Biodata{Nama: nama, Alamat: alamat, Pekerjaan: pekerjaan}
}

func (bd *Biodata) PrintData() {
	fmt.Println("Nama\t\t:", bd.Nama)
	fmt.Println("Alamat\t\t:", bd.Alamat)
	fmt.Println("Pekerjaan\t:", bd.Pekerjaan)
}
