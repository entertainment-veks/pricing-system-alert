package domain

type DataBase interface {
	AddNoteBTC(*PriceNote)
	GetLastNoteBTC() (*PriceNote, error)

	AddNoteETH(*PriceNote)
	GetLastNoteETH() (*PriceNote, error)

	AddNoteBNB(*PriceNote)
	GetLastNoteBNB() (*PriceNote, error)
}
